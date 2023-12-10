package jwt

import (
	"fmt"
	"net/http"
)

func GetToken(w http.ResponseWriter, r *http.Request) (string, error) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)

		return "", fmt.Errorf("missing authorization header")
	}
	tokenString = tokenString[len("Bearer"):]
	return tokenString, nil

}
