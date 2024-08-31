package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/danilovict2/chirpy/internal/database"
	"golang.org/x/crypto/bcrypt"
)

func login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	body := RequestUserBody{}
	decoder.Decode(&body)

	db, err := database.NewDB("database.json")
	if err != nil {
		log.Fatal(err)
	}

	user, err := db.GetUserFromEmail(body.Email)
	if err != nil {
		respondWithError(w, err.Error(), 404)
		return
	} 

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		respondWithError(w, "Invalid password", 401)
		return
	}

	respondWithJSON(w, user, 200)
}