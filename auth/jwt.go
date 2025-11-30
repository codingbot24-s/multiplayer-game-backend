package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/codingbot24-s/helper"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(username string) (string, error) {
	claims := helper.MyClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),

			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// dummy secret not will be used in prod
	ss, err := token.SignedString([]byte("1234"))
	if err != nil {
		return "", fmt.Errorf("error signing token")
	}
	return ss, nil
}
func VerifyTheToken(t string) (*helper.MyClaims, error) {
	token, err := jwt.ParseWithClaims(t, &helper.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("1234"), nil
	})
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %v", err)
	}
	claims, ok := token.Claims.(*helper.MyClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("error token is not valid or cant extract claims")
	}

	fmt.Printf("username is %s", claims.Username)
	return claims, nil
}

func AuthMiddleware(w http.ResponseWriter, r *http.Request) *helper.MyClaims {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Authorization header missing", http.StatusUnauthorized)
		return nil
	}
	parts := strings.Split(authHeader, " ")

	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
		return nil
	}

	token := parts[1]
	claims, err := VerifyTheToken(token)

	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return nil
	}

	return claims
}
