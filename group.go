package interfaces

type Group interface {
	Cond(func(Controller) bool)
	CheckCond(Controller) bool
}
