package faceit

import (
	"fmt"
	"net/http"
)

type ResponseError struct {
	StatusCode int
	Err        error
}

func (r *ResponseError) Error() string {
	return fmt.Sprintf("status %d: %v", r.StatusCode, r.Err)
}

func (r *ResponseError) Temporary() bool {
	return r.StatusCode == http.StatusServiceUnavailable
}

func (r *ResponseError) NotFound() bool {
	return r.StatusCode == http.StatusNotFound
}

func (r *ResponseError) NotAuthorized() bool {
	return r.StatusCode == http.StatusUnauthorized
}
