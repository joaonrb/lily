package old

// Author João Nuno.
//
// joaonrb@gmail.com
//
import (
	"github.com/valyala/fasthttp"
	"testing"
	"time"
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
		request := &fasthttp.Request{}
		request.SetRequestURI("http://127.0.0.1:3333/base")
		request.Header.SetMethod(method)
		response := &fasthttp.Response{}
		err := fasthttp.DoTimeout(request, response, 10*time.Second)

		//ctx := MockRequest(method, "/base")
		if err != nil {
			t.Fatalf("Failed with error %s", err.Error())
		}
		if response.StatusCode() != fasthttp.StatusMethodNotAllowed {
			t.Errorf("Status is %d instead of 405", response.StatusCode())
		}

		// HEAD method doesn't have body
		if method == "HEAD" {
			continue
		}

		body := string(response.Body())
		if body != fasthttp.StatusMessage(fasthttp.StatusMethodNotAllowed) {
			t.Errorf("Body wasn't the expected with method %s. Got %s", method, body)
		}
	}
}

// Test a controller implementation
func TestDummyController(t *testing.T) {
	request := &fasthttp.Request{}
	request.SetRequestURI("http://127.0.0.1:3333/ass?test_param=t")
	request.Header.SetMethod("GET")
	response := &fasthttp.Response{}
	err := fasthttp.DoTimeout(request, response, 10*time.Second)

	//response, err := http.Get("http://127.0.0.1:3333/ass")
	if err != nil {
		t.Fatalf("Failed with error %s", err.Error())
	}
	if response.StatusCode() != 200 {
		t.Errorf("Status is %d instead of 200", response.StatusCode())
	}
	body := response.Body()
	if string(body) != "<h1>I'm a dummy and my name is ass</h1>" {
		t.Errorf("Body wasn't the expected. Got %s", string(body))
	}
	if value := string(response.Header.Peek("x-dummy")); len(value) == 0 {
		t.Error("Response don't have Content-type header")
	} else if value != "dummy" {
		t.Errorf("x-ummy header is not dummy. Is %s instead.", value)
	}
}

// Test a controller implementation
func TestDummyController404(t *testing.T) {
	request := &fasthttp.Request{}
	request.SetRequestURI("http://127.0.0.1:3333/1.234")
	request.Header.SetMethod("GET")
	response := &fasthttp.Response{}
	err := fasthttp.DoTimeout(request, response, 10*time.Second)

	//response, err := http.Get("http://127.0.0.1:3333/1.234")
	if err != nil {
		t.Errorf("Failed with error %s", err.Error())
	}
	if response.StatusCode() != 404 {
		t.Errorf("Status is %d instead of 404", response.StatusCode())
	}
	body := string(response.Body())
	if string(body) != fasthttp.StatusMessage(fasthttp.StatusNotFound) {
		t.Errorf("Body wasn't the expected. Got %s", string(body))
	}
}

// Test a controller implementation
func TestDummyControllerWithError(t *testing.T) {
	request := &fasthttp.Request{}
	request.SetRequestURI("http://127.0.0.1:3333/ass")
	request.Header.SetMethod("POST")
	response := &fasthttp.Response{}
	err := fasthttp.DoTimeout(request, response, 10*time.Second)

	//ctx := MockRequest("POST", "/ass")
	if err != nil {
		t.Errorf("Failed with error %s", err.Error())
	}
	if response.StatusCode() != 500 {
		t.Errorf("Status is %d instead of 500", response.StatusCode())
	}
	if string(response.Body()) != fasthttp.StatusMessage(fasthttp.StatusInternalServerError) {
		t.Errorf("Body wasn't the expected. Got %s", string(response.Body()))
	}
	if value := response.Header.Peek("x-dummy"); len(value) == 0 {
		t.Error("Response don't have Content-type header")
	} else if string(value) != "dummy" {
		t.Errorf("Content-type header is not dummy. Is %s instead.", value)
	}
}

// Test a controller implementation
func TestDummyRegexControllerWithError(t *testing.T) {
	request := &fasthttp.Request{}
	request.SetRequestURI("http://127.0.0.1:3333/ass/hdj")
	request.Header.SetMethod("GET")
	response := &fasthttp.Response{}
	err := fasthttp.DoTimeout(request, response, 10*time.Second)

	//response, err := http.Get("http://127.0.0.1:3333/ass/hdj")
	if err != nil {
		t.Errorf("Failed with error %s", err.Error())
	}
	if response.StatusCode() != 404 {
		t.Errorf("Status is %d instead of 404", response.StatusCode())
	}
	body := string(response.Body())
	if err != nil {
		t.Errorf("Failed with error %s", err.Error())
	}
	if string(body) != fasthttp.StatusMessage(fasthttp.StatusNotFound) {
		t.Errorf("Body wasn't the expected. Got %s", string(body))
	}
}

// Test a controller implementation
func TestDummyControllerFiltered(t *testing.T) {
	request := &fasthttp.Request{}
	request.SetRequestURI("http://127.0.0.1:3333/ass")
	request.Header.SetMethod("PUT")
	response := &fasthttp.Response{}
	err := fasthttp.DoTimeout(request, response, 10*time.Second)

	//ctx := MockRequest("PUT", "/ass")
	if err != nil {
		t.Errorf("Failed with error %s", err.Error())
	}
	if response.StatusCode() != 403 {
		t.Errorf("Status is %d instead of 503", response.StatusCode())
	}
	if string(response.Body()) != "You cannot be here." {
		t.Errorf("Body wasn't the expected. Got %s", string(response.Body()))
	}
	if value := response.Header.Peek("x-dummy"); len(value) == 0 {
		t.Error("Response don't have Content-type header")
	} else if string(value) != "dummy" {
		t.Errorf("Content-type header is not dummy. Is %s instead.", value)
	}
}
