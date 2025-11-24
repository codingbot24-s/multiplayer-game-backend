package router

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"mmo-backend/handler"
)

func StartRouter() {
	app := fiber.New()
	app.Post("/login", handler.LoginHandler)
	fmt.Println("Listening on port 3000")
	log.Fatal(app.Listen(":3000"))
}
