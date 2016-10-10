package lily

import (
	"reflect"
	"testing"
)

//
// Author João Nuno.
//
// joaonrb@gmail.com
//

func TestRouterRoutePath(t *testing.T) {
	controller, _ := getController([]byte("/"))
	if reflect.TypeOf(controller) != reflect.TypeOf(&DummyController{}) {
		t.Error("Contoller is not dummy")
	}
}

func TestRouterRouteRegexPath(t *testing.T) {
	controller, _ := getController([]byte("/ass"))
	if reflect.TypeOf(controller) != reflect.TypeOf(&DummyController{}) {
		t.Error("Contoller is not dummy and my name is ass")
	}
}
