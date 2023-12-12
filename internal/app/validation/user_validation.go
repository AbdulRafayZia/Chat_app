package validation

import (
	"net/http"

	"github.com/AbdulRafayZia/Gorilla-mux/utils"
	database "github.com/AbdulRafayZia/Gorilla-mux/internal/infrastructure/Database"
)

func CheckUserValidity(w http.ResponseWriter, r *http.Request, request utils.Credentials) (bool, error) {
	role,_, err := database.GetRole(request.Username)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, "unauthozied username", http.StatusUnauthorized)
		return false, err

	}

	validRole := CheckUserRole(role)
	if !validRole {
		http.Error(w, "not a user", http.StatusUnauthorized)
		return false, err

	}

	user, err := database.FindByName(request.Username)
	if err != nil {
		http.Error(w, "error finding user", http.StatusInternalServerError)
		return false, err

	}
	dbPassword, err := database.GetPassword(request.Username)
	if err != nil {
		http.Error(w, "error getting hashed password", http.StatusBadRequest)
		return false, err

	}

	// Verify the password
	validPassword := VerifyPassword(dbPassword, request.Password)
	if !validPassword {

		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "incorrect password", http.StatusUnauthorized)
		return false, err

	}
	if user == nil && !validPassword {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, "invalid credentials ", http.StatusUnauthorized)

		return false, err

	}
	return true, nil

}
