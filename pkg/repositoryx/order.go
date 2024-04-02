package repositoryx

import "entgo.io/ent/dialect/sql"

// Order present sql order ACS or DESC.
type Order uint8

// Allowed ordering.
const (
	OrderAsc  Order = iota // ASC
	OrderDesc              // DESC
)

// ToString converts Order to string for given field name.
func (o Order) ToString(field string) string {
	switch o {
	case OrderAsc:
		return sql.Asc(field)
	case OrderDesc:
		return sql.Desc(field)
	default:
		return sql.Asc(field)
	}
}
