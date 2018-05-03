package baseentity

import "github.com/laonsx/statemachines/westworld"

type Miner struct {
	Id          int
	Machine     westworld.StateMachine
	GoldCarried int
	MoneyInBank int
	Thirst      int
	Fatigue     int
}

func (miner *Miner) GetId() int {

	return miner.Id
}

func (miner *Miner) Update() {

	miner.Machine.Update()
}

func (miner *Miner) HandleMessage(msg interface{}) bool {

	return miner.Machine.HandleMessage(msg)
}
