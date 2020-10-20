package main

import (
	"log"
	"net/http"

	"github.com/3nt3/urlshortener/db"
	"github.com/3nt3/urlshortener/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	db.Init()

	r.HandleFunc("/new-url/", handlers.CreateShortURL).Methods("POST")
	r.HandleFunc("/{id}", handlers.AccessShortURL).Methods("POST")
	r.HandleFunc("/user/register", handlers.RegisterUserHandler).Methods("POST")

	log.Printf("[ ~ ] starting server on port 8080")
	log.Panic(http.ListenAndServe(":8080", r))
}
