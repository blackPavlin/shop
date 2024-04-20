package product

import "time"

// ImageID represents an id for product image.
type ImageID int64

// ImageIDs represents a slice of ImageID.
type ImageIDs []ImageID

// Image represents the product image.
type Image struct {
	ID        ImageID
	CreatedAt time.Time
	UpdatedAt time.Time
	Props     *ImageProps
}

// Images represents slice of Product.
type Images []*Image

// ImageProps represents product image editable fields.
type ImageProps struct {
	ProductID ID
	Name      string
}

// ToInt64 convert ImageID to int64.
func (id ImageID) ToInt64() int64 {
	return int64(id)
}

// ToInt64 convert slice of ImageIDs to slice int64.
func (ids ImageIDs) ToInt64() []int64 {
	result := make([]int64, 0, len(ids))
	for _, id := range ids {
		result = append(result, id.ToInt64())
	}
	return result
}

// IDs convert slice of Images to slice ImageIDs.
func (ii Images) IDs() ImageIDs {
	result := make(ImageIDs, 0, len(ii))
	for _, i := range ii {
		result = append(result, i.ID)
	}
	return result
}

// Names convert slice of Images to slice []string.
func (ii Images) Names() []string {
	result := make([]string, 0, len(ii))
	for _, i := range ii {
		result = append(result, i.Props.Name)
	}
	return result
}
