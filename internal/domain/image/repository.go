package image

import "context"

//go:generate mockgen -source $GOFILE -destination "repository_mock.go" -package "image"

// Repository
type Repository interface {
	BulkCreateTx(ctx context.Context, images Images) (Images, error)
}

// QueryCriteria
type QueryCriteria struct {
	Filter Filter
}

// Filter
type Filter struct {
	ID IDFilter
}

// IDFilter
type IDFilter struct {
	Eq  IDs
	Neq IDs
}
