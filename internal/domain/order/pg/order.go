package pg

import (
	"go.uber.org/zap"

	"github.com/blackPavlin/shop/ent"
)

// OrderRepository
type OrderRepository struct {
	client *ent.Client
	logger *zap.Logger
}

// NewOrderRepository
func NewOrderRepository(client *ent.Client, logger *zap.Logger) *OrderRepository {
	return &OrderRepository{client: client, logger: logger}
}
