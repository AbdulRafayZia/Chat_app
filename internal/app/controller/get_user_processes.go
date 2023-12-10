package controller

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/AbdulRafayZia/Gorilla-mux/internal/app/validation"

	database "github.com/AbdulRafayZia/Gorilla-mux/internal/infrastructure/Database"
)

func GetUsersProcessses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "missing authorization header")
		return
	}
	tokenString = tokenString[len("Bearer"):]
	claims, err := validation.VerifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "could not Get claims")
		return
	}
	record := database.GetProcesses(claims)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(record)
}
