package user

import (
	"context"

	"github.com/blackPavlin/shop/app/internal/core/entities"
)

// UserService
type UserService struct {
	userRepository UserRepository
}

// NewUserService
func NewUserService(userRepository UserRepository) UserService {
	return UserService{
		userRepository: userRepository,
	}
}

// Create user in db.
func (s UserService) Create(ctx context.Context, user *entities.User) (entities.UserID, error) {
	return s.userRepository.Create(ctx, user)
}

// FindByEmail user in db.
func (s UserService) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	return s.userRepository.FindByEmail(ctx, email)
}

// FindByID user in db.
func (s UserService) FindByID(ctx context.Context, id entities.UserID) (*entities.User, error) {
	return s.userRepository.FindByID(ctx, id)
}

// ExistsByEmail user in db.
func (s UserService) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	return s.userRepository.ExistsByEmail(ctx, email)
}
