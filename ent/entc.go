//go:build ignore
// +build ignore

package main

import (
	"log"

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

	options := []entc.Option{
		entc.FeatureNames("sql/upsert"),
	}

	if err := entc.Generate("./schema", config, options...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
