package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// Order holds the schema definition for the OrderProduct entity.
type OrderProduct struct {
	ent.Schema
}

// Annotations of the OrderProduct.
func (OrderProduct) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "order_products"},
	}
}

// Mixin of the OrderProduct.
func (OrderProduct) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the OrderProduct.
func (OrderProduct) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the OrderProduct.
func (OrderProduct) Edges() []ent.Edge {
	return []ent.Edge{}
}
