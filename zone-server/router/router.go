package router

import (
	"log"
	"strings"

	"github.com/codingbot24-s/helper"
	"github.com/codingbot24-s/zone-server/handler"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func StartWebSocket() {
	app := fiber.New()
	// websocket middleware request will get upgrade on this route
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

	// handlers
	app.Get("/ping", websocket.New(handler.Pong))

	// TODO: TEST THIS
	app.Get("/ws/:roomid/:userid",websocket.New(handler.Pang))
	// start the websocket fon 4000
	if err := app.Listen(":4000"); err != nil {
		log.Fatalln(err)
	}
}
