package auth

import (
	"fmt"

	"github.com/codingbot24-s/helper"
	"github.com/golang-jwt/jwt/v5"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(username string) (string, error) {
	return "", nil
}

func VerifyTheToken(t string) (*MyClaims, error) {
	secret := helper.GetConfig()
	token, err := jwt.ParseWithClaims(t, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate signing algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret.JwtSecret), nil
	})

	if err != nil {
		return &MyClaims{}, fmt.Errorf("error parsing token: %v", err)
	}

	claims, ok := token.Claims.(MyClaims)
	if !ok || !token.Valid {

	}

	return &claims, nil
}
