package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Address holds the schema definition for the Address entity.
type Address struct {
	ent.Schema
}

// Annotations of the Address.
func (Address) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "addresses"},
	}
}

// Mixin of the Address.
func (Address) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Address.
func (Address) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").Unique(),
		field.String("city").NotEmpty(),
		field.String("country").NotEmpty(),
		field.Int("flat").Optional(),
		field.Int("house"),
		field.String("letter").Optional(),
		field.Int("postcode"),
		field.String("street").NotEmpty(),
	}
}

// Edges of the Address.
func (Address) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).
			Ref("addresses").
			Unique().
			Required().
			Field("user_id"),
	}
}
