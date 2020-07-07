package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

// GorillaMuxRouter implements Router using Gorilla mux.
type GorillaMuxRouter struct {
	mux *mux.Router
}

// NewGorillaMuxRouter returns a new Gorilla mux based Router instance.
func NewGorillaMuxRouter() *GorillaMuxRouter {
	return &GorillaMuxRouter{mux.NewRouter()}
}

// RootPath sets the base path prefix for this router.
func (gm *GorillaMuxRouter) RootPath(path string) Router {
	gm.mux.PathPrefix(path)
	return gm
}

// Handle .
func (gm *GorillaMuxRouter) Handle(method string, path string, handler http.Handler) Router {
	gm.mux.Handle(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	})).Methods(method)
	return gm
}

// HandleFunc .
func (gm *GorillaMuxRouter) HandleFunc(method string, path string, handlerFunc http.HandlerFunc) Router {
	return nil
}

// Get .
func (gm *GorillaMuxRouter) Get(path string, handler http.Handler) Router {
	gm.Handle("GET", path, handler)
	return gm
}

// Post .
func (gm *GorillaMuxRouter) Post(path string, handler http.Handler) Router {
	gm.Handle("POST", path, handler)
	return gm
}

// Put .
func (gm *GorillaMuxRouter) Put(path string, handler http.Handler) Router {
	gm.Handle("PUT", path, handler)
	return gm
}

// Patch .
func (gm *GorillaMuxRouter) Patch(path string, handler http.Handler) Router {
	gm.Handle("PATCH", path, handler)
	return gm
}

// Delete .
func (gm *GorillaMuxRouter) Delete(path string, handler http.Handler) Router {
	gm.Handle("DELETE", path, handler)
	return gm
}

// ServeHTTP .
func (gm *GorillaMuxRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	gm.mux.ServeHTTP(w, r)
}
