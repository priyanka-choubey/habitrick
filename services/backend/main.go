package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/priyanka-choubey/habitrick/services/backend/handlers"
	"github.com/priyanka-choubey/habitrick/services/backend/middleware"
)

func main() {
	r := mux.NewRouter()

	userSubroute := r.PathPrefix("/user").Subrouter()
	userSubroute.HandleFunc("/", handlers.CreateUser).Methods("POST")

	userDeleteSubroute := userSubroute.PathPrefix("/delete").Subrouter()
	userDeleteSubroute.Use(middleware.Authorization)
	userDeleteSubroute.HandleFunc("/", handlers.DeleteUser).Methods("DELETE")

	habitSubroute := r.PathPrefix("/habit").Subrouter()
	habitSubroute.Use(middleware.Authorization)
	// habitSubroute.HandleFunc("/", GetHabits).Methods("GET")
	habitSubroute.HandleFunc("/", handlers.CreateHabit).Methods("POST")
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
