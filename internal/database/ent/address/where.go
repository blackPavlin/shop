// Code generated by ent, DO NOT EDIT.

package address

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/blackPavlin/shop/internal/database/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldUpdatedAt, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldUserID, v))
}

// City applies equality check predicate on the "city" field. It's identical to CityEQ.
func City(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldCity, v))
}

// Country applies equality check predicate on the "country" field. It's identical to CountryEQ.
func Country(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldCountry, v))
}

// Flat applies equality check predicate on the "flat" field. It's identical to FlatEQ.
func Flat(v int) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldFlat, v))
}

// House applies equality check predicate on the "house" field. It's identical to HouseEQ.
func House(v int) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldHouse, v))
}

// Letter applies equality check predicate on the "letter" field. It's identical to LetterEQ.
func Letter(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldLetter, v))
}

// Postcode applies equality check predicate on the "postcode" field. It's identical to PostcodeEQ.
func Postcode(v int) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldPostcode, v))
}

// Street applies equality check predicate on the "street" field. It's identical to StreetEQ.
func Street(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldStreet, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldUpdatedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldUserID, vs...))
}

// CityEQ applies the EQ predicate on the "city" field.
func CityEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldCity, v))
}

// CityNEQ applies the NEQ predicate on the "city" field.
func CityNEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldCity, v))
}

// CityIn applies the In predicate on the "city" field.
func CityIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldCity, vs...))
}

// CityNotIn applies the NotIn predicate on the "city" field.
func CityNotIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldCity, vs...))
}

// CityGT applies the GT predicate on the "city" field.
func CityGT(v string) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldCity, v))
}

// CityGTE applies the GTE predicate on the "city" field.
func CityGTE(v string) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldCity, v))
}

// CityLT applies the LT predicate on the "city" field.
func CityLT(v string) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldCity, v))
}

// CityLTE applies the LTE predicate on the "city" field.
func CityLTE(v string) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldCity, v))
}

// CityContains applies the Contains predicate on the "city" field.
func CityContains(v string) predicate.Address {
	return predicate.Address(sql.FieldContains(FieldCity, v))
}

// CityHasPrefix applies the HasPrefix predicate on the "city" field.
func CityHasPrefix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasPrefix(FieldCity, v))
}

// CityHasSuffix applies the HasSuffix predicate on the "city" field.
func CityHasSuffix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasSuffix(FieldCity, v))
}

// CityEqualFold applies the EqualFold predicate on the "city" field.
func CityEqualFold(v string) predicate.Address {
	return predicate.Address(sql.FieldEqualFold(FieldCity, v))
}

// CityContainsFold applies the ContainsFold predicate on the "city" field.
func CityContainsFold(v string) predicate.Address {
	return predicate.Address(sql.FieldContainsFold(FieldCity, v))
}

// CountryEQ applies the EQ predicate on the "country" field.
func CountryEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldCountry, v))
}

// CountryNEQ applies the NEQ predicate on the "country" field.
func CountryNEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldCountry, v))
}

// CountryIn applies the In predicate on the "country" field.
func CountryIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldCountry, vs...))
}

// CountryNotIn applies the NotIn predicate on the "country" field.
func CountryNotIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldCountry, vs...))
}

// CountryGT applies the GT predicate on the "country" field.
func CountryGT(v string) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldCountry, v))
}

// CountryGTE applies the GTE predicate on the "country" field.
func CountryGTE(v string) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldCountry, v))
}

// CountryLT applies the LT predicate on the "country" field.
func CountryLT(v string) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldCountry, v))
}

// CountryLTE applies the LTE predicate on the "country" field.
func CountryLTE(v string) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldCountry, v))
}

// CountryContains applies the Contains predicate on the "country" field.
func CountryContains(v string) predicate.Address {
	return predicate.Address(sql.FieldContains(FieldCountry, v))
}

// CountryHasPrefix applies the HasPrefix predicate on the "country" field.
func CountryHasPrefix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasPrefix(FieldCountry, v))
}

