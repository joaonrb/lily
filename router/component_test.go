package router

import (
	"github.com/joaonrb/lily"
	"testing"
	"runtime"
	"log"
	"time"
)

var (
	simpleURLPathSamples = []string{
		"/",
		"/foe",
		"/foe",
		"/foe/john-doe",
		"/foe/john-doe/",
		"/john-doe",
		"/john-doe/",
		"/john-doe/foe",
		"/john-doe/foe/",
	}

)

type mockComponent struct {
	lily.Component
	result string
}

func (mock *mockComponent) Resolve(context interface{}) interface{} {
	return mock.result
}

func logMemory()  {
	go func() {
		for {
			var m runtime.MemStats
			runtime.ReadMemStats(&m)
			log.Printf("\nAlloc = %vKB\nTotalAlloc = %vKB\nSys = %v\nNumGC = %v\n\n", m.Alloc / 1024, m.TotalAlloc / 1024, m.Sys / 1024, m.NumGC)
			time.Sleep(100*time.Millisecond)
		}
	}()
	time.Sleep(100*time.Millisecond)
}

// TestRouterSimpleURLPath test simple path
func TestRouterSimpleURLPath(t *testing.T) {
	router := New()
	for _, path := range simpleURLPathSamples {
		router.Add([]byte(path), &mockComponent{result: path})
	}
	for _, path := range simpleURLPathSamples {
		result := router.Resolve([]byte(path))
		if result.(lily.Component).Resolve(nil) != path {
			t.Errorf("Router didn't return the expected path: path(%s)" +
			" is not result(%s)", path,
			result.(lily.Component).Resolve(nil))
		}
	}
}

// TestRouterSimpleURLPath test simple path
func TestRouterRegexURLPath(t *testing.T) {
	router := New()
	for _, path := range regexURLPathSamples {
		router.Add([]byte(path), &mockComponent{result: path})
	}
	for _, path := range regexURLPathSamples {
		result := router.Resolve([]byte(path))
		if result.(lily.Component).Resolve(nil) != path {
			t.Errorf("Router didn't return the expected path: path(%s)" +
				" is not result(%s)", path,
				result.(lily.Component).Resolve(nil))
		}
	}
}
