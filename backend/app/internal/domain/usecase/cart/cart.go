package cart

import (
	"context"
	"errors"

	"github.com/blackPavlin/shop/app/internal/domain/entities"
)

var (
	ErrCartNotFound = errors.New("Cart not found")
)

type CartUseCase struct {
	cartService CartService
}

type CartService interface {
	GetCartByUserID(ctx context.Context, userID string) (*entities.Cart, error)
	AddProductToCart(ctx context.Context, userID string, product *entities.CartProduct) error
}

func NewCartUseCase(cartService CartService) *CartUseCase {
	return &CartUseCase{
		cartService: cartService,
	}
}

func (c *CartUseCase) GetCart(ctx context.Context, userID string) (*entities.Cart, error) {
	cart, err := c.cartService.GetCartByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if cart == nil {
		return nil, ErrCartNotFound
	}

	return cart, nil
}
