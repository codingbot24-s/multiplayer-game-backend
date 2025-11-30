package zoneHelper

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

type WorldMessage struct {
	Type string `json:"type"`
	Data json.RawMessage
}

func SendMessage(conn *websocket.Conn, data []byte) {
	var w_Message WorldMessage
	if err := json.Unmarshal(data, &w_Message); err != nil {
		fmt.Println("error unmarshalling world message")
		return
	}

	switch w_Message.Type {
	case "greeting":
		var greeting Greeting
		if err := json.Unmarshal(w_Message.Data, &greeting); err != nil {
			fmt.Println("error unmarshalling greeting")
			return
		}
		err := conn.WriteJSON(greeting.Message)
		if err != nil {
			fmt.Println("error sending world message")
			return
		}
	case "world_data":
		var snapShot SnapShot
		if err := json.Unmarshal(w_Message.Data, &snapShot); err != nil {
			fmt.Println("error unmarshalling snapshot")
			return
		}
		err := conn.WriteJSON(snapShot)
		if err != nil {
			fmt.Println("error sending world message")
			return
		}
	default:

	}

}
