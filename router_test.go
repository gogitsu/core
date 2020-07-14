package core

import (
	"net/http"
	"testing"
)

func TestRouter(t *testing.T) {
	r := NewRouter(Gorilla)
	g := r.WithRoot("/accounts")
	g.Get("", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	r.Walk(r.LogRoute)
}

/*

package router

import (
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestRouter(t *testing.T) {
	mux := NewRouter(Gorilla)
	mux.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("TEST-RESPONSE"))
	})

	srv := &http.Server{
		Handler:      mux,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		t.Fatal(srv.ListenAndServe())
	}()

	response, err := http.Get("http://127.0.0.1:8000/test")
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
	if bodyString != "TEST-RESPONSE" {
		t.Errorf("expected 'TEST-RESPONSE' response body, got %s", bodyString)
		return
	}
}

func TestRouterWithRootPath(t *testing.T) {
	mux := NewRouterWithRootPath(Gorilla, "/api")
	mux.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("TEST-RESPONSE from /api"))
	})

	srv := &http.Server{
		Handler:      mux,
		Addr:         "127.0.0.1:8001",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		defer srv.Close()
		t.Fatal(srv.ListenAndServe())
	}()

	response, err := http.Get("http://127.0.0.1:8001/api/test")
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
	if bodyString != "TEST-RESPONSE from /api" {
		t.Errorf("expected 'TEST-RESPONSE from /api' response body, got %s", bodyString)
		return
	}
}


*/
