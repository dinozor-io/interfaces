package interfaces

type Route interface {
	GetMethod() int8
	GetPath() string
}
