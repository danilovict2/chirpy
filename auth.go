package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/danilovict2/chirpy/internal/database"
	"github.com/golang-jwt/jwt/v5"
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

	expiresAt := time.Now().Add(24 * time.Hour)
	if body.ExpiresInSeconds != nil && (time.Duration(*body.ExpiresInSeconds) * time.Second).Seconds() < (24 * time.Hour).Seconds() {
		expiresAt = time.Now().Add(time.Duration(*body.ExpiresInSeconds) * time.Second * time.Hour)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer: "chirpy",
		IssuedAt: &jwt.NumericDate{time.Now()},
		ExpiresAt: &jwt.NumericDate{expiresAt},
		Subject: strconv.Itoa(user.ID),
	})

	tkn, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		respondWithError(w, "Something went wrong", 500)
		log.Fatal(err)
	}
	
	ret := struct{
		ID int `json:"id"`
		Email string `json:"email"`
		Token string `json:"token"`
	} {
		user.ID,
		user.Email,
		tkn,
	}

	respondWithJSON(w, ret, 200)
}