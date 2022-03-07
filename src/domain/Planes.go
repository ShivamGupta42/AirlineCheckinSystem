package domain

import "AirlineCheckinSystem/src/errors"

type Plane struct {
	id int `db:"plane_id"`
}

type PlaneRepository interface {
	FindById(id int) (Plane, *errors.AppError)
}
