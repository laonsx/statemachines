package state

import (
	"fmt"
	"github.com/laonsx/statemachines/westworld"
	"github.com/laonsx/statemachines/westworld/baseentity"
)

var ElsaHouseWorkState *elsaHouseWorkState

type elsaHouseWorkState struct {
}

func (state *elsaHouseWorkState) Enter(role westworld.BaseEntity) {

	fmt.Println("Enter ElsaHouseWorkState")
}

func (state *elsaHouseWorkState) Execute(role westworld.BaseEntity) {

	fmt.Println("Execute ElsaHouseWorkState")
}

func (state *elsaHouseWorkState) Exit(role westworld.BaseEntity) {

	fmt.Println("Exit ElsaHouseWorkState")
}

func (state *elsaHouseWorkState) OnMessage(role westworld.BaseEntity, msg interface{}) bool {

	elsa := role.(*baseentity.Elsa)
	elsa.Machine.ChangeState(ElsaCookStewState)
	return true
}
