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
	decoder.Decode(&body);

	if len(body.Body) > 140 {
		type returnError struct {
			Error string `json:"error"`
		}

		ret := returnError{
			"Chirp is too long",
		}

		dat, err := json.Marshal(ret)

		w.WriteHeader(400)
		if err != nil {
			w.Write([]byte("Error - something went wrong"))
			return
		}

		w.Write(dat)
		return
	}
	
	type returnSucess struct {
		Valid bool `json:"valid"`	
	}
	ret := returnSucess{
		true,
	}
	dat, err := json.Marshal(ret)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Error - something went wrong"))
		return
	}

	w.WriteHeader(200)
	w.Write(dat)
}