package utils

import (
	"fmt"
	"github.com/valyala/fasthttp" // Remove the extra quote here
)

// HTTPError represents an HTTP error with a status code and a message.
type HTTPError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

// Error returns the error message and makes HTTPError satisfy the error interface.
func (e *HTTPError) Error() string {
	return e.Message
}

// NewHTTPError creates a new HTTPError.
func NewHTTPError(statusCode int, message string) *HTTPError {
	return &HTTPError{
		StatusCode: statusCode,
		Message:    message,
	}
}

// ErrorHandler is a function that handles an error by setting the status code and returning a JSON error.
func ErrorHandler(ctx *fasthttp.RequestCtx, err error) {
	httpError, ok := err.(*HTTPError)
	if !ok {
		// If the error is not an HTTPError, create a generic HTTP 500 error.
		httpError = &HTTPError{
			StatusCode: fasthttp.StatusInternalServerError,
			Message:    fmt.Sprintf("An unexpected error occurred: %v", err),
		}
	}

	ctx.SetStatusCode(httpError.StatusCode)
	ctx.SetContentType("application/json")
	ctx.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, httpError.Message)))
}
