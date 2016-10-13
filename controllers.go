// Package lily
// Author João Nuno.
//
// joaonrb@gmail.com
//
package lily

import (
	"bytes"
	"github.com/valyala/fasthttp"
)

type Response struct {
	Status  int
	Headers map[string]string
	Body    string
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
	Http500   = func() *Response { return http500 }
)

func sendResponse(ctx *fasthttp.RequestCtx, response *Response) {
	ctx.SetStatusCode(response.Status)
	for header, value := range response.Headers {
		ctx.Response.Header.Add(header, value)
	}
	ctx.SetBodyString(response.Body)
}

type IController interface {
	Init(IController)
	Handle(*fasthttp.RequestCtx, map[string]string)
	Start(*fasthttp.RequestCtx, map[string]string) (bool, *Response)
	Finish(*Response)
	Get(*fasthttp.RequestCtx, map[string]string) *Response
	Head(*fasthttp.RequestCtx, map[string]string) *Response
	Post(*fasthttp.RequestCtx, map[string]string) *Response
	Put(*fasthttp.RequestCtx, map[string]string) *Response
	Patch(*fasthttp.RequestCtx, map[string]string) *Response
	Delete(*fasthttp.RequestCtx, map[string]string) *Response
	Trace(*fasthttp.RequestCtx, map[string]string) *Response
}

type BaseController struct {
	This IController
}

func (c *BaseController) Init(controller IController) {
	c.This = controller
}

// Only touch Handle method if you understand what you are doing.
func (c *BaseController) Handle(ctx *fasthttp.RequestCtx, args map[string]string) {
	ok, response := c.This.Start(ctx, args)
	defer func() {
		if recovery := recover(); recovery != nil {
			Error("Unexpected error on call %s %s: %v", ctx.Method(), ctx.Path(), recovery)
			response = Http500()
		}
		c.This.Finish(response)
		sendResponse(ctx, response)
	}()
	if !ok {
		return
	}
	switch string(bytes.ToUpper(ctx.Method())) {
	case "GET":
		response = c.This.Get(ctx, args)
	case "POST":
		response = c.This.Post(ctx, args)
	case "PUT":
		response = c.This.Put(ctx, args)
	case "PATCH":
		response = c.This.Patch(ctx, args)
	case "DELETE":
		response = c.This.Delete(ctx, args)
	case "HEAD":
		response = c.This.Head(ctx, args)
	case "TRACE":
		response = c.This.Trace(ctx, args)
	default:
		response = Http405()
	}
}

func (c *BaseController) Start(*fasthttp.RequestCtx, map[string]string) (bool, *Response) {
	return true, nil
}

func (c *BaseController) Finish(*Response) {}

func (c *BaseController) Get(request *fasthttp.RequestCtx, args map[string]string) *Response {
	return Http405()
}

func (c *BaseController) Head(request *fasthttp.RequestCtx, args map[string]string) *Response {
	return Http405()
}

func (c *BaseController) Post(request *fasthttp.RequestCtx, args map[string]string) *Response {
	return Http405()
}

func (c *BaseController) Put(request *fasthttp.RequestCtx, args map[string]string) *Response {
	return Http405()
}

func (c *BaseController) Patch(request *fasthttp.RequestCtx, args map[string]string) *Response {
	return Http405()
}

func (c *BaseController) Delete(request *fasthttp.RequestCtx, args map[string]string) *Response {
	return Http405()
}

func (c *BaseController) Trace(request *fasthttp.RequestCtx, args map[string]string) *Response {
	return Http405()
}
