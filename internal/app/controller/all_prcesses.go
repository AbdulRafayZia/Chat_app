package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AbdulRafayZia/Gorilla-mux/internal/app/validation"
	database "github.com/AbdulRafayZia/Gorilla-mux/internal/infrastructure/Database"
	"github.com/AbdulRafayZia/Gorilla-mux/pkg/jwt"
)

func GetAllProcesses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	tokenString, err := jwt.GetToken(w, r)
	if tokenString == "" || err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, "could not provide autherization bearer", http.StatusUnauthorized)
		return
	}
	claims, err := validation.VerifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Could not Get Claims")
		return
	}

	validRole := validation.CheckStaffRole(claims.Role)
	if !validRole {
		http.Error(w, "Not a Staff Member", http.StatusUnauthorized)
		return

	}
	record := database.GetAllProcesses()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(record)

}
