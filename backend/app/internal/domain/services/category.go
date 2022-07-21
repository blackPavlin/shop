package services

import (
	"context"

	"github.com/blackPavlin/shop/app/internal/domain/entities"
)

type CategoryService struct {
	categoryRepository CategoryRepository
}

type CategoryRepository interface {
	Create(ctx context.Context, category *entities.Category) error
	Find(ctx context.Context) ([]*entities.Category, error)
	FindByID(ctx context.Context, id string) (*entities.Category, error)
	FindByName(ctx context.Context, name string) (*entities.Category, error)
}

func NewCategoryService(categoryRepository CategoryRepository) *CategoryService {
	return &CategoryService{
		categoryRepository: categoryRepository,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context) {}

func (c *CategoryService) GetCategories(ctx context.Context) {}

func (c *CategoryService) GetCategoryByID(ctx context.Context) {}

func (c *CategoryService) GetCategoryByName(ctx context.Context) {}
