package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/blackPavlin/shop/internal/domain/cart"
	"github.com/blackPavlin/shop/internal/transport/rest/middleware"
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
func (ctrl *CartController) GetCartHandler(w http.ResponseWriter, r *http.Request) {}

// AddProductHandler define handler for POST /api/cart.
func (ctrl *CartController) AddProductHandler(w http.ResponseWriter, r *http.Request) {}

// UpdateProductHandler define handler for PATCH /api/cart.
func (ctrl *CartController) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {}

// DeleteProductHandler define handler for DELETE /api/cart.
func (ctrl *CartController) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {}
