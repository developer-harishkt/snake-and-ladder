package models

import "fmt"

type Square struct {
	Position  int
	Attribute *Attribute
}

func NewSquare(position int, attribute *Attribute) (*Square, error) {
	// check if position is valid
	if position < 0 {
		return nil, fmt.Errorf("position must be non-negative")
	}

	return &Square{
		Position:  position,
		Attribute: attribute,
	}, nil
}
