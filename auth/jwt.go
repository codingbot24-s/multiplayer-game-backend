package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(username string) (string, error) {
	claims := MyClaims{
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
func VerifyTheToken(t string) (*MyClaims, error) {
    token, err := jwt.ParseWithClaims(t, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return []byte("1234"), nil
    })
    if err != nil {
        return nil, fmt.Errorf("error parsing token: %v", err)
    }
    claims, ok := token.Claims.(*MyClaims)
    if !ok || !token.Valid {
        return nil, fmt.Errorf("error token is not valid or cant extract claims")
    }
    
    fmt.Printf("username is %s", claims.Username)
    return claims, nil
}