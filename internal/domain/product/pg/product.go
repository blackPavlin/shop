package pg

import (
	"context"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/blackPavlin/shop/ent"
	"github.com/blackPavlin/shop/ent/predicate"
	entproduct "github.com/blackPavlin/shop/ent/product"
	"github.com/blackPavlin/shop/internal/domain/category"
	"github.com/blackPavlin/shop/internal/domain/product"
	"github.com/blackPavlin/shop/pkg/errorx"
	"github.com/blackPavlin/shop/pkg/repositoryx/pg"
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
		if pg.IsForeignKeyViolationErr(err, "product_category_fk") {
			return nil, errorx.NewNotFoundError("category not found")
		}
		r.logger.Error("create product error:", zap.Error(err))
		return nil, errorx.ErrInternal
	}
	return mapDomainProductFromRow(row), nil
}

// Get
func (r *ProductRepository) Get(ctx context.Context, filter *product.Filter) (*product.Product, error) {
	row, err := r.client.Product.
		Query().
		Where(makePredicates(filter)...).
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errorx.NewNotFoundError("product not found")
		}
		r.logger.Error("get product error", zap.Error(err))
		return nil, errorx.ErrInternal
	}
	return mapDomainProductFromRow(row), err
}

// Query
func (r *ProductRepository) Query(
	ctx context.Context,
	criteria *product.QueryCriteria,
) (*product.QueryResult, error) {
	var (
		rows  ent.Products
		count int
		err   error
	)
	predicates := makePredicates(&criteria.Filter)
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		rows, err = r.client.Product.Query().
			Where(predicates...).
			Limit(criteria.Pagination.Limit).
			Offset(criteria.Pagination.Offset).
			All(ctx)
		if err != nil {
			r.logger.Error("query products count error:", zap.Error(err))
			return err
		}
		return nil
	})
	g.Go(func() error {
		count, err = r.client.Product.Query().
			Where(predicates...).
			Count(ctx)
		if err != nil {
			r.logger.Error("get products count error:", zap.Error(err))
			return err
		}
		return nil
	})
	if err = g.Wait(); err != nil {
		return nil, errorx.ErrInternal
	}
	result := &product.QueryResult{
		Items: mapDomainProductsFromRows(rows),
		Count: count,
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
		predicates = append(predicates, entproduct.CategoryIDIn(filter.CategoryID.Eq.ToInt64()...))
	}
	if len(filter.CategoryID.Neq) > 0 {
		predicates = append(predicates, entproduct.CategoryIDNotIn(filter.CategoryID.Neq.ToInt64()...))
	}
	return predicates
}

func makeOrderings() {}

func mapDomainProductFromRow(row *ent.Product) *product.Product {
	return &product.Product{
		ID:        product.ID(row.ID),
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
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
