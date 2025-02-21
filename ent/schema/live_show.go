package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type LiveShow struct {
	ent.Schema
}

func (LiveShow) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.CreateTime{},
	}
}

func (LiveShow) Fields() []ent.Field {
	return []ent.Field{
		field.String("legend"),
		field.Int("duration"),
		field.Bool("viewed").Default(false),
		field.Bool("currentPlayed").Default(false),
		field.Time("started_time").Optional(),
		field.Time("ended_time").Optional(),
	}
}

func (LiveShow) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("upload", Upload.Type).Unique(),
	}
}
