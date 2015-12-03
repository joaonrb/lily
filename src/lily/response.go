//
// Copyright (c) Telefonica I+D. All rights reserved.
//
package lily

import (
	"net/http"
	"fmt"
)

type Response struct {
	Status  int
	Headers map[string]string
	Body    string
	RW      http.ResponseWriter
}

func NewResponse(rw http.ResponseWriter) *Response {
	return &Response{200, map[string]string{}, "", rw}
}

type IFinalizer interface {
	Finish(*Request, *Response,http.ResponseWriter)
	RegisterPosController(middleware ResponseMiddleware)
	RegisterFinish(middleware ResponseMiddleware)
}

type Finalizer struct {
	posControllerMiddleware []ResponseMiddleware
	finishMiddleware        []ResponseMiddleware
}

func (self *Finalizer) Finish(request *Request, response *Response, rw http.ResponseWriter) {
	err := recover()
	if err != nil {
		switch err := err.(type) {
		case Http404:
			response = err
		case HttpError:
			response = err
			Error(err.Error())
		case error:
			response = Http500(err, http.StatusInternalServerError, HTTP_500_MESSAGE)
			Error(err.Error())
		default:
			response = Http500(fmt.Sprintf("%x", err), http.StatusInternalServerError, HTTP_500_MESSAGE)
			Error(response.(Http500).Error())
		}
	}
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