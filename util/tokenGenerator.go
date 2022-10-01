package util

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(payload map[string]string) (string, error) {
	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Hour)
	claims["authorized"] = true
	for key, value := range payload {
		claims[key] = value
	}
	tokenString, err := token.SignedString("rohan")
	if err != nil {
		log.Fatalf("Error occurred while signing token %s", err)
		return "", err
	}
	return tokenString, nil
}
