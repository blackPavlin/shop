package errorx

import "net/http"

// Error
type Error struct {
	code    int
	message string
}

// Error
func (e *Error) Error() string {
	return e.message
}

// Code
func (e *Error) Code() int {
	return e.code
}

// NewError
func NewError(code int, message string) *Error {
	return &Error{
		code:    code,
		message: message,
	}
}

var (
	ErrInternal               = NewError(http.StatusInternalServerError, "internal server error")
	ErrUnauthorized           = NewError(http.StatusUnauthorized, "unautorized")
	ErrInvalidLoginOrPassword = NewError(http.StatusUnauthorized, "invalid login or password")
	ErrAllreadyExists         = NewError(http.StatusConflict, "allready exists")
	ErrNotFound               = NewError(http.StatusNotFound, "not found")
)

// NewBadRequestError
func NewBadRequestError(msg string) *Error {
	return NewError(http.StatusBadRequest, msg)
}

// NewNotFoundError
func NewNotFoundError(msg string) *Error {
	return NewError(http.StatusNotExtended, msg)
}
