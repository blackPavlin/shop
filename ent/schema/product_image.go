package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
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
	return []ent.Field{}
}

// Edges of the ProductImage.
func (ProductImage) Edges() []ent.Edge {
	return []ent.Edge{}
}
