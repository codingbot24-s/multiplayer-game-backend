package handler

import (
	"fmt"
	"log"

	roomHelper "github.com/codingbot24-s/room/helper"
	zonehelper "github.com/codingbot24-s/zone-server/helper"
	"github.com/gofiber/contrib/websocket"
)

// Pong is player connection check
func Pong(c *websocket.Conn) {
	name := c.Query("username")
	// add player into the map
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
func Pang (c *websocket.Conn) {
	roomId := c.Params("roomId")
	username := c.Params("username")

	if roomId == "" || username == "" {
		c.Close()
		fmt.Println("missing username or roomId")
		return
	}
	
	rm :=  roomHelper.GetRoomManager()
	room,err := rm.RoomExist(roomId)
	if err != nil {
		c.Close()
		fmt.Println("room dosnt exist")
		return
	}
	user,err := room.FindTheUser(username)
	if err != nil {
		c.Close()
		fmt.Println("user dosnt exist in room")
		return
	}


}