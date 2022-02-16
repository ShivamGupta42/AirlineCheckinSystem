package domain

import "AirlineCheckinSystem/errors"

type Users struct {
	UserId   string `db:"user_id"`
	USerName string `db:"user_name"`
}

type UserRepository interface {
	FindByUserIO(id int) (*Users, *errors.AppError)
}
