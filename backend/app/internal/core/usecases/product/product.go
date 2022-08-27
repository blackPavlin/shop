package product

import (
	"context"

	"github.com/blackPavlin/shop/app/internal/core/entities"
)

//go:generate mockgen -source $GOFILE -destination "service_mock.go" -package "product"

// ProductService
type ProductService interface {
	Create(ctx context.Context, product *entities.Product) (entities.ProductID, error)
}

// ProductUseCase
type ProductUseCase struct {
	productService ProductService
}

// NewProductUseCase
func NewProductUseCase(productService ProductService) ProductUseCase {
	return ProductUseCase{
		productService: productService,
	}
}
