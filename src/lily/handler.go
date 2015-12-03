//
// Copyright (c) João Nuno. All rights reserved.
//
package lily

import (
	"net/http"
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

func NewLilyHandler(init IInitializer, router IRouter, finish IFinalizer) *Handler {
	return &Handler{init, finish, router}
}

func (self *Handler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	var response *Response 
	var lilyRequest *Request
	defer self.finish.Finish(lilyRequest, response, responseWriter)
	lilyRequest = self.init.Start(request)

	controller, params := self.router.Parse(lilyRequest)
	
	response = HandleController(controller, lilyRequest, params)
	
}

func (self *Handler) Initializer() IInitializer { return self.init }

func (self *Handler) Finalizer() IFinalizer { return self.finish }