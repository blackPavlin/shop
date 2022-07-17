package services

import (
	"context"

	"github.com/blackPavlin/shop/app/internal/domain/entities"
)

type OrderService struct {
	orderRepo OrderRepository
}

type OrderRepository interface{
	Create(ctx context.Context, order *entities.Order) error
	FindByUserID(ctx context.Context, userID string) ([]*entities.Order, error)
}

func NewOrderService(orderRepo OrderRepository) *OrderService {
	return &OrderService{
		orderRepo: orderRepo,
	}
}

func (o *OrderService) CreateOrder(ctx context.Context, order *entities.Order) error {
	return o.orderRepo.Create(ctx, order)
}

func (o *OrderService) GetOrdersByUserID(ctx context.Context, userID string) ([]*entities.Order, error) {
	return o.orderRepo.FindByUserID(ctx, userID)
}
