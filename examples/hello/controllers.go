//
// Author João Nuno.
//
package hello

import (
	"lily"
	"fmt"
)

type HelloWorldController struct {
	lily.Controller
}

func (self *HelloWorldController) Get(request *lily.Request, args map[string]string) *lily.Response {
	response := lily.NewResponse()
	response.Body = "<h1>Hello World!</h1>"
	return response
}

type RegexHelloWorldController struct {
	lily.Controller
}

func (self *RegexHelloWorldController) Get(request *lily.Request, args map[string]string) *lily.Response {
	response := lily.NewResponse()
	response.Body = fmt.Sprintf("<h1>Hello %s!</h1>", args["user"])
	return response
}