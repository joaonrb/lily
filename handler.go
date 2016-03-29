//
// Author João Nuno.
//
package lily

import (
	"net/http"
	"fmt"
)

var mainHandler IHandler

func ForceRegisterHandler(handler IHandler) {
	mainHandler = handler
}

func RegisterHandler(handler IHandler) bool {
	if mainHandler != nil {
		ForceRegisterHandler(handler)
		return true
	}
	return false
}

func defaultHandler() IHandler {
	return NewHandler(NewRequestInitializer(), NewFinalizer())
}

type IHandler interface {
	ServeHTTP(responseWriter http.ResponseWriter, request *http.Request)
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
	for _, middleware := range controller.PreMiddleware() {
		middleware(lilyRequest)
	}
	response = controller.Handle(controller, lilyRequest, params)
	for _, middleware := range controller.PosMiddleware() {
		middleware(lilyRequest)
	}
}

func (self *Handler) Initializer() IInitializer { return self.init }

func (self *Handler) Finalizer() IFinalizer { return self.finish }