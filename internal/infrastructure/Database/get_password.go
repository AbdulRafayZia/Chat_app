package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/AbdulRafayZia/Gorilla-mux/dbinit"
)

func GetPassword(username string) (string, error) {
	var hashedPassword string
	db := dbinit.OpenDB()

	defer db.Close()
	err := db.QueryRow("SELECT password FROM users WHERE username = $1", username).Scan(&hashedPassword)
	if err == sql.ErrNoRows {
		return "", fmt.Errorf("user not found")
	} else if err != nil {
		log.Printf("Error retrieving hashed password: %v", err)
		return "", err
	}
	return hashedPassword, nil
}