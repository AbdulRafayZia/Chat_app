package database

import (
	"fmt"

	"github.com/AbdulRafayZia/Gorilla-mux/dbinit"
)

func CreateUser( username string , password string) (error) {
	db := dbinit.OpenDB()
	defer db.Close()
	_, err := db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", username, password)
	if err != nil {
		fmt.Println(err)
		return  err
	}
	return nil
}