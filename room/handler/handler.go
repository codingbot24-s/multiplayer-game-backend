package roomHandler

import (
	"strings"

	"github.com/codingbot24-s/helper"
	roomHelper "github.com/codingbot24-s/room/helper"
	"github.com/gofiber/fiber/v2"
)

type CreateRoomReq struct {
	RoomName string `json:"name"`
}

func CreateRoom(c *fiber.Ctx) error {
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
	_, err := helper.VerifyTheToken(t)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
			"detail":  "cannot verify the token",
		})
	}
	var req CreateRoomReq
	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  err.Error(),
			"detail": "error parsing request body",
		})
	}
	r := roomHelper.NewRoom(req.RoomName)
	rm := roomHelper.GetRoomManager()
	rid := rm.AddInto(r)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"room":    rid,
	})
}

type JoinRoomReq struct {
	RoomId string `json:"roomId"`
}

func JoinRoom(c *fiber.Ctx) error {
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
	username := claims.Username
	// user need to send the roomId
	rm := roomHelper.GetRoomManager()

	var req JoinRoomReq
	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  err.Error(),
			"detail": "error parsing request body",
		})
	}

	r, err := rm.RoomExist(req.RoomId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  err.Error(),
			"detail": "room with this id doesn't exist" + req.RoomId,
		})
	}

	r.AddInto(username)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "joined room with id " + req.RoomId,
	})
}
