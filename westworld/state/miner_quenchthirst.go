package state

import (
	"fmt"

	"github.com/laonsx/statemachines/westworld"
	"github.com/laonsx/statemachines/westworld/baseentity"
)

var MinerQuenchThistState *minerQuenchThirstState

type minerQuenchThirstState struct {
}

func (state *minerQuenchThirstState) Enter(role westworld.BaseEntity) {

	fmt.Println("Enter MinerQuenchThirstState")
}

func (state *minerQuenchThirstState) Execute(role westworld.BaseEntity) {

	fmt.Println("Execute MinerQuenchThirstState")

	miner, _ := role.(*baseentity.Miner)

	miner.Thirst = 0

	miner.Machine.RevertToPreviousState()
}

func (state *minerQuenchThirstState) Exit(role westworld.BaseEntity) {

	fmt.Println("Exit MinerQuenchThirstState")
}

func (state *minerQuenchThirstState) OnMessage(role westworld.BaseEntity, msg interface{}) bool {

	return false
}
