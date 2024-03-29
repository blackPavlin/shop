package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/blackPavlin/shop/internal/domain/auth"
	"github.com/blackPavlin/shop/internal/transport/rest"
	"github.com/blackPavlin/shop/internal/transport/rest/controller/mapping"
	"github.com/blackPavlin/shop/pkg/errorx"
	"github.com/blackPavlin/shop/pkg/restx"
)

// AuthController represents auth controller.
type AuthController struct {
	authService auth.Service
}

// NewAuthController create instance of AuthController.
func NewAuthController(authService auth.Service) *AuthController {
	return &AuthController{authService}
}

// RegisterRoutes register routes to the specified router.
func (ctrl *AuthController) RegisterRoutes(r chi.Router) chi.Router {
	return r.Route("/auth", func(r chi.Router) {
		r.Post("/login", ctrl.LoginHandler)
		r.Post("/signup", ctrl.SignupHandler)
	})
}

// LoginHandler define handler for POST /api/auth/login.
func (ctrl *AuthController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	request := &rest.LoginRequest{}
	if err := render.DecodeJSON(r.Body, request); err != nil {
		restx.HandleError(w, r, errorx.NewBadRequestError(err.Error()))
		return
	}
	if err := request.Validate(); err != nil {
		restx.HandleError(w, r, errorx.NewBadRequestError(err.Error()))
		return
	}
	token, err := ctrl.authService.Login(r.Context(), request.ToDomainEntity())
	if err != nil {
		restx.HandleError(w, r, err)
		return
	}
	render.Respond(w, r, mapping.CreateLoginResponse(token))
}

// SignupHandler define handler for POST /api/auth/signup.
func (ctrl *AuthController) SignupHandler(w http.ResponseWriter, r *http.Request) {
	request := &rest.SignupRequest{}
	if err := render.DecodeJSON(r.Body, request); err != nil {
		restx.HandleError(w, r, errorx.NewBadRequestError(err.Error()))
		return
	}
	if err := request.Validate(); err != nil {
		restx.HandleError(w, r, errorx.NewBadRequestError(err.Error()))
		return
	}
	token, err := ctrl.authService.Signup(r.Context(), request.ToDomainEntity())
	if err != nil {
		restx.HandleError(w, r, err)
		return
	}
	render.Status(r, http.StatusCreated)
	render.Respond(w, r, mapping.CreateLoginResponse(token))
}
