package state

import (
	"fmt"

	"github.com/laonsx/statemachines/westworld"
	"github.com/laonsx/statemachines/westworld/baseentity"
)

var MinerVisitBankAndDepostGoldState *minerVisitBankAndDepostGoldState

type minerVisitBankAndDepostGoldState struct {
}

func (state *minerVisitBankAndDepostGoldState) Enter(role westworld.BaseEntity) {

	fmt.Println("Enter MinerVisitBankAndDepostGoldState")
}

func (state *minerVisitBankAndDepostGoldState) Execute(role westworld.BaseEntity) {

	fmt.Println("Execute MinerVisitBankAndDepostGoldState")

	miner, _ := role.(*baseentity.Miner)

	miner.MoneyInBank += miner.GoldCarried
	miner.GoldCarried = 0

	miner.Machine.RevertToPreviousState()
}

func (state *minerVisitBankAndDepostGoldState) Exit(role westworld.BaseEntity) {

	fmt.Println("Exit MinerVisitBankAndDepostGoldState")
}

func (state *minerVisitBankAndDepostGoldState) OnMessage(role westworld.BaseEntity) bool {

	return false
}
