// Package restx contains utils for rest handlers.
package restx

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"

	"github.com/blackPavlin/shop/pkg/errorx"
)

// ErrorResponse is error struct for response.
type ErrorResponse struct {
	HTTPStatusCode int    `json:"-"`
	Message        string `json:"message"`
}

// Render http error response.
func (e ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

// HandleError return in response the rendered error.
func HandleError(w http.ResponseWriter, r *http.Request, err error) {
	var serviceError *errorx.Error

	if errors.As(err, &serviceError) {
		_ = render.Render(w, r, ErrorResponse{
			HTTPStatusCode: serviceError.Code(),
			Message:        serviceError.Error(),
		})
		return
	}
	render.Status(r, http.StatusInternalServerError)
	_ = render.Render(w, r, ErrorResponse{
		HTTPStatusCode: http.StatusInternalServerError,
		Message:        "internal server error",
	})
}
