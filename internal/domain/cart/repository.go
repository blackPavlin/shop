package cart

import (
	"context"

	"github.com/blackPavlin/shop/internal/domain/user"
)

//go:generate go run go.uber.org/mock/mockgen@v0.4.0 -source $GOFILE -destination "repository_mock.go" -package "cart"

// Repository represents cart repository.
type Repository interface {
	Get(ctx context.Context, filter *Filter) (*Cart, error)
	Query(ctx context.Context, criteria *QueryCriteria) (*QueryResult, error)
	Save(ctx context.Context, props *Props) (*Cart, error)
	Delete(ctx context.Context, filter *Filter) error
}

// QueryCriteria  represents a criteria for cart query.
type QueryCriteria struct {
	Filter Filter
}

// Filter represents cart filter.
type Filter struct {
	ID     IDFilter
	UserID user.IDFilter
}

// IDFilter represents ID filter.
type IDFilter struct {
	Eq  IDs
	Neq IDs
}

// QueryResult represents a result for cart query.
type QueryResult struct {
	Items Carts
}
