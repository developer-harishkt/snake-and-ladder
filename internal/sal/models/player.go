package models

import (
	"fmt"
)

// Player represents a player in the game
type Player struct {
	Id       int32
	Name     string
	Position int
}

func NewPlayer(name string, id int32) (*Player, error) {
	// check if name is empty
	if name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}

	return &Player{
		Id:   id,
		Name: name,
	}, nil
}

func (p *Player) Move(steps int) {
	// check if steps are valid
	if steps < 0 {
		fmt.Println("steps must be non-negative")
		return
	}

	p.Position += steps
}

func (p *Player) GetPosition() int {
	return p.Position
}

func (p *Player) GetName() string {
	return p.Name
}
