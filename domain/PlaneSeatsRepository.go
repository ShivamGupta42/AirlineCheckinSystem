package domain

import (
	"github.com/jmoiron/sqlx"
)

type PlaneSeats struct {
	client *sqlx.DB
}
