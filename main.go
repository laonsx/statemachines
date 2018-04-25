package main

import (
	"fmt"

	"github.com/laonsx/statemachines/westworld/baseentity"
	"github.com/laonsx/statemachines/westworld/statemachine"
	"github.com/laonsx/statemachines/westworld/state"
)

func main() {

	miner := &baseentity.Miner{}
	machine := &statemachine.MinerStateMachine{
		Owner:     miner,
		CurrState: state.MinerEnterMineAndDigForNuggetState,
	}
	miner.Machine = machine

	//miner := miner.InitMiner()
	for i := 0; i < 100; i++ {

		miner.Update()

		fmt.Println("i =", i, "==>", miner)
	}
}
