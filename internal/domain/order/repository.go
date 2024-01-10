package order

import (
	"context"

	"github.com/blackPavlin/shop/internal/domain/user"
	"github.com/blackPavlin/shop/pkg/paging"
)

//go:generate go run go.uber.org/mock/mockgen@v0.4.0 -source $GOFILE -destination "repository_mock.go" -package "order"

// Repository represents order repository.
type Repository interface {
	Create(ctx context.Context, props *Props) (*Order, error)
	Get(ctx context.Context, filter *Filter) (*Order, error)
	Query(ctx context.Context, criteria *QueryCriteria) (*QueryResult, error)
}

// QueryCriteria represents criteria for service query.
type QueryCriteria struct {
	Filter     Filter
	Pagination paging.Pagination
}

// Filter represents order filter.
type Filter struct {
	ID     IDFilter
	UserID user.IDFilter
}

// IDFilter represents ID filter.
type IDFilter struct {
	Eq  IDs
	Neq IDs
}

// QueryResult represents a result for order query.
type QueryResult struct {
	Items Orders
	Count int
}
