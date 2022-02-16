package domain

import (
	"AirlineCheckinSystem/errors"
	"AirlineCheckinSystem/logger"
	"github.com/jmoiron/sqlx"
)

type UsersDb struct {
	client *sqlx.DB
}

func (u *UsersDb) FindByUserIO(id int) (*Users, *errors.AppError) {
	FindByUserID := "Select user_id, user_name from users where user_id= ?"
	var user Users
	err := u.client.Get(&user, FindByUserID, id)

	if err != nil {
		logger.Error("Error while fetching user from db")
		return nil, errors.NewNotFoundError("Error while fetching user from db")
	}

	return &user, nil
}
