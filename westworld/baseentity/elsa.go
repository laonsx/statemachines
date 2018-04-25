package baseentity

import "github.com/laonsx/statemachines/westworld"

type Elsa struct {
	Id          int
	Machine     westworld.StateMachine
}

func (miner *Elsa) Update() {

	miner.Machine.Update()
}

func (miner *Elsa) HandleMessage() bool {

	return miner.Machine.HandleMessage()
}
