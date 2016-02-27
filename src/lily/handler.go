//
// Copyright (c) João Nuno. All rights reserved.
//
package lily

import (
	"net/http"
	"fmt"
	"strings"
)

type IHandler interface {
	Initializer() IInitializer
	Finalizer() IFinalizer
}

/**
 * Implements http.Handler
 */
type Handler struct {
	init   IInitializer
	finish IFinalizer
}

func NewHandler(init IInitializer, finish IFinalizer) *Handler {
	return &Handler{init, finish}
}

func (self *Handler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	var response *Response 
	var lilyRequest *Request
	defer func() {
		err := recover()
		if err != nil {
			switch err := err.(type) {
			case IHttpError:
				response = err.ToResponse()
				Debug("Http error caught with status %d and message %s", response.Status, response.Body)
			case error:
				response = NewHttp500(err.Error()).ToResponse()
				Error(err.Error())
			default:
				msg := fmt.Sprintf("%x", err)
				response = NewHttp500(msg).ToResponse()
				Error(msg)
			}
		}
		self.finish.Finish(lilyRequest, response, responseWriter)
	}()
	lilyRequest = self.init.Start(request)

	controller, params, err := mainRouter.Parse(lilyRequest.URL.Path)
	if err != nil {
		panic(err)
	}
	
	response = self.Handle(controller, lilyRequest, params)
}



func (self *Handler) Handle(controller IController, request *Request, args map[string]string) *Response {
	for _, middleware := range controller.PreMiddleware() {
		middleware(request)
	}
	var response *Response
	switch strings.ToUpper(request.Method) {
	case "GET":
		response = controller.Get(request, args)
	case "POST":
		response = controller.Post(request, args)
	case "PUT":
		response = controller.Put(request, args)
	case "DELETE":
		response = controller.Delete(request, args)
	case "HEAD":
		response = controller.Head(request, args)
	case "TRACE":
		response = controller.Trace(request, args)
	default:
		RaiseHttp400("Wrong method")
	}
	for _, middleware := range controller.PosMiddleware() {
		middleware(request, response)
	}
	return response
}

func (self *Handler) Initializer() IInitializer { return self.init }

func (self *Handler) Finalizer() IFinalizer { return self.finish }