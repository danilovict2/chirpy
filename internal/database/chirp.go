package database

import "slices"

type Chirp struct {
	ID   int `json:"id"`
	Body string `json:"body"`
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

func (db *DB) CreateChirp(body string) (Chirp, error) {
	dbStruct, err := db.loadDB()
	if err != nil {
		return Chirp{}, err
	}

	chirpID := 0
	for id := range dbStruct.Chirps {
		chirpID = id
	}

	chirp := Chirp{chirpID + 1, body}
	dbStruct.Chirps[chirp.ID] = chirp

	err = db.writeDB(dbStruct)
	if err != nil {
		return Chirp{}, err
	}

	return chirp, nil
}