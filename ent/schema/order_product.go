package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// OrderProduct holds the schema definition for the OrderProduct entity.
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
	return []ent.Field{
		field.Int64("order_id"),
		field.Int64("product_id"),
		field.Int64("amount").Min(0).Default(0),
		field.Int64("price").Min(0),
	}
}

// Edges of the OrderProduct.
func (OrderProduct) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("orders", Order.Type).
			Ref("order_products").
			Unique().
			Required().
			Field("order_id"),
		edge.From("products", Product.Type).
			Ref("order_products").
			Unique().
			Required().
			Field("product_id"),
	}
}

// Indexes of the ProductImage.
func (OrderProduct) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("product_id", "order_id").
			Unique().
			StorageKey("order_product_order_id_product_id_unique"),
	}
}
