package error

// Error is an error whose details to be shared with client.
type Error interface {
	// Error method to make Error implement error interface.
	Error() string
	// ResponseBody returns response body.
	ResponseBody() ([]byte, error)
	// ResponseHeaders returns http status code and headers.
	ResponseHeaders() (int, map[string]string)
}
