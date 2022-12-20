package controller

import (
	"net/http"

	"github.com/blackPavlin/shop/internal/domain/auth"
	"github.com/blackPavlin/shop/internal/transport/rest"
	"github.com/blackPavlin/shop/internal/transport/rest/controller/mapping"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// AuthController
type AuthController struct {
	authService auth.AuthService
}

// NewAuthController
func NewAuthController(authService auth.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

// RegisterRoutes
func (ctrl *AuthController) RegisterRoutes(r chi.Router) chi.Router {
	return r.Route("/auth", func(r chi.Router) {
		r.Post("/login", ctrl.LoginHandler)
		r.Post("/signup", ctrl.SignupHandler)
	})
}

// Login
func (ctrl *AuthController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	request := &rest.LoginRequest{}
	if err := render.DecodeJSON(r.Body, request); err != nil {
		return
	}

	if err := request.Validate(); err != nil {
		return
	}

	token, err := ctrl.authService.Login(r.Context(), request.ToDomainEntity())
	if err != nil {
		return
	}

	render.Respond(w, r, mapping.CreateLoginResponse(token))
}

// Signup
func (ctrl *AuthController) SignupHandler(w http.ResponseWriter, r *http.Request) {
	request := &rest.SignupRequest{}
	if err := render.DecodeJSON(r.Body, request); err != nil {
		return
	}

	if err := request.Validate(); err != nil {
		return
	}

	token, err := ctrl.authService.Signup(r.Context(), request.ToDomainEntity())
	if err != nil {
		return
	}

	render.Respond(w, r, mapping.CreateLoginResponse(token))
}
