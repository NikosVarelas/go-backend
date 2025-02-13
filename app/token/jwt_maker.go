package token

import (
	"fmt"
	"go-backed/app/configuration"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTMaker struct {
	secretKey string
	period    int
}

func NewJWTMaker(config *configuration.Config) *JWTMaker {
	return &JWTMaker{
		secretKey: config.JwtToken.SecretKey,
		period:    config.JwtToken.Period,
	}
}

func (maker *JWTMaker) CreateToken(id int, email string, isAdmin bool) (string, *UserClaims, error) {
	claims, err := NewUserClaims(id, email, isAdmin, time.Duration(maker.period)*time.Hour*24)
	if err != nil {
		return "", nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(maker.secretKey))
	if err != nil {
		return "", nil, fmt.Errorf("error signing token: %w", err)
	}

	return tokenStr, claims, nil
}

func (maker *JWTMaker) VerifyToken(tokenStr string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		// verify the signing method
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("invalid token signing method")
		}

		return []byte(maker.secretKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}

func (maker *JWTMaker) GetExpiration() time.Duration {
	return time.Duration(maker.period) * time.Hour * 24
}
