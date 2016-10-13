// Package lily
// Author João Nuno.
//
// joaonrb@gmail.com
//
package lily

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"os"
	"testing"
)

type DummyController struct {
	BaseController
}

// Get for dummy
func (dummy *DummyController) Get(request *fasthttp.RequestCtx, args map[string]string) *Response {
	response := NewResponse()
	if name, ok := args["name"]; ok {
		response.Body = fmt.Sprintf("<h1>I'm a dummy and my name is %s</h1>", name)
	} else {
		response.Body = "<h1>I'm a dummy</h1>"
	}
	return response
}

// Init tests
func TestMain(m *testing.M) {

	var (
		controller IController = &DummyController{}
		base       IController = &BaseController{}
	)

	Url("/", controller)
	Url("/:(?P<name>\\w+)", controller)
	Url("/base", base)
	os.Exit(m.Run())
}
