package user

import "context"

//go:generate mockgen -source $GOFILE -destination "repository_mock.go" -package "user"

// Repository represents user repository.
type Repository interface {
	Create(ctx context.Context, props *Props) (*User, error)
	Get(ctx context.Context, filter *Filter) (*User, error)
	Exist(ctx context.Context, filter *Filter) (bool, error)
}

// QueryCriteria represents a criteria for user query.
type QueryCriteria struct {
	Filter *Filter
}

// Filter represents user filter.
type Filter struct {
	ID    IDFilter
	Email EmailFilter
}

// IDFilter represents ID filter.
type IDFilter struct {
	Eq  IDs
	Neq IDs
}

// EmailFilter represents an email filter.
type EmailFilter struct {
	Eq  []string
	Neq []string
}
