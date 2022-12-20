package controller

import (
	"net/http"

	"github.com/blackPavlin/shop/internal/domain/category"
	"github.com/blackPavlin/shop/internal/transport/rest/controller/mapping"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// CategoryController
type CategoryController struct {
	categoryService category.CategoryService
}

// NewCategoryController
func NewCategoryController(categoryService category.CategoryService) *CategoryController {
	return &CategoryController{
		categoryService: categoryService,
	}
}

// RegisterRoutes
func (ctrl *CategoryController) RegisterRoutes(r chi.Router) chi.Router {
	return r.Route("/category", func(r chi.Router) {
		r.Get("/", ctrl.GetCategoriesHandler)
		r.Post("/", ctrl.CreateCategoryHandler)
	})
}

// GetCategoriesHandler
func (ctrl *CategoryController) GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	categories, err := ctrl.categoryService.Query(r.Context(), &category.QueryCriteria{})
	if err != nil {
		return
	}

	render.Respond(w, r, mapping.CreateGetCategoriesResponse(categories))
}

// CreateCategoryHandler
func (ctrl *CategoryController) CreateCategoryHandler(w http.ResponseWriter, r *http.Request) {}
