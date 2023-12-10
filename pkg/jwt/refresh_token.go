package jwt

import (
	"fmt"
	"time"

	"github.com/AbdulRafayZia/Gorilla-mux/internal/app/validation"
)

func RefreshToken(refreshTokenString string) (string, error) {
	// Validate the refresh token
	claims, err := validation.VerifyToken(refreshTokenString)
	if err != nil {
		return "", err
	}

	// Check if the refresh token has expired
	if time.Now().Unix() > claims.ExpiresAt {
		return "", fmt.Errorf("refresh token has expired")
	}

	// Create a new access token
	newAccessToken, _, err := CreateToken(claims.Username , claims.Role)
	if err != nil {
		return "", err
	}

	return newAccessToken, nil
}