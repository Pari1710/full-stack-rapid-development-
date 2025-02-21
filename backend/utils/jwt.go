package utils

import (
	"errors"
	"time"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("your-secret-key")

func GenerateJWT(email string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, _ := token.SignedString(jwtKey)
	return tokenString
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, errors.New("invalid token")
	}

	return token, nil
}
