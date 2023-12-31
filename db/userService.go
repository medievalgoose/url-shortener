package db

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int
	Username string
	Password string
	Admin    bool
}

func CreateUser(username string, password string) bool {
	db := OpenConnection()
	defer db.Close()

	passwordByte := []byte(password)
	bcryptPassword, err := bcrypt.GenerateFromPassword(passwordByte, 14)
	if err != nil {
		log.Fatal(err)
	}

	createUserQuery := "INSERT INTO users (username, password, admin) VALUES ($1, $2, false);"
	_, err = db.Exec(createUserQuery, username, bcryptPassword)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func ValidateUser(username string, password string) bool {
	db := OpenConnection()
	defer db.Close()

	var requestedUser User

	getUserQuery := "SELECT id, username, password, admin FROM users WHERE username = $1;"
	err := db.QueryRow(getUserQuery, username).Scan(&requestedUser.Id, &requestedUser.Username, &requestedUser.Password, &requestedUser.Admin)
	if err != nil {
		log.Fatal(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(requestedUser.Password), []byte(password))
	return err == nil
}
