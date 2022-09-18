package category

import (
	"context"

	"github.com/blackPavlin/shop/app/internal/entities"
)

//go:generate mockgen -source $GOFILE -destination "service_mock.go" -package "category"

// CategoryService
type CategoryService interface {
	Create(ctx context.Context, category *entities.Category) (entities.CategoryID, error)
	Find(ctx context.Context) (entities.Categories, error)
	FindByID(ctx context.Context, id entities.CategoryID) (*entities.Category, error)
	FindByName(ctx context.Context, name string) (*entities.Category, error)
	DeleteByID(ctx context.Context, id entities.CategoryID) error
	Update(ctx context.Context, category *entities.Category) (*entities.Category, error)
}

// CategoryService
type CategoryUseCase struct {
	categoryService CategoryService
}

// NewCategoryUseCase
func NewCategoryUseCase(categoryService CategoryService) CategoryUseCase {
	return CategoryUseCase{
		categoryService: categoryService,
	}
}

// Create
func (c CategoryUseCase) Create(ctx context.Context, category *entities.Category) (*entities.Category, error) {
	categoryID, err := c.categoryService.Create(ctx, category)
	if err != nil {
		return nil, err
	}

	return c.categoryService.FindByID(ctx, categoryID)
}

// Update
func (c CategoryUseCase) Update(ctx context.Context, category *entities.Category) (*entities.Category, error) {
	return c.categoryService.Update(ctx, category)
}

// Find
func (c CategoryUseCase) Find(ctx context.Context) (entities.Categories, error) {
	return c.categoryService.Find(ctx)
}
