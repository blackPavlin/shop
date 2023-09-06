// Package order contains user order oriented logic.
package order

import (
	"time"

	"github.com/blackPavlin/shop/internal/domain/user"
)

//go:generate enumer -sql -linecomment -type Status -trimprefix Status -output status_string.go order.go

const (
	StatusCreated  Status = iota // CREATED
	StatusAccepted               // ACCEPTED
	StatusCanceled               // CANCELED
)

// ID represents an id for order.
type ID int64

// IDs defines a slice of ID.
type IDs []ID

// Status represents order's status.
type Status uint8

// Order represents the order.
type Order struct {
	ID ID

	CreatedAt time.Time
	UpdatedAt time.Time

	Props *Props
}

// Orders slice of Order.
type Orders []*Order

// Props represents order editable fields.
type Props struct {
	UserID user.ID
	Status Status
}

// ToInt64 convert slice of IDs to slice int64.
func (ids IDs) ToInt64() []int64 {
	result := make([]int64, 0, len(ids))
	for _, id := range ids {
		result = append(result, int64(id))
	}
	return result
}
