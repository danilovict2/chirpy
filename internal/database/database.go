package database

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"slices"
	"sync"
)

type DB struct {
	path          string
	mux           *sync.RWMutex
}

type Chirp struct {
	ID   int `json:"id"`
	Body string `json:"body"`
}

type DBStructure struct {
	Chirps map[int]Chirp `json:"chirps"`
}

var id int = 0

func NewDB(path string) (*DB, error) {
	db := &DB{
		path:          path,
		mux:           &sync.RWMutex{},
	}

	db.ensureDB()
	return db, nil
}

// ensureDB creates a new database file if it doesn't exist
func (db *DB) ensureDB() error {
	db.mux.Lock()
	defer db.mux.Unlock()

	if _, err := os.Stat(db.path); errors.Is(err, os.ErrNotExist) {
		err := os.WriteFile(db.path, []byte("{}"), 0666)
		if err != nil {
			return err
		}
	}

	return nil
}

func (db *DB) CreateChirp(body string) (Chirp, error) {
	chirps, err := db.loadDB()
	if err != nil {
		return Chirp{}, err
	}

	id++
	chirp := Chirp{id, body}
	chirps.Chirps[chirp.ID] = chirp

	err = db.writeDB(chirps)
	if err != nil {
		return Chirp{}, err
	}


	return chirp, nil
}

func (db *DB) loadDB() (DBStructure, error) {
	db.mux.RLock()
	defer db.mux.RUnlock()

	data, err := os.ReadFile(db.path)
	if err != nil {
		return DBStructure{}, err
	}

	decoder := json.NewDecoder(bytes.NewReader(data))
	chirps := DBStructure{
		map[int]Chirp{},
	}
	decoder.Decode(&chirps)

	return chirps, nil
}


func (db *DB) writeDB(dbStructure DBStructure) error {
	db.mux.Lock()
	defer db.mux.Unlock()

	json, err := json.Marshal(dbStructure)
	if err != nil {
		return err
	}

	err = os.WriteFile(db.path, json, 0666)
	if err != nil {
		return err
	}

	return nil
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