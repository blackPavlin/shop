package user

import (
	"context"
	"fmt"
)

//go:generate mockgen -source $GOFILE -destination "service_mock.go" -package "user"

// Service represents user use cases.
type Service interface {
	Create(ctx context.Context, props *Props) (*User, error)
	Get(ctx context.Context, filter *Filter) (*User, error)
}

// UseCase represents user service.
type UseCase struct {
	userRepo Repository
}

// NewUseCase create instance of UseCase.
func NewUseCase(userRepo Repository) *UseCase {
	return &UseCase{userRepo: userRepo}
}

// Create user.
func (s *UseCase) Create(ctx context.Context, props *Props) (*User, error) {
	result, err := s.userRepo.Create(ctx, props)
	if err != nil {
		return nil, fmt.Errorf("create user error: %w", err)
	}
	return result, nil
}

// Get user.
func (s *UseCase) Get(ctx context.Context, filter *Filter) (*User, error) {
	result, err := s.userRepo.Get(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("get user error: %w", err)
	}
	return result, nil
}
