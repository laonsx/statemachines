package state

import (
	"fmt"

	"github.com/laonsx/statemachines/westworld"
	"github.com/laonsx/statemachines/westworld/baseentity"
	"github.com/laonsx/statemachines/westworld/sence"
)

var MinerGoHomeAndSleepTilRestedState *minerGoHomeAndSleepTilRestedState

type minerGoHomeAndSleepTilRestedState struct {
}

func (state *minerGoHomeAndSleepTilRestedState) Enter(role westworld.BaseEntity) {

	sence.DispatchMessage(202, 101, sence.MSG_TEST1)

	fmt.Println("Enter MinerGoHomeAndSleepTilRestedState")
}

func (state *minerGoHomeAndSleepTilRestedState) Execute(role westworld.BaseEntity) {

	fmt.Println("Execute MinerGoHomeAndSleepTilRestedState")

	miner, _ := role.(*baseentity.Miner)

	if miner.Fatigue < 10 {

		miner.Machine.ChangeState(MinerEnterMineAndDigForNuggetState)

		return
	}

	miner.Fatigue -= 5
}

func (state *minerGoHomeAndSleepTilRestedState) Exit(role westworld.BaseEntity) {

	fmt.Println("Exit MinerGoHomeAndSleepTilRestedState")
}

func (state *minerGoHomeAndSleepTilRestedState) OnMessage(role westworld.BaseEntity, msg interface{}) bool {

	return false
}
