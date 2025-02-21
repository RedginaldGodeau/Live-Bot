package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type Upload struct {
	ent.Schema
}

func (Upload) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.CreateTime{},
	}
}

func (Upload) Fields() []ent.Field {
	return []ent.Field{
		field.String("file_path"),
		field.String("name"),
		field.Enum("file_type").Values("image", "video"),
	}
}

func (Upload) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("live_show_upload", LiveShow.Type).Ref("upload"),
	}
}
