// Code generated by ent, DO NOT EDIT.

package address

import (
	"time"
)

const (
	// Label holds the string label denoting the address type in the database.
	Label = "address"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldCity holds the string denoting the city field in the database.
	FieldCity = "city"
	// FieldCountry holds the string denoting the country field in the database.
	FieldCountry = "country"
	// FieldFlat holds the string denoting the flat field in the database.
	FieldFlat = "flat"
	// FieldHouse holds the string denoting the house field in the database.
	FieldHouse = "house"
	// FieldLetter holds the string denoting the letter field in the database.
	FieldLetter = "letter"
	// FieldPostcode holds the string denoting the postcode field in the database.
	FieldPostcode = "postcode"
	// FieldStreet holds the string denoting the street field in the database.
	FieldStreet = "street"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// Table holds the table name of the address in the database.
	Table = "addresses"
	// UsersTable is the table that holds the users relation/edge.
	UsersTable = "addresses"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// UsersColumn is the table column denoting the users relation/edge.
	UsersColumn = "user_id"
)

// Columns holds all SQL columns for address fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldUserID,
	FieldCity,
	FieldCountry,
	FieldFlat,
	FieldHouse,
	FieldLetter,
	FieldPostcode,
	FieldStreet,
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
	// CityValidator is a validator for the "city" field. It is called by the builders before save.
	CityValidator func(string) error
	// CountryValidator is a validator for the "country" field. It is called by the builders before save.
	CountryValidator func(string) error
	// StreetValidator is a validator for the "street" field. It is called by the builders before save.
	StreetValidator func(string) error
)