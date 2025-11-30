package zoneHelper

import "github.com/gorilla/websocket"

type Player struct {
	Name string
	X 	int
	Y 	int
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