//
// Author João Nuno.
// 
// joaonrb@gmail.com
//
package lily

import (
	"github.com/valyala/fasthttp"
	"strings"
)

type Response struct {
	Status           int
	Headers          map[string]string
	Body             string
}

func HttpError(status int) *Response {
	return &Response{Status: status, Body: fasthttp.StatusMessage(status)}
}


type IController interface {
	Handle(IController, *fasthttp.RequestCtx, map[string]string)
	PrepareResponse(IController, *fasthttp.RequestCtx, map[string]string) *Response
	Get(*fasthttp.RequestCtx, map[string]string) *Response
	Head(*fasthttp.RequestCtx, map[string]string) *Response
	Post(*fasthttp.RequestCtx, map[string]string) *Response
	Put(*fasthttp.RequestCtx, map[string]string) *Response
	Delete(*fasthttp.RequestCtx, map[string]string) *Response
	Trace(*fasthttp.RequestCtx, map[string]string) *Response
	Http400() *Response
	Http404() *Response
	Http500() *Response
}

type BaseController struct {}

func (self *BaseController) Handle(controller IController, ctx *fasthttp.RequestCtx, args map[string]string) {
	response := self.PrepareResponse(controller, ctx, args)
	ctx.SetStatusCode(response.Status)
	for header, value := range response.Headers {
		ctx.Response.Header.Add(header, value)
	}
	ctx.SetBodyString(response.Body)
}

func (self *BaseController) PrepareResponse(controller IController, ctx *fasthttp.RequestCtx,
                                            args map[string]string) *Response {
	var response *Response
	switch strings.ToUpper(ctx.Method()) {
	case "GET":
		response = controller.Get(ctx, args)
	case "POST":
		response = controller.Post(ctx, args)
	case "PUT":
		response = controller.Put(ctx, args)
	case "DELETE":
		response = controller.Delete(ctx, args)
	case "HEAD":
		response = controller.Head(ctx, args)
	case "TRACE":
		response = controller.Trace(ctx, args)
	default:
		response = HttpError(fasthttp.StatusMethodNotAllowed)
	}
	return response
}

func (self *BaseController) Get(request *fasthttp.Request, args map[string]string) *Response {
	return HttpError(fasthttp.StatusMethodNotAllowed)
}

func (self *BaseController) Head(request *fasthttp.Request, args map[string]string) *Response {
	return HttpError(fasthttp.StatusMethodNotAllowed)
}

func (self *BaseController) Post(request *fasthttp.Request, args map[string]string) *Response {
	return HttpError(fasthttp.StatusMethodNotAllowed)
}

func (self *BaseController) Put(request *fasthttp.Request, args map[string]string) *Response {
	return HttpError(fasthttp.StatusMethodNotAllowed)
}

func (self *BaseController) Delete(request *fasthttp.Request, args map[string]string) *Response {
	return HttpError(fasthttp.StatusMethodNotAllowed)
}

func (self *BaseController) Trace(request *fasthttp.Request, args map[string]string) *Response {
	return HttpError(fasthttp.StatusMethodNotAllowed)
}

func (self *BaseController) Http400() *Response {
	return HttpError(fasthttp.StatusBadRequest)
}

func (self *BaseController) Http404() *Response {
	return HttpError(fasthttp.StatusNotFound)
}

func (self *BaseController) Http500() *Response {
	return HttpError(fasthttp.StatusInternalServerError)
}

