package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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

func getChirp(w http.ResponseWriter, r *http.Request) {
	db, err := database.NewDB("database.json")
	if err != nil {
		log.Fatal(err)
	}

	chirps, err := db.GetChirps()
	if err != nil {
		log.Fatal(err)
	}

	id, err := strconv.Atoi(r.PathValue("chirpID"))
	if err != nil {
		log.Fatal(err)
	}

	for _, chirp := range chirps {
		if chirp.ID == id {
			respondWithJSON(w, chirp, 200)
			return
		}
	}

	respondWithError(w, "Not found!", 404)
}
