package order

import "time"

// ID
type ID int64

// Order
type Order struct {
	ID ID

	CreatedAt time.Time
	UpdatedAt time.Time

	Props *Props
}

// Props
type Props struct {}
