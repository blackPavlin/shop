package address

import (
	"context"

	"github.com/blackPavlin/shop/app/internal/entities"
)

// AddressService
type AddressService struct {
	addressRepository AddressRepository
}

// NewAddressService
func NewAddressService(addressRepository AddressRepository) AddressService {
	return AddressService{
		addressRepository: addressRepository,
	}
}

// Create
func (s AddressService) Create(ctx context.Context, address *entities.Address) (entities.AddressID, error) {
	return s.addressRepository.Create(ctx, address)
}

// FindByID
func (s AddressService) FindByID(ctx context.Context, userID entities.UserID, id entities.AddressID) (*entities.Address, error) {
	return s.addressRepository.FindByID(ctx, userID, id)
}

// FindByUserID
func (s AddressService) FindByUserID(ctx context.Context, userID entities.UserID) ([]*entities.Address, error) {
	return s.addressRepository.FindByUserID(ctx, userID)
}

// DeleteByID
func (s AddressService) DeleteByID(ctx context.Context, userID entities.UserID, id entities.AddressID) error {
	return s.addressRepository.DeleteByID(ctx, userID, id)
}
