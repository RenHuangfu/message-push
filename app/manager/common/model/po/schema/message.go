package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

type Message struct {
	ent.Schema
}

func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Unique().
			Immutable(),
		field.Int("app_id"),
		field.JSON("client_ids", []int{}),
		field.String("content"),
		field.Int("delay"),
		field.Time("send_time"),
		field.Int("send_count"),
		field.Int("status"),
		field.Int("is_del"),
		field.Time("create_time").
			Default(time.Now).
			Immutable(),
		field.Time("update_time").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

func (Message) Edges() []ent.Edge {
	return nil
}
