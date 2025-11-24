package helper

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func StartRouter() {
	app := fiber.New()
	app.Post("/login")
	fmt.Println("Listening on port 3000")
	log.Fatal(app.Listen(":3000"))
}
