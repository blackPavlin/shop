package auth

import (
	"context"

	"github.com/blackPavlin/shop/app/internal/models"
)

type UseCase interface {
	Signin(ctx context.Context, email, password string) (string, error)
	Signup(ctx context.Context, user *models.User) error
}
