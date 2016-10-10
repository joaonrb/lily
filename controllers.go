//
// Author João Nuno.
// 
// joaonrb@gmail.com
//
package lily

import (
	"github.com/valyala/fasthttp"
	"bytes"
)

type Response struct {
	Status           int
	Headers          map[string]string
	Body             string
}

func NewResponse() *Response {
	return &Response{Status: fasthttp.StatusOK}
}

var (
	HttpError = func(status int) *Response { return &Response{Status: status, Body: fasthttp.StatusMessage(status)} }
	http400   = HttpError(fasthttp.StatusBadRequest)
	http404   = HttpError(fasthttp.StatusNotFound)
	http405   = HttpError(fasthttp.StatusMethodNotAllowed)
	http500   = HttpError(fasthttp.StatusInternalServerError)
	Http400   = func() *Response { return http400 }
	Http404   = func() *Response { return http404 }
	Http405   = func() *Response { return http405 }
	Http500   = func() *Response { return http500}
)

func sendResponse(ctx *fasthttp.RequestCtx, response *Response) {
	ctx.SetStatusCode(response.Status)
	for header, value := range response.Headers {
		ctx.Response.Header.Add(header, value)
	}
	ctx.SetBodyString(response.Body)
}


type IController interface {
	Handle(IController, *fasthttp.RequestCtx, map[string]string)
	Start(*fasthttp.RequestCtx, map[string]string)
	Finish(*Response)
	Get(*fasthttp.RequestCtx, map[string]string) *Response
	Head(*fasthttp.RequestCtx, map[string]string) *Response
	Post(*fasthttp.RequestCtx, map[string]string) *Response
	Put(*fasthttp.RequestCtx, map[string]string) *Response
	Patch(*fasthttp.RequestCtx, map[string]string) *Response
	Delete(*fasthttp.RequestCtx, map[string]string) *Response
	Trace(*fasthttp.RequestCtx, map[string]string) *Response
}

type BaseController struct {}

// Only touch Handle method if you understand what you are doing.
func (self *BaseController) Handle(controller IController, ctx *fasthttp.RequestCtx, args map[string]string) {
	controller.Start(ctx, args)
	var response *Response
	defer func() {
		if recovery := recover(); recovery != nil {
			Error("Unexpected error on call %s %s: %v", ctx.Method(), ctx.Path(), recovery)
			response = Http500()
		}
		controller.Finish(response)
		sendResponse(ctx, response)
	}()
	switch string(bytes.ToUpper(ctx.Method())) {
	case "GET":
		response = controller.Get(ctx, args)
	case "POST":
		response = controller.Post(ctx, args)
	case "PUT":
		response = controller.Put(ctx, args)
	case "PATCH":
		response = controller.Patch(ctx, args)
	case "DELETE":
		response = controller.Delete(ctx, args)
	case "HEAD":
		response = controller.Head(ctx, args)
	case "TRACE":
		response = controller.Trace(ctx, args)
	default:
		response = Http405()
	}
}

func (self *BaseController) Start(*fasthttp.RequestCtx, map[string]string) {}

func (self *BaseController) Finish(*Response) {}

func (self *BaseController) Get(request *fasthttp.RequestCtx, args map[string]string) *Response {
	return Http405()
}

func (self *BaseController) Head(request *fasthttp.RequestCtx, args map[string]string) *Response {
	return Http405()
}

func (self *BaseController) Post(request *fasthttp.RequestCtx, args map[string]string) *Response {
	return Http405()
}

func (self *BaseController) Put(request *fasthttp.RequestCtx, args map[string]string) *Response {
	return Http405()
}

func (self *BaseController) Patch(request *fasthttp.RequestCtx, args map[string]string) *Response {
	return Http405()
}

func (self *BaseController) Delete(request *fasthttp.RequestCtx, args map[string]string) *Response {
	return Http405()
}

func (self *BaseController) Trace(request *fasthttp.RequestCtx, args map[string]string) *Response {
	return Http405()
}

