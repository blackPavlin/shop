package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/blackPavlin/shop/internal/domain/address"
	"github.com/blackPavlin/shop/internal/transport/rest"
	"github.com/blackPavlin/shop/internal/transport/rest/controller/mapping"
	"github.com/blackPavlin/shop/internal/transport/rest/middleware"
	"github.com/blackPavlin/shop/pkg/errorx"
	"github.com/blackPavlin/shop/pkg/restx"
)

// AddressController represents address controller.
type AddressController struct {
	addressService address.Service
	authMiddleware *middleware.Middleware
}

// NewAddressController create instance of AddressController.
func NewAddressController(
	addressService address.Service,
	authMiddleware *middleware.Middleware,
) *AddressController {
	return &AddressController{addressService, authMiddleware}
}

// RegisterRoutes register routes to the specified router.
func (ctrl *AddressController) RegisterRoutes(r chi.Router) chi.Router {
	return r.Route("/address", func(r chi.Router) {
		r.Use(ctrl.authMiddleware.Authorization)
		r.Get("/", ctrl.GetAddressesHandler)
		r.Post("/", ctrl.CreateAddressHandler)
	})
}

// GetAddressesHandler define handler for GET /api/address.
func (ctrl *AddressController) GetAddressesHandler(w http.ResponseWriter, r *http.Request) {
	addresses, err := ctrl.addressService.Query(r.Context(), &address.QueryCriteria{})
	if err != nil {
		restx.HandleError(w, r, err)
		return
	}
	render.Respond(w, r, mapping.CreateAddressesResponse(addresses))
}

// CreateAddressHandler define handler for POST /api/address.
func (ctrl *AddressController) CreateAddressHandler(w http.ResponseWriter, r *http.Request) {
	request := &rest.CreateAddressRequest{}
	if err := render.DecodeJSON(r.Body, request); err != nil {
		restx.HandleError(w, r, errorx.NewBadRequestError(err.Error()))
		return
	}
	if err := request.Validate(); err != nil {
		restx.HandleError(w, r, errorx.NewBadRequestError(err.Error()))
		return
	}
	addr, err := ctrl.addressService.Create(r.Context(), request.ToDomainEntity())
	if err != nil {
		restx.HandleError(w, r, err)
		return
	}
	render.Status(r, http.StatusCreated)
	render.Respond(w, r, mapping.CreateAddressResponse(addr))
}
