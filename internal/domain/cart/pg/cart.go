package pg

import (
	"context"

	"github.com/blackPavlin/shop/ent"
	"github.com/blackPavlin/shop/ent/predicate"
	"github.com/blackPavlin/shop/internal/domain/cart"
	"github.com/blackPavlin/shop/internal/domain/product"
	"github.com/blackPavlin/shop/internal/domain/user"
	"github.com/blackPavlin/shop/pkg/errorx"
	"go.uber.org/zap"
)

// CartRepository
type CartRepository struct {
	client *ent.Client
	logger *zap.Logger
}

// NewCartRepository
func NewCartRepository(client *ent.Client, logger *zap.Logger) *CartRepository {
	return &CartRepository{client: client, logger: logger}
}

// Create
func (r *CartRepository) Create(ctx context.Context, props *cart.Props) (*cart.Cart, error) {
	// TODO: Get UserID from ctx

	row, err := r.client.Cart.Create().
		// SetUserID(int64(crt.UserID)).
		SetProductID(int64(props.ProductID)).
		SetAmount(props.Amount).
		Save(ctx)
	if err != nil {
		r.logger.Error("create cart error", zap.Error(err))
		return nil, errorx.ErrInternal
	}

	return mapDomainCartFromRow(row), nil
}

// Query
func (r *CartRepository) Query(
	ctx context.Context,
	criteria *cart.QueryCriteria,
) (*cart.QueryResult, error) {
	predicates := makePredicates(criteria.Filter)
	// TODO: Get UserID from ctx
	rows, err := r.client.Cart.Query().
		Where(predicates...).
		All(ctx)
	if err != nil {
		r.logger.Error("query carts error", zap.Error(err))
		return nil, errorx.ErrInternal
	}
	result := &cart.QueryResult{
		Items: mapDomainCartsFromRows(rows),
	}
	return result, nil
}

func makePredicates(criteria *cart.Filter) []predicate.Cart {
	predicates := make([]predicate.Cart, 0)
	// TODO
	return predicates
}

func mapDomainCartFromRow(row *ent.Cart) *cart.Cart {
	return &cart.Cart{
		ID:     cart.ID(row.ID),
		UserID: user.ID(row.UserID),
		Props: &cart.Props{
			ProductID: product.ID(row.ProductID),
			Amount:    row.Amount,
		},
	}
}

func mapDomainCartsFromRows(rows ent.Carts) cart.Carts {
	result := make(cart.Carts, 0, len(rows))
	for _, row := range rows {
		result = append(result, mapDomainCartFromRow(row))
	}
	return result
}
