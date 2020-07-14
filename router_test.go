package core

import (
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

const (
	expectedHTTPStatus = http.StatusOK
	expectedHTTPBody   = "OK"
)

func TestRouter(t *testing.T) {
	r := NewRouter(Gorilla).WithRoot("/api/v1")
	g := r.WithRoot("/accounts")
	g.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	g2 := r.WithRoot("/posts")
	g2.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	r.Walk(r.LogRoute)

	srv := &http.Server{
		Handler:      r.Mux(),
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		t.Fatal(srv.ListenAndServe())
	}()

	response, err := http.Get("http://127.0.0.1:8000/api/v1/accounts")
	if err != nil {
		t.Error(err)
		return
	}
	defer response.Body.Close()
	if response.StatusCode != expectedHTTPStatus {
		t.Errorf("expected '%d' response code, got %d", expectedHTTPStatus, response.StatusCode)
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
		return
	}

	bodyString := string(body)
	if bodyString != expectedHTTPBody {
		t.Errorf("expected '%s' response body, got %s", expectedHTTPBody, bodyString)
		return
	}
}
