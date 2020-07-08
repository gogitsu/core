package core

import (
	"net/http"
	"testing"
)

func TestRouter(t *testing.T) {
	r := NewRouter()
	r.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("GET")

	g := r.WithRoot("/accounts")
	g.Get("", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	r.Walk(LogRoute)
}
