package order

import (
	"context"

	"github.com/blackPavlin/shop/internal/domain/user"
)

//go:generate mockgen -source $GOFILE -destination "repository_mock.go" -package "order"

// Repository represents order repository.
type Repository interface {
	Create(ctx context.Context) (*Order, error)
	Query(ctx context.Context, criteria *QueryCriteria) (*QueryResult, error)
}

// QueryCriteria represents criteria for service query.
type QueryCriteria struct {
	Filter Filter
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
}
