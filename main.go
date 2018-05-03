package main

import (
	"github.com/laonsx/statemachines/westworld"
	"github.com/laonsx/statemachines/westworld/baseentity"
	"github.com/laonsx/statemachines/westworld/sence"
	"github.com/laonsx/statemachines/westworld/state"
	"github.com/laonsx/statemachines/westworld/statemachine"
)

func main() {

	miner := &baseentity.Miner{}
	minerMachine := &statemachine.MinerStateMachine{
		Owner:     miner,
		CurrState: state.MinerEnterMineAndDigForNuggetState,
	}
	miner.Id = 101
	miner.Machine = minerMachine

	elsa := &baseentity.Elsa{}
	elsa.Id = 202
	elsaMachine := &statemachine.MinerStateMachine{
		Owner:     elsa,
		CurrState: state.ElsaHouseWorkState,
	}
	elsa.Machine = elsaMachine

	sence.RegistEntity(miner)
	sence.RegistEntity(elsa)

	//miner := miner.InitMiner()
	for i := 0; i < 100; i++ {

		sence.EntityManager.Range(func(key, value interface{}) bool {

			entity := value.(westworld.BaseEntity)
			entity.Update()

			return true
		})
	}
}
