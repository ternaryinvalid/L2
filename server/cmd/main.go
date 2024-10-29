package main

import (
	"L2/server/app"
	"L2/server/handlers"
	"L2/server/middleware"
	"L2/server/repository"
	"log"
	"net/http"
)

var port = "8081"

func main() {
	mux := http.NewServeMux()
	cache := repository.NewCache()

	mux.HandleFunc("/create_event", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateEventHandler(w, r, cache)
	})
	mux.HandleFunc("/events_for_day", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetEventDayHandler(w, r, cache)
	})
	mux.HandleFunc("/delete_event", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteEventHandler(w, r, cache)
	})
	mux.HandleFunc("/events_for_week", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetEventWeekHandler(w, r, cache)
	})
	mux.HandleFunc("/events_for_month", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetEventMonthHandler(w, r, cache)
	})
	mux.HandleFunc("/update_event", func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateEventHandler(w, r, cache)
	})

	handler := middleware.Logger(mux)
	server := app.NewServer(port, handler)

	err := server.Run()

	if err != nil {
		log.Printf("error with listen and serve: : %v", err)
	}

}
