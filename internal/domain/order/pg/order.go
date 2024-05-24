// Package pg contains implementations for order repositories
package pg

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/blackPavlin/shop/ent"
	entorder "github.com/blackPavlin/shop/ent/order"
	"github.com/blackPavlin/shop/ent/predicate"
	"github.com/blackPavlin/shop/internal/domain/order"
	"github.com/blackPavlin/shop/internal/domain/user"
	"github.com/blackPavlin/shop/pkg/errorx"
	"github.com/blackPavlin/shop/pkg/repositoryx/pg"
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
func (r OrderRepository) Create(ctx context.Context, or *order.Order) (*order.Order, error) {
	client := r.client
	if tx := ent.TxFromContext(ctx); tx != nil {
		client = tx.Client()
	}
	row, err := client.Order.Create().
		SetUserID(or.UserID.ToInt64()).
		SetStatus(entorder.StatusCREATED).
		Save(ctx)
	if err != nil {
		if pg.IsForeignKeyViolationErr(err, "order_user_fk") {
			return nil, errorx.ErrNotFound
		}
		r.logger.Error("create order error", zap.Error(err))
		return nil, errorx.ErrInternal
	}
	result, err := mapDomainOrderFromRow(row)
	if err != nil {
		r.logger.Error("map domain order from row error", zap.Error(err))
		return nil, errorx.ErrInternal
	}
	return result, nil
}

// Get order from db.
func (r OrderRepository) Get(ctx context.Context, filter *order.Filter) (*order.Order, error) {
	row, err := r.client.Order.Query().
		Where(makeOrderPredicates(filter)...).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errorx.ErrNotFound
		}
		r.logger.Error("get order error", zap.Error(err))
		return nil, errorx.ErrInternal
	}
	result, err := mapDomainOrderFromRow(row)
	if err != nil {
		r.logger.Error("map domain order from row error", zap.Error(err))
		return nil, errorx.ErrInternal
	}
	return result, nil
}

// Query orders from db.
func (r OrderRepository) Query(ctx context.Context, criteria *order.QueryCriteria) (*order.QueryResult, error) {
	var (
		result = &order.QueryResult{}
		err    error
	)
	predicates := makeOrderPredicates(&criteria.Filter)
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		query := r.client.Order.Query().Where(predicates...)
		if criteria.Pagination.Limit > 0 {
			query.Limit(int(criteria.Pagination.Limit))
		}
		if criteria.Pagination.Offset > 0 {
			query.Offset(int(criteria.Pagination.Offset))
		}
		rows, err := query.All(ctx)
		if err != nil {
			r.logger.Error("query orders error", zap.Error(err))
			return errorx.ErrInternal
		}
		result.Items, err = mapDomainOrdersFromRows(rows)
		if err != nil {
			r.logger.Error("map domain orders from rows error", zap.Error(err))
			return errorx.ErrInternal
		}
		return nil
	})
	g.Go(func() error {
		result.Count, err = r.client.Order.Query().
			Where(predicates...).
			Count(ctx)
		if err != nil {
			r.logger.Error("get orders count error", zap.Error(err))
			return errorx.ErrInternal
		}
		return nil
	})
	if err = g.Wait(); err != nil {
		return nil, errorx.ErrInternal
	}
	return result, nil
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

func mapDomainOrderFromRow(row *ent.Order) (*order.Order, error) {
	status, err := order.StatusString(row.Status.String())
	if err != nil {
		return nil, fmt.Errorf("convertation order status error: %w", err)
	}
	return &order.Order{
		ID:        order.ID(row.ID),
		UserID:    user.ID(row.UserID),
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
		Props: &order.Props{
			Status: status,
		},
	}, nil
}

func mapDomainOrdersFromRows(rows ent.Orders) (order.Orders, error) {
	result := make(order.Orders, 0, len(rows))
	for _, row := range rows {
		res, err := mapDomainOrderFromRow(row)
		if err != nil {
			return nil, fmt.Errorf("map domain orders from rows error: %w", err)
		}
		result = append(result, res)
	}
	return result, nil
}
