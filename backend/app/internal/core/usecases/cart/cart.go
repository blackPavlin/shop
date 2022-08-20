package cart

import (
	"context"

	"github.com/blackPavlin/shop/app/internal/core/entities"
)

//go:generate mockgen -source $GOFILE -destination "service_mock.go" -package "cart"

// CartService
type CartService interface {
	FindByUserID(ctx context.Context, userID entities.UserID) (*entities.Cart, error)
}

// CartUseCase
type CartUseCase struct {
	cartService CartService
}

// NewCartUseCase
func NewCartUseCase(cartService CartService) CartUseCase {
	return CartUseCase{
		cartService: cartService,
	}
}

// GetCartByUserID
func (c *CartUseCase) GetCartByUserID(ctx context.Context, userID entities.UserID) (*entities.Cart, error) {
	return c.cartService.FindByUserID(ctx, userID)
}
