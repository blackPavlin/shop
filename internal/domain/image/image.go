package image

import (
	"time"
)

// ID represents an id for image.
type ID int64

// IDs defines a slice of ID.
type IDs []ID

// Image represents the image.
type Image struct {
	ID ID

	CreatedAt time.Time
	UpdatedAt time.Time

	Props *Props
}

// Images slice Image.
type Images []*Image

// Props represents image editable fields.
type Props struct {
	Name         string
	OriginalName string
}

// ToInt64 convert slice of IDs to slice int64.
func (ids IDs) ToInt64() []int64 {
	result := make([]int64, 0, len(ids))
	for _, id := range ids {
		result = append(result, int64(id))
	}
	return result
}
