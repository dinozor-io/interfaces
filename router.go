package interfaces

type Router interface {
	AddRoute(int8, string, func(Controller), Group)
	Routes() []Route
}
