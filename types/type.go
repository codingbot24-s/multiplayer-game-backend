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

// Player struct
type Player struct {
	Username string `json:"username"`
}

//

func NewPlayer(username string) *Player {
	return &Player{Username: username}
}
