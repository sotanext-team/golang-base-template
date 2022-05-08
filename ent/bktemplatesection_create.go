// Code generated by entc, DO NOT EDIT.

package ent

import (
	"app-api/ent/bktemplatesection"
	"app-api/ent/templatesection"
	"app-api/ent/templatesectionversion"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BkTemplateSectionCreate is the builder for creating a BkTemplateSection entity.
type BkTemplateSectionCreate struct {
	config
	mutation *BkTemplateSectionMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (btsc *BkTemplateSectionCreate) SetCreatedAt(t time.Time) *BkTemplateSectionCreate {
	btsc.mutation.SetCreatedAt(t)
	return btsc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (btsc *BkTemplateSectionCreate) SetNillableCreatedAt(t *time.Time) *BkTemplateSectionCreate {
	if t != nil {
		btsc.SetCreatedAt(*t)
	}
	return btsc
}

// SetUpdatedAt sets the "updated_at" field.
func (btsc *BkTemplateSectionCreate) SetUpdatedAt(t time.Time) *BkTemplateSectionCreate {
	btsc.mutation.SetUpdatedAt(t)
	return btsc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (btsc *BkTemplateSectionCreate) SetNillableUpdatedAt(t *time.Time) *BkTemplateSectionCreate {
	if t != nil {
		btsc.SetUpdatedAt(*t)
	}
	return btsc
}

// SetVersionID sets the "version_id" field.
func (btsc *BkTemplateSectionCreate) SetVersionID(u uint64) *BkTemplateSectionCreate {
	btsc.mutation.SetVersionID(u)
	return btsc
}

// SetNillableVersionID sets the "version_id" field if the given value is not nil.
func (btsc *BkTemplateSectionCreate) SetNillableVersionID(u *uint64) *BkTemplateSectionCreate {
	if u != nil {
		btsc.SetVersionID(*u)
	}
	return btsc
}

// SetThemeTemplateID sets the "theme_template_id" field.
func (btsc *BkTemplateSectionCreate) SetThemeTemplateID(u uint64) *BkTemplateSectionCreate {
	btsc.mutation.SetThemeTemplateID(u)
	return btsc
}

// SetNillableThemeTemplateID sets the "theme_template_id" field if the given value is not nil.
func (btsc *BkTemplateSectionCreate) SetNillableThemeTemplateID(u *uint64) *BkTemplateSectionCreate {
	if u != nil {
		btsc.SetThemeTemplateID(*u)
	}
	return btsc
}

// SetTemplateSectionID sets the "template_section_id" field.
func (btsc *BkTemplateSectionCreate) SetTemplateSectionID(u uint64) *BkTemplateSectionCreate {
	btsc.mutation.SetTemplateSectionID(u)
	return btsc
}

// SetNillableTemplateSectionID sets the "template_section_id" field if the given value is not nil.
func (btsc *BkTemplateSectionCreate) SetNillableTemplateSectionID(u *uint64) *BkTemplateSectionCreate {
	if u != nil {
		btsc.SetTemplateSectionID(*u)
	}
	return btsc
}

// SetThemeID sets the "theme_id" field.
func (btsc *BkTemplateSectionCreate) SetThemeID(u uint64) *BkTemplateSectionCreate {
	btsc.mutation.SetThemeID(u)
	return btsc
}

// SetNillableThemeID sets the "theme_id" field if the given value is not nil.
func (btsc *BkTemplateSectionCreate) SetNillableThemeID(u *uint64) *BkTemplateSectionCreate {
	if u != nil {
		btsc.SetThemeID(*u)
	}
	return btsc
}

// SetThemeLayoutID sets the "theme_layout_id" field.
func (btsc *BkTemplateSectionCreate) SetThemeLayoutID(u uint64) *BkTemplateSectionCreate {
	btsc.mutation.SetThemeLayoutID(u)
	return btsc
}

// SetNillableThemeLayoutID sets the "theme_layout_id" field if the given value is not nil.
func (btsc *BkTemplateSectionCreate) SetNillableThemeLayoutID(u *uint64) *BkTemplateSectionCreate {
	if u != nil {
		btsc.SetThemeLayoutID(*u)
	}
	return btsc
}

// SetData sets the "data" field.
func (btsc *BkTemplateSectionCreate) SetData(s string) *BkTemplateSectionCreate {
	btsc.mutation.SetData(s)
	return btsc
}

// SetID sets the "id" field.
func (btsc *BkTemplateSectionCreate) SetID(u uint64) *BkTemplateSectionCreate {
	btsc.mutation.SetID(u)
	return btsc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (btsc *BkTemplateSectionCreate) SetNillableID(u *uint64) *BkTemplateSectionCreate {
	if u != nil {
		btsc.SetID(*u)
	}
	return btsc
}

// SetTemplateSection sets the "templateSection" edge to the TemplateSection entity.
func (btsc *BkTemplateSectionCreate) SetTemplateSection(t *TemplateSection) *BkTemplateSectionCreate {
	return btsc.SetTemplateSectionID(t.ID)
}

// SetVersion sets the "version" edge to the TemplateSectionVersion entity.
func (btsc *BkTemplateSectionCreate) SetVersion(t *TemplateSectionVersion) *BkTemplateSectionCreate {
	return btsc.SetVersionID(t.ID)
}

// Mutation returns the BkTemplateSectionMutation object of the builder.
func (btsc *BkTemplateSectionCreate) Mutation() *BkTemplateSectionMutation {
	return btsc.mutation
}

// Save creates the BkTemplateSection in the database.
func (btsc *BkTemplateSectionCreate) Save(ctx context.Context) (*BkTemplateSection, error) {
	var (
		err  error
		node *BkTemplateSection
	)
	btsc.defaults()
	if len(btsc.hooks) == 0 {
		if err = btsc.check(); err != nil {
			return nil, err
		}
		node, err = btsc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BkTemplateSectionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = btsc.check(); err != nil {
				return nil, err
			}
			btsc.mutation = mutation
			if node, err = btsc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(btsc.hooks) - 1; i >= 0; i-- {
			if btsc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = btsc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, btsc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (btsc *BkTemplateSectionCreate) SaveX(ctx context.Context) *BkTemplateSection {
	v, err := btsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (btsc *BkTemplateSectionCreate) Exec(ctx context.Context) error {
	_, err := btsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (btsc *BkTemplateSectionCreate) ExecX(ctx context.Context) {
	if err := btsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (btsc *BkTemplateSectionCreate) defaults() {
	if _, ok := btsc.mutation.CreatedAt(); !ok {
		v := bktemplatesection.DefaultCreatedAt()
		btsc.mutation.SetCreatedAt(v)
	}
	if _, ok := btsc.mutation.UpdatedAt(); !ok {
		v := bktemplatesection.DefaultUpdatedAt()
		btsc.mutation.SetUpdatedAt(v)
	}
	if _, ok := btsc.mutation.ID(); !ok {
		v := bktemplatesection.DefaultID()
		btsc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (btsc *BkTemplateSectionCreate) check() error {
	if _, ok := btsc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := btsc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := btsc.mutation.Data(); !ok {
		return &ValidationError{Name: "data", err: errors.New(`ent: missing required field "data"`)}
	}
	return nil
}

func (btsc *BkTemplateSectionCreate) sqlSave(ctx context.Context) (*BkTemplateSection, error) {
	_node, _spec := btsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, btsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	return _node, nil
}

func (btsc *BkTemplateSectionCreate) createSpec() (*BkTemplateSection, *sqlgraph.CreateSpec) {
	var (
		_node = &BkTemplateSection{config: btsc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: bktemplatesection.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: bktemplatesection.FieldID,
			},
		}
	)
	if id, ok := btsc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := btsc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: bktemplatesection.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := btsc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: bktemplatesection.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := btsc.mutation.ThemeTemplateID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: bktemplatesection.FieldThemeTemplateID,
		})
		_node.ThemeTemplateID = value
	}
	if value, ok := btsc.mutation.ThemeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: bktemplatesection.FieldThemeID,
		})
		_node.ThemeID = value
	}
	if value, ok := btsc.mutation.ThemeLayoutID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: bktemplatesection.FieldThemeLayoutID,
		})
		_node.ThemeLayoutID = value
	}
	if value, ok := btsc.mutation.Data(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bktemplatesection.FieldData,
		})
		_node.Data = value
	}
	if nodes := btsc.mutation.TemplateSectionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bktemplatesection.TemplateSectionTable,
			Columns: []string{bktemplatesection.TemplateSectionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: templatesection.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TemplateSectionID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := btsc.mutation.VersionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bktemplatesection.VersionTable,
			Columns: []string{bktemplatesection.VersionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: templatesectionversion.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.VersionID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BkTemplateSectionCreateBulk is the builder for creating many BkTemplateSection entities in bulk.
type BkTemplateSectionCreateBulk struct {
	config
	builders []*BkTemplateSectionCreate
}

// Save creates the BkTemplateSection entities in the database.
func (btscb *BkTemplateSectionCreateBulk) Save(ctx context.Context) ([]*BkTemplateSection, error) {
	specs := make([]*sqlgraph.CreateSpec, len(btscb.builders))
	nodes := make([]*BkTemplateSection, len(btscb.builders))
	mutators := make([]Mutator, len(btscb.builders))
	for i := range btscb.builders {
		func(i int, root context.Context) {
			builder := btscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BkTemplateSectionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, btscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, btscb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, btscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (btscb *BkTemplateSectionCreateBulk) SaveX(ctx context.Context) []*BkTemplateSection {
	v, err := btscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (btscb *BkTemplateSectionCreateBulk) Exec(ctx context.Context) error {
	_, err := btscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (btscb *BkTemplateSectionCreateBulk) ExecX(ctx context.Context) {
	if err := btscb.Exec(ctx); err != nil {
		panic(err)
	}
}