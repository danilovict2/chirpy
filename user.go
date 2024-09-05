package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/danilovict2/chirpy/internal/database"
	"github.com/golang-jwt/jwt/v5"
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

func updateUser(w http.ResponseWriter, r *http.Request) {
	tokenString, found := strings.CutPrefix(r.Header.Get("Authorization"), "Bearer ")
	if !found {
		respondWithError(w, "Please provide your token", 401)
		return
	}

	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		respondWithError(w, err.Error(), 401)
		return
	}

	userID, err := token.Claims.GetSubject()
	if err != nil {
		respondWithError(w, err.Error(), 500)
		return
	}

	ID, err := strconv.Atoi(userID)
	if err != nil {
		respondWithError(w, err.Error(), 500)
		return
	}

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

	user, err := db.UpdateUser(ID, body.Email, string(encryptedPassword))
	if err != nil {
		respondWithError(w, err.Error(), 400)
		return
	}

	respondWithJSON(w, user, 200)
}