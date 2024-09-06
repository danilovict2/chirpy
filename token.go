package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func getUserIDFromRequest(r *http.Request) (int, error) {
	tokenString, found := strings.CutPrefix(r.Header.Get("Authorization"), "Bearer ")
	if !found {
		return -1, fmt.Errorf("Please provide your token")
	}

	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return -1, err
	}

	userID, err := token.Claims.GetSubject()
	if err != nil {
		return -1, err
	}

	ID, err := strconv.Atoi(userID)
	if err != nil {
		return -1, err
	}

	return ID, nil
}