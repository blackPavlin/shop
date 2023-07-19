package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/blackPavlin/shop/internal/domain/user"
	"github.com/blackPavlin/shop/internal/transport/rest/controller/mapping"
	"github.com/blackPavlin/shop/internal/transport/rest/middleware"
	"github.com/blackPavlin/shop/pkg/restx"
)

// UserController
type UserController struct {
	userService    user.UserService
	authMiddleware *middleware.AuthMiddleware
}

// NewUserController
func NewUserController(
	userService user.UserService,
	authMiddleware *middleware.AuthMiddleware,
) *UserController {
	return &UserController{userService, authMiddleware}
}

// RegisterRoutes
func (ctrl *UserController) RegisterRoutes(r chi.Router) chi.Router {
	return r.Route("/user", func(r chi.Router) {
		r.Use(ctrl.authMiddleware.Authorization)
		r.Get("/", ctrl.GetUserHandler)
	})
}

// GetUserHandler
func (ctrl *UserController) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	user, err := ctrl.userService.Get(r.Context(), &user.Filter{})
	if err != nil {
		restx.HandleError(w, r, err)
		return
	}
	render.Respond(w, r, mapping.CreateGetUserResponse(user))
}
