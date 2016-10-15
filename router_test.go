package lily

// Author João Nuno.
//
// joaonrb@gmail.com
//

import (
	"reflect"
	"testing"
)

// TestRouterRoutePath test simple route
func TestRouterRoutePath(t *testing.T) {
	controller, _ := getController([]byte("/"))
	if reflect.TypeOf(controller) != reflect.TypeOf(&DummyController{}) {
		t.Error("Contoller is not dummy")
	}
}

// TestRouterRouteRegexPath test route with parameters
func TestRouterRouteRegexPath(t *testing.T) {
	controller, args := getController([]byte("/ass/"))
	if reflect.TypeOf(controller) != reflect.TypeOf(&DummyController{}) {
		t.Error("Contoller is not dummy")
	}
	if name, ok := args["name"]; !ok {
		t.Error("Name not in arguments")
	} else if name != "ass" {
		t.Errorf("Name is not ass. Is %s instead", name)
	}
}

// TestRouterUrlError test route with parameters
func TestRouterUrlError(t *testing.T) {
	err := Url("/:(?P<name>^\\[a-zA-Z0-9]+)$", &BaseController{})
	if err != nil {
		t.Error("Error is expected in malformed url")
	}
}