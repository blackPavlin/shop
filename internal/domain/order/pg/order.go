// Package pg contains implementations for order repositories
package pg

import (
	"go.uber.org/zap"

	"github.com/blackPavlin/shop/ent"
)

// OrderRepository pg repository implementation.
type OrderRepository struct {
	client *ent.Client
	logger *zap.Logger
}

// NewOrderRepository create instance of OrderRepository.
func NewOrderRepository(client *ent.Client, logger *zap.Logger) *OrderRepository {
	return &OrderRepository{client: client, logger: logger}
}
