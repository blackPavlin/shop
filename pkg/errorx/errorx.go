// Package errorx is a set of errors and error codes
package errorx

import (
	"fmt"
	"net/http"
)

// Error describes a domain-level error.
type Error struct {
	code    int
	message string
}

// Error returns a string representation of domain error.
func (e *Error) Error() string {
	return e.message
}

// Code returns error code for domain error.
func (e *Error) Code() int {
	return e.code
}

// NewError makes new domain-level error with its own error code.
func NewError(code int, message string) *Error {
	return &Error{
		code:    code,
		message: message,
	}
}

var (
	ErrInternal               = NewError(http.StatusInternalServerError, "internal server error")
	ErrUnauthorized           = NewError(http.StatusUnauthorized, "unauthorized")
	ErrInvalidLoginOrPassword = NewError(http.StatusUnauthorized, "invalid login or password")
	ErrAlreadyExists          = NewError(http.StatusConflict, "already exists")
	ErrNotFound               = NewError(http.StatusNotFound, "not found")
)

// NewBadRequestError create new bad request error.
func NewBadRequestError(msg string) *Error {
	return NewError(http.StatusBadRequest, msg)
}

// NewNotFoundError create new not found error.
func NewNotFoundError(msg string) *Error {
	return NewError(http.StatusNotFound, fmt.Sprintf("%s not found", msg))
}

// NewConflictError create new conflict error.
func NewConflictError(msg string) *Error {
	return NewError(http.StatusConflict, msg)
}
