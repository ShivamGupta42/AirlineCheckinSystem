package domain

import "AirlineCheckinSystem/src/errors"

type Users struct {
	UserId   string `db:"user_id"`
	USerName string `db:"user_name"`
	SeatId   string `db:"seat_id"`
}

type UserRepository interface {
	FindByUserIO(id int) (*Users, *errors.AppError)
	UpdateSeat(id, seatId int) *errors.AppError
}
