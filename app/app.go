package app

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"time"
)

//func StartServer() {
//	//wiring
//	client := GetDbClient()
//	//
//	b := BookingHandler{&service.DefaultBookingService{JourneyRepo: &domain.JourneyRepositoryDb{Client: client}}}
//	//routing
//	r := mux.NewRouter()
//	r.HandleFunc("/user/{userId}/book/plane/{plane}/seat/{seat}", b.BookPlaneSeat)
//	log.Fatal(http.ListenAndServe(":70001", r))
//}

func GetDbClient() *sqlx.DB {
	dbUser := "postgres"
	dbPasswd := "postgres"
	dbAddr := "localhost"
	dbPort := "6432"
	dbName := "postgres"

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	client, err := sqlx.Open("postgres", dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
