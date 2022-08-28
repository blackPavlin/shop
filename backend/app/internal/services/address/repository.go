package address

import (
	"context"

	"github.com/blackPavlin/shop/app/internal/entities"
)

//go:generate mockgen -source $GOFILE -destination "repository_mock.go" -package "address"

type AddressRepository interface {
	Create(ctx context.Context, address *entities.Address) (*entities.Address, error)
	FindByID(ctx context.Context, userID entities.UserID, id entities.AddressID) (*entities.Address, error)
	FindByUserID(ctx context.Context, userID entities.UserID) ([]*entities.Address, error)
	DeleteByID(ctx context.Context, userID entities.UserID, id entities.AddressID) error
}
