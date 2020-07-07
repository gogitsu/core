package router

import (
	"net/http"
)

const (
	// Gorilla is the string identifier for Gorilla Mux implemenetation.
	Gorilla string = "gorilla"

	// Chi is the string identifier for Chi mux implementation.
	Chi string = "chi"
)

type (
	// Router .
	Router interface {
		RootPath(path string) Router
		Handle(method string, path string, handler http.Handler) Router
		HandleFunc(method string, path string, handlerFunc http.HandlerFunc) Router
		Get(path string, handlerFunc http.HandlerFunc) Router
		Post(path string, handler http.Handler) Router
		Put(path string, handler http.Handler) Router
		Patch(path string, handler http.Handler) Router
		Delete(path string, handler http.Handler) Router
		ServeHTTP(w http.ResponseWriter, r *http.Request)
	}
)

// NewRouter is the factory function to instantiate a new Router
// according to the input router type.
// Default is the Gorilla mux instance.
func NewRouter(rt string) Router {
	switch rt {
	case Gorilla:
		return NewGorillaMuxRouter()
	default:
		return NewGorillaMuxRouter()
	}
}

// NewRouterWithRootPath is the factory function to instantiate a new Router
// according to the input router type and root path.
// Default is the Gorilla mux instance.
func NewRouterWithRootPath(rt string, path string) Router {
	switch rt {
	case Gorilla:
		return NewGorillaMuxRouterWithRootPath(path)
	default:
		return NewGorillaMuxRouterWithRootPath(path)
	}
}
