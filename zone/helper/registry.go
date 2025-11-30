package zoneHelper

import (
	"fmt"
)

type Registry struct {
	Pmap map[string]*Player
}

var registry Registry

// call this only one time we dont want to create a new registry everytime
// Only create the registry
func NewRegistry() {
	registry = Registry{
		Pmap: make(map[string]*Player),
	}
}

func GetZoneRegistry() *Registry {
	return &registry
}

func (r *Registry) AddPlayer(name string, p *Player) {
	r.Pmap[name] = p
}

func (r *Registry) RemovePlayer(name string) {
	delete(r.Pmap, name)
}

// Get the player by name and update its movement
func (r *Registry) UpdatePlayerMovement (name string, x int, y int ) {
	player := r.Pmap[name]
	player.X = x
	player.Y = y
	fmt.Printf("player with name %s moved to (%d %d)",player.Name, player.X,player.Y)
} 

// TODO: remove this
func (r *Registry) Check() {
	for _, p := range r.Pmap {
		fmt.Println("player exists with name ", p.Name)
		fmt.Printf("player is at (%d %d)\n", p.X, p.Y)

	}
}


func(r *Registry) BuildSnapShot() *[]Player {
	snapShot := make([]Player, 0)
	for _, p := range r.Pmap {
		snapShot = append(snapShot, *p)
	}
	return &snapShot
}