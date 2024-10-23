package services

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	JwtSecret []byte
}

func NewAuthService(jwtSecret string) *AuthService {
	return &AuthService{
		JwtSecret: []byte(jwtSecret),
	}
}

func (s *AuthService) VerifyToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return s.JwtSecret, nil
	})
}
