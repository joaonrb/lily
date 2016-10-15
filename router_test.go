package lily

// Author João Nuno.
//
// joaonrb@gmail.com
//

import (
	"reflect"
	"testing"
)

// Test simple route
func TestRouterRoutePath(t *testing.T) {
	controller, _ := getController([]byte("/"))
	if reflect.TypeOf(controller) != reflect.TypeOf(&DummyController{}) {
		t.Error("Contoller is not dummy")
	}
}

// Test route with parameters
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
