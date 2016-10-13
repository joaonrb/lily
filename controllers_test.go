package lily

import (
	"github.com/valyala/fasthttp"
	"testing"
)

//
// Author João Nuno.
//
// joaonrb@gmail.com
//
func TestController(t *testing.T) {

	for _, method := range []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "TRACE"} {
		ctx := MockRequest(method, "/base")
		if ctx.Response.StatusCode() != fasthttp.StatusMethodNotAllowed {
			t.Errorf("Status is %d instead of 405", ctx.Response.StatusCode())
		}
		if string(ctx.Response.Body()) != fasthttp.StatusMessage(fasthttp.StatusMethodNotAllowed) {
			t.Error("Body wasn't the expected.")
		}
	}
}

func TestDummyController(t *testing.T) {
	ctx := MockRequest("GET", "/ass")
	if ctx.Response.StatusCode() != 200 {
		t.Errorf("Status is %d instead of 200", ctx.Response.StatusCode())
	}
	if string(ctx.Response.Body()) != "<h1>I'm a dummy and my name is ass</h1>" {
		t.Error("Body wasn't the expected.")
	}
}
