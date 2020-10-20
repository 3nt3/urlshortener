package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/3nt3/urlshortener/db"
	"github.com/3nt3/urlshortener/structs"
)

func CreateShortURL(w http.ResponseWriter, r *http.Request) {
	var url structs.ShortURL
	err := json.NewDecoder(r.Body).Decode(&url)

	var response structs.APIResponse

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		response = structs.APIResponse{Errors: []string{"body is not JSON serializable"}}
		log.Printf("[ - ] (handlers.CreateShourtURL) error decoding json: %s\n", err.Error())

		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			log.Printf("[ - ] error encoding error message???: %s\n", err.Error())
		}
		return
	}

	id, err := db.CreateURL(url)
	if err != nil {
		log.Printf("[ - ] error creating url: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		response = structs.APIResponse{nil, []string{"internal server error"}}
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			log.Printf("[ - ] error encoding error message???: %s\n", err.Error())
		}
		return
	}

	url, err = db.GetURLByID(id)
	if err != nil {
		log.Printf("[ - ] error retrieving url from db: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		response = structs.APIResponse{nil, []string{"internal server error"}}
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			log.Printf("[ - ] error encoding error message???: %s\n", err.Error())
		}
		return
	}

	log.Printf("[ * ] url: %+v\n", url)
	response = structs.APIResponse{url, nil}
	json.NewEncoder(w).Encode(response)
}

func AccessShortURL(w http.ResponseWriter, r *http.Request) {
	var response structs.APIResponse

	id, ok := mux.Vars(r)["id"]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	url, err := db.GetURLByID(id)
	if err == sql.ErrNoRows {
		log.Printf("[ - ] short link does not exist.\n")
		w.WriteHeader(404)
		response = structs.APIResponse{nil, []string{"redirection entry does not exist"}}
		json.NewEncoder(w).Encode(response)
		return
	}

	if err != nil {
		log.Printf("[ - ] error retrieving url from db: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		response = structs.APIResponse{nil, []string{"internal server error"}}
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			log.Printf("[ - ] error encoding error message???: %s\n", err.Error())
		}
		return
	}

	log.Printf("[ * ] url: %+v\n", url)
	http.Redirect(w, r, url.OriginalURL, 303)
}
