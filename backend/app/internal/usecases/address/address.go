package address

// AddressService
type AddressService interface{}

// AddressUseCase
type AddressUseCase struct {
	addressService AddressService
}

// NewAddressUseCase
func NewAddressUseCase(addressService AddressService) AddressUseCase {
	return AddressUseCase{
		addressService: addressService,
	}
}
