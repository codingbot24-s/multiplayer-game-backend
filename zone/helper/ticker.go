package zoneHelper

import (
	"time"

)

func StartTicker() {
	ticker := time.NewTicker(time.Millisecond * 100)
	zoneRegistery := GetZoneRegistry()	
	for {
		select {
		case <-ticker.C :
			snapshot := zoneRegistery.BuildSnapShot()
			players := zoneRegistery.GetAllPlayers()
			for _, p := range players {
				p.Conn.WriteJSON(snapshot)
			}
		}
	}
	
}
