package controller

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/blackPavlin/shop/internal/domain/category"
	"github.com/blackPavlin/shop/internal/transport/rest"
	"github.com/blackPavlin/shop/internal/transport/rest/controller/mapping"
	"github.com/blackPavlin/shop/pkg/errorx"
	"github.com/blackPavlin/shop/pkg/restx"
)

// CategoryController
type CategoryController struct {
	categoryService category.Service
}

// NewCategoryController
func NewCategoryController(categoryService category.Service) *CategoryController {
	return &CategoryController{categoryService: categoryService}
}

// RegisterRoutes
func (ctrl *CategoryController) RegisterRoutes(r chi.Router) chi.Router {
	return r.Route("/category", func(r chi.Router) {
		r.Get("/", ctrl.GetCategoriesHandler)
		r.Post("/", ctrl.CreateCategoryHandler)
		r.Patch("/", ctrl.UpdateCategoryHandler)
		r.Delete("/{categoryID}", ctrl.DeleteCategoryHandler)
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
	categ, err := ctrl.categoryService.Create(r.Context(), request.ToDomainEntity())
	if err != nil {
		restx.HandleError(w, r, err)
		return
	}
	render.Status(r, http.StatusCreated)
	render.Respond(w, r, mapping.CreateCategoryResponse(categ))
}

// UpdateCategoryHandler
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
	categ, err := ctrl.categoryService.Update(r.Context(), category.ID(request.Id), &category.Props{
		Name: request.Name,
	})
	if err != nil {
		restx.HandleError(w, r, err)
		return
	}
	render.Respond(w, r, mapping.CreateCategoryResponse(categ))
}

// DeleteCategoryHandler
func (ctrl *CategoryController) DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	categoryID := chi.URLParam(r, "categoryID")
	id, err := strconv.Atoi(categoryID)
	if err != nil {
		restx.HandleError(w, r, errorx.NewBadRequestError(err.Error()))
		return
	}
	if err := ctrl.categoryService.Delete(r.Context(), category.ID(id)); err != nil {
		restx.HandleError(w, r, err)
		return
	}
}
