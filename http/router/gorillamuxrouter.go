package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

// GorillaMuxRouter implements Router using Gorilla mux.
type GorillaMuxRouter struct {
	root string
	mux  *mux.Router
}

// NewGorillaMuxRouter returns a new Gorilla mux based Router instance.
func NewGorillaMuxRouter() Router {
	return &GorillaMuxRouter{mux: mux.NewRouter(), root: ""}
}

// NewGorillaMuxRouterWithRootPath returns a new Gorilla mux based Router instance
// initialized with the root path.
func NewGorillaMuxRouterWithRootPath(path string) Router {
	return &GorillaMuxRouter{mux: mux.NewRouter(), root: path}
}

// RootPath sets the base path prefix for this router.
func (gm *GorillaMuxRouter) RootPath(path string) Router {
	gm.mux.PathPrefix(path)
	return gm
}

// Handle .
func (gm *GorillaMuxRouter) Handle(method string, path string, handler http.Handler) Router {
	gm.mux.Handle(gm.root+path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	})).Methods(method)
	return gm
}

// HandleFunc .
func (gm *GorillaMuxRouter) HandleFunc(method string, path string, handlerFunc http.HandlerFunc) Router {
	gm.Handle(method, path, http.HandlerFunc(handlerFunc))
	return nil
}

// Get .
func (gm *GorillaMuxRouter) Get(path string, handlerFunc http.HandlerFunc) Router {
	gm.HandleFunc("GET", path, handlerFunc)
	return gm
}

// Post .
func (gm *GorillaMuxRouter) Post(path string, handlerFunc http.HandlerFunc) Router {
	gm.HandleFunc("POST", path, handlerFunc)
	return gm
}

// Put .
func (gm *GorillaMuxRouter) Put(path string, handlerFunc http.HandlerFunc) Router {
	gm.HandleFunc("PUT", path, handlerFunc)
	return gm
}

// Patch .
func (gm *GorillaMuxRouter) Patch(path string, handlerFunc http.HandlerFunc) Router {
	gm.HandleFunc("PATCH", path, handlerFunc)
	return gm
}

// Delete .
func (gm *GorillaMuxRouter) Delete(path string, handlerFunc http.HandlerFunc) Router {
	gm.HandleFunc("DELETE", path, handlerFunc)
	return gm
}

// ServeHTTP .
func (gm *GorillaMuxRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	gm.mux.ServeHTTP(w, r)
}
