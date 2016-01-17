//
// Copyright (c) João Nuno. All rights reserved.
//
package lily

import (
	"strings"
)

type IController interface {
	Handle(*Request, map[string]string) *Response
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

func (self *Controller) Handle(request *Request, args map[string]string) *Response {
	for _, middleware := range self.PreMiddleware() {
		middleware(request)
	}
	var response *Response
	switch strings.ToUpper(request.Method) {
	case "GET":
		response = self.Get(request, args)
	case "POST":
		response = self.Post(request, args)
	case "PUT":
		response = self.Put(request, args)
	case "DELETE":
		response = self.Delete(request, args)
	case "HEAD":
		response = self.Head(request, args)
	case "TRACE":
		response = self.Trace(request, args)
	default:
		RaiseHttp400("Wrong method")
	}
	for _, middleware := range self.PosMiddleware() {
		middleware(request, response)
	}
	return response
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