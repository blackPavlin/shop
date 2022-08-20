package user

import (
	"context"

	"github.com/blackPavlin/shop/app/internal/core/entities"
)

//go:generate mockgen -source $GOFILE -destination "service_mock.go" -package "user"

// UserService
type UserService interface {
	FindByID(ctx context.Context, id entities.UserID) (*entities.User, error)
}

// UserUseCase
type UserUseCase struct {
	userService UserService
}

// NewUserUseCase
func NewUserUseCase(userService UserService) UserUseCase {
	return UserUseCase{
		userService: userService,
	}
}

// GetUserByID
func (u UserUseCase) GetUserByID(ctx context.Context, id entities.UserID) (*entities.User, error) {
	return u.userService.FindByID(ctx, id)
}
