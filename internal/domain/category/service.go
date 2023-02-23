package category

import "context"

// CategoryService
type CategoryService interface {
	Create(ctx context.Context, props *Props) (*Category, error)
	Query(ctx context.Context, criteria *QueryCriteria) (Categories, error)
	Update(ctx context.Context, id ID, props *Props) (*Category, error)
	Delete(ctx context.Context, id ID) error
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

// Update
func (s *CategoryUseCase) Update(ctx context.Context, id ID, props *Props) (*Category, error) {
	return s.categoryRepo.Update(ctx, id, props)
}

// Delete
func (s *CategoryUseCase) Delete(ctx context.Context, id ID) error {
	return s.categoryRepo.Delete(ctx, id)
}
