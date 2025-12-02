package zoneHelper

import "github.com/gorilla/websocket"
// 1 unit means one meter
// 2 server tick interval is 100ms 

type Player struct {
	Name string
	// position x , y 
	X 	int
	Y 	int
	// last time when clients input was processed
	LastUpdate int64
	// last time when update was sent to the client 
	LastSent  int64``
	isMoving 	bool
	Conn  *websocket.Conn `json:"-"`
}

func NewPlayer(name string,conn *websocket.Conn) *Player {
	p := Player{
		Name: name,
		X: 0,
		Y: 0,
		Conn: conn,
	}
	return &p
}