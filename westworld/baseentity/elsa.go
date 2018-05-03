package baseentity

import (
	"github.com/laonsx/statemachines/westworld"
)

type Elsa struct {
	Id       int
	Machine  westworld.StateMachine
	CookTime int
}

func (elsa *Elsa) GetId() int {

	return elsa.Id
}

func (elsa *Elsa) Update() {

	elsa.Machine.Update()
}

func (elsa *Elsa) HandleMessage(msg interface{}) bool {

	return elsa.Machine.HandleMessage(msg)
}
