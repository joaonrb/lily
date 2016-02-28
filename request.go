//
// Author João Nuno.
//
// joaonrb@gmail.com
//
package lily

import (
	"net/http"
)

type Context map[string]interface{}

type Request struct {
	http.Request
	Context      Context
}

type IInitializer interface {
	Start(*http.Request) *Request
	Register(middleware RequestMiddleware)
}

type RequestInitializer struct {
	middleware []RequestMiddleware
}

func NewRequestInitializer() *RequestInitializer {
	return &RequestInitializer{[]RequestMiddleware{}}
}

func (self *RequestInitializer) Start(request *http.Request) *Request {
	lilyRequest := &Request{
		Request: *request,
		Context: Context {},
	}
	for _, middleware := range self.middleware {
		middleware(lilyRequest)
	}
	return lilyRequest
}

func (self *RequestInitializer) Register(middleware RequestMiddleware) {
	self.middleware = append(self.middleware, middleware)
}