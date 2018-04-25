package miner

import (
	"fmt"

	"github.com/laonsx/statemachines/westworld"
)

var MinerGoHomeAndSleepTilRestedState *minerGoHomeAndSleepTilRestedState

type minerGoHomeAndSleepTilRestedState struct {
}

func (state *minerGoHomeAndSleepTilRestedState) Enter(role westworld.BaseEntity) {

	fmt.Println("Enter MinerGoHomeAndSleepTilRestedState")
}

func (state *minerGoHomeAndSleepTilRestedState) Execute(role westworld.BaseEntity) {

	fmt.Println("Execute MinerGoHomeAndSleepTilRestedState")

	miner, _ := role.(*Miner)

	if miner.Fatigue < 10 {

		miner.Machine.ChangeState(MinerEnterMineAndDigForNuggetState)

		return
	}

	miner.Fatigue -= 5
}

func (state *minerGoHomeAndSleepTilRestedState) Exit(role westworld.BaseEntity) {

	fmt.Println("Exit MinerGoHomeAndSleepTilRestedState")
}

func (state *minerGoHomeAndSleepTilRestedState) OnMessage(role westworld.BaseEntity) bool {

	return false
}
