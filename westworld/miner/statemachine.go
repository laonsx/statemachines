package miner

import "github.com/laonsx/statemachines/westworld"

type MinerStateMachine struct {
	Owner         westworld.BaseEntity
	GlobalState   westworld.State
	CurrState     westworld.State
	PreviousState westworld.State
}

func (machine *MinerStateMachine) Update() {

	if machine.GlobalState != nil {

		machine.GlobalState.Execute(machine.Owner)
	}

	if machine.CurrState != nil {

		machine.CurrState.Execute(machine.Owner)
	}
}

func (machine *MinerStateMachine) ChangeState(newState westworld.State) {

	machine.PreviousState = machine.CurrState
	machine.CurrState.Exit(machine.Owner)

	machine.CurrState = newState
	machine.CurrState.Enter(machine.Owner)
}

func (machine *MinerStateMachine) RevertToPreviousState() {

	machine.ChangeState(machine.PreviousState)
}

func (machine *MinerStateMachine) HandleMessage() bool {

	if machine.CurrState != nil && machine.CurrState.OnMessage(machine.Owner) {

		return true
	}

	if machine.GlobalState != nil && machine.GlobalState.OnMessage(machine.Owner) {

		return true
	}

	return false
}
