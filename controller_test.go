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

func (mc *MyController) Route(r Router) {
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

func (mc *FooController) Route(r Router) {
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

func (sc *SimpleController) Route(r Router) {
	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}
func TestController(t *testing.T) {
	r := NewGorillaRouter().WithRoot("/api/v1")
	c := NewMyController("/todos")

	r.WithControllers(c)
	r.Walk(r.LogRoute)

	r2 := NewGorillaRouter()
	r2.WithControllers(NewFooController())
	r2.WithRoot("/api/v2").
		WithControllers(
			NewMyController("/todos"),
			NewSimpleController("/notes"),
		)
	r2.WithRoot("/api/v3").WithControllers(NewFooController())
	r2.Walk(r.LogRoute)
}

/*

package http

import (
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/gogitsu/core/http/router"
)

type ControllerOne struct {
}

func NewControllerOne() Controller {
	return &ControllerOne{}
}

func (co *ControllerOne) Route(r router.Router) Controller {
	r.Get("/one", co.GetOne)
	return co
}

func (co *ControllerOne) GetOne(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("one"))
}

func TestController(t *testing.T) {
	router := router.NewRouterWithRootPath(router.Gorilla, "/api/v1")
	NewControllerOne().Route(router)

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8002",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		defer srv.Close()
		t.Fatal(srv.ListenAndServe())
	}()

	response, err := http.Get("http://127.0.0.1:8002/api/v1/one")
	if err != nil {
		t.Error(err)
		return
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		t.Errorf("expected 200 response code, got %d", response.StatusCode)
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
		return
	}

	bodyString := string(body)
	if bodyString != "one" {
		t.Errorf("expected 'one' response body, got %s", bodyString)
		return
	}
}


*/
