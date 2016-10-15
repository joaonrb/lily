package lily

// Author João Nuno.
//
// joaonrb@gmail.com
//
import (
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"net/http"
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
	for _, method := range []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "TRACE"} {
		ctx := MockRequest(method, "/base")
		if ctx.Response.StatusCode() != fasthttp.StatusMethodNotAllowed {
			t.Errorf("Status is %d instead of 405", ctx.Response.StatusCode())
		}
		if string(ctx.Response.Body()) != fasthttp.StatusMessage(fasthttp.StatusMethodNotAllowed) {
			t.Errorf("Body wasn't the expected. Got %s", string(ctx.Response.Body()))
		}
	}
}

// Test a controller implementation
func TestDummyController(t *testing.T) {
	response, err := http.Get("http://127.0.0.1:3333/ass")
	if err != nil {
		t.Errorf("Failed with error %s", err.Error())
	}
	if response.StatusCode != 200 {
		t.Errorf("Status is %d instead of 200", response.StatusCode)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Failed with error %s", err.Error())
	}
	if string(body) != "<h1>I'm a dummy and my name is ass</h1>" {
		t.Errorf("Body wasn't the expected. Got %s", string(body))
	}
	if value := response.Header.Get("x-dummy"); len(value) == 0 {
		t.Error("Response don't have Content-type header")
	} else if value != "dummy" {
		t.Errorf("x-ummy header is not dummy. Is %s instead.", value)
	}
}

// Test a controller implementation
func TestDummyController404(t *testing.T) {
	response, err := http.Get("http://127.0.0.1:3333/1.234")
	if err != nil {
		t.Errorf("Failed with error %s", err.Error())
	}
	if response.StatusCode != 404 {
		t.Errorf("Status is %d instead of 404", response.StatusCode)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Failed with error %s", err.Error())
	}
	if string(body) != fasthttp.StatusMessage(fasthttp.StatusNotFound) {
		t.Errorf("Body wasn't the expected. Got %s", string(body))
	}
}

// Test a controller implementation
func TestDummyControllerWithError(t *testing.T) {
	ctx := MockRequest("POST", "/ass")
	if ctx.Response.StatusCode() != 500 {
		t.Errorf("Status is %d instead of 500", ctx.Response.StatusCode())
	}
	if string(ctx.Response.Body()) != fasthttp.StatusMessage(fasthttp.StatusInternalServerError) {
		t.Errorf("Body wasn't the expected. Got %s", string(ctx.Response.Body()))
	}
	if value := ctx.Response.Header.Peek("x-dummy"); len(value) == 0 {
		t.Error("Response don't have Content-type header")
	} else if string(value) != "dummy" {
		t.Errorf("Content-type header is not dummy. Is %s instead.", value)
	}
}

// Test a controller implementation
func TestDummyRegexControllerWithError(t *testing.T) {
	response, err := http.Get("http://127.0.0.1:3333/ass/hdj")
	if err != nil {
		t.Errorf("Failed with error %s", err.Error())
	}
	if response.StatusCode != 404 {
		t.Errorf("Status is %d instead of 404", response.StatusCode)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Failed with error %s", err.Error())
	}
	if string(body) != fasthttp.StatusMessage(fasthttp.StatusNotFound) {
		t.Errorf("Body wasn't the expected. Got %s", string(body))
	}
}

// Test a controller implementation
func TestDummyControllerFiltered(t *testing.T) {
	ctx := MockRequest("PUT", "/ass")
	if ctx.Response.StatusCode() != 403 {
		t.Errorf("Status is %d instead of 503", ctx.Response.StatusCode())
	}
	if string(ctx.Response.Body()) != "You cannot be here." {
		t.Errorf("Body wasn't the expected. Got %s", string(ctx.Response.Body()))
	}
	if value := ctx.Response.Header.Peek("x-dummy"); len(value) == 0 {
		t.Error("Response don't have Content-type header")
	} else if string(value) != "dummy" {
		t.Errorf("Content-type header is not dummy. Is %s instead.", value)
	}
}
