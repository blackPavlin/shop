package usecase

import (
	"context"

	"github.com/blackPavlin/shop/app/internal/models"
	"github.com/blackPavlin/shop/app/internal/user"
)

type UserUseCase struct {
	userRepo user.Repository
}

func NewUserUseCase(userRepo user.Repository) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
	}
}

func (u *UserUseCase) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	return u.userRepo.FindByEmail(ctx, email)
}

func (u *UserUseCase) Create(ctx context.Context, user *models.User) error {
	return u.userRepo.Create(ctx, user)
}
