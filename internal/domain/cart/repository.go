package cart

import (
	"context"

	"github.com/blackPavlin/shop/internal/domain/user"
)

//go:generate mockgen -source $GOFILE -destination "repository_mock.go" -package "cart"

// Repository represents cart repository.
type Repository interface {
	Get(ctx context.Context, filter *Filter) (*Cart, error)
	Save(ctx context.Context, props *Props) (*Cart, error)
	Query(ctx context.Context, criteria *QueryCriteria) (*QueryResult, error)
}

// QueryCriteria  represents a criteria for cart query.
type QueryCriteria struct {
	Filter Filter
}

// Filter represents cart filter.
type Filter struct {
	UserID user.IDFilter
}

// QueryResult represents a result for cart query.
type QueryResult struct {
	Items Carts
}
