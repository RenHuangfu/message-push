// Code generated by ent, DO NOT EDIT.

package ent

import (
	"message-push/app/receiver/common/model/po/ent/demo"
	"message-push/app/receiver/common/model/po/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	demoFields := schema.Demo{}.Fields()
	_ = demoFields
	// demoDescName is the schema descriptor for name field.
	demoDescName := demoFields[1].Descriptor()
	// demo.NameValidator is a validator for the "name" field. It is called by the builders before save.
	demo.NameValidator = demoDescName.Validators[0].(func(string) error)
	// demoDescID is the schema descriptor for id field.
	demoDescID := demoFields[0].Descriptor()
	// demo.IDValidator is a validator for the "id" field. It is called by the builders before save.
	demo.IDValidator = demoDescID.Validators[0].(func(int) error)
}
