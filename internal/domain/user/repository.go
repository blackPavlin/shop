package user

import "context"

//go:generate go run go.uber.org/mock/mockgen@v0.4.0 -source $GOFILE -destination "repository_mock.go" -package "user"

// Repository represents user repository.
type Repository interface {
	Get(ctx context.Context, filter *Filter) (*User, error)
	Create(ctx context.Context, props *Props) (*User, error)
	Exist(ctx context.Context, filter *Filter) (bool, error)
}

// QueryCriteria represents a criteria for user query.
type QueryCriteria struct {
	Filter Filter
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
