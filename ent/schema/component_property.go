package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
)

// ComponentProperty holds the schema definition for the ComponentProperty entity.
type ComponentProperty struct {
    ent.Schema
}

func (ComponentProperty) Mixin() []ent.Mixin {
    return []ent.Mixin{
        SonyFlakeIDMixin{},
        TimeMixin{},
    }
}

// Fields of the ComponentProperty.
func (ComponentProperty) Fields() []ent.Field {
    return []ent.Field{
        field.Uint64("component_id").StructTag(`json:"componentID"`),
        field.String("name"),
        field.String("value"),
    }
}

// Edges of the ComponentProperty.
func (ComponentProperty) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("component", CustomComponent.Type).
            Field("component_id").
            Ref("props").
            Unique().
            Required(),
    }
}
