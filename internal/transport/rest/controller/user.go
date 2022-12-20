package controller

import (
	"net/http"

	"github.com/blackPavlin/shop/internal/domain/user"
	"github.com/blackPavlin/shop/internal/transport/rest/controller/mapping"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// UserController
type UserController struct {
	userService user.UserService
}

// NewUserController
func NewUserController(userService user.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// RegisterRoutes
func (ctrl *UserController) RegisterRoutes(r chi.Router) chi.Router {
	return r.Route("/user", func(r chi.Router) {
		r.Get("/", ctrl.GetUserHandler)
	})
}

// GetUserHandler
func (ctrl *UserController) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	user, err := ctrl.userService.Get(r.Context(), &user.Filter{})
	if err != nil {
		return
	}

	render.Respond(w, r, mapping.CreateGetUserResponse(user))
}
