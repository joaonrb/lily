//
// Copyright (c) João Nuno. All rights reserved.
//
package lily

import (
	"net/http"
	"time"
)

var (
	startRequestQueue = []RequestMiddleware{}
)

func RegisterStartRequestMiddleware(middleware RequestMiddleware) {
	startRequestQueue = append(startRequestQueue, middleware)
}

type Context map[string]interface{}

type Request struct {
	http.Request
	Context      Context
}

func startRequest(request *http.Request) *Request {
	lilyRequest := &Request{
		Request: request,
		Context: Context {},
	}
	for _, middleware := range startRequestQueue {
		middleware(request)
	}
	return lilyRequest
}