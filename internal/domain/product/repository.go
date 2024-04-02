package product

import (
	"context"

	"github.com/blackPavlin/shop/internal/domain/category"
	"github.com/blackPavlin/shop/pkg/paging"
	"github.com/blackPavlin/shop/pkg/repositoryx"
)

//go:generate go run go.uber.org/mock/mockgen@v0.4.0 -source $GOFILE -destination "repository_mock.go" -package "product"

// Repository represents product repository.
type Repository interface {
	Create(ctx context.Context, props *Props) (*Product, error)
	Update(ctx context.Context, productID ID, props *Props) (*Product, error)
	Delete(ctx context.Context, filter *Filter) error
	Get(ctx context.Context, filter *Filter) (*Product, error)
	Query(ctx context.Context, criteria *QueryCriteria) (*QueryResult, error)
}

// QueryCriteria represents criteria for products query.
type QueryCriteria struct {
	Filter     Filter
	Ordering   Ordering
	Pagination paging.Pagination
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
type Ordering struct {
	Price *repositoryx.Order
}
