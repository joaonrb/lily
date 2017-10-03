package router


var (
	NoRouteException = &e{"No route for this context"}
	EmptyComponentException = &e{"No component at the end of this route"}
)


type e struct {
	kind string
}

func (err *e) String() string  {
	return err.kind
}

func (err *e) Error() string {
	return err.String()
}

func (err *e) Resolve(_ interface{}) interface {} {
	return err
}