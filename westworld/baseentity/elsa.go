package baseentity

import (
	"github.com/laonsx/statemachines/westworld"
)

type Elsa struct {
	Id          int
	Machine     westworld.StateMachine
}

func (slsa *Elsa) GetId() int {

	return slsa.Id
}

func (slsa *Elsa) Update() {

	slsa.Machine.Update()
}

func (slsa *Elsa) HandleMessage() bool {

	return slsa.Machine.HandleMessage()
}
