package gateway

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func StartRouter() {
	app := fiber.New()
	app.Post("/login", Login)
	app.Get("/profile",LoginCheck)
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("erorr cant start server: %v", err)
	}
}
