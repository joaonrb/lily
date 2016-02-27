//
// Copyright (c) João Nuno. All rights reserved.
//
package hello

import (
	"lily"
)

type HelloWorldController struct {
	lily.Controller
}

func (self *HelloWorldController) Get(request *lily.Request, args map[string]string) *lily.Response {
	response := lily.NewResponse()
	response.Body = "<h1>Hello World!</h1>"
	return response
}