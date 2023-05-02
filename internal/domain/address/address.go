package address

import "github.com/blackPavlin/shop/internal/domain/user"

// ID
type ID int64

// IDs
type IDs []ID

// Address
type Address struct {
	ID     ID
	UserID user.ID

	Props *Props
}

// Addresses
type Addresses []*Address

// Props
type Props struct {
	City     string
	Country  string
	Flat     *int
	House    int
	Letter   *string
	Postcode int
	Street   string
}

// ToInt64
func (ids IDs) ToInt64() []int64 {
	result := make([]int64, 0, len(ids))
	for _, id := range ids {
		result = append(result, int64(id))
	}
	return result
}
