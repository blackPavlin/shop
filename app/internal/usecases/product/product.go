package product

import (
	"context"

	"github.com/blackPavlin/shop/app/internal/entities"
)

//go:generate mockgen -source $GOFILE -destination "service_mock.go" -package "product"

// ProductService
type ProductService interface {
	Create(ctx context.Context, product *entities.Product) (entities.ProductID, error)
	FindByID(ctx context.Context, id entities.ProductID) (*entities.Product, error)
}

// CategoryService
type CategoryService interface {
	FindByName(ctx context.Context, name string) (*entities.Category, error)
}

// ProductUseCase
type ProductUseCase struct {
	productService  ProductService
	categoryService CategoryService
}

// NewProductUseCase
func NewProductUseCase(productService ProductService) ProductUseCase {
	return ProductUseCase{
		productService: productService,
	}
}

// Create
func (p ProductUseCase) Create(ctx context.Context, product *entities.Product) (*entities.Product, error) {
	productID, err := p.productService.Create(ctx, product)
	if err != nil {
		return nil, err
	}

	return p.productService.FindByID(ctx, productID)
}

// Find
func (p ProductUseCase) Find(ctx context.Context) {}

// FindByCategory
func (p ProductUseCase) FindByCategory(ctx context.Context, categoryName string) {}
