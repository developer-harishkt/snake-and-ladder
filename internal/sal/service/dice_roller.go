package service

import "math/rand"

type DiceRollerStrategy int

const (
	RandomDiceRollerStrategy = iota
)

type DiceRoller interface {
	Roll() int
}

type RandomDiceRoller struct {
	MinV, MaxV int
}

func (r *RandomDiceRoller) Roll() int {
	return rand.Intn(r.MaxV-r.MinV+1) + 1
}

func NewRandomDiceRoller(minV, maxV int) *RandomDiceRoller {
	return &RandomDiceRoller{
		MinV: minV,
		MaxV: maxV,
	}
}

func GetDiceRoller(minV, maxV int, dRStrategy DiceRollerStrategy) DiceRoller {
	switch dRStrategy {
	default:
		return &RandomDiceRoller{
			MinV: 1,
			MaxV: 6,
		}
	}
}
