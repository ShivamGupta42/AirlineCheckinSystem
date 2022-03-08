package main

import (
	app "AirlineCheckinSystem/src/app"
	domain "AirlineCheckinSystem/src/domain"
	service "AirlineCheckinSystem/src/service"
	"fmt"
	"sync"
)

func main() {

	//wiring
	client := app.GetDbClient()
	//
	s := &service.DefaultBookingService{JourneyRepo: &domain.JourneyRepositoryDb{Client: client}}
	var wg sync.WaitGroup

	n := 10
	wg.Add(n)
	for i := 1; i <= n; i++ {
		go s.BookASeat(i, 1, 1, &wg)
	}
	wg.Wait()
	fmt.Println(s.AllJourneyStats(1))
}
