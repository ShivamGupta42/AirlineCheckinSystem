package service

import (
	"AirlineCheckinSystem/src/domain"
	"AirlineCheckinSystem/src/errors"
	"AirlineCheckinSystem/src/logger"
	"fmt"
	"strconv"
	"sync"
)

type BookingService interface {
	BookASeat(userId int, planeId int, seatId int, wg sync.WaitGroup) (*domain.Journey, *errors.AppError)
	AllJourneyStats(planeId int) (*domain.JourneyStats, *errors.AppError)
}

type DefaultBookingService struct {
	JourneyRepo *domain.JourneyRepositoryDb
	UsersDb     *domain.UsersDb
}

func (d *DefaultBookingService) Reset() {
	err := d.JourneyRepo.Reset()
	if err != nil {
		logger.Error("Error while resetting database")
		panic(err)
	}

	err = d.UsersDb.Reset()
	if err != nil {
		logger.Error("Error while resetting database")
		panic(err)
	}
}

func (d *DefaultBookingService) BookASeat(userId int, planeId int, seatId int, wg *sync.WaitGroup) (*domain.Journey, *errors.AppError) {
	defer wg.Done()
	journey, err := d.JourneyRepo.AddUserJourney(planeId, seatId, userId)
	if err != nil {
		logger.Error(err.Message + " User: " + strconv.FormatInt(int64(userId), 10) + " Seat: " + strconv.FormatInt(int64(seatId), 10))
		return nil, errors.NewNotFoundError("Error while fetching seat from db")
	} else {
		fmt.Printf("%d booked seat %d \n", userId, seatId)
		d.UsersDb.UpdateSeat(journey.UserId, journey.SeatId)
	}
	return journey, nil
}

func (d *DefaultBookingService) AllJourneyStats(planeId int) (*domain.JourneyStats, *errors.AppError) {
	journeyStats, err := d.JourneyRepo.AllJourneyStats(planeId)
	if err != nil {
		logger.Error("Error while adding journey for user ")
		return nil, errors.NewNotFoundError("Error while fetching seat from db")
	}
	return journeyStats, nil
}
