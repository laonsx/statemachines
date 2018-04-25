package miner

import (
	"fmt"

	"github.com/laonsx/statemachines/westworld"
)

var MinerQuenchThistState *minerQuenchThirstState

type minerQuenchThirstState struct {
}

func (state *minerQuenchThirstState) Enter(role westworld.BaseEntity) {

	fmt.Println("Enter MinerQuenchThirstState")
}

func (state *minerQuenchThirstState) Execute(role westworld.BaseEntity) {

	fmt.Println("Execute MinerQuenchThirstState")

	miner, _ := role.(*Miner)

	miner.Thirst = 0

	miner.Machine.RevertToPreviousState()
}

func (state *minerQuenchThirstState) Exit(role westworld.BaseEntity) {

	fmt.Println("Exit MinerQuenchThirstState")
}

func (state *minerQuenchThirstState) OnMessage(role westworld.BaseEntity) bool {

	return false
}
