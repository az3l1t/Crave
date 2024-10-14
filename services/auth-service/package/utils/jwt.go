package utils

import (
	"auth-service/configs"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(userID uint) (string, error) {
	cfg, _ := configs.LoadConfig()
	claims := jwt.MapClaims{
		"id":  userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWT.Secret))
}
