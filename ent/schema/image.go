package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Order holds the schema definition for the Image entity.
type Image struct {
	ent.Schema
}

// Annotations of the Image.
func (Image) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "images"},
	}
}

// Mixin of the Image.
func (Image) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Image.
func (Image) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("original_name").NotEmpty(),
	}
}

// Edges of the Image.
func (Image) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("product_images", ProductImage.Type).
			Annotations(
				entsql.Annotation{OnDelete: entsql.Restrict},
			).
			StorageKey(
				edge.Symbol("product_image_image_fk"),
			),
	}
}
