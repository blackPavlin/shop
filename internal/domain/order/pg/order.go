// Package pg contains implementations for order repositories
package pg

import (
	"context"

	"go.uber.org/zap"

	"github.com/blackPavlin/shop/ent"
	entorder "github.com/blackPavlin/shop/ent/order"
	"github.com/blackPavlin/shop/ent/predicate"
	"github.com/blackPavlin/shop/internal/domain/order"
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

// Create order in db.
func (r *OrderRepository) Create(ctx context.Context) (*order.Order, error) {
	return nil, nil
}

// Get order from db.
func (r *OrderRepository) Get(ctx context.Context, filter *order.Filter) (*order.Order, error) {
	return nil, nil
}

// Query orders from db.
func (r *OrderRepository) Query(ctx context.Context, criteria *order.QueryCriteria) (*order.QueryResult, error) {
	return nil, nil
}

func makeOrderPredicates(filter *order.Filter) []predicate.Order {
	predicates := make([]predicate.Order, 0)
	if len(filter.ID.Eq) > 0 {
		predicates = append(predicates, entorder.IDIn(filter.ID.Eq.ToInt64()...))
	}
	if len(filter.ID.Neq) > 0 {
		predicates = append(predicates, entorder.IDNotIn(filter.ID.Neq.ToInt64()...))
	}
	if len(filter.UserID.Eq) > 0 {
		predicates = append(predicates, entorder.UserIDIn(filter.UserID.Eq.ToInt64()...))
	}
	if len(filter.UserID.Neq) > 0 {
		predicates = append(predicates, entorder.UserIDNotIn(filter.UserID.Neq.ToInt64()...))
	}
	return predicates
}

func mapDomainOrderFromRow(row *ent.Order) *order.Order {
	return &order.Order{
		ID:        order.ID(row.ID),
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
		Props:     &order.Props{},
	}
}

func mapDomainOrdersFromRows(rows ent.Orders) order.Orders {
	result := make(order.Orders, 0, len(rows))
	return result
}
