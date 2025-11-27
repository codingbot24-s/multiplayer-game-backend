package zonehelper

import (
	"strings"

	"github.com/codingbot24-s/helper"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func StartWebSocket() {
	app := fiber.New()
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
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
			c.Locals("claims", claims)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		mes := []byte("hello websocket")
		if err := c.WriteMessage(websocket.TextMessage, mes); err != nil {
			c.Close()
			return
		}
	}))

	// start the websocket on 4000
	app.Listen(":4000")
}
