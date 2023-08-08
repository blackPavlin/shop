// Package pg contains implementations for cart repositories
package pg

import (
	"context"

	"go.uber.org/zap"

	"github.com/blackPavlin/shop/ent"
	entcart "github.com/blackPavlin/shop/ent/cart"
	"github.com/blackPavlin/shop/ent/predicate"
	"github.com/blackPavlin/shop/internal/domain/cart"
	"github.com/blackPavlin/shop/internal/domain/product"
	"github.com/blackPavlin/shop/internal/domain/user"
	"github.com/blackPavlin/shop/pkg/errorx"
	"github.com/blackPavlin/shop/pkg/repositoryx/pg"
)

// CartRepository pg repository implementation.
type CartRepository struct {
	client *ent.Client
	logger *zap.Logger
}

// NewCartRepository create instance of CartRepository.
func NewCartRepository(client *ent.Client, logger *zap.Logger) *CartRepository {
	return &CartRepository{client: client, logger: logger}
}

// Create cart in db.
func (r *CartRepository) Create(ctx context.Context, props *cart.Props) (*cart.Cart, error) {
	userFromCtx, ok := user.GetUser(ctx)
	if !ok {
		return nil, errorx.ErrUnauthorized
	}
	row, err := r.client.Cart.Create().
		SetUserID(int64(userFromCtx.ID)).
		SetProductID(int64(props.ProductID)).
		SetAmount(props.Amount).
		Save(ctx)
	if err != nil {
		if pg.IsForeignKeyViolationErr(err, "cart_user_fk") {
			return nil, errorx.ErrNotFound
		}
		if pg.IsForeignKeyViolationErr(err, "cart_product_fk") {
			return nil, errorx.ErrNotFound
		}
		r.logger.Error("create cart error", zap.Error(err))
		return nil, errorx.ErrInternal
	}
	return mapDomainCartFromRow(row), nil
}

// Get cart from db.
func (r *CartRepository) Get(ctx context.Context, filter *cart.Filter) (*cart.Cart, error) {
	row, err := r.client.Cart.
		Query().
		Where(makePredicates(filter)...).
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errorx.ErrNotFound
		}
		r.logger.Error("get cart error", zap.Error(err))
		return nil, errorx.ErrInternal
	}
	return mapDomainCartFromRow(row), err
}

// Query carts from db based on criteria.
func (r *CartRepository) Query(
	ctx context.Context,
	criteria *cart.QueryCriteria,
) (*cart.QueryResult, error) {
	if userFromCtx, ok := user.GetUser(ctx); ok {
		criteria.Filter.UserID.Eq = user.IDs{userFromCtx.ID}
	}
	predicates := makePredicates(&criteria.Filter)
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

func makePredicates(filter *cart.Filter) []predicate.Cart {
	predicates := make([]predicate.Cart, 0)
	if len(filter.UserID.Eq) > 0 {
		predicates = append(predicates, entcart.UserIDIn(filter.UserID.Eq.ToInt64()...))
	}
	if len(filter.UserID.Neq) > 0 {
		predicates = append(predicates, entcart.UserIDNotIn(filter.UserID.Neq.ToInt64()...))
	}
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
