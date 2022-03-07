package domain

import (
	"AirlineCheckinSystem/src/errors"
	"AirlineCheckinSystem/src/logger"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type JourneyRepositoryDb struct {
	Client *sqlx.DB
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
		logger.Error("Error while fetching journey : " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}

	FindByID := "SELECT id, plane_id, seat_id from journey where user_id is null"

	var rows *sql.Rows
	rows, err = tx.Query(FindByID)
	defer rows.Close()

	var journeys []Journey

	for rows.Next() {
		var journey Journey
		if err = rows.Scan(&journey.Id, &journey.PlaneId, &journey.SeatId); err != nil {
			tx.Rollback()
			logger.Error("Error while fetching journey : " + err.Error())
			return nil, errors.NewUnexpectedError("Unexpected database error")
		}
	}

	if err = rows.Err(); err != nil {
		tx.Rollback()
		logger.Error("Error while fetching journey : " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}

	UpdateById := "UPDATE table journey set user_id= ? AND set status = 'RESERVED' where id = ?"

	var userJourney Journey

	for _, journey := range journeys {
		if journey.Status != "FILLED" {
			_, err = tx.Exec(UpdateById, userId, journey.Id)
			if err != nil {
				tx.Rollback()
				logger.Error("Error while updating journey : " + err.Error())
				return nil, errors.NewUnexpectedError("Unexpected database error")
			}
			userJourney = journey
			break
		}
	}

	tx.Commit()
	return &userJourney, nil

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
		var status string
		if err = rows.Scan(&status); err != nil {
			logger.Error("Error while getting journey stats: " + err.Error())
			return nil, errors.NewUnexpectedError("Unexpected database error")
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
