package image

import (
	"time"
)

// ID
type ID int64

// IDs
type IDs []ID

// Image
type Image struct {
	ID ID

	CreatedAt time.Time
	UpdatedAt time.Time

	Props *Props
}

// Images
type Images []*Image

// Props
type Props struct {
	Name         string
	OriginalName string
}

// ToInt64
func (ids IDs) ToInt64() []int64 {
	result := make([]int64, 0, len(ids))
	for _, id := range ids {
		result = append(result, int64(id))
	}
	return result
}
