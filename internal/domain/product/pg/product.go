package pg

import (
	"context"

	"github.com/blackPavlin/shop/ent"
	"github.com/blackPavlin/shop/ent/predicate"
	"github.com/blackPavlin/shop/internal/domain/product"
	"github.com/blackPavlin/shop/pkg/errorx"
	"go.uber.org/zap"
)

// ProductRepository
type ProductRepository struct {
	client *ent.Client
	logger *zap.Logger
}

// NewProductRepository
func NewProductRepository(client *ent.Client, logger *zap.Logger) *ProductRepository {
	return &ProductRepository{client: client, logger: logger}
}

// Create
func (r *ProductRepository) Create(
	ctx context.Context,
	props *product.Props,
) (*product.Product, error) {
	row, err := r.client.Product.Create().
		SetCategoryID(int64(props.CategoryID)).
		SetName(props.Name).
		SetDescription(props.Description).
		SetAmount(props.Amount).
		SetPrice(props.Price).
		Save(ctx)
	if err != nil {
		r.logger.Error("create product error:", zap.Error(err))
		return nil, errorx.ErrInternal
	}
	return mapDomainProductFromRow(row), nil
}

// Query
func (r *ProductRepository) Query(
	ctx context.Context,
	criteria *product.QueryCriteria,
) (*product.QueryResult, error) {
	predicates := makePredicates(criteria)
	rows, err := r.client.Product.Query().
		Where(predicates...).
		Limit(criteria.Pagination.Limit).
		Offset(criteria.Pagination.Offset).
		All(ctx)
	if err != nil {
		r.logger.Error("query products count error:", zap.Error(err))
		return nil, errorx.ErrInternal
	}
	count, err := r.client.Product.Query().
		Where(predicates...).
		Count(ctx)
	if err != nil {
		r.logger.Error("get products count error:", zap.Error(err))
		return nil, errorx.ErrInternal
	}
	result := &product.QueryResult{
		Items: mapDomainProductsFromRows(rows),
		Count: count,
	}
	return result, nil
}

func makePredicates(criteria *product.QueryCriteria) []predicate.Product {
	predicates := make([]predicate.Product, 0)
	// TODO
	return predicates
}

func makeOrderings() {}

func mapDomainProductFromRow(row *ent.Product) *product.Product {
	return &product.Product{
		ID: product.ID(row.ID),
	}
}

func mapDomainProductsFromRows(rows ent.Products) product.Products {
	result := make(product.Products, 0, len(rows))
	for _, row := range rows {
		result = append(result, mapDomainProductFromRow(row))
	}
	return result
}
