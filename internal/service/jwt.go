package service

import (
	"github.com/golang-jwt/jwt/v5"
	"rest_api/domain"
	"time"
)

type JWTAuthService struct {
}

func NewJWTAuthService() *JWTAuthService {
	return &JWTAuthService{}
}

func (jas *JWTAuthService) CreateToken(id uint64) (string, error) {
	claims := &domain.JWTCustomClaims{
		id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(domain.JWTKey))
	if err != nil {
		return "", err
	}
	return t, nil
}
