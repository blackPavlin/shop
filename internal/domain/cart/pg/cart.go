// Package pg contains implementations for cart repositories
package pg

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"go.uber.org/zap"

	"github.com/blackPavlin/shop/ent"
	entcart "github.com/blackPavlin/shop/ent/cart"
	"github.com/blackPavlin/shop/ent/predicate"
	"github.com/blackPavlin/shop/internal/domain/cart"
	"github.com/blackPavlin/shop/internal/domain/category"
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

// Save cart in db.
func (r *CartRepository) Save(ctx context.Context, props *cart.Props) (*cart.Cart, error) {
	userFromCtx, ok := user.GetUser(ctx)
	if !ok {
		return nil, errorx.ErrUnauthorized
	}
	err := r.client.Cart.Create().
		SetUserID(int64(userFromCtx.ID)).
		SetProductID(int64(props.ProductID)).
		SetAmount(props.Amount).
		OnConflict(
			sql.ConflictConstraint(""),
		).
		UpdateAmount().
		UpdateUpdatedAt().
		Exec(ctx)
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
	return r.Get(ctx, &cart.Filter{})
}

// Delete carts from db.
func (r *CartRepository) Delete(ctx context.Context, filter *cart.Filter) error {
	if userFromCtx, ok := user.GetUser(ctx); ok {
		filter.UserID.Eq = user.IDs{userFromCtx.ID}
	}
	_, err := r.client.Cart.Delete().
		Where(makePredicates(filter)...).
		Exec(ctx)
	if err != nil {
		r.logger.Error("delete carts error", zap.Error(err))
		return errorx.ErrInternal
	}
	return nil
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
		WithProducts(func(pq *ent.ProductQuery) {
			pq.WithProductImages()
		}).
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
	if len(filter.ID.Eq) > 0 {
		predicates = append(predicates, entcart.IDIn(filter.ID.Eq.ToInt64()...))
	}
	if len(filter.ID.Neq) > 0 {
		predicates = append(predicates, entcart.IDNotIn(filter.ID.Neq.ToInt64()...))
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
		Product: mapDomainProductFromRow(row.Edges.Products),
	}
}

func mapDomainCartsFromRows(rows ent.Carts) cart.Carts {
	result := make(cart.Carts, 0, len(rows))
	for _, row := range rows {
		result = append(result, mapDomainCartFromRow(row))
	}
	return result
}

func mapDomainProductFromRow(row *ent.Product) *product.Product {
	return &product.Product{
		ID:        product.ID(row.ID),
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
		Images:    mapDomainImagesFromRows(row.Edges.ProductImages),
		Props: &product.Props{
			CategoryID:  category.ID(row.CategoryID),
			Name:        row.Name,
			Description: row.Description,
			Amount:      row.Amount,
			Price:       row.Price,
		},
	}
}

func mapDomainImageFromRow(row *ent.ProductImage) *product.Image {
	return &product.Image{
		ID:        product.ImageID(row.ID),
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
		Props: &product.ImageProps{
			ProductID: product.ID(row.ProductID),
			Name:      row.Name,
		},
	}
}

func mapDomainImagesFromRows(rows ent.ProductImages) product.Images {
	result := make(product.Images, 0, len(rows))
	for _, row := range rows {
		result = append(result, mapDomainImageFromRow(row))
	}
	return result
}
