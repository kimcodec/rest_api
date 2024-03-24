package domain

import "github.com/golang-jwt/jwt/v5"

type JWTCustomClaims struct {
	ID uint64 `json:"name"`
	jwt.RegisteredClaims
}

const JWTKey = "secret"
