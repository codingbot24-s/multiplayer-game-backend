package handler

import (
	"fmt"

	zonehelper "github.com/codingbot24-s/zone-server/helper"
	"github.com/gofiber/contrib/websocket"
)

// Pong is player connection check
func Pong(c *websocket.Conn) {
	name := c.Query("username")
	if name == "" {
		if err := c.WriteMessage(websocket.TextMessage, []byte("pong send the username in param")); err != nil {
			fmt.Println(err)
			return
		}
	}

	zonehelper.AddPlayer(name)

	//TODO: check this handler and implement on disconnect remove player from map

}
