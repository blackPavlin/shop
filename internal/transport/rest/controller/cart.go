package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/blackPavlin/shop/internal/domain/cart"
	"github.com/blackPavlin/shop/internal/transport/rest"
	"github.com/blackPavlin/shop/internal/transport/rest/middleware"
	"github.com/blackPavlin/shop/pkg/errorx"
	"github.com/blackPavlin/shop/pkg/restx"
)

// CartController represents cart controller.
type CartController struct {
	cartService    cart.Service
	authMiddleware *middleware.Middleware
}

// NewCartController create instance of CartController.
func NewCartController(
	cartService cart.Service,
	authMiddleware *middleware.Middleware,
) *CartController {
	return &CartController{cartService, authMiddleware}
}

// RegisterRoutes register routes to the specified router.
func (ctrl *CartController) RegisterRoutes(r chi.Router) chi.Router {
	return r.Route("/cart", func(r chi.Router) {
		r.Use(ctrl.authMiddleware.Authorization)
		r.Get("/", ctrl.GetCartHandler)
		r.Post("/", ctrl.AddProductHandler)
		r.Patch("/", ctrl.UpdateProductHandler)
		r.Delete("/", ctrl.DeleteProductHandler)
	})
}

// GetCartHandler define handler for GET /api/cart.
func (ctrl *CartController) GetCartHandler(w http.ResponseWriter, r *http.Request) {
	_, err := ctrl.cartService.Query(r.Context(), &cart.QueryCriteria{})
	if err != nil {
		restx.HandleError(w, r, err)
		return
	}
}

// AddProductHandler define handler for POST /api/cart.
func (ctrl *CartController) AddProductHandler(w http.ResponseWriter, r *http.Request) {
	request := &rest.CartProduct{}
	if err := render.DecodeJSON(r.Body, request); err != nil {
		restx.HandleError(w, r, errorx.NewBadRequestError(err.Error()))
		return
	}
	if err := request.Validate(); err != nil {
		restx.HandleError(w, r, errorx.NewBadRequestError(err.Error()))
		return
	}
	_, err := ctrl.cartService.Save(r.Context(), request.ToDomainEntity())
	if err != nil {
		restx.HandleError(w, r, err)
		return
	}
	render.Status(r, http.StatusCreated)
}

// UpdateProductHandler define handler for PATCH /api/cart.
func (ctrl *CartController) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	request := &rest.CartProduct{}
	if err := render.DecodeJSON(r.Body, request); err != nil {
		restx.HandleError(w, r, errorx.NewBadRequestError(err.Error()))
		return
	}
	if err := request.Validate(); err != nil {
		restx.HandleError(w, r, errorx.NewBadRequestError(err.Error()))
		return
	}
	_, err := ctrl.cartService.Save(r.Context(), request.ToDomainEntity())
	if err != nil {
		restx.HandleError(w, r, err)
		return
	}
}

// DeleteProductHandler define handler for DELETE /api/cart.
func (ctrl *CartController) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {}
