package handler

import (
	"fmt"
	"log"

	zonehelper "github.com/codingbot24-s/zone-server/helper"
	"github.com/gofiber/contrib/websocket"
)

// Pong is player connection check
func Pong(c *websocket.Conn) {
	name := c.Query("username")

	zonehelper.AddPlayer(name)

	if err := c.WriteMessage(websocket.TextMessage, []byte("user with name "+name+" connected")); err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		zonehelper.RemovePlayer(name)
		c.Close()
	}()

	for {
		_, _, err := c.ReadMessage()
		if err != nil {
			log.Printf("Player %s disconnected: %v", name, err)
			return
		}
	}
}
