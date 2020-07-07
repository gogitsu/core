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
		BasePath(bp string) Router
		Handle(method string, path string, handler http.Handler)
		HandleFunc(method string, path string, handlerFunc http.HandlerFunc)
		Get(path string, handler http.Handler)
		Post(path string, handler http.Handler)
		Put(path string, handler http.Handler)
		Patch(path string, handler http.Handler)
		Delete(path string, handler http.Handler)
		ServeHTTP(w http.ResponseWriter, r *http.Request)
	}
)

// NewRouter is the factory function to instantiate a new Router
// according to the input router type.
// Default is the Gorilla mux instance.
func NewRouter(rt string) Router {
	switch rt {
	case Gorilla:
		return &GorillaMuxRouter{}
	default:
		return &GorillaMuxRouter{}
	}
}
