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
