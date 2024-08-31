package database

import "math"

type User struct {
	ID   int `json:"id"`
	Email string `json:"email"`
	Password string `json:"password"`
}


func (db *DB) CreateUser(email, password string) (User, error) {
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