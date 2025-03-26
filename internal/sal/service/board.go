package service

import (
	"fmt"
	"sal/internal/sal/models"
)

// this is a board service which handles the logic for the board moving positions

type BoardService struct {
	Board *models.Board
}

func NewBoardService(dimension int, attributes []*models.Attribute) (*BoardService, error) {
	attributePositionMap := make(map[int]*models.Attribute)
	for _, attribute := range attributes {
		attributePositionMap[attribute.StartPos] = attribute
	}

	board, err := models.NewBoard(dimension, attributePositionMap)
	if err != nil {
		return nil, err
	}

	return &BoardService{
		Board: board,
	}, nil
}

func (bs *BoardService) GetSquare(position int) (*models.Square, error) {
	square, err := bs.Board.GetSquare(position)
	if err != nil {
		return nil, err
	}
	return square, nil
}

func (bs *BoardService) MovePlayer(player *models.Player, steps int) (bool, string, error) {
	// check if steps are valid
	if steps < 0 {
		return false, "", fmt.Errorf("steps must be non-negative")
	}

	if player.GetPosition()+steps > bs.Board.Dimension {
		fmt.Println(fmt.Sprintf("%s cannot move %d steps from position %d", player.GetName(), steps, player.GetPosition()))
		return false, "", nil
	}

	player.Move(steps)
	newPosition := player.GetPosition()

	square, err := bs.GetSquare(newPosition)
	if err != nil {
		return false, "", err
	}

	if square.Attribute != nil {
		fmt.Println(square.Attribute)
		// this subtractions takes care of moving backward & forward
		player.Move(square.Attribute.EndPos - square.Attribute.StartPos)
		return player.GetPosition() == bs.Board.Dimension, square.Attribute.Name, nil
	}

	return player.GetPosition() == bs.Board.Dimension, "", nil
}

func (bs *BoardService) GetAttributesLocation() map[int]*models.Attribute {
	return bs.Board.GetAttributesLocation()
}
