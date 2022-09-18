package category

import (
	"context"

	"github.com/blackPavlin/shop/app/internal/entities"
)

// CategoryService
type CategoryService struct {
	categoryRepository CategoryRepository
}

// NewCategoryService
func NewCategoryService(categoryRepository CategoryRepository) CategoryService {
	return CategoryService{
		categoryRepository: categoryRepository,
	}
}

// Create
func (c CategoryService) Create(ctx context.Context, category *entities.Category) (entities.CategoryID, error) {
	return c.categoryRepository.Create(ctx, category)
}

// Find
func (c CategoryService) Find(ctx context.Context) (entities.Categories, error) {
	return c.categoryRepository.Find(ctx)
}

// FindByID
func (c CategoryService) FindByID(ctx context.Context, id entities.CategoryID) (*entities.Category, error) {
	return c.categoryRepository.FindByID(ctx, id)
}

// FindByName
func (c CategoryService) FindByName(ctx context.Context, name string) (*entities.Category, error) {
	return c.categoryRepository.FindByName(ctx, name)
}

// DeleteByID
func (c CategoryService) DeleteByID(ctx context.Context, id entities.CategoryID) error {
	return c.categoryRepository.DeleteByID(ctx, id)
}

// Update
func (c CategoryService) Update(ctx context.Context, category *entities.Category) (*entities.Category, error) {
	return c.categoryRepository.Update(ctx, category)
}
