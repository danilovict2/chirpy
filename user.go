package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/danilovict2/chirpy/internal/database"
	"golang.org/x/crypto/bcrypt"
)

type RequestUserBody struct{
	Email string `json:"email"`
	Password string `json:"password"`
	ExpiresInSeconds *int `json:"expires_in_seconds"`
}

func createUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	body := RequestUserBody{}
	decoder.Decode(&body)

	db, err := database.NewDB("database.json")
	if err != nil {
		log.Fatal(err)
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		log.Fatal(err)
	}

	ret, err := db.CreateUser(body.Email, string(encryptedPassword))
	if err != nil {
		respondWithError(w, err.Error(), 400)
		return
	}

	respondWithJSON(w, ret, 201)
}