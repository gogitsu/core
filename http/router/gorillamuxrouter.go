package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

// GorillaMuxRouter implements Router using Gorilla mux.
type GorillaMuxRouter struct {
	basePath string
	mux      *mux.Router
}

// BasePath .
func (gm *GorillaMuxRouter) BasePath(bp string) Router {
	return nil
}

// Handle .
func (gm *GorillaMuxRouter) Handle(method string, path string, handler http.Handler) {}

// HandleFunc .
func (gm *GorillaMuxRouter) HandleFunc(method string, path string, handlerFunc http.HandlerFunc) {}

// Get .
func (gm *GorillaMuxRouter) Get(path string, handler http.Handler) {}

// Post .
func (gm *GorillaMuxRouter) Post(path string, handler http.Handler) {}

// Put .
func (gm *GorillaMuxRouter) Put(path string, handler http.Handler) {}

// Patch .
func (gm *GorillaMuxRouter) Patch(path string, handler http.Handler) {}

// Delete .
func (gm *GorillaMuxRouter) Delete(path string, handler http.Handler) {}

// ServeHTTP .
func (gm *GorillaMuxRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
