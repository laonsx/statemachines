package westworld

type BaseEntity interface {
	GetId() int
	Update()
	HandleMessage(msg interface{}) bool
}

type StateMachine interface {
	Update()
	ChangeState(newState State)
	RevertToPreviousState()
	HandleMessage(msg interface{}) bool
}

type State interface {
	Enter(role BaseEntity)
	Execute(role BaseEntity)
	Exit(role BaseEntity)
	OnMessage(role BaseEntity, msg interface{}) bool
}
