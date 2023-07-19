package restx

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"

	"github.com/blackPavlin/shop/pkg/errorx"
)

// ErrorResponse
type ErrorResponse struct {
	HTTPStatusCode int    `json:"-"`
	Message        string `json:"message"`
}

// Render
func (e ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

// HandleError
func HandleError(w http.ResponseWriter, r *http.Request, err error) {
	var serviceError *errorx.Error

	if errors.As(err, &serviceError) {
		render.Render(w, r, ErrorResponse{
			HTTPStatusCode: serviceError.Code(),
			Message:        serviceError.Error(),
		})
		return
	}

	render.Status(r, http.StatusInternalServerError)
	render.Render(w, r, ErrorResponse{
		HTTPStatusCode: http.StatusInternalServerError,
		Message:        "internal server error",
	})
}
