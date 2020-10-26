package main

import (
	"log"
	"net/http"

	"github.com/3nt3/urlshortener/db"
	"github.com/3nt3/urlshortener/handlers"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("-~- starting API v0.1 -~-")
	r := mux.NewRouter()

	err := db.Init()
	if err != nil {
		return
	}

	r.HandleFunc("/url", handlers.CreateShortURL).Methods("POST", "OPTIONS")
	r.HandleFunc("/url/{id}", handlers.AccessShortURL).Methods("GET")

	r.HandleFunc("/user/register", handlers.RegisterUserHandler).Methods("POST", "OPTIONS")

	r.HandleFunc("/metrics", handlers.MetricsHandler).Methods("GET")

	log.Printf("[ ~ ] starting server on port 8080")
	log.Panic(http.ListenAndServe(":8080", r))
}
