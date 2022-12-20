package controller

import (
	"net/http"

	"github.com/blackPavlin/shop/internal/domain/address"
	"github.com/go-chi/chi/v5"
)

// AddressController
type AddressController struct {
	addressService address.AddressService
}

// NewAddressController
func NewAddressController(addressService address.AddressService) *AddressController {
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
func (ctrl *AddressController) GetAddressesHandler(w http.ResponseWriter, r *http.Request) {}

// CreateAddressHandler
func (ctrl *AddressController) CreateAddressHandler(w http.ResponseWriter, r *http.Request) {}
