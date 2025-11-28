package roomRouter

import (
	"fmt"

	roomHandler "github.com/codingbot24-s/room/handler"
	"github.com/gofiber/fiber/v2"
)

func StartRouter() {
	app := fiber.New()

	app.Post("/room/create", roomHandler.CreateRoom)
	app.Post("room/join", roomHandler.JoinRoom)

	fmt.Println("room server running on 5000")
	app.Listen(":5000")
}
