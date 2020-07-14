// Copyright 2020 Luca Stasio. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package core implements core components of gogitsu lib.
//
// router.go implements Router interface and concrete implementations.
// Actually only Gorilla based concrete Router is implemented.
package core

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	// Gorilla is the string identifier for Gorilla Mux implemenetation.
	Gorilla string = "gorilla"

	// Chi is the string identifier for Chi mux implementation.
	// Chi string = "chi"
)

type (
	// WalkFn is the type for function to be called from the Walk function.
	// Heach concrete router implementation will cast the input interfcae{}
	// to thw right input param type (a.e. for Gorilla mux will be a Route).
	WalkFn func(interface{})

	// Router defines interface to work with concrete routers.
	Router interface {
		Mux() http.Handler
		WithRoot(path string) Router
		WithControllers(ctrls ...Controller) Router
		Walk(wfn WalkFn)
		HandleFunc(path string, fn http.HandlerFunc) interface{}
		Get(path string, fn http.HandlerFunc) Router
		Post(path string, fn http.HandlerFunc) Router
		Put(path string, fn http.HandlerFunc) Router
		Patch(path string, fn http.HandlerFunc) Router
		Delete(path string, fn http.HandlerFunc) Router
		PrintRoute(interface{})
		LogRoute(interface{})
	}

	// GorillaRouter is the main routing structure.
	// It holds a Gorilla mux to be used for routing configuration.
	GorillaRouter struct {
		mux *mux.Router
	}
)

// NewRouter is the factory function to instantiate a new Router
// according to the input router type.
// Default is the Gorilla mux instance.
//
// Actually only Gorilla based concrete Router.
func NewRouter(rt string) Router {
	switch rt {
	case Gorilla:
		return NewGorillaRouter()
	default:
		return NewGorillaRouter()
	}
}

// NewRouterWithRootPath is the factory function to instantiate a new Router
// according to the input router type and root path.
// Default is the Gorilla mux instance.
func NewRouterWithRootPath(rt string, path string) Router {
	switch rt {
	case Gorilla:
		return NewGorillaRouterWithRoot(path)
	default:
		return NewGorillaRouterWithRoot(path)
	}
}

// NewGorillaRouter return a new GorillaRouter instance.
func NewGorillaRouter() *GorillaRouter {
	return &GorillaRouter{mux.NewRouter()}
}

// NewGorillaRouterWithRoot return a new GorillaRouter with root path instance.
func NewGorillaRouterWithRoot(path string) *GorillaRouter {
	return NewGorillaRouter().WithRoot(path).(*GorillaRouter)
}

// Mux return the http.Handler implementation.
func (r *GorillaRouter) Mux() http.Handler {
	return r.mux
}

// WithRoot .
// Here the Subrouter is the root for Groups defined by the controller paths.
func (r *GorillaRouter) WithRoot(path string) Router {
	return &GorillaRouter{r.mux.PathPrefix(path).Subrouter().StrictSlash(true)}
}

// WithControllers .
// Here each controller's base path will define the relative Group into the root path.
func (r *GorillaRouter) WithControllers(ctrls ...Controller) Router {
	for _, c := range ctrls {
		c.Route(r.WithRoot(c.BasePath()))
	}
	return r
}

// Walk walks on routes and execute the input callback.
func (r *GorillaRouter) Walk(wfn WalkFn) {
	r.mux.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		if route.GetHandler() != nil {
			wfn(route)
		}
		return nil
	})
}

// HandleFunc .
func (r *GorillaRouter) HandleFunc(path string, fn http.HandlerFunc) interface{} {
	return r.mux.HandleFunc(path, fn)
}

// Get .
func (r *GorillaRouter) Get(path string, fn http.HandlerFunc) Router {
	r.HandleFunc(path, fn).(*mux.Route).Methods("GET")
	return r
}

// Post .
func (r *GorillaRouter) Post(path string, fn http.HandlerFunc) Router {
	r.HandleFunc(path, fn).(*mux.Route).Methods("POST")
	return r
}

// Put .
func (r *GorillaRouter) Put(path string, fn http.HandlerFunc) Router {
	r.HandleFunc(path, fn).(*mux.Route).Methods("PUT")
	return r
}

// Patch .
func (r *GorillaRouter) Patch(path string, fn http.HandlerFunc) Router {
	r.HandleFunc(path, fn).(*mux.Route).Methods("PATCH")
	return r
}

// Delete .
func (r *GorillaRouter) Delete(path string, fn http.HandlerFunc) Router {
	r.HandleFunc(path, fn).(*mux.Route).Methods("DELETE")
	return r
}

// PrintRoute .
// func PrintRoute(route *mux.Route) {
func (r *GorillaRouter) PrintRoute(i interface{}) {
	route := i.(*mux.Route)
	tmpl, _ := route.GetPathTemplate()
	methods, _ := route.GetMethods()
	for _, v := range methods {
		fmt.Printf("route: %s %s initialized", v, tmpl)
	}

}

// LogRoute .
// func LogRoute(route *mux.Route) {
func (r *GorillaRouter) LogRoute(i interface{}) {
	route := i.(*mux.Route)
	tmpl, _ := route.GetPathTemplate()
	methods, _ := route.GetMethods()
	for _, v := range methods {
		log.Printf("route: %s %s initialized", v, tmpl)
	}
}
