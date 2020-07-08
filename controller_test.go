package core

import (
	"net/http"
	"testing"
)

type MyController struct {
	BaseController
}

func NewMyController(path string) Controller {
	return &MyController{BaseController{basePath: path}}
}

func (mc *MyController) Route(r *Router) {
	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}

type FooController struct {
	BaseController
}

func NewFooController() Controller {
	return &MyController{BaseController{basePath: "/foo"}}
}

func (mc *FooController) Route(r *Router) {
	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}

type SimpleController struct {
	BaseController
}

func NewSimpleController(path string) Controller {
	return &SimpleController{BaseController{basePath: path}}
}

func (sc *SimpleController) Route(r *Router) {
	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}
func TestController(t *testing.T) {
	r := NewRouter().WithRoot("/api/v1")
	c := NewMyController("/todos")

	r.WithControllers(c)
	r.Walk(LogRoute)

	r2 := NewRouter()
	r2.WithControllers(NewFooController())
	r2.WithRoot("/api/v2").
		WithControllers(
			NewMyController("/todos"),
			NewSimpleController("/notes"),
		)
	r2.WithRoot("/api/v3").WithControllers(NewFooController())
	r2.Walk(LogRoute)
}
