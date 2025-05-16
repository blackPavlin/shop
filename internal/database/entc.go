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
		IDType:   &field.TypeInfo{Type: field.TypeInt64},
		Features: []gen.Feature{gen.FeatureUpsert, gen.FeatureLock, gen.FeatureVersionedMigration},
		Target:   "./ent",
		Package:  "github.com/blackPavlin/shop/internal/database/ent",
	}

	if err := entc.Generate("./schema", config); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
