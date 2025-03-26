package main

import (
	"fmt"
	"sal/internal/sal/models"
	"sal/internal/sal/service"
)

func main() {
	dim := 30
	attributes := []*models.Attribute{
		{Name: "Snake", StartPos: 28, EndPos: 3, Direction: models.Backward},
		{Name: "Ladder", StartPos: 2, EndPos: 29, Direction: models.Forward},
		{Name: "Snake", StartPos: 20, EndPos: 10, Direction: models.Backward},
		{Name: "Ladder", StartPos: 5, EndPos: 15, Direction: models.Forward},
		{Name: "Snake", StartPos: 12, EndPos: 4, Direction: models.Backward},
		{Name: "Ladder", StartPos: 8, EndPos: 18, Direction: models.Forward},
	}
	newGame, err := service.NewGameService(dim, attributes)
	if err != nil {
		panic(err)
	}
	fmt.Println(newGame.Status)
	// Add players
	err = newGame.AddPlayer("Player 1")
	if err != nil {
		panic(err)
	}

	for {
		nextPlayer := newGame.GetNextPlayer()
		newGame.Roll(nextPlayer)
		if newGame.Status == service.Complete {
			break
		}
	}
}
