package product

import "context"

//go:generate mockgen -source $GOFILE -destination "product_image_repository_mock.go" -package "product"

// ImageRepository represents product image repository.
type ImageRepository interface {
	BulkCreateTx(ctx context.Context, images Images) (Images, error)
	Query(ctx context.Context, criteria *ImageQueryCriteria) (Images, error)
}

// ImageQueryCriteria  represents a criteria for image product query.
type ImageQueryCriteria struct {
	Filter ImageFilter
}

// ImageFilter represents product image filter.
type ImageFilter struct {
	ProductID IDFilter
}
