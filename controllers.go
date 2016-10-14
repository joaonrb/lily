package lily

// Author João Nuno.
//
// joaonrb@gmail.com
//

import (
	"bytes"
	"github.com/valyala/fasthttp"
)

type Response struct {
	Status  int
	Headers map[string]string
	Body    string
}

// Creates a new response with 200 status as default
func NewResponse() *Response {
	return &Response{Status: fasthttp.StatusOK, Headers: map[string]string{}}
}

var (
	// Default HttpError processor
	HttpError = func(status int) *Response { return &Response{
		Status: status,
		Headers: map[string]string{},
		Body: fasthttp.StatusMessage(status),
	} }
	http400   = HttpError(fasthttp.StatusBadRequest)
	http404   = HttpError(fasthttp.StatusNotFound)
	http405   = HttpError(fasthttp.StatusMethodNotAllowed)
	http500   = HttpError(fasthttp.StatusInternalServerError)
	// Default Http 400 error
	Http400 = func() *Response { return http400 }
	// Default Http 404 error
	Http404 = func() *Response { return http404 }
	// Default Http 405 error
	Http405 = func() *Response { return http405 }
	// Default Http 500 error
	Http500 = func() *Response { return http500 }
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
	Finish(*fasthttp.RequestCtx, map[string]string, *Response)
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

// Initiates the controller
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
		c.This.Finish(ctx, args, response)
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

// Check request and initiates any process required before handling
func (c *BaseController) Start(*fasthttp.RequestCtx, map[string]string) (bool, *Response) {
	return true, nil
}

// Close the response. Add any header or so.
func (c *BaseController) Finish(request *fasthttp.RequestCtx, args map[string]string, response *Response) {}

// Get method implementation
func (c *BaseController) Get(request *fasthttp.RequestCtx, args map[string]string) *Response {
	return Http405()
}

// Get method implementation
func (c *BaseController) Head(request *fasthttp.RequestCtx, args map[string]string) *Response {
	return Http405()
}

// Post method implementation
func (c *BaseController) Post(request *fasthttp.RequestCtx, args map[string]string) *Response {
	return Http405()
}

// Put method implementation
func (c *BaseController) Put(request *fasthttp.RequestCtx, args map[string]string) *Response {
	return Http405()
}

// Patch method implementation
func (c *BaseController) Patch(request *fasthttp.RequestCtx, args map[string]string) *Response {
	return Http405()
}

// Delete method implementation
func (c *BaseController) Delete(request *fasthttp.RequestCtx, args map[string]string) *Response {
	return Http405()
}

// Trace method implementation
func (c *BaseController) Trace(request *fasthttp.RequestCtx, args map[string]string) *Response {
	return Http405()
}
