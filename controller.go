//
// Author João Nuno.
//
// joaonrb@gmail.com
//
package lily

type IController interface {
	Get(*Request, map[string]string) *Response
	Head(*Request, map[string]string) *Response
	Post(*Request, map[string]string) *Response
	Put(*Request, map[string]string) *Response
	Delete(*Request, map[string]string) *Response
	Trace(*Request, map[string]string) *Response
	RegisterPre(RequestMiddleware)
	RegisterPos(ResponseMiddleware)
	PreMiddleware() []RequestMiddleware
	PosMiddleware() []ResponseMiddleware
}

type Controller struct {
	preMiddleware []RequestMiddleware
	posMiddleware []ResponseMiddleware
}

func NewController() IController {
	return &Controller{[]RequestMiddleware{}, []ResponseMiddleware{}}
}

func (self *Controller) Get(request *Request, args map[string]string) *Response {
	RaiseHttp404()
	return nil
}

func (self *Controller) Head(request *Request, args map[string]string) *Response {
	RaiseHttp404()
	return nil
}

func (self *Controller) Post(request *Request, args map[string]string) *Response {
	RaiseHttp404()
	return nil
}

func (self *Controller) Put(request *Request, args map[string]string) *Response {
	RaiseHttp404()
	return nil
}

func (self *Controller) Delete(request *Request, args map[string]string) *Response {
	RaiseHttp404()
	return nil
}

func (self *Controller) Trace(request *Request, args map[string]string) *Response {
	RaiseHttp404()
	return nil
}

func (self *Controller) RegisterPre(middleware RequestMiddleware) {
	self.preMiddleware = append(self.preMiddleware, middleware)
}

func (self *Controller) RegisterPos(middleware ResponseMiddleware) {
	self.posMiddleware = append(self.posMiddleware, middleware)
}

func (self *Controller) PreMiddleware() []RequestMiddleware {
	return self.preMiddleware
}

func (self *Controller) PosMiddleware() []ResponseMiddleware {
	return self.posMiddleware
}