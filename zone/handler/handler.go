package zoneHandler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} 
// connect to websocket server
func Connect(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w,r,nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	fmt.Println("player connected ")
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
	}
}
