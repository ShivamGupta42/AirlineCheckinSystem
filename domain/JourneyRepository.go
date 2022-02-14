package domain

import (
	"AirlineCheckinSystem/errors"
	"AirlineCheckinSystem/logger"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type JourneyRepositoryDb struct {
	client *sqlx.DB
}

func (j *JourneyRepositoryDb) FindById(id int) (*Journey, *errors.AppError) {
	FindByID := "SELECT id, plane_id, seat_id, user_id, status from journey where id =?"
	var journey Journey
	err := j.client.Get(&journey, FindByID, id)

	if err != nil {
		logger.Error("Error while fetching plane : " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}

	return &journey, nil
}

func (j *JourneyRepositoryDb) AddUserJourney(planeId, seatId, userId int) *errors.AppError {

	tx, err := j.client.Begin()

	if err != nil {
		logger.Error("Error while fetching journey : " + err.Error())
		return errors.NewUnexpectedError("Unexpected database error")
	}

	FindByID := "SELECT id, plane_id, seat_id, user_id, status from journey For Update"

	var rows *sql.Rows
	rows, err = tx.Query(FindByID)
	defer rows.Close()

	var journeys []Journey

	for rows.Next() {
		var journey Journey
		if err = rows.Scan(&journey.Id, &journey.PlaneId, &journey.SeatId, &journey.UserId, &journey.Status); err != nil {
			tx.Rollback()
			logger.Error("Error while fetching journey : " + err.Error())
			return errors.NewUnexpectedError("Unexpected database error")
		}
	}

	if err = rows.Err(); err != nil {
		tx.Rollback()
		logger.Error("Error while fetching journey : " + err.Error())
		return errors.NewUnexpectedError("Unexpected database error")
	}

	UpdateById := "UPDATE table journey set user_id= ? AND set status = 'RESERVED' where id = ?"

	for _, journey := range journeys {
		if journey.Status != "FILLED" {
			_, err = tx.Exec(UpdateById, userId, journey.Id)
			if err != nil {
				tx.Rollback()
				logger.Error("Error while updating journey : " + err.Error())
				return errors.NewUnexpectedError("Unexpected database error")
			}
			break
		}
	}

	tx.Commit()
	return nil
}
