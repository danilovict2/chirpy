package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/danilovict2/chirpy/internal/database"
)

func createUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	body := struct {
		Email string `json:"email"`
	}{}
	decoder.Decode(&body)

	db, err := database.NewDB("database.json")
	if err != nil {
		log.Fatal(err)
	}

	ret, err := db.CreateUser(body.Email)
	if err != nil {
		log.Fatal(err)
	}

	respondWithJSON(w, ret, 201)
}