package product

import (
	"context"

	"github.com/blackPavlin/shop/internal/domain/category"
)

//go:generate mockgen -source $GOFILE -destination "repository_mock.go" -package "product"

// Repository
type Repository interface {
	CreateTx(ctx context.Context, props *Props) (*Product, error)
	Get(ctx context.Context, filter *Filter) (*Product, error)
	Query(ctx context.Context, criteria *QueryCriteria) (*QueryResult, error)
}

// QueryCriteria
type QueryCriteria struct {
	Filter     Filter
	Ordering   Ordering
	Pagination Pagination
}

// QueryResult
type QueryResult struct {
	Items Products
	Count int
}

// Filter
type Filter struct {
	ID         IDFilter
	CategoryID category.IDFilter
}

// IDFilter
type IDFilter struct {
	Eq  IDs
	Neq IDs
}

// Ordering
type Ordering struct{}

// Pagination
type Pagination struct {
	Limit  int
	Offset int
}
