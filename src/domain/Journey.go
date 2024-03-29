package domain

import "fmt"

type Status int

const (
	RESERVED Status = iota + 1
	FILLED
	EMPTY
)

func (s Status) String() string {
	return [...]string{"RESERVED", "FILLED", "EMPTY"}[s-1]
}

func (s Status) EnumIndex() int {
	return int(s)
}

type Journey struct {
	Id      int    `db:"id"`
	PlaneId int    `db:"plane_id"`
	SeatId  int    `db:"seat_id"`
	UserId  int    `db:"user_id"`
	Status  string `db:"status"`
}

type JourneyStats struct {
	Empty    int
	Filled   int
	Reserved int
}

func (j *JourneyStats) String() string {
	return fmt.Sprintf("Empty Seats : %d\nReserved Seats : %d\n", j.Empty, j.Reserved)
}
