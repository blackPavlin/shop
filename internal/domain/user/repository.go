package user

import "context"

//go:generate mockgen -source $GOFILE -destination "repository_mock.go" -package "user"

// Repository
type Repository interface {
	Create(ctx context.Context, props *Props) (*User, error)
	Get(ctx context.Context, filter *Filter) (*User, error)
	Exist(ctx context.Context, filter *Filter) (bool, error)
}

// QueryCriteria
type QueryCriteria struct {
	Filter *Filter
}

// Filter
type Filter struct {
	Email string
}
