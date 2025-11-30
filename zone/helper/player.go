package zoneHelper

import "github.com/gorilla/websocket"

type Player struct {
	Name string
	X 	int
	Y 	int
	Conn  *websocket.Conn `json:"-"`
}

func NewPlayer(name string) *Player {
	p := Player{
		Name: name,
		X: 0,
		Y: 0,
		Conn: &websocket.Conn{},
	}
	return &p
}