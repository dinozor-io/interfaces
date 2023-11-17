package interfaces

type Response interface {
	Init(Controller)
	Contr() Controller
	JSON(int, map[string]any)
	StatusCode() int
	Data() map[string]any
}
