package category

import "context"

//go:generate mockgen -source $GOFILE -destination "repository_mock.go" -package "category"

// Repository
type Repository interface {
	Create(ctx context.Context, props *Props) (*Category, error)
	Query(ctx context.Context, criteria *QueryCriteria) (Categories, error)
	Get(ctx context.Context, filter *Filter) (*Category, error)
}

// QueryCriteria
type QueryCriteria struct {
	Filter *Filter
}

// Filter
type Filter struct{}
