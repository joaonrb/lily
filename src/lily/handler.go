//
// Copyright (c) João Nuno. All rights reserved.
//
package lily

import (
	"net/http"
	"fmt"
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
	router IRouter
}

func NewHandler(init IInitializer, router IRouter, finish IFinalizer) *Handler {
	return &Handler{init, finish, router}
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

	controller, params := self.router.Parse(lilyRequest.URL.Path)
	
	response = HandleController(controller, lilyRequest, params)
}

func (self *Handler) Initializer() IInitializer { return self.init }

func (self *Handler) Finalizer() IFinalizer { return self.finish }