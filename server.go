package interfaces

type Server interface {
	Router() Router
	Port() string
}
