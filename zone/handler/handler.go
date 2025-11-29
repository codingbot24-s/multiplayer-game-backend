package zoneHandler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codingbot24-s/helper"
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
	fmt.Println("player connected ")
	ch := make(chan string)
	go func () {
		helper.HandleConnection(c,ch)
	} ()	
	// is this blocking yes it is 
	message := <-ch
	if message == "done" {
		fmt.Println("player disconnected")
		return
	}
}
