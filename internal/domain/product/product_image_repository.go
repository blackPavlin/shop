package product

import "context"

//go:generate mockgen -source $GOFILE -destination "product_image_repository_mock.go" -package "product"

// ImageRepository
type ImageRepository interface {
	BulkCreateTx(ctx context.Context, images Images) (Images, error)
	Query(ctx context.Context, criteria *ImageQueryCriteria) (Images, error)
}

// ImageQueryCriteria
type ImageQueryCriteria struct {
	Filter ImageFilter
}

// ImageFilter
type ImageFilter struct {
	ProductID IDFilter
}
