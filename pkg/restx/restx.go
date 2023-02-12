package restx

import (
	"errors"
	"net/http"

	"github.com/blackPavlin/shop/pkg/errorx"
	"github.com/go-chi/render"
)

// ErrorResponse
type ErrorResponse struct {
	Message string `json:"message"`
}

// Render
func (e ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// HandleError
func HandleError(w http.ResponseWriter, r *http.Request, err error) {
	var serviceError *errorx.Error

	if errors.As(err, &serviceError) {
		render.Status(r, serviceError.Code())
		render.Render(w, r, ErrorResponse{Message: serviceError.Error()})
		return
	}

	render.Status(r, http.StatusInternalServerError)
	render.Render(w, r, ErrorResponse{Message: "internal server error"})
}
