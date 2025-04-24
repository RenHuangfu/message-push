package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

type BusinessApp struct {
	ent.Schema
}

func (BusinessApp) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			StorageKey("id"),
		field.String("name").
			MaxLen(255).
			NotEmpty().
			Comment("应用名称"),
		field.String("app_key").
			MaxLen(255).
			NotEmpty().
			Comment("应用key"),
		field.String("app_id").
			MaxLen(255).
			NotEmpty().
			Comment("应用id"),
		field.String("app_type").
			MaxLen(255).
			NotEmpty().
			Comment("应用类型"),
		field.Int("is_del").
			Comment("是否删除"),
		field.Time("create_time").
			Default(time.Now).
			Immutable(),
		field.Time("update_time").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

func (BusinessApp) Edges() []ent.Edge {
	return nil
}
