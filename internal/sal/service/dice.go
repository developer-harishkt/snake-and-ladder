package service

import (
	"math/rand"
)

type DiceService struct {
	MinValue, MaxValue int
}

func NewDiceService(minV, maxV int) *DiceService {
	return &DiceService{MinValue: minV, MaxValue: maxV}
}

func (ds *DiceService) Roll() int {
	return rand.Intn(ds.MaxValue-ds.MinValue+1) + ds.MinValue
}
