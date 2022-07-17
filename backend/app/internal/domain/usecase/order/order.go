package order

import (
	"context"

	"github.com/blackPavlin/shop/app/internal/domain/entities"
)

type OrderUseCase struct {
	orderService OrderService
	cartService  CartService
}

type OrderService interface{}

type CartService interface {
	GetCartByUserID(ctx context.Context, userID string) (*entities.Cart, error)
}

func NewOrderUseCase(orderService OrderService, cartService CartService) *OrderUseCase {
	return &OrderUseCase{
		orderService: orderService,
		cartService:  cartService,
	}
}
