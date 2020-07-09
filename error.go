package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Error is an error whose details to be shared with client.
// Check also this https://github.com/google/jsonapi/blob/master/errors.go
type Error interface {
	// Error method to make Error implement error interface.
	Error() string
	// ResponseBody returns response body.
	ResponseBody() ([]byte, error)
	// ResponseHeaders returns http status code and headers.
	ResponseHeaders() (int, map[string]string)
}

// HTTPError implements Error interface.
type HTTPError struct {
	Code      int         `json:"code"`
	Err       string      `json:"error"`
	Detail    interface{} `json:"detail"`
	RequestID string      `json:"requestId"`
	Timestamp time.Time   `json:"timestamp"`
}

// Error return a formatted description of the error.
func (e *HTTPError) Error() string {
	return fmt.Sprintf("%v : %s", e.Detail, e.Err)
}

// ResponseBody returns JSON response body.
func (e *HTTPError) ResponseBody() ([]byte, error) {
	body, err := json.Marshal(e)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing response body: %v", err)
	}
	return body, nil
}

// ResponseHeaders returns http status code and headers.
func (e *HTTPError) ResponseHeaders() (int, map[string]string) {
	return e.Code, map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	}
}

// NewHTTPError returns a new HTTPError instance
func NewHTTPError(err error, status int, detail interface{}, requestID string) error {

	return &HTTPError{
		Err:       err.Error(),
		Detail:    detail,
		Code:      status,
		RequestID: requestID,
		Timestamp: time.Now(),
	}
}

// RenderError exported to be used in this lib (a.e. in middlewares) returns error to the client.
func RenderError(w http.ResponseWriter, err error) {
	clientError, ok := err.(Error)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, err := clientError.ResponseBody()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	status, headers := clientError.ResponseHeaders()
	for k, v := range headers {
		w.Header().Set(k, v)
	}
	w.WriteHeader(status)
	w.Write(body)
}

// Errors

// BadRequest http error.
func BadRequest(err error, detail interface{}, requestID string) error {
	return NewHTTPError(err, http.StatusBadRequest, detail, requestID)
}

// Unauthorized http error.
func Unauthorized(err error, detail interface{}, requestID string) error {
	return NewHTTPError(err, http.StatusUnauthorized, detail, requestID)
}

// Forbidden http error.
func Forbidden(err error, detail interface{}, requestID string) error {
	return NewHTTPError(err, http.StatusForbidden, detail, requestID)
}

// NotFound http error.
func NotFound(err error, detail interface{}, requestID string) error {
	return NewHTTPError(err, http.StatusNotFound, detail, requestID)
}

// NoContent http error.
func NoContent(err error, detail interface{}, requestID string) error {
	return NewHTTPError(err, http.StatusNoContent, detail, requestID)
}

// MethodNotAllowed http error.
func MethodNotAllowed(err error, detail interface{}, requestID string) error {
	return NewHTTPError(err, http.StatusMethodNotAllowed, detail, requestID)
}

// NotAcceptable http error.
func NotAcceptable(err error, detail interface{}, requestID string) error {
	return NewHTTPError(err, http.StatusNotAcceptable, detail, requestID)
}

// RequestTimeout http error.
func RequestTimeout(err error, detail interface{}, requestID string) error {
	return NewHTTPError(err, http.StatusRequestTimeout, detail, requestID)
}

// Conflict http error.
func Conflict(err error, detail interface{}, requestID string) error {
	return NewHTTPError(err, http.StatusConflict, detail, requestID)
}

// UnprocessableEntity http error.
func UnprocessableEntity(err error, detail interface{}, requestID string) error {
	return NewHTTPError(err, http.StatusUnprocessableEntity, detail, requestID)
}

// InternalServerError http error.
func InternalServerError(err error, detail interface{}, requestID string) error {
	return NewHTTPError(err, http.StatusInternalServerError, detail, requestID)
}

// NotImplemented http error.
func NotImplemented(err error, detail interface{}, requestID string) error {
	return NewHTTPError(err, http.StatusNotImplemented, detail, requestID)
}

// ServiceUnavailable http error.
func ServiceUnavailable(err error, detail interface{}, requestID string) error {
	return NewHTTPError(err, http.StatusServiceUnavailable, detail, requestID)
}

// BadGateway http error.
func BadGateway(err error, detail interface{}, requestID string) error {
	return NewHTTPError(err, http.StatusBadGateway, detail, requestID)
}

// GatewayTimeout http error.
func GatewayTimeout(err error, detail interface{}, requestID string) error {
	return NewHTTPError(err, http.StatusGatewayTimeout, detail, requestID)
}

// NetworkAuthenticationRequired http error.
func NetworkAuthenticationRequired(err error, detail interface{}, requestID string) error {
	return NewHTTPError(err, http.StatusNetworkAuthenticationRequired, detail, requestID)
}
