package interfaces

type Response interface {
	Init(Controller)
	Type() int
	Contr() Controller
	JSON(int, map[string]any)
	HTML(int, string)
	StatusCode() int
	Data() map[string]any
}
