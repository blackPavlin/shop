// Package pg contains implementations for product repositories
package pg

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/blackPavlin/shop/internal/database/ent"
	"github.com/blackPavlin/shop/internal/database/ent/predicate"
	entproduct "github.com/blackPavlin/shop/internal/database/ent/product"
	"github.com/blackPavlin/shop/internal/domain/category"
	"github.com/blackPavlin/shop/internal/domain/product"
	"github.com/blackPavlin/shop/pkg/errorx"
)

// ProductRepository pg repository implementation.
type ProductRepository struct {
	client *ent.Client
	logger *zap.Logger
}

// NewProductRepository create instance of ProductRepository.
func NewProductRepository(client *ent.Client, logger *zap.Logger) *ProductRepository {
	return &ProductRepository{client: client, logger: logger}
}

// Create product in db.
func (r ProductRepository) Create(
	ctx context.Context,
	props *product.Props,
) (*product.Product, error) {
	client := r.client
	if tx := ent.TxFromContext(ctx); tx != nil {
		client = tx.Client()
	}
	row, err := client.Product.Create().
		SetCategoryID(int64(props.CategoryID)).
		SetName(props.Name).
		SetDescription(props.Description).
		SetAmount(props.Amount).
		SetPrice(props.Price).
		Save(ctx)
	if err != nil {
		r.logger.Error("create product error", zap.Error(err))
		return nil, errorx.ErrInternal
	}
	return mapDomainProductFromRow(row), nil
}

// Update product in db.
func (r ProductRepository) Update(
	ctx context.Context,
	productID product.ID,
	props *product.Props,
) (*product.Product, error) {
	row, err := r.client.Product.UpdateOneID(int64(productID)).
		SetCategoryID(int64(props.CategoryID)).
		SetName(props.Name).
		SetDescription(props.Description).
		SetAmount(props.Amount).
		SetPrice(props.Price).
		Save(ctx)
	if err != nil {
		if ent.IsConstraintError(err) {
			return nil, errorx.ErrNotFound
		}
		r.logger.Error("update product error", zap.Error(err))
		return nil, errorx.ErrInternal
	}
	return mapDomainProductFromRow(row), nil
}

// Delete product in db.
func (r ProductRepository) Delete(ctx context.Context, filter *product.Filter) error {
	client := r.client
	if tx := ent.TxFromContext(ctx); tx != nil {
		client = tx.Client()
	}
	_, err := client.Product.
		Delete().
		Where(makePredicates(filter)...).
		Exec(ctx)
	if err != nil {
		r.logger.Error("delete product error", zap.Error(err))
		return errorx.ErrInternal
	}
	return nil
}

// Get product from db.
func (r ProductRepository) Get(ctx context.Context, filter *product.Filter) (*product.Product, error) {
	row, err := r.client.Product.
		Query().
		Where(makePredicates(filter)...).
		WithProductImages().
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errorx.ErrNotFound
		}
		r.logger.Error("get product error", zap.Error(err))
		return nil, errorx.ErrInternal
	}
	return mapDomainProductFromRow(row), err
}

// Query products from db.
func (r ProductRepository) Query(
	ctx context.Context,
	criteria *product.QueryCriteria,
) (*product.QueryResult, error) {
	var (
		result = &product.QueryResult{}
		err    error
	)
	predicates := makePredicates(&criteria.Filter)
	group, groupCtx := errgroup.WithContext(ctx)
	group.Go(func() error {
		rows, err := r.client.Product.Query().
			Where(predicates...).
			WithProductImages().
			Limit(int(criteria.Pagination.Limit)).
			Offset(int(criteria.Pagination.Offset)).
			Order(func(selector *sql.Selector) {
				selector.OrderBy(makeOrdering(&criteria.Ordering))
			}).
			All(groupCtx)
		if err != nil {
			r.logger.Error("query products count error:", zap.Error(err))
			return errorx.ErrInternal
		}
		result.Items = mapDomainProductsFromRows(rows)
		return nil
	})
	group.Go(func() error {
		result.Count, err = r.client.Product.Query().
			Where(predicates...).
			Count(groupCtx)
		if err != nil {
			r.logger.Error("get products count error", zap.Error(err))
			return errorx.ErrInternal
		}
		return nil
	})
	if err = group.Wait(); err != nil {
		return nil, errorx.ErrInternal
	}
	return result, nil
}

func makePredicates(filter *product.Filter) []predicate.Product {
	predicates := make([]predicate.Product, 0)
	if len(filter.ID.Eq) > 0 {
		predicates = append(predicates, entproduct.IDIn(filter.ID.Eq.ToInt64()...))
	}
	if len(filter.ID.Neq) > 0 {
		predicates = append(predicates, entproduct.IDNotIn(filter.ID.Neq.ToInt64()...))
	}
	if len(filter.CategoryID.Eq) > 0 {
		predicates = append(
			predicates,
			entproduct.CategoryIDIn(filter.CategoryID.Eq.ToInt64()...),
		)
	}
	if len(filter.CategoryID.Neq) > 0 {
		predicates = append(
			predicates,
			entproduct.CategoryIDNotIn(filter.CategoryID.Neq.ToInt64()...),
		)
	}
	return predicates
}

func makeOrdering(ordering *product.Ordering) string {
	if ordering.Price != nil {
		return ordering.Price.ToString(entproduct.FieldPrice)
	}
	return sql.Asc(entproduct.FieldID)
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

func mapDomainProductsFromRows(rows ent.Products) product.Products {
	result := make(product.Products, 0, len(rows))
	for _, row := range rows {
		result = append(result, mapDomainProductFromRow(row))
	}
	return result
}
