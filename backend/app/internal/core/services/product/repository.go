package product

import (
	"context"

	"github.com/blackPavlin/shop/app/internal/core/entities"
)

//go:generate mockgen -source $GOFILE -destination "repository_mock.go" -package "product"

type ProductRepository interface {
	Create(ctx context.Context, product *entities.Product) (entities.ProductID, error)
	FindByID(ctx context.Context, id entities.ProductID) (*entities.Product, error)
}
