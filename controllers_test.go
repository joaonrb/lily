package lily

// Author João Nuno.
//
// joaonrb@gmail.com
//
import (
	"github.com/valyala/fasthttp"
	"testing"
)

// Test status of std errors
func TestStatusOfHttpErrors(t *testing.T) {
	if Http400().Status != 400 {
		t.Error("Http 400 status is not 400")
	}
	if Http404().Status != 404 {
		t.Error("Http 404 status is not 404")
	}
	if Http405().Status != 405 {
		t.Error("Http 405 status is not 405")
	}
	if Http500().Status != 500 {
		t.Error("Http 500 status is not 500")
	}
}

// Test base controller
func TestController(t *testing.T) {

	for _, method := range []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "TRACE", "get", "post"} {
		ctx := MockRequest(method, "/base")
		if ctx.Response.StatusCode() != fasthttp.StatusMethodNotAllowed {
			t.Errorf("Status is %d instead of 405", ctx.Response.StatusCode())
		}
		if string(ctx.Response.Body()) != fasthttp.StatusMessage(fasthttp.StatusMethodNotAllowed) {
			t.Error("Body wasn't the expected.")
		}
	}
}

// Test a controller implementation
func TestDummyController(t *testing.T) {
	ctx := MockRequest("GET", "/ass")
	if ctx.Response.StatusCode() != 200 {
		t.Errorf("Status is %d instead of 200", ctx.Response.StatusCode())
	}
	if string(ctx.Response.Body()) != "<h1>I'm a dummy and my name is ass</h1>" {
		t.Error("Body wasn't the expected.")
	}
	if value := ctx.Response.Header.Peek("x-dummy"); len(value) == 0 {
		t.Error("Response don't have Content-type header")
	} else if string(value) != "dummy" {
		t.Errorf("Content-type header is not dummy. Is %s instead.", value)
	}
}

// Test a controller implementation
func TestDummyControllerWithError(t *testing.T) {
	ctx := MockRequest("POST", "/ass")
	if ctx.Response.StatusCode() != 500 {
		t.Errorf("Status is %d instead of 500", ctx.Response.StatusCode())
	}
	if string(ctx.Response.Body()) != fasthttp.StatusMessage(fasthttp.StatusInternalServerError) {
		t.Error("Body wasn't the expected.")
	}
	if value := ctx.Response.Header.Peek("x-dummy"); len(value) == 0 {
		t.Error("Response don't have Content-type header")
	} else if string(value) != "dummy" {
		t.Errorf("Content-type header is not dummy. Is %s instead.", value)
	}
}