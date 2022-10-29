package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Annotations of the User.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "users"},
	}
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("phone").NotEmpty(),
		field.String("email").Unique().NotEmpty(),
		field.Enum("role").
			Values("ADMIN", "USER").
			Default("USER").
			SchemaType(map[string]string{
				dialect.Postgres: "user_role",
			}),
		field.String("password").NotEmpty().Sensitive(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("addresses", Address.Type).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}).
			StorageKey(
				edge.Symbol("address_user_fk"),
			),
		edge.To("carts", Cart.Type).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}).
			StorageKey(
				edge.Symbol("cart_user_fk"),
			),
		edge.To("orders", Order.Type).
			Annotations(entsql.Annotation{OnDelete: entsql.Restrict}).
			StorageKey(
				edge.Symbol("order_user_fk"),
			),
	}
}
