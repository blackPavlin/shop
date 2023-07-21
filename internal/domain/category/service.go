package category

import (
	"context"
	"fmt"
)

// Service represents category use cases.
type Service interface {
	Create(ctx context.Context, props *Props) (*Category, error)
	Query(ctx context.Context, criteria *QueryCriteria) (Categories, error)
	Update(ctx context.Context, id ID, props *Props) (*Category, error)
	Delete(ctx context.Context, id ID) error
}

// UseCase represents category service.
type UseCase struct {
	categoryRepo Repository
}

// NewUseCase create instance of UseCase.
func NewUseCase(categoryRepo Repository) *UseCase {
	return &UseCase{categoryRepo: categoryRepo}
}

// Create category.
func (s *UseCase) Create(ctx context.Context, props *Props) (*Category, error) {
	result, err := s.categoryRepo.Create(ctx, props)
	if err != nil {
		return nil, fmt.Errorf("create category error: %w", err)
	}
	return result, nil
}

// Query categories.
func (s *UseCase) Query(ctx context.Context, criteria *QueryCriteria) (Categories, error) {
	result, err := s.categoryRepo.Query(ctx, criteria)
	if err != nil {
		return nil, fmt.Errorf("query categories error: %w", err)
	}
	return result, nil
}

// Update category.
func (s *UseCase) Update(ctx context.Context, id ID, props *Props) (*Category, error) {
	result, err := s.categoryRepo.Update(ctx, id, props)
	if err != nil {
		return nil, fmt.Errorf("update category error: %w", err)
	}
	return result, nil
}

// Delete category.
func (s *UseCase) Delete(ctx context.Context, id ID) error {
	if err := s.categoryRepo.Delete(ctx, id); err != nil {
		return fmt.Errorf("delete category error: %w", err)
	}
	return nil
}
