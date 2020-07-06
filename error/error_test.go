package error

import (
	"errors"
	"net/http"
	"testing"
)

func TestError(t *testing.T) {
	err := NewHTTPError(
		errors.New("Test error"),
		http.StatusInternalServerError,
		"Internal Error",
		"fake-ID",
	)

	if _, ok := err.(Error); !ok {
		t.Error("err is not a Error implementation")
	}

	httpErr, ok := err.(*HTTPError)
	if !ok {
		t.Error("err is not a *HTTPError")
	}

	if httpErr == nil {
		t.Error("err is nil")
	} else if httpErr.Detail != "Internal Error" {
		t.Error("detail != 'Internal Error'")
	} else if httpErr.Code != http.StatusInternalServerError {
		t.Errorf("Code != %d\n", http.StatusInternalServerError)
	}
}

func TestBadRequest(t *testing.T) {
	err := BadRequest(
		errors.New("errTestBadRequest"),
		"TestBadRequest",
		"TestBadRequest-ID",
	)
	if _, ok := err.(Error); !ok {
		t.Error("err is not a Error implementation")
	}

	httpErr, ok := err.(*HTTPError)
	if !ok {
		t.Error("err is not a *HTTPError")
	}
	if httpErr == nil {
		t.Error("err is nil")
	} else if httpErr.Detail != "TestBadRequest" {
		t.Error("detail != 'TestBadRequest'")
	} else if httpErr.Code != http.StatusBadRequest {
		t.Errorf("Code != %d\n", http.StatusInternalServerError)
	}
}
