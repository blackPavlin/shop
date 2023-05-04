package order

import "time"

// ID
type ID int64

// IDs
type IDs []ID

// Order
type Order struct {
	ID ID

	CreatedAt time.Time
	UpdatedAt time.Time

	Props *Props
}

// Orders
type Orders []Order

// Props
type Props struct{}
