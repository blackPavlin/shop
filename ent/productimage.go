// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/blackPavlin/shop/ent/product"
	"github.com/blackPavlin/shop/ent/productimage"
)

// ProductImage is the model entity for the ProductImage schema.
type ProductImage struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// ProductID holds the value of the "product_id" field.
	ProductID int64 `json:"product_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProductImageQuery when eager-loading is set.
	Edges        ProductImageEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ProductImageEdges holds the relations/edges for other nodes in the graph.
type ProductImageEdges struct {
	// Products holds the value of the products edge.
	Products *Product `json:"products,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ProductsOrErr returns the Products value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProductImageEdges) ProductsOrErr() (*Product, error) {
	if e.Products != nil {
		return e.Products, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: product.Label}
	}
	return nil, &NotLoadedError{edge: "products"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ProductImage) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case productimage.FieldID, productimage.FieldProductID:
			values[i] = new(sql.NullInt64)
		case productimage.FieldName:
			values[i] = new(sql.NullString)
		case productimage.FieldCreatedAt, productimage.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ProductImage fields.
func (pi *ProductImage) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case productimage.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pi.ID = int64(value.Int64)
		case productimage.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pi.CreatedAt = value.Time
			}
		case productimage.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pi.UpdatedAt = value.Time
			}
		case productimage.FieldProductID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field product_id", values[i])
			} else if value.Valid {
				pi.ProductID = value.Int64
			}
		case productimage.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pi.Name = value.String
			}
		default:
			pi.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ProductImage.
// This includes values selected through modifiers, order, etc.
func (pi *ProductImage) Value(name string) (ent.Value, error) {
	return pi.selectValues.Get(name)
}

// QueryProducts queries the "products" edge of the ProductImage entity.
func (pi *ProductImage) QueryProducts() *ProductQuery {
	return NewProductImageClient(pi.config).QueryProducts(pi)
}

// Update returns a builder for updating this ProductImage.
// Note that you need to call ProductImage.Unwrap() before calling this method if this ProductImage
// was returned from a transaction, and the transaction was committed or rolled back.
func (pi *ProductImage) Update() *ProductImageUpdateOne {
	return NewProductImageClient(pi.config).UpdateOne(pi)
}

// Unwrap unwraps the ProductImage entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pi *ProductImage) Unwrap() *ProductImage {
	_tx, ok := pi.config.driver.(*txDriver)
	if !ok {
		panic("ent: ProductImage is not a transactional entity")
	}
	pi.config.driver = _tx.drv
	return pi
}

// String implements the fmt.Stringer.
func (pi *ProductImage) String() string {
	var builder strings.Builder
	builder.WriteString("ProductImage(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pi.ID))
	builder.WriteString("created_at=")
	builder.WriteString(pi.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pi.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("product_id=")
	builder.WriteString(fmt.Sprintf("%v", pi.ProductID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pi.Name)
	builder.WriteByte(')')
	return builder.String()
}

// ProductImages is a parsable slice of ProductImage.
type ProductImages []*ProductImage
