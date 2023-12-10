package controller

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/AbdulRafayZia/Gorilla-mux/internal/app/validation"
	database "github.com/AbdulRafayZia/Gorilla-mux/internal/infrastructure/Database"
	"github.com/AbdulRafayZia/Gorilla-mux/pkg/jwt"
	"github.com/AbdulRafayZia/Gorilla-mux/utils"
)

func Statistics(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var request utils.StatsRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "unable to get data", http.StatusBadRequest)
		return
	}
	tokenString, err := jwt.GetToken(w, r)
	if tokenString == "" || err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, "could not provide autherization bearer", http.StatusUnauthorized)
		return
	}

	Claims, err := validation.VerifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, " could not get Claims ", http.StatusUnauthorized)

		return
	}
	validrole := validation.CheckStaffRole(Claims.Role)
	if !validrole {
		http.Error(w, "not a Staff Member", http.StatusUnauthorized)
		return

	}

	average_execution_time, err := database.GetAvergeExeTime(request)
	if err != nil {
		http.Error(w, "could not fetch average execution time from database", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(average_execution_time)
}
