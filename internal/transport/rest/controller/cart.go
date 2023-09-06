package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/blackPavlin/shop/internal/domain/cart"
	"github.com/blackPavlin/shop/internal/transport/rest"
	"github.com/blackPavlin/shop/internal/transport/rest/controller/mapping"
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
		r.Delete("/", ctrl.DeleteProductsHandler)
		r.Route("/{cartID}", func(r chi.Router) {
			r.Delete("/", ctrl.DeleteProductHandler)
		})
	})
}

// GetCartHandler define handler for GET /api/cart.
func (ctrl *CartController) GetCartHandler(w http.ResponseWriter, r *http.Request) {
	carts, err := ctrl.cartService.Query(r.Context(), &cart.QueryCriteria{})
	if err != nil {
		restx.HandleError(w, r, err)
		return
	}
	render.Respond(w, r, mapping.CreateCartListResponse(carts.Items))
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
	c, err := ctrl.cartService.Save(r.Context(), request.ToDomainEntity())
	if err != nil {
		restx.HandleError(w, r, err)
		return
	}
	render.Status(r, http.StatusCreated)
	render.Respond(w, r, mapping.CreateCartResponse(c))
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
	c, err := ctrl.cartService.Save(r.Context(), request.ToDomainEntity())
	if err != nil {
		restx.HandleError(w, r, err)
		return
	}
	render.Respond(w, r, mapping.CreateCartResponse(c))
}

// DeleteProductsHandler define handler for DELETE /api/cart.
func (ctrl *CartController) DeleteProductsHandler(w http.ResponseWriter, r *http.Request) {
	if err := ctrl.cartService.Delete(r.Context(), &cart.Filter{}); err != nil {
		restx.HandleError(w, r, err)
		return
	}
	render.Status(r, http.StatusNoContent)
	render.Respond(w, r, nil)
}

// DeleteProductHandler define handler for DELETE /api/cart/{cartId}.
func (ctrl *CartController) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	cartID, err := restx.GetIDFromURLParams(r, "cartId")
	if err != nil {
		restx.HandleError(w, r, errorx.NewBadRequestError(err.Error()))
		return
	}
	if err := ctrl.cartService.Delete(r.Context(), &cart.Filter{
		ID: cart.IDFilter{Eq: cart.IDs{cart.ID(cartID)}},
	}); err != nil {
		restx.HandleError(w, r, err)
		return
	}
	render.Status(r, http.StatusNoContent)
	render.Respond(w, r, nil)
}
