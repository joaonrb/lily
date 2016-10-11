package lily

//
// Author João Nuno.
//
// joaonrb@gmail.com
//

import (
	"os"
	"testing"
	"github.com/valyala/fasthttp"
	"fmt"
)


type DummyController struct {
	BaseController
}

func (self *DummyController) Get(request *fasthttp.RequestCtx, args map[string]string) *Response {
	response := NewResponse()
	if name, ok := args["name"]; ok {
		response.Body = fmt.Sprintf("<h1>I'm a dummy and my name is %s</h1>", name)
	} else {
		response.Body = "<h1>I'm a dummy</h1>"
	}
	return response
}

func TestMain(m *testing.M) {

	controller := &DummyController{}

	Url("/", controller)
	Url("/:(?P<name>\\w+)", controller)

	os.Exit(m.Run())
}
