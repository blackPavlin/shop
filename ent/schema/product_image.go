package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ProductImage holds the schema definition for the ProductImage entity.
type ProductImage struct {
	ent.Schema
}

// Annotations of the ProductImage.
func (ProductImage) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "product_images"},
	}
}

// Mixin of the ProductImage.
func (ProductImage) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the ProductImage.
func (ProductImage) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("product_id").Unique(),
		field.Int64("image_id").Unique(),
	}
}

// Edges of the ProductImage.
func (ProductImage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("products", Product.Type).
			Ref("product_images").
			Unique().
			Required().
			Field("product_id"),
		edge.From("images", Image.Type).
			Ref("product_images").
			Unique().
			Required().
			Field("image_id"),
	}
}
