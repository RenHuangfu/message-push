package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

type BusinessClient struct {
	ent.Schema
}

func (BusinessClient) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			StorageKey("id"),
		field.String("name").
			MaxLen(255).
			NotEmpty().
			Comment("客户端名称"),
		field.String("client_key").
			MaxLen(255).
			NotEmpty().
			Comment("客户端key"),
		field.String("app_id").
			MaxLen(255).
			NotEmpty().
			Comment("应用id"),
		field.String("client_id").
			MaxLen(255).
			NotEmpty().
			Comment("客户端id"),
		field.String("client_type").
			MaxLen(255).
			NotEmpty().
			Comment("客户端类型"),
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

func (BusinessClient) Edges() []ent.Edge {
	return nil
}
