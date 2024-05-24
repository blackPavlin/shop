package address

import (
	"context"

	"github.com/blackPavlin/shop/internal/domain/user"
)

//go:generate go run go.uber.org/mock/mockgen@v0.4.0 -source $GOFILE -destination "repository_mock.go" -package "address"

// Repository represents address repository.
type Repository interface {
	Get(ctx context.Context, filter *Filter) (*Address, error)
	Query(ctx context.Context, criteria *QueryCriteria) (Addresses, error)
	Create(ctx context.Context, props *Props) (*Address, error)
}

// QueryCriteria represents a criteria for address query.
type QueryCriteria struct {
	Filter Filter
}

// Filter represents address filter.
type Filter struct {
	ID     IDFilter
	UserID user.IDFilter
}

// IDFilter represents ID filter.
type IDFilter struct {
	Eq  IDs
	Neq IDs
}
