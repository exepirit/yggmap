//go:build tools
// +build tools

package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {

	if err := entc.Generate("./data/ent/schema", &gen.Config{
		Target: "./data/ent",
		Features: []gen.Feature{
			gen.FeatureUpsert,
		},
	}); err != nil {
		log.Fatal("running ent codegen:", err)
	}
	log.Println("ent boilerplate generated")
}
