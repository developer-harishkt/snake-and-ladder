package models

import "fmt"

type Attribute struct {
	Name      string
	StartPos  int
	EndPos    int
	Direction Direction
}

type Direction int

const (
	Forward Direction = iota
	Backward
)

func NewAttribute(name string, startPos int, endPos int, direction Direction) (*Attribute, error) {
	// check if name is empty
	if name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}

	// check if start and end positions are valid
	if startPos < 0 || endPos < 0 {
		return nil, fmt.Errorf("start and end positions must be non-negative for forward direction")
	}

	// validate start and end positions based on the direction
	switch direction {
	case Forward:
		if startPos >= endPos {
			return nil, fmt.Errorf("start position must be less than end position for forward direction")
		}
	case Backward:
		if startPos <= endPos {
			return nil, fmt.Errorf("start position must be greater than end position for backward direction")
		}
	default:
		return nil, fmt.Errorf("invalid direction")
	}

	return &Attribute{
		Name:      name,
		StartPos:  startPos,
		EndPos:    endPos,
		Direction: direction,
	}, nil
}
