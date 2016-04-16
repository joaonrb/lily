//
// Author João Nuno.
//
// joaonrb@gmail.com
//
package lily

import (
	"github.com/valyala/fasthttp"
)

const contentType = "ContentType"

type Context map[string]interface{}

type Request struct {
	*fasthttp.Request
	Context Context
	ctx     *fasthttp.RequestCtx
}

func (self *Request) Method() string {
	return string(self.Header.Method())
}

func (self *Request) RemoteAddr() string {
	return self.ctx.RemoteIP().String()
}

type IInitializer interface {
	Start(*fasthttp.RequestCtx) *Request
	Register(middleware RequestMiddleware)
}

type RequestInitializer struct {
	middleware []RequestMiddleware
}

func NewRequestInitializer() *RequestInitializer {
	return &RequestInitializer{[]RequestMiddleware{}}
}

func (self *RequestInitializer) Start(request *fasthttp.RequestCtx) *Request {
	lilyRequest := &Request{
		Request: &request.Request,
		Context: Context{contentType: "text/html"},
		ctx:     request,
	}
	for _, middleware := range self.middleware {
		middleware(lilyRequest)
	}
	return lilyRequest
}

func (self *RequestInitializer) Register(middleware RequestMiddleware) {
	self.middleware = append(self.middleware, middleware)
}
