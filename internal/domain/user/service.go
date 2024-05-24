package user

import (
	"context"
	"fmt"
)

//go:generate go run go.uber.org/mock/mockgen@v0.4.0 -source $GOFILE -destination "service_mock.go" -package "user"

// Service represents user use cases.
type Service interface {
	Get(ctx context.Context, filter *Filter) (*User, error)
	Create(ctx context.Context, props *Props) (*User, error)
}

// UseCase represents user service.
type UseCase struct {
	repository Repository
}

// NewUseCase create instance of UseCase.
func NewUseCase(repository Repository) *UseCase {
	return &UseCase{repository: repository}
}

// Get user.
func (uc UseCase) Get(ctx context.Context, filter *Filter) (*User, error) {
	result, err := uc.repository.Get(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("get user error: %w", err)
	}
	return result, nil
}

// Create user.
func (uc UseCase) Create(ctx context.Context, props *Props) (*User, error) {
	result, err := uc.repository.Create(ctx, props)
	if err != nil {
		return nil, fmt.Errorf("create user error: %w", err)
	}
	return result, nil
}
