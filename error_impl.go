package errors

import (
	"fmt"
)

type ErrorImpl struct {
	msg        string
	statusCode int
	errImpl    *ErrorImpl
}

func (e *ErrorImpl) StatusCode() int {
	return e.statusCode
}

func (e *ErrorImpl) Error() string {
	return e.msg
}

func (e *ErrorImpl) Is(err Error) bool {
	return e.errImpl == err
}

func (e *ErrorImpl) WithArgument(args ...any) Error {
	ne := &ErrorImpl{
		msg:        fmt.Sprintf(e.msg, args...),
		statusCode: e.statusCode,
		errImpl:    e,
	}
	return ne
}
