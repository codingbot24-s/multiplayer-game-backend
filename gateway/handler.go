package gateway

import (
	"strings"

	"github.com/codingbot24-s/auth"
	"github.com/gofiber/fiber/v2"
)

type LoginReq struct {
	Username string `json:"username"`
}

func Login(ctx *fiber.Ctx) error {
	var req LoginReq
	if err := ctx.BodyParser(&req); err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": "false",
			"detail":  "Invalid request body",
		})
	}
	token, err := auth.GenerateToken(req.Username)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": "false",
			"detail":  "Internal server error",
		})

	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": "true",
		"token":   token,
	})
}
// TODO: define a midlleware for all the protected routes
func LoginCheck(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")

	if authHeader == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{

			"message": "Missing Authorization header",
		})
	}

	parts := strings.Split(authHeader, " ")

	if len(parts) != 2 || parts[0] != "Bearer" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{

			"message": "Invalid Authorization header format",
		})
	}

	t := parts[1]

	claims, err := auth.VerifyTheToken(t)

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
			"detail": "cannot verify the token",
		})

	}

	return ctx.Status(200).JSON(fiber.Map{	
		"message": "success",
		"claims": claims.Username,

	})
}
