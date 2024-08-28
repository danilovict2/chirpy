package main

import "strings"

func cleanOfProfanity(msg string) string {
	words := []string{"kerfuffle", "sharbert", "fornax"}
	splitted := strings.Split(msg, " ")

	for i, str := range splitted {
		lower := strings.ToLower(str)
		for _, word := range words {
			if lower == word {
				splitted[i] = "****"
				break
			} 
		}
	}

	return strings.Join(splitted, " ")
}
