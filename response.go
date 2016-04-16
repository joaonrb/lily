package lily
//
// Author João Nuno.
//
// joaonrb@gmail.com
//

import (
	"github.com/valyala/fasthttp"
)

type Response struct {
	Status           int
	Headers          map[string]string
	Body             string
	FastHttpResponse *fasthttp.Response
}

func NewResponse() *Response {
	return &Response{fasthttp.StatusOK, map[string]string{}, "", nil}
}

type IFinalizer interface {
	Finish(*Request, *Response)
	RegisterPosController(middleware ResponseMiddleware)
	RegisterFinish(middleware ResponseMiddleware)
}

type Finalizer struct {
	posControllerMiddleware []ResponseMiddleware
	finishMiddleware        []ResponseMiddleware
}

func NewFinalizer() *Finalizer {
	return &Finalizer{[]ResponseMiddleware{}, []ResponseMiddleware{}}
}

func (self *Finalizer) Finish(request *Request, response *Response) {
	for _, middleware := range self.posControllerMiddleware {
		middleware(request, response)
	}
	if response.FastHttpResponse != nil {
		for header, value := range response.Headers {
			response.FastHttpResponse.Header.Add(header, value)
		}
		response.FastHttpResponse.SetStatusCode(response.Status)
		response.FastHttpResponse.SetBodyString(response.Body)
	}
	for _, middleware := range self.finishMiddleware {
		middleware(request, response)
	}
}

func (self *Finalizer) RegisterPosController(middleware ResponseMiddleware) {
	self.posControllerMiddleware = append(self.posControllerMiddleware, middleware)
}
func (self *Finalizer) RegisterFinish(middleware ResponseMiddleware) {
	self.finishMiddleware = append(self.finishMiddleware, middleware)
}
