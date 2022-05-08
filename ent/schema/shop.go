package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Shop holds the schema definition for the Shop entity.
type Shop struct {
	ent.Schema
}

func (Shop) Mixin() []ent.Mixin {
	return []ent.Mixin{
		SonyFlakeIDMixin{},
		TimeMixin{},
		DeleteTimeMixin{},
	}
}

// Fields of the Shop.
func (Shop) Fields() []ent.Field {
	return []ent.Field{
		field.Text("shop_name").
			Default("").
			MaxLen(200).
			Annotations(
				entgql.OrderField("SHOP_NAME"),
			),
		field.Text("default_domain").
			NotEmpty().
			MaxLen(200).
			Unique().
			Annotations(
				entgql.OrderField("DEFAULT_DOMAIN"),
			),
		field.Text("custom_domain").
			MaxLen(200).
			Annotations(
				entgql.OrderField("CUSTOM_DOMAIN"),
			),
	}
}

// Edges of the Shop.
func (Shop) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("themes", Theme.Type).Annotations(entgql.Bind()),
		edge.To("themeTemplates", ThemeTemplate.Type),
	}
}
