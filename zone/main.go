package main

import (
	"fmt"
	"net/http"

	zoneHandler "github.com/codingbot24-s/zone/handler"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	// connect to websocket server
	r.HandleFunc("/ws", zoneHandler.Connect)
	fmt.Println("starting http server on :4000")
	
	http.ListenAndServe(":4000", r)
}
