package mapping

import (
	"github.com/blackPavlin/shop/internal/domain/address"
	"github.com/blackPavlin/shop/internal/transport/rest"
)

// CreateAddressResponse
func CreateAddressResponse(address *address.Address) rest.Address {
	return rest.Address{
		Id:       int64(address.ID),
		UserId:   int64(address.UserID),
		City:     address.Props.City,
		Country:  address.Props.Country,
		Flat:     address.Props.Flat,
		House:    address.Props.House,
		Letter:   address.Props.Letter,
		Postcode: address.Props.Postcode,
		Street:   address.Props.Street,
	}
}

// CreateAddressesResponse
func CreateAddressesResponse(addresses address.Addresses) rest.AddressList {
	result := make(rest.AddressList, 0, len(addresses))
	for _, address := range addresses {
		result = append(result, CreateAddressResponse(address))
	}
	return result
}
