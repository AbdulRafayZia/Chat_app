package utils

import "github.com/dgrijalva/jwt-go"

type MyClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Id       uint 	`json:"id"`
	jwt.StandardClaims
}
