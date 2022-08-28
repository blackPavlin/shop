package category

import (
	"context"

	"github.com/blackPavlin/shop/app/internal/entities"
)

//go:generate mockgen -source $GOFILE -destination "repository_mock.go" -package "category"

type CategoryRepository interface {
	Create(ctx context.Context, category *entities.Category) (entities.CategoryID, error)
	Find(ctx context.Context) ([]*entities.Category, error)
	FindByID(ctx context.Context, id entities.CategoryID) (*entities.Category, error)
	FindByName(ctx context.Context, name string) (*entities.Category, error)
	DeleteByID(ctx context.Context, id entities.CategoryID) error
	Update(ctx context.Context, category *entities.Category) (*entities.Category, error)
}
