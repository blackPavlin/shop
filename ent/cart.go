// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/blackPavlin/shop/ent/cart"
	"github.com/blackPavlin/shop/ent/product"
	"github.com/blackPavlin/shop/ent/user"
)

// Cart is the model entity for the Cart schema.
type Cart struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int64 `json:"user_id,omitempty"`
	// ProductID holds the value of the "product_id" field.
	ProductID int64 `json:"product_id,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount int64 `json:"amount,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CartQuery when eager-loading is set.
	Edges CartEdges `json:"edges"`
}

// CartEdges holds the relations/edges for other nodes in the graph.
type CartEdges struct {
	// Users holds the value of the users edge.
	Users *User `json:"users,omitempty"`
	// Products holds the value of the products edge.
	Products *Product `json:"products,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CartEdges) UsersOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Users == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// ProductsOrErr returns the Products value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CartEdges) ProductsOrErr() (*Product, error) {
	if e.loadedTypes[1] {
		if e.Products == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: product.Label}
		}
		return e.Products, nil
	}
	return nil, &NotLoadedError{edge: "products"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Cart) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case cart.FieldID, cart.FieldUserID, cart.FieldProductID, cart.FieldAmount:
			values[i] = new(sql.NullInt64)
		case cart.FieldCreatedAt, cart.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Cart", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Cart fields.
func (c *Cart) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case cart.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int64(value.Int64)
		case cart.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case cart.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case cart.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				c.UserID = value.Int64
			}
		case cart.FieldProductID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field product_id", values[i])
			} else if value.Valid {
				c.ProductID = value.Int64
			}
		case cart.FieldAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				c.Amount = value.Int64
			}
		}
	}
	return nil
}

// QueryUsers queries the "users" edge of the Cart entity.
func (c *Cart) QueryUsers() *UserQuery {
	return (&CartClient{config: c.config}).QueryUsers(c)
}

// QueryProducts queries the "products" edge of the Cart entity.
func (c *Cart) QueryProducts() *ProductQuery {
	return (&CartClient{config: c.config}).QueryProducts(c)
}

// Update returns a builder for updating this Cart.
// Note that you need to call Cart.Unwrap() before calling this method if this Cart
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Cart) Update() *CartUpdateOne {
	return (&CartClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Cart entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Cart) Unwrap() *Cart {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Cart is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Cart) String() string {
	var builder strings.Builder
	builder.WriteString("Cart(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", c.UserID))
	builder.WriteString(", ")
	builder.WriteString("product_id=")
	builder.WriteString(fmt.Sprintf("%v", c.ProductID))
	builder.WriteString(", ")
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", c.Amount))
	builder.WriteByte(')')
	return builder.String()
}

// Carts is a parsable slice of Cart.
type Carts []*Cart

func (c Carts) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
