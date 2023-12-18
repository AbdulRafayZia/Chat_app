package controller

import (
	"fmt"
	"net/http"

	"github.com/AbdulRafayZia/Gorilla-mux/dbinit"
)

func CreateTable(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement this method to create a new table in the database
	err := dbinit.CreateTables()
	if err != nil {
		fmt.Fprintln(w, err)
		return

	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User created successfully")

}
