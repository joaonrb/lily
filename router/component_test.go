package router

import (
	"github.com/joaonrb/lily"
	"testing"
)

var simpleURLPathSamples = []string{
	"/",
	"/foe",
	"/foe",
	"/foe/john-doe",
	"/foe/john-doe/",
}

type mockComponent struct {
	lily.Component
	result string
}

func (mock *mockComponent) Resolve(context interface{}) interface{} {
	return mock.result
}

// TestRouterSimpleURLPath test simple path
func TestRouterSimpleURLPath(t *testing.T) {
	router := New()
	for _, path := range simpleURLPathSamples {
		router.Add([]byte(path), &mockComponent{result: path})
	}
	for _, path := range simpleURLPathSamples {
		result := router.Resolve([]byte(path))
		if result.(mockComponent).result != path {
			t.Errorf("Router didn't return the expected path: path(%s)" +
			" is not result(%s)", path, result.(mockComponent).result)
		}
	}
}
