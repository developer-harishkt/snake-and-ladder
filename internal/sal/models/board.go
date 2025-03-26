package models

import (
	"fmt"
)

// Board represents the game board
type Board struct {
	Dimension int
	Squares   []*Square
}

func NewBoard(dimension int, attributePositionMap map[int]*Attribute) (*Board, error) {
	// check if dimension is valid
	if dimension <= 0 {
		return nil, fmt.Errorf("dimension must be greater than 0")
	}

	squares := make([]*Square, dimension)
	for i := 0; i < dimension; i++ {
		square, err := NewSquare(i+1, attributePositionMap[i+1])
		if err != nil {
			return nil, err
		}
		squares[i] = square
	}

	return &Board{
		Dimension: dimension,
		Squares:   squares,
	}, nil
}

func (b *Board) GetSquare(position int) (*Square, error) {
	// check if position is valid
	if position < 1 || position > b.Dimension*b.Dimension {
		return nil, fmt.Errorf("position must be between 1 and %d", b.Dimension*b.Dimension)
	}

	return b.Squares[position-1], nil
}

func (b *Board) GetAttributesLocation() map[int]*Attribute {
	attributeLocation := make(map[int]*Attribute)
	for _, square := range b.Squares {
		if square.Attribute != nil {
			attributeLocation[square.Attribute.StartPos] = square.Attribute
		}
	}
	return attributeLocation
}
