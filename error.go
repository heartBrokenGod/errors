package errors

import "net/http"

type Error interface {
	error
	StatusCode() int
	Is(err Error) bool
	WithArgument(args ...any) Error
}

// default constructor
func New(msg string) Error {
	return &ErrorImpl{
		msg:        msg,
		statusCode: http.StatusInternalServerError,
	}
}
