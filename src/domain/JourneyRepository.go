package domain

import (
	"AirlineCheckinSystem/src/errors"
	"AirlineCheckinSystem/src/logger"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type JourneyRepositoryDb struct {
	Client *sqlx.DB
}

func (j *JourneyRepositoryDb) Reset() *errors.AppError {
	ResetJourneyTable := "update journey set user_id =null, status=null"
	tx, err := j.Client.Begin()
	_, err = tx.Exec(ResetJourneyTable)
	if err != nil {
		tx.Rollback()
		//logger.Error("Error while updating journey : " + err.Error())
		return errors.NewUnexpectedError("Unable to update journey")
	}
	tx.Commit()
	return nil
}

func (j *JourneyRepositoryDb) FindById(id int) (*Journey, *errors.AppError) {
	FindByID := "SELECT id, plane_id, seat_id, user_id, status from journey where id =?"
	var journey Journey
	err := j.Client.Get(&journey, FindByID, id)

	if err != nil {
		logger.Error("Error while fetching plane : " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}

	return &journey, nil
}

func (j *JourneyRepositoryDb) AddUserJourney(planeId, seatId, userId int) (*Journey, *errors.AppError) {

	tx, err := j.Client.Begin()

	if err != nil {
		//logger.Error("Error while fetching transaction for journey : " + err.Error())
		return nil, errors.NewUnexpectedError("Error while getting transaction lock")
	}

	FindByID := "SELECT id, plane_id, seat_id from journey where user_id is null and seat_id=%d for share"

	var rows *sql.Rows
	rows, err = tx.Query(fmt.Sprintf(FindByID, seatId))

	if err != nil {
		tx.Rollback()
		//logger.Error("Error while updating journey : " + err.Error())
		return nil, errors.NewUnexpectedError("Error while executing query " + err.Error())
	}

	defer rows.Close()

	var journey Journey

	rowRead := false
	for rows.Next() {
		if err = rows.Scan(&journey.Id, &journey.PlaneId, &journey.SeatId); err != nil {
			tx.Rollback()
			//logger.Error("Error while fetching journey : " + err.Error())
			return nil, errors.NewUnexpectedError("Error while reading rows")
		} else {
			rowRead = true
		}
	}

	if err = rows.Err(); !rowRead || err != nil {
		tx.Rollback()
		//logger.Error("Error while fetching journey : " + GetNotNilString(err))
		return nil, errors.NewUnexpectedError("No rows to read error")
	}

	UpdateById := "UPDATE journey set user_id = %d, status = 'RESERVED' where id = %d"

	_, err = tx.Exec(fmt.Sprintf(UpdateById, userId, journey.Id))
	if err != nil {
		tx.Rollback()
		//logger.Error("Error while updating journey : " + err.Error())
		return nil, errors.NewUnexpectedError("Unable to update journey")
	}

	journey.UserId = userId
	tx.Commit()
	return &journey, nil

}

func GetNotNilString(err error) string {
	if err == nil {
		return ""
	} else {
		return err.Error()
	}
}

func (j *JourneyRepositoryDb) AllJourneyStats(planeId int) (*JourneyStats, *errors.AppError) {
	GetAllJourneys := "Select status from journey"

	journeyStats := JourneyStats{}

	var rows *sql.Rows
	var err error
	rows, err = j.Client.Query(GetAllJourneys)
	defer rows.Close()

	if err != nil {
		logger.Error("Error while getting journey stats: " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}

	for rows.Next() {
		var dbStatus sql.NullString
		if err = rows.Scan(&dbStatus); err != nil {
			logger.Error("Error while getting journey stats: " + err.Error())
			return nil, errors.NewUnexpectedError("Unexpected database error")
		}

		var status string
		if !dbStatus.Valid {
			status = "EMPTY"
		} else {
			status = dbStatus.String
		}

		switch status {
		case "FILLED":
			journeyStats.Filled = journeyStats.Filled + 1
		case "RESERVED":
			journeyStats.Reserved = journeyStats.Reserved + 1
		case "EMPTY":
			journeyStats.Empty = journeyStats.Empty + 1
		default:
		}
	}

	if err = rows.Err(); err != nil {
		logger.Error("Error while fetching journey : " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}

	return &journeyStats, nil

}
