package mapping

import (
	"github.com/blackPavlin/shop/internal/domain/address"
	"github.com/blackPavlin/shop/internal/transport/rest"
)

// CreateAddressResponse transform domain entity to rest response.
func CreateAddressResponse(addr *address.Address) rest.Address {
	return rest.Address{
		Id:       int64(addr.ID),
		UserId:   int64(addr.UserID),
		City:     addr.Props.City,
		Country:  addr.Props.Country,
		Flat:     addr.Props.Flat,
		House:    addr.Props.House,
		Letter:   addr.Props.Letter,
		Postcode: addr.Props.Postcode,
		Street:   addr.Props.Street,
	}
}

// CreateAddressesResponse transform domain entity to rest response.
func CreateAddressesResponse(addresses address.Addresses) rest.AddressList {
	result := make(rest.AddressList, 0, len(addresses))
	for _, addr := range addresses {
		result = append(result, CreateAddressResponse(addr))
	}
	return result
}
