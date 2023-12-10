package controller

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/AbdulRafayZia/Gorilla-mux/internal/app/validation"
	"github.com/AbdulRafayZia/Gorilla-mux/pkg/jwt"
	"github.com/AbdulRafayZia/Gorilla-mux/utils"

	database "github.com/AbdulRafayZia/Gorilla-mux/internal/infrastructure/Database"
)

func GetProcessByName(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var request utils.Username
	json.NewDecoder(r.Body).Decode(&request)
	username := request.Username
	tokenString, err := jwt.GetToken(w, r)
	if tokenString == "" || err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, "could not provide autherization bearer", http.StatusUnauthorized)
		return
	}
	claims, err := validation.VerifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, " could not get claims ", http.StatusUnauthorized)

		return
	}
	validRole := validation.CheckStaffRole(claims.Role)
	if !validRole {
		http.Error(w, "not a Staff Member", http.StatusUnauthorized)
		return

	}
	record, err := database.GetProcessesByUserName(username)
	if err != nil {
		http.Error(w, "could not get proceseses from database", http.StatusUnauthorized)
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(record)

}
