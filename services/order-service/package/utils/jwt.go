package utils

import (
	"errors"
	"order-service/configs"

	"github.com/dgrijalva/jwt-go"
)

func ValidateToken(tokenString string) (uint, error) {
	cfg, _ := configs.LoadConfig()
	claims := &jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWT.Secret), nil
	})

	if err != nil || !token.Valid {
		return 0, errors.New("invalid token")
	}

	return uint((*claims)["id"].(float64)), nil
}
