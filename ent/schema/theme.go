package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Theme holds the schema definition for the Theme entity.
type Theme struct {
	ent.Schema
}

func (Theme) Mixin() []ent.Mixin {
	return []ent.Mixin{
		SonyFlakeIDMixin{},
		TimeMixin{},
		DeleteTimeMixin{},
		ShopMixin{},
	}
}

// Fields of the Theme.
func (Theme) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			MaxLen(200).
			NotEmpty().StructTag(`json:"name"`),
		field.String("thumbnail").
			MaxLen(2084).
			NotEmpty().StructTag(`json:"thumbnail"`),
		field.Bool("publish").
			Default(false).StructTag(`json:"publish"`),
	}
}

// Edges of the Theme.
func (Theme) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("themeTemplates", ThemeTemplate.Type),
		edge.From("shop", Shop.Type).
			Ref("themes").
			Unique().
			Field("shop_id"),
	}
}
