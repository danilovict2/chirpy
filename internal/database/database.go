package database

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"sync"
)

type DB struct {
	path          string
	mux           *sync.RWMutex
}

type DBStructure struct {
	Chirps map[int]Chirp `json:"chirps"`
	Users map[int]User `json:"users"`
}


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
		map[int]User{},
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