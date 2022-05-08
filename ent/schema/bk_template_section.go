package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// BkTemplateSection holds the schema definition for the BkTemplateSection entity.
type BkTemplateSection struct {
	ent.Schema
}

func (BkTemplateSection) Mixin() []ent.Mixin {
	return []ent.Mixin{
		SonyFlakeIDMixin{},
		TimeMixin{},
	}
}

// Fields of the BkTemplateSection.
func (BkTemplateSection) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("version_id").
			StructTag(`json:"versionId"`).
			Optional(),
		// Annotations(entgql.Skip()),
		field.Uint64("theme_template_id").
			StructTag(`json:"themeTemplateId"`).
			Optional().
			Annotations(entgql.Skip()),
		field.Uint64("template_section_id").
			StructTag(`json:"templateSectionId"`).
			Optional().
			Annotations(entgql.Skip()),
		field.Uint64("theme_id").
			StructTag(`json:"themeId"`).
			Optional().
			Annotations(entgql.Skip()),
		field.Uint64("theme_layout_id").
			StructTag(`json:"themeLayoutId"`).
			Optional().
			Annotations(entgql.Skip()),
		field.Text("data").
			StructTag(`json:"data"`),
	}
}

// Edges of the BkTemplateSection.
func (BkTemplateSection) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("templateSection", TemplateSection.Type).
			Ref("revisions").
			Unique().
			Field("template_section_id"),
		edge.From("version", TemplateSectionVersion.Type).
			Ref("bkTemplateSections").
			Unique().
			Annotations(
				entgql.Bind(),
				entsql.Annotation{
					OnDelete: entsql.Cascade,
				},
			).
			Field("version_id"),
	}
}
