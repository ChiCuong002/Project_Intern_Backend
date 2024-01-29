package helper

import (
	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	UserId    uint   `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	IsAdmin   uint   `json:"role_id"`
	jwt.RegisteredClaims
}

