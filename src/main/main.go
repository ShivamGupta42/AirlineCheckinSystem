package main

import (
	app "AirlineCheckinSystem/src/app"
	domain "AirlineCheckinSystem/src/domain"
	"AirlineCheckinSystem/src/logger"
	service "AirlineCheckinSystem/src/service"
	"fmt"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

func main() {

	//Profiling server
	go func() {
		err := http.ListenAndServe(":5040", nil)
		logger.Fatal(err.Error())
	}()

	//app wiring
	client := app.GetDbClient()
	s := &service.DefaultBookingService{JourneyRepo: &domain.JourneyRepositoryDb{Client: client},
		UsersDb: &domain.UsersDb{Client: client}}

	//resetting database
	s.Reset()
	fmt.Println("DB RESET SUCCESS")

	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())
	n := 100
	wg.Add(n)
	for i := 1; i <= n; i++ {
		go s.BookASeat(i, 1, rand.Intn(100)+1, &wg)
	}
	wg.Wait()
	fmt.Println(s.AllJourneyStats(1))
}
