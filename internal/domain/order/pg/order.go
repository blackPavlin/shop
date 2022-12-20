package pg

import (
	"github.com/blackPavlin/shop/ent"
	"go.uber.org/zap"
)

// OrderRepository
type OrderRepository struct {
	client *ent.Client
	logger *zap.Logger
}

// NewOrderRepository
func NewOrderRepository(client *ent.Client, logger *zap.Logger) *OrderRepository {
	return &OrderRepository{
		client: client,
		logger: logger,
	}
}
