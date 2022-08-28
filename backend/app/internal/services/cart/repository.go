package cart

import (
	"context"

	"github.com/blackPavlin/shop/app/internal/entities"
)

//go:generate mockgen -source $GOFILE -destination "repository_mock.go" -package "cart"

type CartRepository interface {
	Create(ctx context.Context, cart *entities.Cart) (entities.CartID, error)
	FindByUserID(ctx context.Context, userID entities.UserID) (*entities.Cart, error)
	DeleteProducts(ctx context.Context, userID entities.UserID) (*entities.Cart, error)
	DeleteProductByID(ctx context.Context, userID entities.UserID, id entities.ProductID) (*entities.Cart, error)
	AppendProduct(ctx context.Context, userID entities.UserID, product *entities.CartProduct) (*entities.Cart, error)
	UpdateProduct(ctx context.Context, userID entities.UserID, product *entities.CartProduct) (*entities.Cart, error)
}
