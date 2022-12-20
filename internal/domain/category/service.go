package category

import "context"

// CategoryService
type CategoryService interface {
	Create(ctx context.Context, props *Props) (*Category, error)
	Query(ctx context.Context, criteria *QueryCriteria) (Categories, error)
}

// CategoryUseCase
type CategoryUseCase struct {
	categoryRepo Repository
}

// NewCategoryUseCase
func NewCategoryUseCase(categoryRepo Repository) *CategoryUseCase {
	return &CategoryUseCase{categoryRepo: categoryRepo}
}

// Create
func (s *CategoryUseCase) Create(ctx context.Context, props *Props) (*Category, error) {
	return s.categoryRepo.Create(ctx, props)
}

// Query
func (s *CategoryUseCase) Query(ctx context.Context, criteria *QueryCriteria) (Categories, error) {
	return s.categoryRepo.Query(ctx, criteria)
}
