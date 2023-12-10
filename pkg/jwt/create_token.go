package jwt

import (
	"time"

	
	"github.com/AbdulRafayZia/Gorilla-mux/utils"
	"github.com/dgrijalva/jwt-go"
)

func CreateToken(username string, role string) (string,string, error) {
	accessTokenExpiration := time.Now().Add(1 * time.Minute).Unix()
	refreshTokenExpiration := time.Now().Add(7 * 24 * time.Hour).Unix()

	accesstoken := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"role":     role,
			"exp":      accessTokenExpiration,
		})
	accesstokenString, err := accesstoken.SignedString(utils.SecretKey)
	if err != nil {
		return "","", err
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"role":     role,
			"exp":      refreshTokenExpiration,
		})
	refreshtokenString, err := refreshToken.SignedString(utils.SecretKey)
	if err != nil {
		return "","", err
	}

	return accesstokenString,refreshtokenString, nil
}
