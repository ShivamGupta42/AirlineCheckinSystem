package domain

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
