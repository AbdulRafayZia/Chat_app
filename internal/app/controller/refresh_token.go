package controller

import (
	"encoding/json"
	"net/http"

	"github.com/AbdulRafayZia/Gorilla-mux/pkg/jwt"
)

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	refreshToken := r.FormValue("refreshtoken")

	newAccessToken, err := jwt.RefreshToken(refreshToken)
	if err != nil {
		http.Error(w, "cannot generate new access token", http.StatusUnauthorized)
		return
	}
	tokenMap := make(map[string]string)
	tokenMap["access_token"] = newAccessToken

	json.NewEncoder(w).Encode(tokenMap)

}
