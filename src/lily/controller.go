//
// Copyright (c) João Nuno. All rights reserved.
//
package lily

type IController interface {
	Get(request Request, args ...string) *Response
	Head(request Request, args ...string) *Response
	Post(request Request, args ...string) *Response
	Put(request Request, args ...string) *Response
	Delete(request Request, args ...string) *Response
	Trace(request Request, args ...string) *Response
	Handle(request Request, args ...string) *Response
	RegisterPreMiddleware(RequestMiddleware)
	RegisterPosMiddleware(ResponseMiddleware)
}