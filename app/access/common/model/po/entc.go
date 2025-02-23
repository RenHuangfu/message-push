//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	target := "./ent"
	pkgPath := "message-push/app/access/common/model/po/ent"
	err := entc.Generate("./schema", &gen.Config{
		Target:  target,
		Package: pkgPath,
		Features: []gen.Feature{
			gen.FeatureModifier,
			gen.FeatureUpsert,
			gen.FeatureLock,
			gen.FeatureExecQuery,
		},
	})
	if err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
