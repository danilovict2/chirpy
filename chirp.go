package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/danilovict2/chirpy/internal/database"
)

func createChirp(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	body := struct {
		Body string `json:"body"`
	}{}
	decoder.Decode(&body)

	if len(body.Body) > 140 {
		respondWithError(w, "Chirp is too long", 400)
		return
	}

	db, err := database.NewDB("database.json")
	if err != nil {
		log.Fatal(err)
	}

	ret, err := db.CreateChirp(cleanOfProfanity(body.Body))
	if err != nil {
		log.Fatal(err)
	}

	respondWithJSON(w, ret, 201)
}

func getChirps(w http.ResponseWriter, r *http.Request) {
	db, err := database.NewDB("database.json")
	if err != nil {
		log.Fatal(err)
	}

	ret, err := db.GetChirps()
	if err != nil {
		log.Fatal(err)
	}

	respondWithJSON(w, ret, 201)
}
