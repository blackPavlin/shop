package product

import "time"

// ImageID represents an id for product image.
type ImageID int64

// ImageIDs defines a slice of ImageID.
type ImageIDs []ImageID

// Image represents the product image.
type Image struct {
	ID ImageID

	CreatedAt time.Time
	UpdatedAt time.Time

	Props *ImageProps
}

// Images slice of Product.
type Images []*Image

// ImageProps represents product image editable fields.
type ImageProps struct {
	ProductID ID
	Name      string
}

// ToInt64 convert slice of IDs to slice int64.
func (ids ImageIDs) ToInt64() []int64 {
	result := make([]int64, 0, len(ids))
	for _, id := range ids {
		result = append(result, int64(id))
	}
	return result
}
