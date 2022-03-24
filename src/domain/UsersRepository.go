package domain

import (
	"AirlineCheckinSystem/src/errors"
	"AirlineCheckinSystem/src/logger"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UsersDb struct {
	Client *sqlx.DB
}

func (u *UsersDb) Reset() *errors.AppError {
	ResetJourneyTable := "update users set seat_id =null"
	tx, err := u.Client.Begin()
	_, err = tx.Exec(ResetJourneyTable)
	if err != nil {
		tx.Rollback()
		//logger.Error("Error while updating journey : " + err.Error())
		return errors.NewUnexpectedError("Unable to update users table")
	}
	tx.Commit()
	return nil
}

func (u *UsersDb) FindByUserIO(id int) (*Users, *errors.AppError) {
	FindByUserID := "Select user_id, user_name from users where user_id= ?"
	var user Users
	err := u.Client.Get(&user, FindByUserID, id)

	if err != nil {
		logger.Error("Error while fetching user from db")
		return nil, errors.NewNotFoundError("Error while fetching user from db")
	}

	return &user, nil
}

func (u *UsersDb) UpdateSeat(userId, seatId int) *errors.AppError {
	tx, _ := u.Client.Begin()

	UpdateUserSeat := "Update users set seat_id=%d where user_id=%d"
	_, err := tx.Exec(fmt.Sprintf(UpdateUserSeat, seatId, userId))
	if err != nil {
		logger.Error(err.Error())
		return errors.NewUnexpectedError("error while updating user seat")
	}

	tx.Commit()
	return nil
}
