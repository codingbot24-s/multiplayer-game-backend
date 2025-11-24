package handler

import (
	"mmo-backend/helper"
	atypes "mmo-backend/types"
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
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Unix(1516239022, 0)),
		Issuer:    req.Username,
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
