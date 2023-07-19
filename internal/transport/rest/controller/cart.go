package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/blackPavlin/shop/internal/domain/cart"
	"github.com/blackPavlin/shop/internal/transport/rest/middleware"
)

// CartController
type CartController struct {
	cartService    cart.Service
	authMiddleware *middleware.AuthMiddleware
}

// NewCartController
func NewCartController(
	cartService cart.Service,
	authMiddleware *middleware.AuthMiddleware,
) *CartController {
	return &CartController{cartService, authMiddleware}
}

// RegisterRoutes
func (ctrl *CartController) RegisterRoutes(r chi.Router) chi.Router {
	return r.Route("/cart", func(r chi.Router) {
		r.Use(ctrl.authMiddleware.Authorization)
		r.Get("/", ctrl.GetCartHandler)
		r.Post("/", ctrl.AddProductHandler)
		r.Patch("/", ctrl.UpdateProductHandler)
		r.Delete("/", ctrl.DeleteProductHandler)
	})
}

// GetCartHandler
func (ctrl *CartController) GetCartHandler(w http.ResponseWriter, r *http.Request) {}

// AddProductHandler
func (ctrl *CartController) AddProductHandler(w http.ResponseWriter, r *http.Request) {}

// UpdateProductHandler
func (ctrl *CartController) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {}

// DeleteProductHandler
func (ctrl *CartController) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {}
