package state

import (
	"github.com/laonsx/statemachines/westworld"
	"fmt"
)

var ElsaCookStewState *elsaCookStewState

type elsaCookStewState struct {
}

func (state *elsaCookStewState) Enter(role westworld.BaseEntity) {

	fmt.Println("Enter ElsaCookStewState")
}

func (state *elsaCookStewState) Execute(role westworld.BaseEntity) {

	fmt.Println("Execute ElsaCookStewState")
}

func (state *elsaCookStewState) Exit(role westworld.BaseEntity) {

	fmt.Println("Exit ElsaCookStewState")
}

func (state *elsaCookStewState) OnMessage(role westworld.BaseEntity) bool {

	return false
}
