package database

import (
	"fmt"
	"math"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (db *DB) CreateUser(email, password string) (User, error) {
	if _, err := db.GetUserFromEmail(email); err == nil {
		return User{}, fmt.Errorf("User with this email already exists")
	}
	
	dbStruct, err := db.loadDB()
	if err != nil {
		return User{}, err
	}

	userID := 0
	for id := range dbStruct.Users {
		userID = int(math.Max(float64(id), float64(userID)))
	}

	user := User{userID + 1, email, password}
	dbStruct.Users[user.ID] = user

	err = db.writeDB(dbStruct)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (db *DB) GetUserFromEmail(email string) (User, error) {
	dbStruct, err := db.loadDB()
	if err != nil {
		return User{}, err
	}

	for _, user := range dbStruct.Users {
		if user.Email == email {
			return user, nil
		}
	}

	return User{}, fmt.Errorf("User does not exist")
}

func (db *DB) UpdateUser(ID int, email, password string) (User, error) {
	dbStruct, err := db.loadDB()
	if err != nil {
		return User{}, err
	}

	user, exists := dbStruct.Users[ID]
	if !exists {
		return User{}, fmt.Errorf("User does not exist")
	}

	user.Email = email
	user.Password = password
	dbStruct.Users[user.ID] = user

	err = db.writeDB(dbStruct)
	if err != nil {
		return User{}, err
	}

	return user, nil
}