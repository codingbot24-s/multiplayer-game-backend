package router

import (
	"fmt"
	"log"
	"mmo-backend/handler"

	"github.com/gofiber/fiber/v2"
)

func StartRouter() {
	app := fiber.New()
	app.Post("/login", handler.LoginHandler)
	app.Get("/me", handler.LoginCheck)
	fmt.Println("Listening on port 3000")
	log.Fatal(app.Listen(":3000"))
}
