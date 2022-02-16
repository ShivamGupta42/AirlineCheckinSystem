package domain

import (
	"AirlineCheckinSystem/errors"
	"AirlineCheckinSystem/logger"
	"github.com/jmoiron/sqlx"
)

type SeatsDb struct {
	client *sqlx.DB
}

func (s *SeatsDb) FindById(id int) (*Seats, *errors.AppError) {
	selectSeatsSb := "Select seat_id from Seats"
	var seat Seats
	err := s.client.Get(&seat, selectSeatsSb)

	if err != nil {
		logger.Error("Error while fetching seat from db")
		return nil, errors.NewNotFoundError("Error while fetching seat from db")
	}
	return &seat, nil
}
