package router

import (
	"fmt"
	"log"

	"github.com/codingbot24-s/player/handler"
	"github.com/gofiber/fiber/v2"
)

func StartRouter() {
	app := fiber.New()
	// GET
	app.Get("/me", handler.LoginCheck)
	app.Get("/session", handler.SessionHandler)

	app.Post("/login", handler.LoginHandler)
	fmt.Println("Listening on port 3000")
	log.Fatal(app.Listen(":3000"))
}
