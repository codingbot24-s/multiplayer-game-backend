package roomHelper

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type RoomManager struct {
	RoomMap map[string]*Room
}

type Room struct {
	RoomName string
	// Players name only
	Players []string
}

var (
	instance *RoomManager
	once     sync.Once
)

// GetRoomManager returns the singleton instance of RoomManager
func GetRoomManager() *RoomManager {
	once.Do(func() {
		instance = &RoomManager{
			RoomMap: make(map[string]*Room),
		}
	})
	return instance
}

// AddInto will add the room and return the room id
func (rm *RoomManager) AddInto(r *Room) string {
	// generate the id and insert it into the room map
	id := uuid.New().String()
	rm.RoomMap[id] = r
	return id
}

func (rm *RoomManager) RoomExist(roomId string) (*Room, error) {
	// if room exist return the room
	r, ok := rm.RoomMap[roomId]
	if !ok {
		return nil, fmt.Errorf("room %s not found", roomId)
	}

	return r, nil
}

// NewRoom will create a room
func NewRoom(name string) *Room {
	r := Room{
		RoomName: name,
	}

	return &r
}

// AddInto will add player into room
func (r *Room) AddInto(player string) {
	r.Players = append(r.Players, player)
}
