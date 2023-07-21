// Package address contains user address oriented logic.
package address

import "github.com/blackPavlin/shop/internal/domain/user"

// ID represents an id for address.
type ID int64

// IDs defines a slice of ID.
type IDs []ID

// Address represents the address.
type Address struct {
	ID     ID
	UserID user.ID

	Props *Props
}

// Addresses slice of Address.
type Addresses []*Address

// Props represents address editable fields.
type Props struct {
	City     string
	Country  string
	Flat     *int
	House    int
	Letter   *string
	Postcode int
	Street   string
}

// ToInt64 convert slice of IDs to slice int64.
func (ids IDs) ToInt64() []int64 {
	result := make([]int64, 0, len(ids))
	for _, id := range ids {
		result = append(result, int64(id))
	}
	return result
}
