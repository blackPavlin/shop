package product

import (
	"context"

	"github.com/blackPavlin/shop/app/internal/domain/entities"
	"github.com/blackPavlin/shop/app/internal/domain/usecase/product/dto"
)

type ProductUseCase struct {
	productService ProductService
}

type ProductService interface {
	CreateProduct(ctx context.Context, product *entities.Product) error
}

func NewProductUseCase(productService ProductService) *ProductUseCase {
	return &ProductUseCase{
		productService: productService,
	}
}

func (p *ProductUseCase) CreateProduct(ctx context.Context, dto dto.CreateProductDTO) error {
	product := &entities.Product{
		Name:        dto.Name,
		Description: dto.Description,
		Category:    dto.Category,
	}

	return p.productService.CreateProduct(ctx, product)
}
