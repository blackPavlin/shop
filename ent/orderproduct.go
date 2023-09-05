// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/blackPavlin/shop/ent/order"
	"github.com/blackPavlin/shop/ent/orderproduct"
	"github.com/blackPavlin/shop/ent/product"
)

// OrderProduct is the model entity for the OrderProduct schema.
type OrderProduct struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// OrderID holds the value of the "order_id" field.
	OrderID int64 `json:"order_id,omitempty"`
	// ProductID holds the value of the "product_id" field.
	ProductID int64 `json:"product_id,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount int64 `json:"amount,omitempty"`
	// Price holds the value of the "price" field.
	Price int64 `json:"price,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrderProductQuery when eager-loading is set.
	Edges        OrderProductEdges `json:"edges"`
	selectValues sql.SelectValues
}

// OrderProductEdges holds the relations/edges for other nodes in the graph.
type OrderProductEdges struct {
	// Orders holds the value of the orders edge.
	Orders *Order `json:"orders,omitempty"`
	// Products holds the value of the products edge.
	Products *Product `json:"products,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// OrdersOrErr returns the Orders value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderProductEdges) OrdersOrErr() (*Order, error) {
	if e.loadedTypes[0] {
		if e.Orders == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: order.Label}
		}
		return e.Orders, nil
	}
	return nil, &NotLoadedError{edge: "orders"}
}

// ProductsOrErr returns the Products value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderProductEdges) ProductsOrErr() (*Product, error) {
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
func (*OrderProduct) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case orderproduct.FieldID, orderproduct.FieldOrderID, orderproduct.FieldProductID, orderproduct.FieldAmount, orderproduct.FieldPrice:
			values[i] = new(sql.NullInt64)
		case orderproduct.FieldCreatedAt, orderproduct.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrderProduct fields.
func (op *OrderProduct) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case orderproduct.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			op.ID = int64(value.Int64)
		case orderproduct.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				op.CreatedAt = value.Time
			}
		case orderproduct.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				op.UpdatedAt = value.Time
			}
		case orderproduct.FieldOrderID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_id", values[i])
			} else if value.Valid {
				op.OrderID = value.Int64
			}
		case orderproduct.FieldProductID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field product_id", values[i])
			} else if value.Valid {
				op.ProductID = value.Int64
			}
		case orderproduct.FieldAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				op.Amount = value.Int64
			}
		case orderproduct.FieldPrice:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				op.Price = value.Int64
			}
		default:
			op.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the OrderProduct.
// This includes values selected through modifiers, order, etc.
func (op *OrderProduct) Value(name string) (ent.Value, error) {
	return op.selectValues.Get(name)
}

// QueryOrders queries the "orders" edge of the OrderProduct entity.
func (op *OrderProduct) QueryOrders() *OrderQuery {
	return NewOrderProductClient(op.config).QueryOrders(op)
}

// QueryProducts queries the "products" edge of the OrderProduct entity.
func (op *OrderProduct) QueryProducts() *ProductQuery {
	return NewOrderProductClient(op.config).QueryProducts(op)
}

// Update returns a builder for updating this OrderProduct.
// Note that you need to call OrderProduct.Unwrap() before calling this method if this OrderProduct
// was returned from a transaction, and the transaction was committed or rolled back.
func (op *OrderProduct) Update() *OrderProductUpdateOne {
	return NewOrderProductClient(op.config).UpdateOne(op)
}

// Unwrap unwraps the OrderProduct entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (op *OrderProduct) Unwrap() *OrderProduct {
	_tx, ok := op.config.driver.(*txDriver)
	if !ok {
		panic("ent: OrderProduct is not a transactional entity")
	}
	op.config.driver = _tx.drv
	return op
}

// String implements the fmt.Stringer.
func (op *OrderProduct) String() string {
	var builder strings.Builder
	builder.WriteString("OrderProduct(")
	builder.WriteString(fmt.Sprintf("id=%v, ", op.ID))
	builder.WriteString("created_at=")
	builder.WriteString(op.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(op.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("order_id=")
	builder.WriteString(fmt.Sprintf("%v", op.OrderID))
	builder.WriteString(", ")
	builder.WriteString("product_id=")
	builder.WriteString(fmt.Sprintf("%v", op.ProductID))
	builder.WriteString(", ")
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", op.Amount))
	builder.WriteString(", ")
	builder.WriteString("price=")
	builder.WriteString(fmt.Sprintf("%v", op.Price))
	builder.WriteByte(')')
	return builder.String()
}

// OrderProducts is a parsable slice of OrderProduct.
type OrderProducts []*OrderProduct
