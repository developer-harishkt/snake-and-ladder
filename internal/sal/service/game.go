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
	Dice                   DiceRoller
	Observers              []Observer
}

func NewGameService(dimension int, attributePositionMap []*models.Attribute) (*GameService, error) {
	boardService, err := NewBoardService(dimension, attributePositionMap)
	if err != nil {
		return nil, err
	}

	gs := &GameService{
		Status:       Initialised,
		Players:      make(map[int32]*models.Player),
		Board:        boardService,
		PlayerId:     0,
		LastPlayerId: 0,
		Dice:         GetDiceRoller(1, 6, RandomDiceRollerStrategy),
		Observers:    []Observer{},
	}
	gs.RegisterObserver(NewConsoleObserver())
	return gs, nil
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
	playerReachedEnd, attribute, err := gs.Board.MovePlayer(player, steps)
	if err != nil {
		fmt.Errorf("Error moving player %s: %v", player.GetName(), err)
		return
	}
	gs.NotifyPlayerPositionUpdates(player.Id, steps, player.GetPosition(), attribute)
	gs.LastPlayerId = player.Id
	if playerReachedEnd {
		gs.Status = Complete
		gs.NotifyGameOver(player.Id)
	}
}

func (gs *GameService) PrintBoard() {
	gs.NotifyPrintBoard(gs.Players, gs.Board.GetAttributesLocation())
}

func (gs *GameService) RegisterObserver(observer Observer) {
	gs.Observers = append(gs.Observers, observer)
}

func (gs *GameService) NotifyPlayerPositionUpdates(playerId int32, roll, position int, attribute string) {
	for _, observer := range gs.Observers {
		observer.PlayerPosition(playerId, roll, position, attribute)
	}
}

func (gs *GameService) NotifyGameOver(playerId int32) {
	for _, observer := range gs.Observers {
		observer.GameOver(playerId)
	}
}

func (gs *GameService) NotifyPrintBoard(players map[int32]*models.Player, attributes map[int]*models.Attribute) {
	for _, observer := range gs.Observers {
		observer.PrintBoard(players, attributes)
	}
}
