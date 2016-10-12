package lily

import (
	"reflect"
	"testing"
	"github.com/valyala/fasthttp"
)

//
// Author João Nuno.
//
// joaonrb@gmail.com
//

func TestController(t *testing.T) {
	controller, args := getController([]byte("/ass"))
	if reflect.TypeOf(controller) != reflect.TypeOf(&DummyController{}) {
		t.Error("Contoller is not dummy")
	}
	r := fasthttp.RequestHeader{}
	r.SetMethod("GET")
	ctx := &fasthttp.RequestCtx{Request: fasthttp.Request{Header: r}}
	controller.Handle(ctx, args)
	if ctx.Response.StatusCode() != 200 {
		t.Errorf("Status is %d instead of 200", ctx.Response.StatusCode())
	}
	if string(ctx.Response.Body()) != "<h1>I'm a dummy and my name is ass</h1>" {
		t.Error("Body wasn't the expected.")
	}
}
