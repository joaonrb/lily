//
// Copyright (c) João Nuno. All rights reserved.
//
package lily

import (
	"net/http"
)

type Response struct {
	Status  int
	Headers map[string]string
	Body    string
	RW      http.ResponseWriter
}

func NewResponse() *Response {
	return &Response{http.StatusOK, map[string]string{}, "", nil}
}

type IFinalizer interface {
	Finish(*Request, *Response, http.ResponseWriter)
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

func (self *Finalizer) Finish(request *Request, response *Response, rw http.ResponseWriter) {
	response.RW = rw
	for _, middleware := range self.posControllerMiddleware {
		middleware(request, response)
	}
	for header, value := range response.Headers { response.RW.Header().Add(header, value) }
	response.RW.WriteHeader(response.Status)
	response.RW.Write([]byte(response.Body))
	
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