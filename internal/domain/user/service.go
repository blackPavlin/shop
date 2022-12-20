package user

import "context"

//go:generate mockgen -source $GOFILE -destination "service_mock.go" -package "user"

// UserService
type UserService interface {
	Create(ctx context.Context, props *Props) (*User, error)
	Get(ctx context.Context, filter *Filter) (*User, error)
}

// UserUseCase
type UserUseCase struct {
	userRepo Repository
}

// NewUserUseCase
func NewUserUseCase(userRepo Repository) *UserUseCase {
	return &UserUseCase{userRepo: userRepo}
}

// Create
func (s *UserUseCase) Create(ctx context.Context, props *Props) (*User, error) {
	return s.userRepo.Create(ctx, props)
}

// Get
func (s UserUseCase) Get(ctx context.Context, filter *Filter) (*User, error) {
	return s.userRepo.Get(ctx, filter)
}
