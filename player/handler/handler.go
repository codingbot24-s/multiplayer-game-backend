package handler

import (
	"strings"
	"time"

	"github.com/codingbot24-s/helper"
	atypes "github.com/codingbot24-s/types"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var validate *validator.Validate

var zoneMap = make(map[string]string)

// Login and generate the jwt
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
	claims := &atypes.MyClaims{
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

	zoneMap[req.Username] = "ws://localhost:4000"

	return c.Status(200).JSON(fiber.Map{
		"message": "success",
		"token":   ss,
	})

}

func LoginCheck(c *fiber.Ctx) error {
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
	claims, err := helper.VerifyTheToken(t)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
			"detail":  "cannot verify the token",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "success",
		"claims":  claims.Username,
	})

}

func SessionHandler(c *fiber.Ctx) error {
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
	claims, err := helper.VerifyTheToken(t)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
			"detail":  "cannot verify the token",
		})
	}

	zone, ok := zoneMap[claims.Username]
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Username not found in zone",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "success",
		"claims":  claims.Username,
		"zone":    zone,
	})
}
