package errs

import (
	"net/http"
)

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
	// Internal
	ErrInternal               = NewError(http.StatusInternalServerError, "internal server error")
	ErrFailidTypecastObjectID = NewError(http.StatusInternalServerError, "failid typecast ObjectID")

	// Auth
	ErrUnauthorized           = NewError(http.StatusUnauthorized, "unautorized")
	ErrInvalidLoginOrPassword = NewError(http.StatusUnauthorized, "invalid login or password")

	//	User
	ErrUserNotFound       = NewError(http.StatusNotFound, "user not found")
	ErrUserAllreadyExists = NewError(http.StatusConflict, "user allready exists")

	// Cart
	ErrCartNotFound    = NewError(http.StatusNotFound, "cart not found")
	ErrTooManyProducts = NewError(http.StatusConflict, "too many products")

	// Product
	ErrProductNotFound = NewError(http.StatusNotFound, "product not found")

	// Category
	ErrCategoryNotFound = NewError(http.StatusNotFound, "category not found")

	// Address
	ErrAddressNotFound = NewError(http.StatusNotFound, "address not found")
)
