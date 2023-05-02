package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/blackPavlin/shop/internal/domain/address"
	"github.com/blackPavlin/shop/internal/transport/rest"
	"github.com/blackPavlin/shop/internal/transport/rest/controller/mapping"
	"github.com/blackPavlin/shop/pkg/errorx"
	"github.com/blackPavlin/shop/pkg/restx"
)

// AddressController
type AddressController struct {
	addressService address.Service
}

// NewAddressController
func NewAddressController(addressService address.Service) *AddressController {
	return &AddressController{
		addressService: addressService,
	}
}

// RegisterRoutes
func (ctrl *AddressController) RegisterRoutes(r chi.Router) chi.Router {
	return r.Route("/address", func(r chi.Router) {
		r.Get("/", ctrl.GetAddressesHandler)
		r.Post("/", ctrl.CreateAddressHandler)
	})
}

// GetAddressesHandler
func (ctrl *AddressController) GetAddressesHandler(w http.ResponseWriter, r *http.Request) {
	addresses, err := ctrl.addressService.Query(r.Context(), &address.QueryCriteria{})
	if err != nil {
		restx.HandleError(w, r, err)
		return
	}

	render.Respond(w, r, mapping.CreateAddressesResponse(addresses))
}

// CreateAddressHandler
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
