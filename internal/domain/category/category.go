package category

import "time"

// ID
type ID int64

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
