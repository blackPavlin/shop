package order

import (
	"context"

	"github.com/blackPavlin/shop/internal/domain/user"
)

//go:generate mockgen -source $GOFILE -destination "repository_mock.go" -package "order"

// Repository
type Repository interface {
	Create(ctx context.Context) (*Order, error)
	Query(ctx context.Context, criteria *QueryCriteria) (*QueryResult, error)
}

// QueryCriteria
type QueryCriteria struct {
	Filter Filter
}

// Filter
type Filter struct {
	ID     IDFilter
	UserID user.IDFilter
}

// IDFilter
type IDFilter struct {
	Eq  IDs
	Neq IDs
}

// QueryResult
type QueryResult struct {
	Items Orders
}
