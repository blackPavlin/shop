package services

import (
	"context"

	"github.com/blackPavlin/shop/app/internal/domain/entities"
)

type ProductService struct {
	productRepo ProductRepository
}

type ProductRepository interface{
	Create(ctx context.Context, product *entities.Product) error
	FindByID(ctx context.Context, id string) (*entities.Product, error)
}

func NewProductService(productRepo ProductRepository) *ProductService {
	return &ProductService{
		productRepo: productRepo,
	}
}

func (p *ProductService) CreateProduct(ctx context.Context, product *entities.Product) error {
	return p.productRepo.Create(ctx, product)
}

func (p *ProductService) GetProductByID(ctx context.Context, id string) (*entities.Product, error) {
	return p.productRepo.FindByID(ctx, id)
}

func (p *ProductService) GetProductsByCategory(ctx context.Context, category string, offset, limit int32) {}
