//
// Copyright (c) João Nuno. All rights reserved.
//
package lily


var (
	preRouting    = map[string]RequestMiddleware{}
	posRouting    = map[string]RequestMiddleware{}
	preController = map[string]RequestMiddleware{}
	posController = map[string]ResponseMiddleware{}
	closeResponse = map[string]ResponseMiddleware{}
)

type RequestMiddleware func(*Request)

type ResponseMiddleware func(*Request, *Response)
