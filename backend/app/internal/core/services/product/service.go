package product

import (
	"context"

	"github.com/blackPavlin/shop/app/internal/core/entities"
)

// ProductService
type ProductService struct {
	productRepository ProductRepository
}

// NewProductService
func NewProductService(productRepository ProductRepository) ProductService {
	return ProductService{
		productRepository: productRepository,
	}
}

// Create
func (s ProductService) Create(ctx context.Context, product *entities.Product) (entities.ProductID, error) {
	return s.productRepository.Create(ctx, product)
}

// FindByID
func (s ProductService) FindByID(ctx context.Context, id entities.ProductID) (*entities.Product, error) {
	return s.productRepository.FindByID(ctx, id)
}
