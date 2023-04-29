package category

import "context"

//go:generate mockgen -source $GOFILE -destination "repository_mock.go" -package "category"

// Repository
type Repository interface {
	Create(ctx context.Context, props *Props) (*Category, error)
	Query(ctx context.Context, criteria *QueryCriteria) (Categories, error)
	Get(ctx context.Context, filter *Filter) (*Category, error)
	Update(ctx context.Context, id ID, props *Props) (*Category, error)
	Delete(ctx context.Context, id ID) error
}

// QueryCriteria
type QueryCriteria struct {
	Filter Filter
}

// Filter
type Filter struct {
	ID   IDFilter
	Name NameFilter
}

// IDFilter
type IDFilter struct {
	Eq  IDs
	Neq IDs
}

// NameFilter
type NameFilter struct {
	Eq  []string
	Neq []string
}
