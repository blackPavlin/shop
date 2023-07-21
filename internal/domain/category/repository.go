package category

import "context"

//go:generate mockgen -source $GOFILE -destination "repository_mock.go" -package "category"

// Repository represents category repository.
type Repository interface {
	Create(ctx context.Context, props *Props) (*Category, error)
	Query(ctx context.Context, criteria *QueryCriteria) (Categories, error)
	Get(ctx context.Context, filter *Filter) (*Category, error)
	Update(ctx context.Context, id ID, props *Props) (*Category, error)
	Delete(ctx context.Context, id ID) error
}

// QueryCriteria represents a criteria for category query.
type QueryCriteria struct {
	Filter Filter
}

// Filter represents category filter.
type Filter struct {
	ID   IDFilter
	Name NameFilter
}

// IDFilter represents ID filter.
type IDFilter struct {
	Eq  IDs
	Neq IDs
}

// NameFilter represents name filter.
type NameFilter struct {
	Eq  []string
	Neq []string
}
