package domain

import (
	"AirlineCheckinSystem/errors"
	"AirlineCheckinSystem/logger"
	"github.com/jmoiron/sqlx"
)

type PlaneRepositoryDb struct {
	client *sqlx.DB
}

func (p *PlaneRepositoryDb) FindById(id int) (*Plane, *errors.AppError) {
	FindByID := "Select id from plane where id =?"
	var plane Plane
	err := p.client.Get(&plane, FindByID, id)

	if err != nil {
		logger.Error("Error while fetching plane : " + err.Error())
		return nil, errors.NewUnexpectedError("Unexpected database error")
	}

	return &plane, nil
}
