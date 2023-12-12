package controller

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/AbdulRafayZia/Gorilla-mux/internal/app/validation"
	"github.com/AbdulRafayZia/Gorilla-mux/pkg/jwt"

	database "github.com/AbdulRafayZia/Gorilla-mux/internal/infrastructure/Database"
	"github.com/gorilla/mux"
)

func PrivateChatHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	tokenString, err := jwt.GetToken(w, r)
	if tokenString == "" || err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, "authorized token is not given", http.StatusUnauthorized)
		return
	}

	claims, err := validation.VerifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, " Could not get claims ", http.StatusUnauthorized)

		return
	}

	record, err := database.GetProcessesById(claims, id)
	if err != nil {
		http.Error(w, "could not get proceseses from database", http.StatusUnauthorized)
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(record)

}
