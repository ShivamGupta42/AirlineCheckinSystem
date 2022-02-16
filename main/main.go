package main

import (
	"AirlineCheckinSystem/app"
	"AirlineCheckinSystem/domain"
	service "AirlineCheckinSystem/service"
	"fmt"
	"sync"
)

func main() {

	//wiring
	client := app.GetDbClient()
	//
	s := &service.DefaultBookingService{JourneyRepo: &domain.JourneyRepositoryDb{Client: client}}
	var wg sync.WaitGroup
	wg.Add(100)

	for i := 1; i <= 100; i++ {
		go s.BookASeat(i, 1, 1, wg)
	}

	wg.Wait()
	fmt.Println(s.AllJourneyStats(1))
}
