package schema

import (
    "entgo.io/contrib/entgql"
    "entgo.io/ent"
    "entgo.io/ent/dialect/entsql"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
)

// CustomComponent holds the schema definition for the CustomComponent entity.
type CustomComponent struct {
    ent.Schema
}

func (CustomComponent) Mixin() []ent.Mixin {
    return []ent.Mixin{
        SonyFlakeIDMixin{},
        TimeMixin{},
    }
}

// Fields of the CustomComponent.
func (CustomComponent) Fields() []ent.Field {
    return []ent.Field{
        field.Uint64("shop_id").Optional().Nillable().
            StructTag(`json:"shopID"`).
            Annotations(entgql.Skip()),
        field.Uint64("created_by").StructTag(`json:"createdBy"`).
            Annotations(entgql.Skip()),
        field.String("name"),
        field.String("handle"),
        field.Text("content").Optional(),
        field.String("dist_url").Optional().Nillable().Unique().StructTag(`json:"distURL"`),
        field.Bool("is_draft").Optional().Default(true).StructTag(`json:"isDraft"`), // determines user has ever saved the component or not
        field.String("entry_file_name").Optional().StructTag(`json:"entryFileName"`),
        field.Uint64("session_id").Optional().Nillable().
            StructTag(`json:"sessionID"`).
            Annotations(entgql.Skip()),
    }
}

// Edges of the CustomComponent.
func (CustomComponent) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("props", ComponentProperty.Type).
            Annotations(
                entsql.Annotation{
                    OnDelete: entsql.Cascade,
                },
            ),
    }
}
