// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AddressesColumns holds the columns for the "addresses" table.
	AddressesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp"}},
		{Name: "city", Type: field.TypeString},
		{Name: "country", Type: field.TypeString},
		{Name: "flat", Type: field.TypeInt, Nullable: true},
		{Name: "house", Type: field.TypeInt},
		{Name: "letter", Type: field.TypeString, Nullable: true},
		{Name: "postcode", Type: field.TypeInt},
		{Name: "street", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeInt64},
	}
	// AddressesTable holds the schema information for the "addresses" table.
	AddressesTable = &schema.Table{
		Name:       "addresses",
		Columns:    AddressesColumns,
		PrimaryKey: []*schema.Column{AddressesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "address_user_fk",
				Columns:    []*schema.Column{AddressesColumns[10]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// CartsColumns holds the columns for the "carts" table.
	CartsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp"}},
		{Name: "amount", Type: field.TypeInt64},
		{Name: "product_id", Type: field.TypeInt64},
		{Name: "user_id", Type: field.TypeInt64},
	}
	// CartsTable holds the schema information for the "carts" table.
	CartsTable = &schema.Table{
		Name:       "carts",
		Columns:    CartsColumns,
		PrimaryKey: []*schema.Column{CartsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "cart_product_fk",
				Columns:    []*schema.Column{CartsColumns[4]},
				RefColumns: []*schema.Column{ProductsColumns[0]},
				OnDelete:   schema.Restrict,
			},
			{
				Symbol:     "cart_user_fk",
				Columns:    []*schema.Column{CartsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// CategoriesColumns holds the columns for the "categories" table.
	CategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp"}},
		{Name: "name", Type: field.TypeString},
	}
	// CategoriesTable holds the schema information for the "categories" table.
	CategoriesTable = &schema.Table{
		Name:       "categories",
		Columns:    CategoriesColumns,
		PrimaryKey: []*schema.Column{CategoriesColumns[0]},
	}
	// OrdersColumns holds the columns for the "orders" table.
	OrdersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp"}},
		{Name: "user_id", Type: field.TypeInt64},
	}
	// OrdersTable holds the schema information for the "orders" table.
	OrdersTable = &schema.Table{
		Name:       "orders",
		Columns:    OrdersColumns,
		PrimaryKey: []*schema.Column{OrdersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "order_user_fk",
				Columns:    []*schema.Column{OrdersColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Restrict,
			},
		},
	}
	// ProductsColumns holds the columns for the "products" table.
	ProductsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp"}},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "category_id", Type: field.TypeInt64},
	}
	// ProductsTable holds the schema information for the "products" table.
	ProductsTable = &schema.Table{
		Name:       "products",
		Columns:    ProductsColumns,
		PrimaryKey: []*schema.Column{ProductsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "product_category_fk",
				Columns:    []*schema.Column{ProductsColumns[5]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.Restrict,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamp"}},
		{Name: "name", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"ADMIN", "USER"}, Default: "USER", SchemaType: map[string]string{"postgres": "user_role"}},
		{Name: "password", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AddressesTable,
		CartsTable,
		CategoriesTable,
		OrdersTable,
		ProductsTable,
		UsersTable,
	}
)

func init() {
	AddressesTable.ForeignKeys[0].RefTable = UsersTable
	AddressesTable.Annotation = &entsql.Annotation{
		Table: "addresses",
	}
	CartsTable.ForeignKeys[0].RefTable = ProductsTable
	CartsTable.ForeignKeys[1].RefTable = UsersTable
	CartsTable.Annotation = &entsql.Annotation{
		Table: "carts",
	}
	CategoriesTable.Annotation = &entsql.Annotation{
		Table: "categories",
	}
	OrdersTable.ForeignKeys[0].RefTable = UsersTable
	OrdersTable.Annotation = &entsql.Annotation{
		Table: "orders",
	}
	ProductsTable.ForeignKeys[0].RefTable = CategoriesTable
	ProductsTable.Annotation = &entsql.Annotation{
		Table: "products",
	}
	UsersTable.Annotation = &entsql.Annotation{
		Table: "users",
	}
}
