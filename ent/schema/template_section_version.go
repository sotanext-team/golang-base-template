package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TemplateSectionVersion holds the schema definition for the TemplateSectionVersion entity.
type TemplateSectionVersion struct {
	ent.Schema
}

func (TemplateSectionVersion) Mixin() []ent.Mixin {
	return []ent.Mixin{
		SonyFlakeIDMixin{},
		TimeMixin{},
	}
}

// Fields of the TemplateSectionVersion.
func (TemplateSectionVersion) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("theme_template_id").
			StructTag(`json:"themeTemplateId"`).
			Optional().
			Annotations(entgql.Skip()),
		field.String("version").
			MaxLen(10).
			NotEmpty().
			StructTag(`json:"version"`),
		field.String("name").
			MaxLen(100).
			StructTag(`json:"customName"`),
	}
}

// Edges of the TemplateSectionVersion.
func (TemplateSectionVersion) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("bkTemplateSections", BkTemplateSection.Type).
			Annotations(
				entgql.Bind(),
				entsql.Annotation{
					OnDelete: entsql.Cascade,
				},
			),
		// Field("version_id"),
	}
}
