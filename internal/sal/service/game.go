package service

import (
	"fmt"
	"sal/internal/sal/models"
	"sort"
	"sync/atomic"
)

type GameStatus int

const (
	Initialised GameStatus = iota
	InProgress
	Complete
)

type GameService struct {
	Status                 GameStatus
	Players                map[int32]*models.Player
	Board                  *BoardService
	PlayerId, LastPlayerId int32
	Dice                   *DiceService
}

func NewGameService(dimension int, attributePositionMap []*models.Attribute) (*GameService, error) {
	boardService, err := NewBoardService(dimension, attributePositionMap)
	if err != nil {
		return nil, err
	}

	return &GameService{
		Status:       Initialised,
		Players:      make(map[int32]*models.Player),
		Board:        boardService,
		PlayerId:     0,
		LastPlayerId: 0,
		Dice:         NewDiceService(1, 6),
	}, nil
}

func (gs *GameService) AddPlayer(name string) error {
	playerId := atomic.AddInt32(&gs.PlayerId, 1)

	player, err := models.NewPlayer(name, playerId)
	if err != nil {
		return err
	}

	gs.Players[player.Id] = player
	return nil
}

func (gs *GameService) GetNextPlayer() *models.Player {
	var playerList []*models.Player
	for _, player := range gs.Players {
		playerList = append(playerList, player)
	}

	sort.Slice(playerList, func(i, j int) bool {
		return playerList[i].Id < playerList[j].Id
	})

	for _, player := range playerList {
		if player.Id > gs.LastPlayerId {
			return player
		}
	}
	return playerList[0]
}

func (gs *GameService) Roll(player *models.Player) {
	if gs.Status == Complete {
		fmt.Println("Game is already complete")
		return
	}

	if gs.Status == Initialised {
		gs.Status = InProgress
	}

	steps := gs.Dice.Roll()
	fmt.Println(fmt.Sprintf("Player %s rolled a %d", player.GetName(), steps))
	playerReachedEnd, err := gs.Board.MovePlayer(player, steps)
	if err != nil {
		fmt.Errorf("Error moving player %s: %v", player.GetName(), err)
		return
	}
	gs.LastPlayerId = player.Id
	if playerReachedEnd {
		fmt.Println(fmt.Sprintf("Player %s has reached the end", player.GetName()))
		gs.Status = Complete
	}
}

func (gs *GameService) PrintBoard() {
	for _, player := range gs.Players {
		position := player.GetPosition()
		fmt.Println(fmt.Sprintf("Player %s is at position %d", player.GetName(), position))
	}

	for position, attribute := range gs.Board.GetAttributesLocation() {
		fmt.Println(fmt.Sprintf("Attribute %s is at position %d", attribute.Name, position))
	}

	fmt.Println(fmt.Sprintf("Game Status: %d", gs.Status))
	fmt.Println(fmt.Sprintf("Last Player: %d", gs.LastPlayerId))
}
