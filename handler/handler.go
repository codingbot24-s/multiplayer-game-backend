package handler

import "github.com/gofiber/fiber/v2"

func LoginHandler(c *fiber.Ctx) error {
	var req type.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
			"detail" : "Cannot parse request body",
		})
	}
}
