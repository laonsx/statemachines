package miner

import (
	"fmt"

	"github.com/laonsx/statemachines/westworld"
)

var MinerEnterMineAndDigForNuggetState *minerEnterMineAndDigForNuggetState

type minerEnterMineAndDigForNuggetState struct {
}

func (state *minerEnterMineAndDigForNuggetState) Enter(role westworld.BaseEntity) {

	fmt.Println("Enter MinerEnterMineAndDigForNuggetState")
}

func (state *minerEnterMineAndDigForNuggetState) Execute(role westworld.BaseEntity) {

	fmt.Println("Execute MinerEnterMineAndDigForNuggetState")

	miner, _ := role.(*Miner)

	miner.Thirst++
	miner.Fatigue++
	miner.GoldCarried++

	if miner.Thirst > 5 {

		miner.Machine.ChangeState(MinerQuenchThistState)

		return
	}

	if miner.Fatigue > 30 {

		miner.Machine.ChangeState(MinerGoHomeAndSleepTilRestedState)

		return
	}

	if miner.GoldCarried > 3 {

		miner.Machine.ChangeState(MinerVisitBankAndDepostGoldState)

		return
	}
}

func (state *minerEnterMineAndDigForNuggetState) Exit(role westworld.BaseEntity) {

	fmt.Println("Exit MinerEnterMineAndDigForNuggetState")
}

func (state *minerEnterMineAndDigForNuggetState) OnMessage(role westworld.BaseEntity) bool {

	return false
}
