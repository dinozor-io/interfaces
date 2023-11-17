package interfaces

type Route interface {
	Init(int8, string, func(Controller))
	Method() int8
	Path() string
	Callback() func(Controller)
}
