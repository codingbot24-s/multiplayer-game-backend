package zoneHelper

import "github.com/gorilla/websocket"

type Player struct {
	Name string
	Conn  *websocket.Conn
}

func NewPlayer(name string) *Player {
	p := Player{
		Name: name,
		Conn: &websocket.Conn{},
	}
	return &p
}