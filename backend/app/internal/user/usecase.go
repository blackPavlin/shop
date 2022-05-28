package user

import (
	"context"

	"github.com/blackPavlin/shop/app/internal/models"
)

type UseCase interface {
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	Create(ctx context.Context, user *models.User) error
}
