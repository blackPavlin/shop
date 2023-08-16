package product

import (
	"context"

	"github.com/blackPavlin/shop/internal/domain/category"
)

//go:generate mockgen -source $GOFILE -destination "repository_mock.go" -package "product"

// Repository represents product repository.
type Repository interface {
	CreateTx(ctx context.Context, props *Props) (*Product, error)
	Update(ctx context.Context, productID ID, props *Props) (*Product, error)
	Get(ctx context.Context, filter *Filter) (*Product, error)
	Query(ctx context.Context, criteria *QueryCriteria) (*QueryResult, error)
}

// QueryCriteria represents criteria for product query.
type QueryCriteria struct {
	Filter     Filter
	Ordering   Ordering
	Pagination Pagination
}

// QueryResult represents a result for products query.
type QueryResult struct {
	Items Products
	Count int
}

// Filter represents product filter.
type Filter struct {
	ID         IDFilter
	CategoryID category.IDFilter
}

// IDFilter represents ID filter.
type IDFilter struct {
	Eq  IDs
	Neq IDs
}

// Ordering represents order criteria.
type Ordering struct{}

// Pagination is a pagination parameters.
type Pagination struct {
	Limit  uint64
	Offset uint64
}
