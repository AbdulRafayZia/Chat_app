package validation

import (
	"net/http"

	database "github.com/AbdulRafayZia/Gorilla-mux/internal/infrastructure/Database"
	"github.com/AbdulRafayZia/Gorilla-mux/utils"
	// "github.com/AbdulRafayZia/Gorilla-mux/utils"
)

func CheckStaffValidity(w http.ResponseWriter, r *http.Request, request utils.Credentials) (bool, error) {

	role,_, err := database.GetRole(request.Username)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, "unauthozied username", http.StatusUnauthorized)

		return false, err
	}
	validRole := CheckStaffRole(role)
	if !validRole {
		http.Error(w, "not a staff Member", http.StatusUnauthorized)
		return false, err

	}
	user, err := database.FindByName(request.Username)
	if err != nil {
		http.Error(w, "error finding user", http.StatusInternalServerError)
		return false, err

	}
	dbPassword, err := database.GetPassword(request.Username)
	if err != nil {
		http.Error(w, "error in getting  from db", http.StatusBadRequest)
		return false, err

	}
	validPassword := VerifyPassword(dbPassword, request.Password)
	if !validPassword {

		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "incorrect password ", http.StatusUnauthorized)
		return false, err

	}
	if user == nil && !validPassword && !validRole {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, "invalid credentails ", http.StatusUnauthorized)
	}

	return true, nil
}
