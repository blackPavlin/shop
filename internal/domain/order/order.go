package order

import "time"

//go:generate enumer -sql -linecomment -type Status -trimprefix Status -output status_string.go order.go

const (
	StatusCreated  Status = iota // CREATED
	StatusAccepted               // ACCEPTED
	StatusCanceled               // CANCELED
)

// ID
type ID int64

// IDs
type IDs []ID

// Status
type Status uint8

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
