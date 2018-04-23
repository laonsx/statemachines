package entity

type BaseEntity interface {
	Update()
	HandleMessage() bool
}

type StateMachine interface {
	Update()
	ChangeState(newState State)
	RevertToPreviousState()
	HandleMessage() bool
}

type State interface {
	Enter(role BaseEntity)
	Execute(role BaseEntity)
	Exit(role BaseEntity)
	OnMessage(role BaseEntity) bool
}
