package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ThemeTemplate holds the schema definition for the ThemeTemplate entity.
type ThemeTemplate struct {
	ent.Schema
}

func (ThemeTemplate) Mixin() []ent.Mixin {
	return []ent.Mixin{
		SonyFlakeIDMixin{},
		TimeMixin{},
		DeleteTimeMixin{},
		ShopMixin{},
	}
}

// Fields of the ThemeTemplate.
func (ThemeTemplate) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("theme_id").
			StructTag(`json:"themeId"`).
			Optional().
			Annotations(entgql.Skip()),
		// field.Int("theme_global_style_id").
		// 	Optional(),
		// field.Int("template_id").
		// 	Optional(),
		field.String("name").
			MaxLen(200).
			NotEmpty().
			Annotations(
				entgql.OrderField("NAME"),
			).
			StructTag(`json:"name"`),
		field.Enum("page_type").
			NamedValues(
				"index", "INDEX",
				"product", "PRODUCT",
				"collection", "COLLECTION",
				"article", "ARTICLE",
				"page", "PAGE",
			).
			StructTag(`json:"pageType"`).
			Default("INDEX").
			Annotations(
				entgql.OrderField("PAGE_TYPE"),
			),
		field.Bool("default").
			Default(false).
			StructTag(`json:"default"`),
	}
}

// Edges of the ThemeTemplate.
func (ThemeTemplate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("shop", Shop.Type).
			Ref("themeTemplates").
			Unique().
			Field("shop_id"),
		edge.From("theme", Theme.Type).
			Ref("themeTemplates").
			Unique().
			Field("theme_id"),
	}
}
