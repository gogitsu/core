package core

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// RouteCallback is a function on *mux.Route.
type RouteCallback func(*mux.Route)

// Router is the main routing structure.
// It holds a Gorilla mux to be used for routing configuration.
type Router struct {
	// rootPath string
	Mux *mux.Router
}

// NewRouter return a new Router instance.
func NewRouter() *Router {
	return &Router{Mux: mux.NewRouter()}
}

// WithRoot .
func (r *Router) WithRoot(path string) *Router {
	r.Mux = r.Mux.PathPrefix(path).Subrouter().StrictSlash(true)
	return r
}

// WithControllers .
func (r *Router) WithControllers(ctrls ...Controller) *Router {
	for _, c := range ctrls {
		c.Route(r.Group(c.BasePath()))
	}
	return r
}

// Walk walks on routes and execute callback.
func (r *Router) Walk(rcb RouteCallback) {
	r.Mux.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		if route.GetHandler() != nil {
			rcb(route)
		}
		return nil
	})
}

// Group .
func (r *Router) Group(path string) *Router {
	return &Router{Mux: r.Mux.PathPrefix(path).Subrouter()}
}

// HandleFunc .
func (r *Router) HandleFunc(path string, fn http.HandlerFunc) *mux.Route {
	return r.Mux.HandleFunc(path, fn)
}

// Get .
func (r *Router) Get(path string, fn http.HandlerFunc) *Router {
	r.HandleFunc(path, fn).Methods("GET")
	return r
}

// Post .
func (r *Router) Post(path string, fn http.HandlerFunc) *Router {
	r.HandleFunc(path, fn).Methods("POST")
	return r
}

// PrintRoute .
func PrintRoute(route *mux.Route) {
	tmpl, _ := route.GetPathTemplate()
	methods, _ := route.GetMethods()
	for _, v := range methods {
		fmt.Printf("route: %s %s initialized", v, tmpl)
	}

}

// LogRoute .
func LogRoute(route *mux.Route) {
	tmpl, _ := route.GetPathTemplate()
	methods, _ := route.GetMethods()
	for _, v := range methods {
		log.Printf("route: %s %s initialized", v, tmpl)
	}
}
