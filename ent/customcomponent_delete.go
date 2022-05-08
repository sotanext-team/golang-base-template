// Code generated by entc, DO NOT EDIT.

package ent

import (
	"app-api/ent/customcomponent"
	"app-api/ent/predicate"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CustomComponentDelete is the builder for deleting a CustomComponent entity.
type CustomComponentDelete struct {
	config
	hooks    []Hook
	mutation *CustomComponentMutation
}

// Where appends a list predicates to the CustomComponentDelete builder.
func (ccd *CustomComponentDelete) Where(ps ...predicate.CustomComponent) *CustomComponentDelete {
	ccd.mutation.Where(ps...)
	return ccd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ccd *CustomComponentDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ccd.hooks) == 0 {
		affected, err = ccd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CustomComponentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ccd.mutation = mutation
			affected, err = ccd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ccd.hooks) - 1; i >= 0; i-- {
			if ccd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ccd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ccd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccd *CustomComponentDelete) ExecX(ctx context.Context) int {
	n, err := ccd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ccd *CustomComponentDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: customcomponent.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: customcomponent.FieldID,
			},
		},
	}
	if ps := ccd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ccd.driver, _spec)
}

// CustomComponentDeleteOne is the builder for deleting a single CustomComponent entity.
type CustomComponentDeleteOne struct {
	ccd *CustomComponentDelete
}

// Exec executes the deletion query.
func (ccdo *CustomComponentDeleteOne) Exec(ctx context.Context) error {
	n, err := ccdo.ccd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{customcomponent.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ccdo *CustomComponentDeleteOne) ExecX(ctx context.Context) {
	ccdo.ccd.ExecX(ctx)
}