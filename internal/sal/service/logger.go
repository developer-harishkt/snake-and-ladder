package service

import (
	"fmt"
	"sal/internal/sal/models"
)

type Observer interface {
	PlayerPosition(player int32, roll int, position int, attribute string)
	GameOver(player int32)
	PrintBoard(players map[int32]*models.Player, attributes map[int]*models.Attribute)
}

type ConsoleObserver struct{}

func NewConsoleObserver() *ConsoleObserver {
	return &ConsoleObserver{}
}

func (co *ConsoleObserver) PlayerPosition(player int32, roll int, position int, attribute string) {
	if attribute != "" {
		fmt.Printf("Player %d rolled %d and moved to position %d with attribute %s\n", player, roll, position, attribute)
	} else {
		fmt.Printf("Player %d rolled %d and moved to position %d\n", player, roll, position)
	}
}

func (co *ConsoleObserver) GameOver(player int32) {
	fmt.Printf("Player %d has won the game!\n", player)
}

func (co *ConsoleObserver) PrintBoard(players map[int32]*models.Player, attributes map[int]*models.Attribute) {
	fmt.Println("Current Board State:")
	for _, player := range players {
		fmt.Printf("Player %d is at position %d\n", player.Id, player.GetPosition())
	}
	for _, attribute := range attributes {
		fmt.Printf("Attribute %s is at position %d\n", attribute.Name, attribute.StartPos)
	}
}
