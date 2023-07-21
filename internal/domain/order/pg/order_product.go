package pg

import (
	"go.uber.org/zap"

	"github.com/blackPavlin/shop/ent"
)

// OrderProductRepository pg repository implementation.
type OrderProductRepository struct {
	client *ent.Client
	logger *zap.Logger
}

// NewOrderProductRepository create instance of OrderProductRepository.
func NewOrderProductRepository(client *ent.Client, logger *zap.Logger) *OrderProductRepository {
	return &OrderProductRepository{client: client, logger: logger}
}
