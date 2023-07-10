package interfaces

type Engine interface {
	Run()
	SetPort(string)
	SetRoutes([]Route)
}
