package utils

import (
	"MillyPhoto/models"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
	"time"
)

func GenerateJwt(login string) string {
	expirationTime := time.Now().Add(24 * time.Hour) // Токен на 24 часа
	claims := &models.Claims{
		Login: login,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := os.Getenv("JWT_SECRET")

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		log.Println("ошибка в GenerateJwt")
	}
	return tokenString
}
