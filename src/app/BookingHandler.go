package app

//type BookingHandler struct {
//	service service.BookingService
//}

//func (b *BookingHandler) BookPlaneSeat(w http.ResponseWriter, r *http.Request) {
//	m := mux.Vars(r)
//	userId, ok := m["userId"]
//	valueExists("userId", userId, ok)
//
//	plane, ok1 := m["plane"]
//	valueExists("planeId", plane, ok1)
//
//	seat, ok2 := m["seat"]
//	valueExists("seat", seat, ok2)
//
//}
//
//func valueExists(k string, v string, ok bool) {
//	if !ok || strings.TrimSpace(v) == "" {
//		logger.Error("Request Doesn't have " + k)
//		errors.NewNotFoundError("UserId not found")
//	}
//}
