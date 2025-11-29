package gateway

import "github.com/gofiber/fiber/v2"

func Login(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{})
}
