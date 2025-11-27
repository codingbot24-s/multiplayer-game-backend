package atypes

import "github.com/golang-jwt/jwt/v5"

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// Claims type
type MyClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
