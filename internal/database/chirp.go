package database

import (
	"math"
	"slices"
)

type Chirp struct {
	ID   int `json:"id"`
	Body string `json:"body"`
	AuthorID int `json:"author_id"`
}

func (db *DB) GetChirps() ([]Chirp, error) {
	dbStruct, err := db.loadDB()
	if err != nil {
		return nil, err
	}

	chirps := make([]Chirp, 0, len(dbStruct.Chirps))
	for _, chirp := range dbStruct.Chirps {
		chirps = append(chirps, chirp)
	}

	slices.SortFunc(chirps, func (a, b Chirp) int {
		switch {
		case a.ID > b.ID:
			return 1
		case a.ID < b.ID:
			return -1
		default:
			return 0
		}
	})

	return chirps, nil
}

func (db *DB) CreateChirp(body string, authorID int) (Chirp, error) {
	dbStruct, err := db.loadDB()
	if err != nil {
		return Chirp{}, err
	}

	chirpID := 0
	for id := range dbStruct.Chirps {
		chirpID = int(math.Max(float64(id), float64(chirpID)))
	}

	chirp := Chirp{chirpID + 1, body, authorID}
	dbStruct.Chirps[chirp.ID] = chirp

	err = db.writeDB(dbStruct)
	if err != nil {
		return Chirp{}, err
	}

	return chirp, nil
}