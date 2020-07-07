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
	t.Log("\n\n******* TestController ******\n")
	r := NewRouter().WithRoot("/api/v1")
	c := NewMyController("/todos")
	t.Logf("MyController.basePath: %s", c.BasePath())

	// c.Route(r.Group(c.BasePath()))
	r.WithControllers(c)
	r.Walk(LogRoute)

	r2 := NewRouter().
		WithRoot("/api/v2").
		WithControllers(
			NewMyController("/todos"),
			NewSimpleController("/notes"),
		)
	r2.Walk(LogRoute)
}
