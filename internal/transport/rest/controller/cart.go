package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/blackPavlin/shop/internal/domain/cart"
)

// CartController
type CartController struct {
	cartService cart.Service
}

// NewCartController
func NewCartController(cartService cart.Service) *CartController {
	return &CartController{
		cartService: cartService,
	}
}

// RegisterRoutes
func (ctrl *CartController) RegisterRoutes(r chi.Router) chi.Router {
	return r.Route("/cart", func(r chi.Router) {
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
