package domain

import "github.com/golang-jwt/jwt/v5"

type JWTCustomClaims struct {
	UserID uint64 `json:"user_id"`
	jwt.RegisteredClaims
}

const JWTKey = "secret"
