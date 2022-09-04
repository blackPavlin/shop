package address

import (
	"context"

	"github.com/blackPavlin/shop/app/internal/entities"
)

//go:generate mockgen -source $GOFILE -destination "service_mock.go" -package "address"

// AddressService
type AddressService interface {
	Create(ctx context.Context, address *entities.Address) (entities.AddressID, error)
	FindByID(ctx context.Context, userID entities.UserID, id entities.AddressID) (*entities.Address, error)
	FindByUserID(ctx context.Context, userID entities.UserID) ([]*entities.Address, error)
	DeleteByID(ctx context.Context, userID entities.UserID, id entities.AddressID) error
}

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

// Create
func (a AddressUseCase) Create(ctx context.Context, userID entities.UserID, address *entities.Address) (*entities.Address, error) {
	addressID, err := a.addressService.Create(ctx, address)
	if err != nil {
		return nil, err
	}

	return a.addressService.FindByID(ctx, userID, addressID)
}

// FindByUserID
func (a AddressUseCase) FindByUserID(ctx context.Context, userID entities.UserID) ([]*entities.Address, error) {
	return a.addressService.FindByUserID(ctx, userID)
}
