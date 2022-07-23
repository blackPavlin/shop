package services

import (
	"context"
	"errors"

	"github.com/blackPavlin/shop/app/internal/domain/entities"
)

var (
	ErrCategoryAllreadyExista = errors.New("Category allready exists")
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

func (c *CategoryService) CreateCategory(ctx context.Context, category *entities.Category) error {
	existsCategory, err := c.GetCategoryByName(ctx, category.Name)
	if err != nil {
		return err
	}

	if existsCategory != nil {
		return ErrCategoryAllreadyExista
	}

	return c.categoryRepository.Create(ctx, category)
}

func (c *CategoryService) GetCategories(ctx context.Context) ([]*entities.Category, error) {
	return c.categoryRepository.Find(ctx)
}

func (c *CategoryService) GetCategoryByID(ctx context.Context, id string) (*entities.Category, error) {
	return c.categoryRepository.FindByID(ctx, id)
}

func (c *CategoryService) GetCategoryByName(ctx context.Context, name string) (*entities.Category, error) {
	return c.categoryRepository.FindByName(ctx, name)
}
