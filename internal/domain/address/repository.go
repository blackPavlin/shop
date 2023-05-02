package address

import (
	"context"

	"github.com/blackPavlin/shop/internal/domain/user"
)

//go:generate mockgen -source $GOFILE -destination "repository_mock.go" -package "address"

// Repository
type Repository interface {
	Create(ctx context.Context, props *Props) (*Address, error)
	Get(ctx context.Context, filter *Filter) (*Address, error)
	Query(ctx context.Context, criteria *QueryCriteria) (Addresses, error)
}

// QueryCriteria
type QueryCriteria struct {
	Filter *Filter
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
