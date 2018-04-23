package baseent

import "github.com/laonsx/statemachines/entity"

type Miner struct {
	Id          int
	Machine     entity.StateMachine
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
