package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Cart holds the schema definition for the Cart entity.
type Cart struct {
	ent.Schema
}

// Annotations of the Cart.
func (Cart) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "carts"},
	}
}

// Mixin of the Cart.
func (Cart) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Cart.
func (Cart) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id"),
		field.Int64("product_id"),
		field.Int64("amount").Min(0),
	}
}

// Edges of the Cart.
func (Cart) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).
			Ref("carts").
			Unique().
			Required().
			Field("user_id"),
		edge.From("products", Product.Type).
			Ref("carts").
			Unique().
			Required().
			Field("product_id"),
	}
}

// Indexes of the Cart.
func (Cart) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "product_id").Unique(),
	}
}
