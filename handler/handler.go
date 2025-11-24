package handler

import (
	"fmt"
	"mmo-backend/helper"
	atypes "mmo-backend/types"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var validate *validator.Validate

func LoginHandler(c *fiber.Ctx) error {
	var req atypes.LoginRequest
	// parse the request body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"detail":  "error parsing request body",
		})
	}

	validate = validator.New()
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"detail":  "error validating request body",
		})
	}
	claims := &MyClaims{
		Username: req.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := helper.GetConfig()
	ss, err := token.SignedString([]byte(secret.JwtSecret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"detail":  "error signing token",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "success",
		"token":   ss,
	})

}

type MyClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func LoginCheck(c *fiber.Ctx) error {
	secret := helper.GetConfig()
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Missing Authorization header",
		})
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid Authorization header format",
		})
	}

	t := parts[1]
	token, err := jwt.ParseWithClaims(t, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate signing algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret.JwtSecret), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "token parsing error",
			"error":   err.Error(),
		})
	}

	claims, ok := token.Claims.(*MyClaims)
	if !ok || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid token",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "success",
		"claims":  claims.Username,
	})

}
