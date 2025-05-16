package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Annotations of the Order.
func (Order) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "orders"},
	}
}

// Mixin of the Order.
func (Order) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id").Unique(),
		field.Enum("status").
			Values(
				"CREATED",
				"ACCEPTED",
				"CANCELED",
			).
			Default("CREATED").
			SchemaType(map[string]string{
				dialect.Postgres: "order_status",
			}),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).
			Ref("orders").
			Unique().
			Required().
			Field("user_id"),
		edge.To("order_products", OrderProduct.Type).
			Annotations(
				entsql.Annotation{OnDelete: entsql.Restrict},
			).
			StorageKey(
				edge.Symbol("order_product_order_fk"),
			),
	}
}
