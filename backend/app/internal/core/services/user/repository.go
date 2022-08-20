package user

import (
	"context"

	"github.com/blackPavlin/shop/app/internal/core/entities"
)

//go:generate mockgen -source $GOFILE -destination "repository_mock.go" -package "user"

// UserRepository
type UserRepository interface {
	Create(ctx context.Context, user *entities.User) (entities.UserID, error)
	FindByID(ctx context.Context, id entities.UserID) (*entities.User, error)
	FindByEmail(ctx context.Context, email string) (*entities.User, error)
	ExistsByEmail(ctx context.Context, email string) (bool, error)
}
