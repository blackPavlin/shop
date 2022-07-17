package services

import (
	"context"

	"github.com/blackPavlin/shop/app/internal/domain/entities"
)

type UserService struct {
	userRepo UserRepository
}

type UserRepository interface {
	Create(ctx context.Context, user *entities.User) ([12]byte, error)
	FindByID(ctx context.Context, id string) (*entities.User, error)
	FindByEmail(ctx context.Context, email string) (*entities.User, error)
}

func NewUserService(userRepo UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (u *UserService) CreateUser(ctx context.Context, user *entities.User) ([12]byte, error) {
	return u.userRepo.Create(ctx, user)
}

func (u *UserService) GetUserByID(ctx context.Context, id string) (*entities.User, error) {
	return u.userRepo.FindByID(ctx, id)
}

func (u *UserService) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	return u.userRepo.FindByEmail(ctx, email)
}
