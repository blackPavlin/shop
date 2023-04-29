package category

import "time"

// ID
type ID int64

// IDs
type IDs []ID

// Category
type Category struct {
	ID        ID
	
	CreatedAt time.Time
	UpdatedAt time.Time

	Props *Props
}

// Categories
type Categories []*Category

// Props
type Props struct {
	Name string
}

// ToInt64
func (ids IDs) ToInt64() []int64 {
	result := make([]int64, 0, len(ids))
	for _, id := range ids {
		result = append(result, int64(id))
	}
	return result
}
