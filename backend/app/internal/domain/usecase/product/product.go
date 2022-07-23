package product

import (
	"context"
	"errors"

	"github.com/blackPavlin/shop/app/internal/domain/entities"
	usecase_dto "github.com/blackPavlin/shop/app/internal/domain/usecase/product/dto"
)

var (
	ErrCategoryNotFound = errors.New("Category not found")
	ErrProductNotFound  = errors.New("Product not found")
)

type ProductUseCase struct {
	productService  ProductService
	categoryService CategoryService
}

type ProductService interface {
	CreateProduct(ctx context.Context, product *entities.Product) error
	GetProductByID(ctx context.Context, id string) (*entities.Product, error)
	GetProductsByCategory(ctx context.Context, categoryID string, offset, limit int64) ([]*entities.Product, error)
}

type CategoryService interface {
	GetCategoryByName(ctx context.Context, name string) (*entities.Category, error)
	GetCategoryByID(ctx context.Context, id string) (*entities.Category, error)
}

func NewProductUseCase(productService ProductService, categoryService CategoryService) *ProductUseCase {
	return &ProductUseCase{
		productService:  productService,
		categoryService: categoryService,
	}
}

func (p *ProductUseCase) CreateProduct(ctx context.Context, dto *usecase_dto.CreateProductDTO) error {
	category, err := p.categoryService.GetCategoryByName(ctx, dto.Name)
	if err != nil {
		return err
	}

	if category == nil {
		return ErrCategoryNotFound
	}

	product := &entities.Product{
		Name:        dto.Name,
		Description: dto.Description,
		CategoryID:  category.ID,
		Quantity:    dto.Quantity,
	}

	return p.productService.CreateProduct(ctx, product)
}

func (p *ProductUseCase) GetProductsByCategory(ctx context.Context, dto *usecase_dto.GetProductDTO) ([]*usecase_dto.GetProductOutDTO, error) {
	category, err := p.categoryService.GetCategoryByName(ctx, dto.Category)
	if err != nil {
		return nil, err
	}

	if category == nil {
		return nil, ErrCategoryNotFound
	}

	categoryID := category.ID.Hex()

	products, err := p.productService.GetProductsByCategory(ctx, categoryID, dto.Offset, dto.Limit)
	if err != nil {
		return nil, err
	}

	outProducts := []*usecase_dto.GetProductOutDTO{}

	for _, product := range products {
		outProduct := &usecase_dto.GetProductOutDTO{
			ID:          product.ID.Hex(),
			Name:        product.Name,
			Description: product.Description,
			Category:    category.Name,
			Quantity:    product.Quantity,
		}

		outProducts = append(outProducts, outProduct)
	}

	return outProducts, nil
}

func (p *ProductUseCase) GetProductByID(ctx context.Context, id string) (*usecase_dto.GetProductOutDTO, error) {
	category, err := p.categoryService.GetCategoryByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if category == nil {
		return nil, ErrCategoryNotFound
	}

	product, err := p.productService.GetProductByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if product == nil {
		return nil, ErrProductNotFound
	}

	outProduct := &usecase_dto.GetProductOutDTO{
		ID:          product.ID.Hex(),
		Name:        product.Name,
		Description: product.Description,
		Category:    category.Name,
		Quantity:    product.Quantity,
	}

	return outProduct, nil
}
