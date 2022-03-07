package domain

import "AirlineCheckinSystem/src/errors"

type Seats struct {
	id int `db:"seat_id"`
}

type SeatRepository interface {
	FindById(id int) (*Seats, *errors.AppError)
}
