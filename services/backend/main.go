package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/priyanka-choubey/habitrick/handlers"
)

func main() {
	r := mux.NewRouter()

	userSubroute := r.PathPrefix("/user").Subrouter()
	userSubroute.HandleFunc("/", handlers.CreateUser).Methods("POST")

	// s := r.PathPrefix("/habit").Subrouter()
	// s.Use(middleware.Authorization)
	// s.HandleFunc("/", GetHabits).Methods("GET")
	// s.HandleFunc("/", CreateHabit).Methods("POST")
	// s.HandleFunc("/{id}", HabitTrack).Methods("GET")
	// http.Handle("/", r)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
