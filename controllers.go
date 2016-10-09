//
// Author João Nuno.
// 
// joaonrb@gmail.com
//
package lily

import (
	"github.com/valyala/fasthttp"
)

type Response struct {
	Status           int
	Headers          map[string]string
	Body             string
}

type IController interface {
	Handle(IController, *fasthttp.RequestCtx, map[string]string)
	Get(*fasthttp.Request, map[string]string) *Response
	Head(*fasthttp.Request, map[string]string) *Response
	Post(*fasthttp.Request, map[string]string) *Response
	Put(*fasthttp.Request, map[string]string) *Response
	Delete(*fasthttp.Request, map[string]string) *Response
	Trace(*fasthttp.Request, map[string]string) *Response
}