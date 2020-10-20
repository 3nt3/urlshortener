package handlers

import (
	"encoding/json"
	"github.com/3nt3/urlshortener/db"
	"github.com/3nt3/urlshortener/structs"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type loginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var user structs.User
	var loginData loginData
	var response structs.APIResponse

	if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
		log.Printf("[ - ] (handlers/user/RegisterUserHandler) error decoding: %s\n", err.Error())

		response = structs.APIResponse{Errors: []string{"error decoding JSON body"}}
		w.WriteHeader(http.StatusBadRequest)

		if err = json.NewEncoder(w).Encode(response); err != nil {
			// why would this even happen???
			log.Printf("[ - ] (handlers/user/RegisterUserHandler) error encoding response: %s\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(loginData.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("[ - ] (handlers/user/RegisterUserHandler) error hashing password: %s\n", err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		response = structs.APIResponse{Errors: []string{"internal server error"}}

		if err = json.NewEncoder(w).Encode(response); err != nil {
			// why would this even happen???
			log.Printf("[ - ] (handlers/user/RegisterUserHandler) error encoding response: %s\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		return
	}

	id, err := db.CreateUser(structs.User{Username: loginData.Username, Email: loginData.Email, Permission: 0, PasswordHash: string(hash)})
	if err != nil {
		log.Printf("[ - ] (handlers/user/RegisterUserHandler) error adding user to db: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		response = structs.APIResponse{Errors: []string{"internal server error"}}

		if err = json.NewEncoder(w).Encode(response); err != nil {
			// why would this even happen???
			log.Printf("[ - ] (handlers/user/RegisterUserHandler) error encoding response: %s\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		return
	}

	user, err = db.GetUserById(id)
	if err != nil {
		log.Printf("[ - ] (handlers/user/RegisterUserHandler) error retrieving user from db: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		response = structs.APIResponse{Errors: []string{"internal server error"}}

		if err = json.NewEncoder(w).Encode(response); err != nil {
			// why would this even happen???
			log.Printf("[ - ] (handlers/user/RegisterUserHandler) error encoding response: %s\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		return
	}

	user.PasswordHash = ""
	response = structs.APIResponse{Content: user}
	_ = json.NewEncoder(w).Encode(response)
	return
}

func Login(w http.ResponseWriter, r *http.Request) {
	var loginData loginData
	var response structs.APIResponse
	const identifier = "handlers/user/LoginHandler"

	if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
		response = structs.APIResponse{Errors: []string{"error decoding JSON body"}}
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(response)
		return
	}

	user, err := db.GetUserByUsername(loginData.Username)
	if err != nil {
		log.Printf("[ - ] (%s) error retrieving user '%s' from database: %s\n", identifier, loginData.Username, err.Error())
		response = structs.APIResponse{Errors: []string{"invalid credentials"}}
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(response)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(loginData.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("[ - ] (%s) error hashing password: %s\n", identifier, err.Error())
		w.WriteHeader(500)
		response = structs.APIResponse{Errors: []string{"internal server error"}}
		_ = json.NewEncoder(w).Encode(response)
		return
	}

	if err = bcrypt.CompareHashAndPassword(hash, []byte(user.PasswordHash)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			log.Printf("[ - ] (%s) invalid password!\n", identifier)
		} else {
			log.Printf("[ - ] (%s) error comparing hash and password: %s\n", identifier, err.Error())
		}

		// return with `invalid credentials` error no matter what to be more secure against brute-force?
		response = structs.APIResponse{Errors: []string{"invalid credentials"}}
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(response)
		return
	}

	// generate safe (without password hash and other sensitive data) map to return
	// in the future i guess there will be some helper function to strip the user struct of sensitive data??
	var safeUser = make(map[string]interface{})
	safeUser["username"] = user.Username
	safeUser["id"] = user.ID
	safeUser["permission"] = user.Permission
	safeUser["email"] = user.Email

	response = structs.APIResponse{Content: safeUser, Errors: nil}
	json.NewEncoder(w).Encode(response)

	// i think this is redundant but it gives me a good feeling at the end of the function :)
	w.WriteHeader(http.StatusOK)
}
