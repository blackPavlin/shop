package category

import (
	"context"

	"github.com/blackPavlin/shop/app/internal/domain/entities"
	usecase_dto "github.com/blackPavlin/shop/app/internal/domain/usecase/category/dto"
)

type CategoryUseCase struct {
	categoryService CategoryService
}

type CategoryService interface {
	CreateCategory(ctx context.Context, category *entities.Category) error
	GetCategories(ctx context.Context) ([]*entities.Category, error)
}

func NewCategoryUseCase(categoryService CategoryService) *CategoryUseCase {
	return &CategoryUseCase{
		categoryService: categoryService,
	}
}

func (c *CategoryUseCase) CreateCategory(ctx context.Context, dto *usecase_dto.CreateCategoryDTO) error {
	category := &entities.Category{
		Name: dto.Name,
	}

	return c.categoryService.CreateCategory(ctx, category)
}

func (c *CategoryUseCase) GetCategoryList(ctx context.Context) ([]string, error) {
	categories, err := c.categoryService.GetCategories(ctx)
	if err != nil {
		return nil, err
	}

	categoryList := []string{}

	for _, category := range categories {
		categoryList = append(categoryList, category.Name)
	}

	return categoryList, nil
}
