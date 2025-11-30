package zoneHelper

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)



type WorldMessage struct {
    Type string          `json:"type"`
    Data json.RawMessage 
}

func SendMessage(conn *websocket.Conn, data []byte) {
	var w_Message WorldMessage
	if err := json.Unmarshal(data, &w_Message); err != nil {
		fmt.Println("error unmarshalling world message")
		return
	}
	var greeting Greeting
	switch w_Message.Type {
	case "greeting":
		if err := json.Unmarshal(w_Message.Data, &greeting); err != nil {
			fmt.Println("error unmarshalling greeting")
			return
		}
		fmt.Println("greeting ", greeting.Message)
		break
	default:
		fmt.Println("message type is unknown")
	}
	err := conn.WriteJSON(w_Message)
	if err != nil {
		fmt.Println("error sending world message")
		return
	}
}
