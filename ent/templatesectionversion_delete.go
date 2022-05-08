// Code generated by entc, DO NOT EDIT.

package ent

import (
	"app-api/ent/predicate"
	"app-api/ent/templatesectionversion"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TemplateSectionVersionDelete is the builder for deleting a TemplateSectionVersion entity.
type TemplateSectionVersionDelete struct {
	config
	hooks    []Hook
	mutation *TemplateSectionVersionMutation
}

// Where appends a list predicates to the TemplateSectionVersionDelete builder.
func (tsvd *TemplateSectionVersionDelete) Where(ps ...predicate.TemplateSectionVersion) *TemplateSectionVersionDelete {
	tsvd.mutation.Where(ps...)
	return tsvd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tsvd *TemplateSectionVersionDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tsvd.hooks) == 0 {
		affected, err = tsvd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TemplateSectionVersionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tsvd.mutation = mutation
			affected, err = tsvd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tsvd.hooks) - 1; i >= 0; i-- {
			if tsvd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tsvd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tsvd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (tsvd *TemplateSectionVersionDelete) ExecX(ctx context.Context) int {
	n, err := tsvd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tsvd *TemplateSectionVersionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: templatesectionversion.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: templatesectionversion.FieldID,
			},
		},
	}
	if ps := tsvd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, tsvd.driver, _spec)
}

// TemplateSectionVersionDeleteOne is the builder for deleting a single TemplateSectionVersion entity.
type TemplateSectionVersionDeleteOne struct {
	tsvd *TemplateSectionVersionDelete
}

// Exec executes the deletion query.
func (tsvdo *TemplateSectionVersionDeleteOne) Exec(ctx context.Context) error {
	n, err := tsvdo.tsvd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{templatesectionversion.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tsvdo *TemplateSectionVersionDeleteOne) ExecX(ctx context.Context) {
	tsvdo.tsvd.ExecX(ctx)
}
