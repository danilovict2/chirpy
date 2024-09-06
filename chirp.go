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

	ID, err := getUserIDFromRequest(r)
	if err != nil {
		respondWithError(w, err.Error(), 401)
		return
	}

	db, err := database.NewDB("database.json")
	if err != nil {
		log.Fatal(err)
	}

	ret, err := db.CreateChirp(cleanOfProfanity(body.Body), ID)
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

	id, err := strconv.Atoi(r.PathValue("chirpID"))
	if err != nil {
		log.Fatal(err)
	}

	chirp, err := db.GetChirp(id)
	if err != nil {
		log.Fatal(err)
	}

	if chirp == (database.Chirp{}) {
		respondWithError(w, "Not found!", 404)
		return
	}

	respondWithJSON(w, chirp, 200)
}

func deleteChirp(w http.ResponseWriter, r *http.Request) {
	db, err := database.NewDB("database.json")
	if err != nil {
		log.Fatal(err)
	}

	chirpID, err := strconv.Atoi(r.PathValue("chirpID"))
	if err != nil {
		log.Fatal(err)
	}

	chirp, err := db.GetChirp(chirpID)
	if err != nil {
		log.Fatal(err)
	}

	if chirp == (database.Chirp{}) {
		respondWithError(w, "Not found!", 404)
		return
	}

	authorID, err := getUserIDFromRequest(r)
	if err != nil {
		respondWithError(w, err.Error(), 500)
		return
	}

	if authorID != chirp.AuthorID {
		respondWithError(w, "Forbidden", 403)
		return
	}

	err = db.DeleteChirp(chirpID)
	if err != nil {
		respondWithError(w, err.Error(), 500)
	}

	respondWithJSON(w, "", 204);
}
