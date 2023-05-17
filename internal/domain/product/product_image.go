package product

import (
	"time"

	"github.com/blackPavlin/shop/internal/domain/image"
)

// ImageID
type ImageID int64

// ImageIDs
type ImageIDs []ImageID

// Image
type Image struct {
	ID ImageID

	CreatedAt time.Time
	UpdatedAt time.Time

	Props *ImageProps
}

// Images
type Images []*Image

// ImageProps
type ImageProps struct {
	ProductID ID
	ImageID   image.ID
}

// ToInt64
func (ids ImageIDs) ToInt64() []int64 {
	result := make([]int64, 0, len(ids))
	for _, id := range ids {
		result = append(result, int64(id))
	}
	return result
}
