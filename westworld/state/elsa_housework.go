package state

import (
	"github.com/laonsx/statemachines/westworld"
	"fmt"
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

func (state *elsaHouseWorkState) OnMessage(role westworld.BaseEntity) bool {

	return false
}
