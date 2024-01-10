package product

import "context"

//go:generate go run go.uber.org/mock/mockgen@v0.4.0 -source $GOFILE -destination "image_repository_mock.go" -package "product"

// ImageRepository represents product image repository.
type ImageRepository interface {
	BulkCreateTx(ctx context.Context, images Images) (Images, error)
	Get(ctx context.Context, filter *ImageFilter) (*Image, error)
	Query(ctx context.Context, criteria *ImageQueryCriteria) (Images, error)
	Delete(ctx context.Context, filter *ImageFilter) error
}

// ImageQueryCriteria represents a criteria for image product query.
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
