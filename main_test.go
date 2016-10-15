package lily

// Author João Nuno.
//
// joaonrb@gmail.com
//

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"os"
	"testing"
	"time"
)

type DummyController struct {
	Controller
}

func (dummy *DummyController) Get(request *fasthttp.RequestCtx, args map[string]string) *Response {
	response := NewResponse()
	if name, ok := args["name"]; ok {
		response.Body = fmt.Sprintf("<h1>I'm a dummy and my name is %s</h1>", name)
	} else {
		response.Body = "<h1>I'm a dummy</h1>"
	}
	return response
}

func (dummy *DummyController) Post(ctx *fasthttp.RequestCtx, args map[string]string) *Response {
	panic("Dummy on fire")
}

func (dummy *DummyController) Start(ctx *fasthttp.RequestCtx, args map[string]string) (bool, *Response) {
	if string(ctx.Method()) == "PUT" {
		return false, &Response{Status: 403, Headers: map[string]string{}, Body: "You cannot be here."}
	}
	return true, nil
}

func (dummy *DummyController) Finish(request *fasthttp.RequestCtx, args map[string]string, response *Response) {
	dummy.Controller.Finish(request, args, response)
	response.Headers["x-dummy"] = "dummy"
}

func TestMain(m *testing.M) {

	var (
		controller     IController = &DummyController{}
		base           IController     = &BaseController{}
		jsonController IController = &JsonController{}
	)

	Url("/", jsonController)
	Url("/base/", base)
	Url("/:(?P<name>^[a-zA-Z0-9]+)$", controller)

	server := fasthttp.Server{Handler: CoreHandler, Name: "Dummy Server 0.69 Alpha"}
	go server.ListenAndServe("0.0.0.0:3333")
	<-time.After(20 * time.Millisecond)
	os.Exit(m.Run())
}
