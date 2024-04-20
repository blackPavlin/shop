// Package order contains user order oriented logic.
package order

import (
	"time"

	"github.com/blackPavlin/shop/internal/domain/user"
)

//go:generate go run github.com/dmarkham/enumer@v1.5.9 -sql -linecomment -type Status -trimprefix Status -output status_string.go order.go

const (
	StatusCreated  Status = iota + 1 // CREATED
	StatusAccepted                   // ACCEPTED
	StatusCanceled                   // CANCELED
)

// ID represents an id for order.
type ID int64

// IDs represents a slice of ID.
type IDs []ID

// Status represents order's status.
type Status uint8

// Order represents the order.
type Order struct {
	ID        ID
	UserID    user.ID
	CreatedAt time.Time
	UpdatedAt time.Time
	Props     *Props
}

// Orders represents slice of Order.
type Orders []*Order

// Props represents order editable fields.
type Props struct {
	Status Status
}

// ToInt64 convert ID to int64.
func (id ID) ToInt64() int64 {
	return int64(id)
}

// ToInt64 convert slice of IDs to slice int64.
func (ids IDs) ToInt64() []int64 {
	result := make([]int64, 0, len(ids))
	for _, id := range ids {
		result = append(result, id.ToInt64())
	}
	return result
}
