package zoneHelper

import "fmt"

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
// TODO: remove this
func (r *Registry) Check() {
	for _, p := range r.Pmap {
		fmt.Println("player exists with name ", p.Name)
	}
}
