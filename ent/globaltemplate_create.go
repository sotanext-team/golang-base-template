// Code generated by entc, DO NOT EDIT.

package ent

import (
	"app-api/ent/globaltemplate"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GlobalTemplateCreate is the builder for creating a GlobalTemplate entity.
type GlobalTemplateCreate struct {
	config
	mutation *GlobalTemplateMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (gtc *GlobalTemplateCreate) SetCreatedAt(t time.Time) *GlobalTemplateCreate {
	gtc.mutation.SetCreatedAt(t)
	return gtc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gtc *GlobalTemplateCreate) SetNillableCreatedAt(t *time.Time) *GlobalTemplateCreate {
	if t != nil {
		gtc.SetCreatedAt(*t)
	}
	return gtc
}

// SetUpdatedAt sets the "updated_at" field.
func (gtc *GlobalTemplateCreate) SetUpdatedAt(t time.Time) *GlobalTemplateCreate {
	gtc.mutation.SetUpdatedAt(t)
	return gtc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (gtc *GlobalTemplateCreate) SetNillableUpdatedAt(t *time.Time) *GlobalTemplateCreate {
	if t != nil {
		gtc.SetUpdatedAt(*t)
	}
	return gtc
}

// SetDeletedAt sets the "deleted_at" field.
func (gtc *GlobalTemplateCreate) SetDeletedAt(t time.Time) *GlobalTemplateCreate {
	gtc.mutation.SetDeletedAt(t)
	return gtc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gtc *GlobalTemplateCreate) SetNillableDeletedAt(t *time.Time) *GlobalTemplateCreate {
	if t != nil {
		gtc.SetDeletedAt(*t)
	}
	return gtc
}

// SetShopID sets the "shop_id" field.
func (gtc *GlobalTemplateCreate) SetShopID(u uint64) *GlobalTemplateCreate {
	gtc.mutation.SetShopID(u)
	return gtc
}

// SetNillableShopID sets the "shop_id" field if the given value is not nil.
func (gtc *GlobalTemplateCreate) SetNillableShopID(u *uint64) *GlobalTemplateCreate {
	if u != nil {
		gtc.SetShopID(*u)
	}
	return gtc
}

// SetName sets the "name" field.
func (gtc *GlobalTemplateCreate) SetName(s string) *GlobalTemplateCreate {
	gtc.mutation.SetName(s)
	return gtc
}

// SetViewCount sets the "view_count" field.
func (gtc *GlobalTemplateCreate) SetViewCount(i int) *GlobalTemplateCreate {
	gtc.mutation.SetViewCount(i)
	return gtc
}

// SetInstallCount sets the "install_count" field.
func (gtc *GlobalTemplateCreate) SetInstallCount(i int) *GlobalTemplateCreate {
	gtc.mutation.SetInstallCount(i)
	return gtc
}

// SetID sets the "id" field.
func (gtc *GlobalTemplateCreate) SetID(u uint64) *GlobalTemplateCreate {
	gtc.mutation.SetID(u)
	return gtc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (gtc *GlobalTemplateCreate) SetNillableID(u *uint64) *GlobalTemplateCreate {
	if u != nil {
		gtc.SetID(*u)
	}
	return gtc
}

// Mutation returns the GlobalTemplateMutation object of the builder.
func (gtc *GlobalTemplateCreate) Mutation() *GlobalTemplateMutation {
	return gtc.mutation
}

// Save creates the GlobalTemplate in the database.
func (gtc *GlobalTemplateCreate) Save(ctx context.Context) (*GlobalTemplate, error) {
	var (
		err  error
		node *GlobalTemplate
	)
	gtc.defaults()
	if len(gtc.hooks) == 0 {
		if err = gtc.check(); err != nil {
			return nil, err
		}
		node, err = gtc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GlobalTemplateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gtc.check(); err != nil {
				return nil, err
			}
			gtc.mutation = mutation
			if node, err = gtc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(gtc.hooks) - 1; i >= 0; i-- {
			if gtc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gtc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gtc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gtc *GlobalTemplateCreate) SaveX(ctx context.Context) *GlobalTemplate {
	v, err := gtc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gtc *GlobalTemplateCreate) Exec(ctx context.Context) error {
	_, err := gtc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gtc *GlobalTemplateCreate) ExecX(ctx context.Context) {
	if err := gtc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gtc *GlobalTemplateCreate) defaults() {
	if _, ok := gtc.mutation.CreatedAt(); !ok {
		v := globaltemplate.DefaultCreatedAt()
		gtc.mutation.SetCreatedAt(v)
	}
	if _, ok := gtc.mutation.UpdatedAt(); !ok {
		v := globaltemplate.DefaultUpdatedAt()
		gtc.mutation.SetUpdatedAt(v)
	}
	if _, ok := gtc.mutation.ID(); !ok {
		v := globaltemplate.DefaultID()
		gtc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gtc *GlobalTemplateCreate) check() error {
	if _, ok := gtc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := gtc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := gtc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if v, ok := gtc.mutation.Name(); ok {
		if err := globaltemplate.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "name": %w`, err)}
		}
	}
	if _, ok := gtc.mutation.ViewCount(); !ok {
		return &ValidationError{Name: "view_count", err: errors.New(`ent: missing required field "view_count"`)}
	}
	if _, ok := gtc.mutation.InstallCount(); !ok {
		return &ValidationError{Name: "install_count", err: errors.New(`ent: missing required field "install_count"`)}
	}
	return nil
}

func (gtc *GlobalTemplateCreate) sqlSave(ctx context.Context) (*GlobalTemplate, error) {
	_node, _spec := gtc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gtc.driver, _spec); err != nil {
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

func (gtc *GlobalTemplateCreate) createSpec() (*GlobalTemplate, *sqlgraph.CreateSpec) {
	var (
		_node = &GlobalTemplate{config: gtc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: globaltemplate.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: globaltemplate.FieldID,
			},
		}
	)
	if id, ok := gtc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := gtc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: globaltemplate.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := gtc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: globaltemplate.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := gtc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: globaltemplate.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if value, ok := gtc.mutation.ShopID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: globaltemplate.FieldShopID,
		})
		_node.ShopID = value
	}
	if value, ok := gtc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: globaltemplate.FieldName,
		})
		_node.Name = value
	}
	if value, ok := gtc.mutation.ViewCount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: globaltemplate.FieldViewCount,
		})
		_node.ViewCount = value
	}
	if value, ok := gtc.mutation.InstallCount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: globaltemplate.FieldInstallCount,
		})
		_node.InstallCount = value
	}
	return _node, _spec
}

// GlobalTemplateCreateBulk is the builder for creating many GlobalTemplate entities in bulk.
type GlobalTemplateCreateBulk struct {
	config
	builders []*GlobalTemplateCreate
}

// Save creates the GlobalTemplate entities in the database.
func (gtcb *GlobalTemplateCreateBulk) Save(ctx context.Context) ([]*GlobalTemplate, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gtcb.builders))
	nodes := make([]*GlobalTemplate, len(gtcb.builders))
	mutators := make([]Mutator, len(gtcb.builders))
	for i := range gtcb.builders {
		func(i int, root context.Context) {
			builder := gtcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GlobalTemplateMutation)
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
					_, err = mutators[i+1].Mutate(root, gtcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gtcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, gtcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gtcb *GlobalTemplateCreateBulk) SaveX(ctx context.Context) []*GlobalTemplate {
	v, err := gtcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gtcb *GlobalTemplateCreateBulk) Exec(ctx context.Context) error {
	_, err := gtcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gtcb *GlobalTemplateCreateBulk) ExecX(ctx context.Context) {
	if err := gtcb.Exec(ctx); err != nil {
		panic(err)
	}
}