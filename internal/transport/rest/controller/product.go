package controller

import (
	"net/http"

	"github.com/blackPavlin/shop/internal/domain/product"
	"github.com/go-chi/chi/v5"
)

// ProductController
type ProductController struct {
	productService product.ProductService
}

// NewProductController
func NewProductController(productService product.ProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

// RegisterRoutes
func (ctrl *ProductController) RegisterRoutes(r chi.Router) chi.Router {
	return r.Route("/product", func(r chi.Router) {
		r.Get("/", ctrl.GetProductsHandler)
		r.Post("/", ctrl.CreateProductHandler)
		r.Get("/:id", ctrl.GetProductHandler)
		r.Patch("/:id", ctrl.UpdateProductHandler)
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
