package interfaces

type Request interface {
	Init(Controller)
	Contr() Controller
	// TODO: NO COMPLETED
}
