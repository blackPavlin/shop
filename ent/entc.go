//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/field"
)

func main() {
	config := &gen.Config{
		IDType: &field.TypeInfo{
			Type: field.TypeInt64,
		},
	}

	ex, err := entgql.NewExtension(
		entgql.WithConfigPath("../gqlgen.yml"),
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("../api/ent.graphql"),
		entgql.WithWhereInputs(true),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}

	options := []entc.Option{
		entc.Extensions(ex),
	}

	if err := entc.Generate("./schema", config, options...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
