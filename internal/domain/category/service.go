package category

import "context"

// Service
type Service interface {
	Create(ctx context.Context, props *Props) (*Category, error)
	Query(ctx context.Context, criteria *QueryCriteria) (Categories, error)
	Update(ctx context.Context, id ID, props *Props) (*Category, error)
	Delete(ctx context.Context, id ID) error
}

// UseCase
type UseCase struct {
	categoryRepo Repository
}

// NewUseCase
func NewUseCase(categoryRepo Repository) *UseCase {
	return &UseCase{categoryRepo: categoryRepo}
}

// Create
func (s *UseCase) Create(ctx context.Context, props *Props) (*Category, error) {
	return s.categoryRepo.Create(ctx, props)
}

// Query
func (s *UseCase) Query(ctx context.Context, criteria *QueryCriteria) (Categories, error) {
	return s.categoryRepo.Query(ctx, criteria)
}

// Update
func (s *UseCase) Update(ctx context.Context, id ID, props *Props) (*Category, error) {
	return s.categoryRepo.Update(ctx, id, props)
}

// Delete
func (s *UseCase) Delete(ctx context.Context, id ID) error {
	return s.categoryRepo.Delete(ctx, id)
}
