// Code generated by ent, DO NOT EDIT.

package product

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the product type in the database.
	Label = "product"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCategoryID holds the string denoting the category_id field in the database.
	FieldCategoryID = "category_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// EdgeCategories holds the string denoting the categories edge name in mutations.
	EdgeCategories = "categories"
	// EdgeCarts holds the string denoting the carts edge name in mutations.
	EdgeCarts = "carts"
	// Table holds the table name of the product in the database.
	Table = "products"
	// CategoriesTable is the table that holds the categories relation/edge.
	CategoriesTable = "products"
	// CategoriesInverseTable is the table name for the Category entity.
	// It exists in this package in order to avoid circular dependency with the "category" package.
	CategoriesInverseTable = "categories"
	// CategoriesColumn is the table column denoting the categories relation/edge.
	CategoriesColumn = "category_id"
	// CartsTable is the table that holds the carts relation/edge.
	CartsTable = "carts"
	// CartsInverseTable is the table name for the Cart entity.
	// It exists in this package in order to avoid circular dependency with the "cart" package.
	CartsInverseTable = "carts"
	// CartsColumn is the table column denoting the carts relation/edge.
	CartsColumn = "product_id"
)

// Columns holds all SQL columns for product fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldCategoryID,
	FieldName,
	FieldDescription,
	FieldAmount,
	FieldPrice,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultAmount holds the default value on creation for the "amount" field.
	DefaultAmount int64
	// AmountValidator is a validator for the "amount" field. It is called by the builders before save.
	AmountValidator func(int64) error
	// PriceValidator is a validator for the "price" field. It is called by the builders before save.
	PriceValidator func(int64) error
)

// OrderOption defines the ordering options for the Product queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByCategoryID orders the results by the category_id field.
func ByCategoryID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCategoryID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByAmount orders the results by the amount field.
func ByAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAmount, opts...).ToFunc()
}

// ByPrice orders the results by the price field.
func ByPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrice, opts...).ToFunc()
}

// ByCategoriesField orders the results by categories field.
func ByCategoriesField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCategoriesStep(), sql.OrderByField(field, opts...))
	}
}

// ByCartsCount orders the results by carts count.
func ByCartsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCartsStep(), opts...)
	}
}

// ByCarts orders the results by carts terms.
func ByCarts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCartsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newCategoriesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CategoriesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CategoriesTable, CategoriesColumn),
	)
}
func newCartsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CartsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CartsTable, CartsColumn),
	)
}
