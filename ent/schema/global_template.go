package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// GlobalTemplate holds the schema definition for the GlobalTemplate entity.
type GlobalTemplate struct {
	ent.Schema
}

func (GlobalTemplate) Mixin() []ent.Mixin {
	return []ent.Mixin{
		SonyFlakeIDMixin{},
		TimeMixin{},
		DeleteTimeMixin{},
		ShopMixin{},
	}
}

// Fields of the GlobalTemplate.
func (GlobalTemplate) Fields() []ent.Field {
	return []ent.Field{
		field.Text("name").
			NotEmpty().
			MaxLen(200),
		// field.Uint64("theme_template_id").
		// 	Positive().
		// 	StructTag(`json:"themeTemplateId"`),
		field.Int("view_count").
			StructTag(`json:"viewCount"`),
		field.Int("install_count").
			StructTag(`json:"installCount"`),
	}
}

// Edges of the GlobalTemplate.
func (GlobalTemplate) Edges() []ent.Edge {
	return []ent.Edge{}
}
