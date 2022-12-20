package address

import "github.com/blackPavlin/shop/internal/domain/user"

// ID
type ID int64

// Address
type Address struct {
	ID     ID
	UserID user.ID

	Props *Props
}

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

type Addresses []*Address
