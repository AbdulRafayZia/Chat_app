package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	database "github.com/AbdulRafayZia/Gorilla-mux/internal/infrastructure/Database"
	"github.com/AbdulRafayZia/Gorilla-mux/utils"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request parameters
	w.Header().Set("Content-Type", "application/json")
	fmt.Printf("The request body is %v\n", r.Body)

	var request utils.Credentials
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "unable to get data", http.StatusBadRequest)
	}

	//extract data from request
	username := request.Username
	password := request.Password

	//Insert into database
	err = database.CreateUser(username, password)
	if err != nil {
		http.Error(w, "unable to create user", http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User created successfully")
}
