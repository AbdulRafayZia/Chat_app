package database

import (
	"database/sql"
	"fmt"

	"github.com/AbdulRafayZia/Gorilla-mux/dbinit"
)

func GetRole(name string) (string,uint, error) {
	var Role string
	var id uint

	db := dbinit.OpenDB()

	defer db.Close()

	err := db.QueryRow("SELECT role,id  FROM users WHERE username = $1", name).Scan(&Role , &id) 
	if err == sql.ErrNoRows {

		return "",0, fmt.Errorf("no role for this name")
	} else if err != nil {

		return "",0, fmt.Errorf("error retrieving Role ")
	}
	

	return Role,id, nil

}
