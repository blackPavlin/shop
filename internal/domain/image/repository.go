package image

import "context"

//go:generate mockgen -source $GOFILE -destination "repository_mock.go" -package "image"

// Repository represents image repository.
type Repository interface {
	Get(ctx context.Context, filter *Filter) (*Image, error)
	BulkCreateTx(ctx context.Context, images Images) (Images, error)
}

// QueryCriteria represents a criteria for image query.
type QueryCriteria struct {
	Filter Filter
}

// Filter represents image filter.
type Filter struct {
	ID IDFilter
}

// IDFilter represents ID filter.
type IDFilter struct {
	Eq  IDs
	Neq IDs
}
