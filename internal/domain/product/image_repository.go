package product

import "context"

//go:generate mockgen -source $GOFILE -destination "image_repository_mock.go" -package "product"

// ImageRepository represents product image repository.
type ImageRepository interface {
	BulkCreateTx(ctx context.Context, images Images) (Images, error)
	Get(ctx context.Context, filter *ImageFilter) (*Image, error)
	Query(ctx context.Context, criteria *ImageQueryCriteria) (Images, error)
	DeleteTx(ctx context.Context, imageID ImageID) error
}

// ImageQueryCriteria  represents a criteria for image product query.
type ImageQueryCriteria struct {
	Filter ImageFilter
}

// ImageFilter represents product image filter.
type ImageFilter struct {
	ImageID   ImageIDFilter
	ProductID IDFilter
}

// ImageIDFilter represents ImageID filter.
type ImageIDFilter struct {
	Eq  ImageIDs
	Neq ImageIDs
}
