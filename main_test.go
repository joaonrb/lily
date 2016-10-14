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
)

type DummyController struct {
	BaseController
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
	request.SetContentType("text/html")
	response.Headers["x-dummy"] = "dummy"
}


func TestMain(m *testing.M) {

	var (
		controller IController = &DummyController{}
		base       IController = &BaseController{}
	)

	Url("/", controller)
	Url("/:(?P<name>\\w+)", controller)
	Url("/base", base)

	server := fasthttp.Server{Handler: CoreHandler, Name: "Dummy Server 0.69 Alpha"}
	go os.Exit(m.Run())
	panic(server.ListenAndServe(":3333"))
}