// CountryHasSuffix applies the HasSuffix predicate on the "country" field.
func CountryHasSuffix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasSuffix(FieldCountry, v))
}

// CountryEqualFold applies the EqualFold predicate on the "country" field.
func CountryEqualFold(v string) predicate.Address {
	return predicate.Address(sql.FieldEqualFold(FieldCountry, v))
}

// CountryContainsFold applies the ContainsFold predicate on the "country" field.
func CountryContainsFold(v string) predicate.Address {
	return predicate.Address(sql.FieldContainsFold(FieldCountry, v))
}

// FlatEQ applies the EQ predicate on the "flat" field.
func FlatEQ(v int) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldFlat, v))
}

// FlatNEQ applies the NEQ predicate on the "flat" field.
func FlatNEQ(v int) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldFlat, v))
}

// FlatIn applies the In predicate on the "flat" field.
func FlatIn(vs ...int) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldFlat, vs...))
}

// FlatNotIn applies the NotIn predicate on the "flat" field.
func FlatNotIn(vs ...int) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldFlat, vs...))
}

// FlatGT applies the GT predicate on the "flat" field.
func FlatGT(v int) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldFlat, v))
}

// FlatGTE applies the GTE predicate on the "flat" field.
func FlatGTE(v int) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldFlat, v))
}

// FlatLT applies the LT predicate on the "flat" field.
func FlatLT(v int) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldFlat, v))
}

// FlatLTE applies the LTE predicate on the "flat" field.
func FlatLTE(v int) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldFlat, v))
}

// FlatIsNil applies the IsNil predicate on the "flat" field.
func FlatIsNil() predicate.Address {
	return predicate.Address(sql.FieldIsNull(FieldFlat))
}

// FlatNotNil applies the NotNil predicate on the "flat" field.
func FlatNotNil() predicate.Address {
	return predicate.Address(sql.FieldNotNull(FieldFlat))
}

// HouseEQ applies the EQ predicate on the "house" field.
func HouseEQ(v int) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldHouse, v))
}

// HouseNEQ applies the NEQ predicate on the "house" field.
func HouseNEQ(v int) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldHouse, v))
}

// HouseIn applies the In predicate on the "house" field.
func HouseIn(vs ...int) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldHouse, vs...))
}

// HouseNotIn applies the NotIn predicate on the "house" field.
func HouseNotIn(vs ...int) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldHouse, vs...))
}

// HouseGT applies the GT predicate on the "house" field.
func HouseGT(v int) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldHouse, v))
}

// HouseGTE applies the GTE predicate on the "house" field.
func HouseGTE(v int) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldHouse, v))
}

// HouseLT applies the LT predicate on the "house" field.
func HouseLT(v int) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldHouse, v))
}

// HouseLTE applies the LTE predicate on the "house" field.
func HouseLTE(v int) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldHouse, v))
}

// LetterEQ applies the EQ predicate on the "letter" field.
func LetterEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldLetter, v))
}

// LetterNEQ applies the NEQ predicate on the "letter" field.
func LetterNEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldLetter, v))
}

// LetterIn applies the In predicate on the "letter" field.
func LetterIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldLetter, vs...))
}

// LetterNotIn applies the NotIn predicate on the "letter" field.
func LetterNotIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldLetter, vs...))
}

// LetterGT applies the GT predicate on the "letter" field.
func LetterGT(v string) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldLetter, v))
}

// LetterGTE applies the GTE predicate on the "letter" field.
func LetterGTE(v string) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldLetter, v))
}

// LetterLT applies the LT predicate on the "letter" field.
func LetterLT(v string) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldLetter, v))
}

// LetterLTE applies the LTE predicate on the "letter" field.
func LetterLTE(v string) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldLetter, v))
}

// LetterContains applies the Contains predicate on the "letter" field.
func LetterContains(v string) predicate.Address {
	return predicate.Address(sql.FieldContains(FieldLetter, v))
}

// LetterHasPrefix applies the HasPrefix predicate on the "letter" field.
func LetterHasPrefix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasPrefix(FieldLetter, v))
}

