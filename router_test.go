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

	p := r.Mux.PathPrefix("/products").Subrouter()

	p.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("GET")
	p.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("GET")
	// r.Walk(LogRoute)

	g := r.Group("/accounts")
	g.Get("", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	r.Walk(LogRoute)
}
