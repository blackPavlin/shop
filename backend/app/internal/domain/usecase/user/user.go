package user

import (
	"context"
	"errors"

	"github.com/blackPavlin/shop/app/internal/domain/entities"
)

var (
	ErrUserNotFound = errors.New("User not found")
)

type UserUseCase struct {
	userService UserService
}

type UserService interface {
	GetUserByID(ctx context.Context, id string) (*entities.User, error)
}

func NewUserUseCase(userService UserService) *UserUseCase {
	return &UserUseCase{
		userService: userService,
	}
}

func (u *UserUseCase) GetUser(ctx context.Context, id string) (*entities.User, error) {
	user, err := u.userService.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErrUserNotFound
	}

	return user, nil
}
