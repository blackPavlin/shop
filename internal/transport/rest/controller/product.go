package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/blackPavlin/shop/internal/domain/product"
	"github.com/blackPavlin/shop/internal/transport/rest/middleware"
)

// ProductController represents product controller.
type ProductController struct {
	productService product.Service
	authMiddleware *middleware.Middleware
}

// NewProductController create instance of ProductController.
func NewProductController(
	productService product.Service,
	authMiddleware *middleware.Middleware,
) *ProductController {
	return &ProductController{productService, authMiddleware}
}

// RegisterRoutes register routes to the specified router.
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

// GetProductsHandler define handler for GET /api/product.
func (ctrl *ProductController) GetProductsHandler(w http.ResponseWriter, r *http.Request) {}

// CreateProductHandler define handler for POST /api/product.
func (ctrl *ProductController) CreateProductHandler(w http.ResponseWriter, r *http.Request) {}

// GetProductHandler define handler for GET /api/product/{productId}.
func (ctrl *ProductController) GetProductHandler(w http.ResponseWriter, r *http.Request) {}

// UpdateProductHandler  define handler for PATCH /api/product/{productId}.
func (ctrl *ProductController) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {}
