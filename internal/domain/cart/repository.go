package cart

import (
	"context"

	"github.com/blackPavlin/shop/internal/domain/user"
)

//go:generate mockgen -source $GOFILE -destination "repository_mock.go" -package "cart"

// Repository
type Repository interface {
	Create(ctx context.Context, props *Props) (*Cart, error)
	Query(ctx context.Context, criteria *QueryCriteria) (*QueryResult, error)
}

// QueryCriteria
type QueryCriteria struct {
	Filter *Filter
}

// Filter
type Filter struct {
	UserID user.ID
}

// QueryResult
type QueryResult struct {
	Items Carts
}
