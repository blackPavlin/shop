package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/blackPavlin/shop/internal/domain/product"
	"github.com/blackPavlin/shop/internal/transport/rest/middleware"
)

// ProductController
type ProductController struct {
	productService product.Service
	authMiddleware *middleware.AuthMiddleware
}

// NewProductController
func NewProductController(
	productService product.Service,
	authMiddleware *middleware.AuthMiddleware,
) *ProductController {
	return &ProductController{productService, authMiddleware}
}

// RegisterRoutes
func (ctrl *ProductController) RegisterRoutes(r chi.Router) chi.Router {
	return r.Route("/product", func(r chi.Router) {
		r.Get("/", ctrl.GetProductsHandler)
		r.Get("/:id", ctrl.GetProductHandler)
		r.Group(func(r chi.Router) {
			r.Use(ctrl.authMiddleware.Authorization)
			r.Post("/", ctrl.CreateProductHandler)
			r.Patch("/:id", ctrl.UpdateProductHandler)
		})
	})
}

// GetProductsHandler
func (ctrl *ProductController) GetProductsHandler(w http.ResponseWriter, r *http.Request) {}

// CreateProductHandler
func (ctrl *ProductController) CreateProductHandler(w http.ResponseWriter, r *http.Request) {}

// GetProductHandler
func (ctrl *ProductController) GetProductHandler(w http.ResponseWriter, r *http.Request) {}

// UpdateProductHandler
func (ctrl *ProductController) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {}
