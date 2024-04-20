// Package category contains product category oriented logic.
package category

import "time"

// ID represents an id for category.
type ID int64

// IDs represents a slice of ID.
type IDs []ID

// Category represents the category.
type Category struct {
	ID        ID
	CreatedAt time.Time
	UpdatedAt time.Time
	Props     *Props
}

// Categories represents slice of Category.
type Categories []*Category

// Props represents category editable fields.
type Props struct {
	Name string
}

// ToInt64 convert slice of IDs to slice int64.
func (ids IDs) ToInt64() []int64 {
	result := make([]int64, 0, len(ids))
	for _, id := range ids {
		result = append(result, int64(id))
	}
	return result
}
