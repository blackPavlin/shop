package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Annotations of the Product.
func (Product) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "products"},
	}
}

// Mixin of the Product.
func (Product) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("category_id").Unique(),
		field.String("name").NotEmpty(),
		field.String("description").Optional(),
		field.Int64("amount").Min(0).Default(0),
		field.Int64("price").Min(0),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("categories", Category.Type).
			Ref("products").
			Unique().
			Required().
			Field("category_id"),
		edge.To("carts", Cart.Type).
			Annotations(
				entsql.Annotation{OnDelete: entsql.Restrict},
			).
			StorageKey(
				edge.Symbol("cart_product_fk"),
			),
		edge.To("product_images", ProductImage.Type).
			Annotations(
				entsql.Annotation{OnDelete: entsql.Cascade},
			).
			StorageKey(
				edge.Symbol("product_image_product_fk"),
			),
		edge.To("order_products", OrderProduct.Type).
			Annotations(
				entsql.Annotation{OnDelete: entsql.Restrict},
			).
			StorageKey(
				edge.Symbol("order_product_product_fk"),
			),
	}
}
