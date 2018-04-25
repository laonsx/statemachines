package miner

import "github.com/laonsx/statemachines/westworld"

func InitMiner() *Miner {
	miner := &Miner{}
	machine := &MinerStateMachine{
		Owner:     miner,
		CurrState: MinerEnterMineAndDigForNuggetState,
	}
	miner.Machine = machine

	return miner
}

type Miner struct {
	Id          int
	Machine     westworld.StateMachine
	GoldCarried int
	MoneyInBank int
	Thirst      int
	Fatigue     int
}

func (miner *Miner) Update() {

	miner.Machine.Update()
}

func (miner *Miner) HandleMessage() bool {

	return miner.Machine.HandleMessage()
}
