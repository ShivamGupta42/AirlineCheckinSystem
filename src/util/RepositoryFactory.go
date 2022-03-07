package util

import "AirlineCheckinSystem/src/domain"

type RepositoryFactory struct {
	planeDb   *domain.PlaneRepositoryDb
	usersDb   *domain.UsersDb
	seatsDb   *domain.SeatsDb
	journeyDb *domain.JourneyRepositoryDb
}

func (r *RepositoryFactory) init() {
	r.planeDb = &domain.PlaneRepositoryDb{}
	r.usersDb = &domain.UsersDb{}
	r.seatsDb = &domain.SeatsDb{}
	r.journeyDb = &domain.JourneyRepositoryDb{}
}
