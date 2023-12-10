package database

import (
	"database/sql"
	"fmt"

	"github.com/AbdulRafayZia/Gorilla-mux/dbinit"
)

func GetRole(name string) (string, error) {
	var Role string

	db := dbinit.OpenDB()

	defer db.Close()

	err := db.QueryRow("SELECT role FROM users WHERE username = $1", name).Scan(&Role)
	if err == sql.ErrNoRows {

		return "", fmt.Errorf("no role for this name")
	} else if err != nil {

		return "", fmt.Errorf("error retrieving Role ")
	}

	return Role, nil

}
