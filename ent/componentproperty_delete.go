// Code generated by entc, DO NOT EDIT.

package ent

import (
	"app-api/ent/componentproperty"
	"app-api/ent/predicate"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ComponentPropertyDelete is the builder for deleting a ComponentProperty entity.
type ComponentPropertyDelete struct {
	config
	hooks    []Hook
	mutation *ComponentPropertyMutation
}

// Where appends a list predicates to the ComponentPropertyDelete builder.
func (cpd *ComponentPropertyDelete) Where(ps ...predicate.ComponentProperty) *ComponentPropertyDelete {
	cpd.mutation.Where(ps...)
	return cpd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cpd *ComponentPropertyDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cpd.hooks) == 0 {
		affected, err = cpd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ComponentPropertyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cpd.mutation = mutation
			affected, err = cpd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cpd.hooks) - 1; i >= 0; i-- {
			if cpd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cpd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cpd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (cpd *ComponentPropertyDelete) ExecX(ctx context.Context) int {
	n, err := cpd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cpd *ComponentPropertyDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: componentproperty.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: componentproperty.FieldID,
			},
		},
	}
	if ps := cpd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, cpd.driver, _spec)
}

// ComponentPropertyDeleteOne is the builder for deleting a single ComponentProperty entity.
type ComponentPropertyDeleteOne struct {
	cpd *ComponentPropertyDelete
}

// Exec executes the deletion query.
func (cpdo *ComponentPropertyDeleteOne) Exec(ctx context.Context) error {
	n, err := cpdo.cpd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{componentproperty.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cpdo *ComponentPropertyDeleteOne) ExecX(ctx context.Context) {
	cpdo.cpd.ExecX(ctx)
}