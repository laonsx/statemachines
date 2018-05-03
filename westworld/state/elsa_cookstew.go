package state

import (
	"github.com/laonsx/statemachines/westworld"
	"github.com/laonsx/statemachines/westworld/baseentity"

	"fmt"
)

var ElsaCookStewState *elsaCookStewState

type elsaCookStewState struct {
}

func (state *elsaCookStewState) Enter(role westworld.BaseEntity) {

	fmt.Println("Enter ElsaCookStewState")
}

func (state *elsaCookStewState) Execute(role westworld.BaseEntity) {

	elsa := role.(*baseentity.Elsa)

	elsa.CookTime++

	if elsa.CookTime == 10 {

		elsa.Machine.RevertToPreviousState()

		return
	}

	fmt.Println("Execute ElsaCookStewState", elsa.CookTime)
}

func (state *elsaCookStewState) Exit(role westworld.BaseEntity) {

	elsa := role.(*baseentity.Elsa)
	elsa.CookTime = 0

	fmt.Println("Exit ElsaCookStewState")
}

func (state *elsaCookStewState) OnMessage(role westworld.BaseEntity, msg interface{}) bool {

	return false
}
