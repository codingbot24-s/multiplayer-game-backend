package zoneHelper

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func HandleConnection(c *websocket.Conn, ch chan string, name string) {
	r := GetZoneRegistry()
	var msgReq Message
	defer func() {
		c.Close()
		ch <- "done"
	}()
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		fmt.Printf("recv: %s\n", message)

		if err := json.Unmarshal(message, &msgReq); err != nil {
			c.WriteMessage(websocket.TextMessage, []byte("error unmarshalling message: "+err.Error()))
			continue
		}

		switch msgReq.Type {
		case "move":
			fmt.Println("move request")
			var moveReq MoveReq
			if err := json.Unmarshal(msgReq.Data, &moveReq); err != nil {
				log.Printf("error unmarshalling move request: %v", err)
				continue 
			}
			r.UpdatePlayerMovement(name, moveReq.X, moveReq.Y)
			r.Check()
		default:
			fmt.Println("message type is unknown")
		}
	}
}
