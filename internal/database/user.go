package database

type User struct {
	ID   int `json:"id"`
	Email string `json:"email"`
}

var userID int = 0

func (db *DB) CreateUser(email string) (User, error) {
	dbStruct, err := db.loadDB()
	if err != nil {
		return User{}, err
	}

	userID++
	user := User{userID, email}
	dbStruct.Users[user.ID] = user

	err = db.writeDB(dbStruct)
	if err != nil {
		return User{}, err
	}

	return user, nil
}