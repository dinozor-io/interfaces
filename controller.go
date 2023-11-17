package interfaces

type Controller interface {
	Init(Request, Response)
	Req() Request
	Res() Response
}
