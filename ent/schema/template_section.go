package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TemplateSection holds the schema definition for the TemplateSection entity.
type TemplateSection struct {
	ent.Schema
}

func (TemplateSection) Mixin() []ent.Mixin {
	return []ent.Mixin{
		SonyFlakeIDMixin{},
		TimeMixin{},
		DeleteTimeMixin{},
		ShopMixin{},
	}
}

// Fields of the TemplateSection.
func (TemplateSection) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("global_section_id").
			StructTag(`json:"globalSectionId"`).
			Optional().
			Annotations(entgql.Skip()),
		field.Uint64("theme_template_id").
			StructTag(`json:"themeTemplateId"`).
			Optional().
			Annotations(entgql.Skip()),
		field.Uint64("current_version_id").
			StructTag(`json:"currentVersionId"`).
			Optional().
			Annotations(entgql.Skip()),
		field.String("cid").
			MaxLen(50).
			NotEmpty().
			StructTag(`json:"cid"`),
		field.String("name").
			MaxLen(100).
			NotEmpty().
			Annotations(
				entgql.OrderField("NAME"),
			).
			StructTag(`json:"name"`),
		field.Enum("area").
			NamedValues(
				"header", "HEADER",
				"main", "MAIN",
				"footer", "FOOTER",
			).
			Annotations(
				entgql.OrderField("AREA"),
			).
			StructTag(`json:"area"`).
			Default("MAIN"),
		field.Text("component").
			StructTag(`json:"component"`),
		field.Int("position").
			Default(1).StructTag(`json:"position"`).
			Annotations(
				entgql.OrderField("POSITION"),
			),
		field.Bool("display").
			Default(true).
			StructTag(`json:"display"`),
	}
}

// Edges of the TemplateSection.
func (TemplateSection) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("revisions", BkTemplateSection.Type).
			Annotations(
				entgql.Bind(),
				entsql.Annotation{
					OnDelete: entsql.Cascade,
				},
			),
		// edge.From("shop", Shop.Type).
		// 	Ref("TemplateSections").
		// 	Unique().
		// 	Field("shop_id"),
		// edge.From("theme", Theme.Type).
		// 	Ref("TemplateSections").
		// 	Unique().
		// 	Field("theme_id"),
	}
}
