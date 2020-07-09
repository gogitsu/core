package core

import (
	"net/http"
	"testing"
)

func TestRouter(t *testing.T) {
	r := NewGorillaRouter()
	g := r.WithRoot("/accounts")
	g.Get("", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	r.Walk(LogRoute)
}
