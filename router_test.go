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
	controller, _, err := mainRouter.Parse("/")
	if err != nil {
		t.Errorf(err.Error())
	}
	if reflect.TypeOf(controller) != reflect.TypeOf(&DummyController{}) {
		t.Errorf("Contoller is not dummy")
	}
}
