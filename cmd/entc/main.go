package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	if err := entc.Generate(
		"./ent/schema",
		&gen.Config{
			Target:  "./ent/gen",
			Package: "backend/ent/gen",
			Features: []gen.Feature{
				gen.FeatureIntercept,
				gen.FeatureModifier,
			},
		},
		entc.TemplateDir("./ent/template"),
	); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
