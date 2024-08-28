package main

import (
	"encoding/json"
	"net/http"
)

func validateChirp(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	body := struct {
		Body string `json:"body"`
	}{}
	decoder.Decode(&body)

	if len(body.Body) > 140 {
		respondWithError(w, "Chirp is too long", 400)
		return
	}

	ret := struct {
		CleanedBody string `json:"cleaned_body"`
	}{cleanOfProfanity(body.Body)}

	respondWithJSON(w, ret, 200)
}

func respondWithError(w http.ResponseWriter, msg string, code int) {
	type returnError struct {
		Error string `json:"error"`
	}

	dat, err := json.Marshal(returnError{msg})

	w.WriteHeader(code)

	if err != nil {
		w.Write([]byte("Error - something went wrong"))
		return
	}

	w.Write(dat)
}

func respondWithJSON(w http.ResponseWriter, payload interface{}, code int) {
	dat, err := json.Marshal(payload)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Error - something went wrong"))
		return
	}

	w.WriteHeader(code)
	w.Write(dat)
}
