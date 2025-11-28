package zonehelper

import (
	"fmt"

	atypes "github.com/codingbot24-s/types"
)

// global player map
var PMap = make(map[string]*atypes.Player)

func AddPlayer(name string) {
	PMap[name] = atypes.NewPlayer(name)
	fmt.Printf("Player with name %s connected", name)
}

func RemovePlayer(name string) {
	delete(PMap, name)
	fmt.Printf("Player with name %s disconnected", name)
}
