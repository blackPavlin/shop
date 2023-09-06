package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/blackPavlin/shop/internal/domain/category"
	"github.com/blackPavlin/shop/internal/transport/rest"
	"github.com/blackPavlin/shop/internal/transport/rest/controller/mapping"
	"github.com/blackPavlin/shop/internal/transport/rest/middleware"
	"github.com/blackPavlin/shop/pkg/errorx"
	"github.com/blackPavlin/shop/pkg/restx"
)

// CategoryController represents category controller.
type CategoryController struct {
	categoryService category.Service
	authMiddleware  *middleware.Middleware
}

// NewCategoryController create instance of CategoryController.
func NewCategoryController(
	categoryService category.Service,
	authMiddleware *middleware.Middleware,
) *CategoryController {
	return &CategoryController{categoryService, authMiddleware}
}

// RegisterRoutes register routes to the specified router.
func (ctrl *CategoryController) RegisterRoutes(r chi.Router) chi.Router {
	return r.Route("/category", func(r chi.Router) {
		r.Get("/", ctrl.GetCategoriesHandler)
		r.Group(func(r chi.Router) {
			r.Use(ctrl.authMiddleware.Authorization)
			r.Post("/", ctrl.CreateCategoryHandler)
			r.Patch("/", ctrl.UpdateCategoryHandler)
			r.Route("/{categoryID}", func(r chi.Router) {
				r.Delete("/", ctrl.DeleteCategoryHandler)
			})
		})
	})
}

// GetCategoriesHandler define handler for GET /api/category.
func (ctrl *CategoryController) GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	categories, err := ctrl.categoryService.Query(r.Context(), &category.QueryCriteria{})
	if err != nil {
		restx.HandleError(w, r, err)
		return
	}
	render.Respond(w, r, mapping.CreateGetCategoriesResponse(categories))
}

// CreateCategoryHandler define handler for POST /api/category.
func (ctrl *CategoryController) CreateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	request := &rest.CreateCategoryRequest{}
	if err := render.DecodeJSON(r.Body, request); err != nil {
		restx.HandleError(w, r, errorx.NewBadRequestError(err.Error()))
		return
	}
	if err := request.Validate(); err != nil {
		restx.HandleError(w, r, errorx.NewBadRequestError(err.Error()))
		return
	}
	c, err := ctrl.categoryService.Create(r.Context(), request.ToDomainEntity())
	if err != nil {
		restx.HandleError(w, r, err)
		return
	}
	render.Status(r, http.StatusCreated)
	render.Respond(w, r, mapping.CreateCategoryResponse(c))
}

// UpdateCategoryHandler define handler for PATCH /api/category.
func (ctrl *CategoryController) UpdateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	request := &rest.Category{}
	if err := render.DecodeJSON(r.Body, request); err != nil {
		restx.HandleError(w, r, errorx.NewBadRequestError(err.Error()))
		return
	}
	if err := request.Validate(); err != nil {
		restx.HandleError(w, r, errorx.NewBadRequestError(err.Error()))
		return
	}
	c, err := ctrl.categoryService.Update(r.Context(), category.ID(request.Id), &category.Props{
		Name: request.Name,
	})
	if err != nil {
		restx.HandleError(w, r, err)
		return
	}
	render.Respond(w, r, mapping.CreateCategoryResponse(c))
}

// DeleteCategoryHandler define handler for DELETE /api/category/{categoryId}.
func (ctrl *CategoryController) DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	categoryID, err := restx.GetIDFromURLParams(r, "categoryID")
	if err != nil {
		restx.HandleError(w, r, errorx.NewBadRequestError(err.Error()))
		return
	}
	if err := ctrl.categoryService.Delete(r.Context(), category.ID(categoryID)); err != nil {
		restx.HandleError(w, r, err)
		return
	}
	render.Status(r, http.StatusNoContent)
	render.Respond(w, r, nil)
}
