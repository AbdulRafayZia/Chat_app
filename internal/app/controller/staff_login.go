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

func StaffLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request utils.Credentials
	json.NewDecoder(r.Body).Decode(&request)

	role,id, err := database.GetRole(request.Username)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, "unauthozied username", http.StatusUnauthorized)

		return
	}

	validStaf, err := validation.CheckStaffValidity(w, r, request)
	if !validStaf {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Println(err)
		return

	}

	accessToken, refreshToken, err := jwt.CreateToken(request.Username, role , id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "error in generating toke ", http.StatusInternalServerError)
		return
	}

	token := utils.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)

}
