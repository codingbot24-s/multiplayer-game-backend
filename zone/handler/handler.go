package zoneHandler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codingbot24-s/helper"
	zoneHelper "github.com/codingbot24-s/zone/helper"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

// connect to websocket server
func Connect(w http.ResponseWriter, r *http.Request) {
	//TODO: how can we put the middleware here?

	// temporary solution get from query from query param or we can get the token also 
	name := r.URL.Query().Get("username")
	fmt.Println("name is ", name)
	if name == "" {
		fmt.Println("no name")
		return
	}

	
	// get the registry
	registry := zoneHelper.GetZoneRegistry()
	// add the player in the registry
	registry.AddPlayer(name, zoneHelper.NewPlayer(name))
	fmt.Println("player connected with name ", name)
	// check the registry 
	registry.Check()


	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	//TODO: extract player name from claims
	// this is a goroutine channel
	ch := make(chan string)
	go func() {
		helper.HandleConnection(c, ch)
	}()

	message := <-ch
	if message == "done" {
		// remove the player from the registry
		fmt.Println("player disconnected")
		registry.RemovePlayer(name)
		// check the registry
		fmt.Println("removing player from registry with name ", name)
		registry.Check()
		return
	}
}