// LetterHasSuffix applies the HasSuffix predicate on the "letter" field.
func LetterHasSuffix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasSuffix(FieldLetter, v))
}

// LetterIsNil applies the IsNil predicate on the "letter" field.
func LetterIsNil() predicate.Address {
	return predicate.Address(sql.FieldIsNull(FieldLetter))
}

// LetterNotNil applies the NotNil predicate on the "letter" field.
func LetterNotNil() predicate.Address {
	return predicate.Address(sql.FieldNotNull(FieldLetter))
}

// LetterEqualFold applies the EqualFold predicate on the "letter" field.
func LetterEqualFold(v string) predicate.Address {
	return predicate.Address(sql.FieldEqualFold(FieldLetter, v))
}

// LetterContainsFold applies the ContainsFold predicate on the "letter" field.
func LetterContainsFold(v string) predicate.Address {
	return predicate.Address(sql.FieldContainsFold(FieldLetter, v))
}

// PostcodeEQ applies the EQ predicate on the "postcode" field.
func PostcodeEQ(v int) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldPostcode, v))
}

// PostcodeNEQ applies the NEQ predicate on the "postcode" field.
func PostcodeNEQ(v int) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldPostcode, v))
}

// PostcodeIn applies the In predicate on the "postcode" field.
func PostcodeIn(vs ...int) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldPostcode, vs...))
}

// PostcodeNotIn applies the NotIn predicate on the "postcode" field.
func PostcodeNotIn(vs ...int) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldPostcode, vs...))
}

// PostcodeGT applies the GT predicate on the "postcode" field.
func PostcodeGT(v int) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldPostcode, v))
}

// PostcodeGTE applies the GTE predicate on the "postcode" field.
func PostcodeGTE(v int) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldPostcode, v))
}

// PostcodeLT applies the LT predicate on the "postcode" field.
func PostcodeLT(v int) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldPostcode, v))
}

// PostcodeLTE applies the LTE predicate on the "postcode" field.
func PostcodeLTE(v int) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldPostcode, v))
}

// StreetEQ applies the EQ predicate on the "street" field.
func StreetEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldStreet, v))
}

// StreetNEQ applies the NEQ predicate on the "street" field.
func StreetNEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldStreet, v))
}

// StreetIn applies the In predicate on the "street" field.
func StreetIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldStreet, vs...))
}

// StreetNotIn applies the NotIn predicate on the "street" field.
func StreetNotIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldStreet, vs...))
}

// StreetGT applies the GT predicate on the "street" field.
func StreetGT(v string) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldStreet, v))
}

// StreetGTE applies the GTE predicate on the "street" field.
func StreetGTE(v string) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldStreet, v))
}

// StreetLT applies the LT predicate on the "street" field.
func StreetLT(v string) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldStreet, v))
}

// StreetLTE applies the LTE predicate on the "street" field.
func StreetLTE(v string) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldStreet, v))
}

// StreetContains applies the Contains predicate on the "street" field.
func StreetContains(v string) predicate.Address {
	return predicate.Address(sql.FieldContains(FieldStreet, v))
}

// StreetHasPrefix applies the HasPrefix predicate on the "street" field.
func StreetHasPrefix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasPrefix(FieldStreet, v))
}

// StreetHasSuffix applies the HasSuffix predicate on the "street" field.
func StreetHasSuffix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasSuffix(FieldStreet, v))
}

// StreetEqualFold applies the EqualFold predicate on the "street" field.
func StreetEqualFold(v string) predicate.Address {
	return predicate.Address(sql.FieldEqualFold(FieldStreet, v))
}

// StreetContainsFold applies the ContainsFold predicate on the "street" field.
func StreetContainsFold(v string) predicate.Address {
	return predicate.Address(sql.FieldContainsFold(FieldStreet, v))
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UsersTable, UsersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		step := newUsersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Address) predicate.Address {
	return predicate.Address(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Address) predicate.Address {
	return predicate.Address(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Address) predicate.Address {
	return predicate.Address(sql.NotPredicates(p))
}
