// Code generated by entc, DO NOT EDIT.

package ent

import (
	"app-api/ent/predicate"
	"app-api/ent/shop"
	"app-api/ent/theme"
	"app-api/ent/themetemplate"
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ThemeUpdate is the builder for updating Theme entities.
type ThemeUpdate struct {
	config
	hooks    []Hook
	mutation *ThemeMutation
}

// Where appends a list predicates to the ThemeUpdate builder.
func (tu *ThemeUpdate) Where(ps ...predicate.Theme) *ThemeUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetUpdatedAt sets the "updated_at" field.
func (tu *ThemeUpdate) SetUpdatedAt(t time.Time) *ThemeUpdate {
	tu.mutation.SetUpdatedAt(t)
	return tu
}

// SetDeletedAt sets the "deleted_at" field.
func (tu *ThemeUpdate) SetDeletedAt(t time.Time) *ThemeUpdate {
	tu.mutation.SetDeletedAt(t)
	return tu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tu *ThemeUpdate) SetNillableDeletedAt(t *time.Time) *ThemeUpdate {
	if t != nil {
		tu.SetDeletedAt(*t)
	}
	return tu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (tu *ThemeUpdate) ClearDeletedAt() *ThemeUpdate {
	tu.mutation.ClearDeletedAt()
	return tu
}

// SetShopID sets the "shop_id" field.
func (tu *ThemeUpdate) SetShopID(u uint64) *ThemeUpdate {
	tu.mutation.SetShopID(u)
	return tu
}

// SetNillableShopID sets the "shop_id" field if the given value is not nil.
func (tu *ThemeUpdate) SetNillableShopID(u *uint64) *ThemeUpdate {
	if u != nil {
		tu.SetShopID(*u)
	}
	return tu
}

// ClearShopID clears the value of the "shop_id" field.
func (tu *ThemeUpdate) ClearShopID() *ThemeUpdate {
	tu.mutation.ClearShopID()
	return tu
}

// SetName sets the "name" field.
func (tu *ThemeUpdate) SetName(s string) *ThemeUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetThumbnail sets the "thumbnail" field.
func (tu *ThemeUpdate) SetThumbnail(s string) *ThemeUpdate {
	tu.mutation.SetThumbnail(s)
	return tu
}

// SetPublish sets the "publish" field.
func (tu *ThemeUpdate) SetPublish(b bool) *ThemeUpdate {
	tu.mutation.SetPublish(b)
	return tu
}

// SetNillablePublish sets the "publish" field if the given value is not nil.
func (tu *ThemeUpdate) SetNillablePublish(b *bool) *ThemeUpdate {
	if b != nil {
		tu.SetPublish(*b)
	}
	return tu
}

// AddThemeTemplateIDs adds the "themeTemplates" edge to the ThemeTemplate entity by IDs.
func (tu *ThemeUpdate) AddThemeTemplateIDs(ids ...uint64) *ThemeUpdate {
	tu.mutation.AddThemeTemplateIDs(ids...)
	return tu
}

// AddThemeTemplates adds the "themeTemplates" edges to the ThemeTemplate entity.
func (tu *ThemeUpdate) AddThemeTemplates(t ...*ThemeTemplate) *ThemeUpdate {
	ids := make([]uint64, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.AddThemeTemplateIDs(ids...)
}

// SetShop sets the "shop" edge to the Shop entity.
func (tu *ThemeUpdate) SetShop(s *Shop) *ThemeUpdate {
	return tu.SetShopID(s.ID)
}

// Mutation returns the ThemeMutation object of the builder.
func (tu *ThemeUpdate) Mutation() *ThemeMutation {
	return tu.mutation
}

// ClearThemeTemplates clears all "themeTemplates" edges to the ThemeTemplate entity.
func (tu *ThemeUpdate) ClearThemeTemplates() *ThemeUpdate {
	tu.mutation.ClearThemeTemplates()
	return tu
}

// RemoveThemeTemplateIDs removes the "themeTemplates" edge to ThemeTemplate entities by IDs.
func (tu *ThemeUpdate) RemoveThemeTemplateIDs(ids ...uint64) *ThemeUpdate {
	tu.mutation.RemoveThemeTemplateIDs(ids...)
	return tu
}

// RemoveThemeTemplates removes "themeTemplates" edges to ThemeTemplate entities.
func (tu *ThemeUpdate) RemoveThemeTemplates(t ...*ThemeTemplate) *ThemeUpdate {
	ids := make([]uint64, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.RemoveThemeTemplateIDs(ids...)
}

// ClearShop clears the "shop" edge to the Shop entity.
func (tu *ThemeUpdate) ClearShop() *ThemeUpdate {
	tu.mutation.ClearShop()
	return tu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *ThemeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	tu.defaults()
	if len(tu.hooks) == 0 {
		if err = tu.check(); err != nil {
			return 0, err
		}
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ThemeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tu.check(); err != nil {
				return 0, err
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *ThemeUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *ThemeUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *ThemeUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *ThemeUpdate) defaults() {
	if _, ok := tu.mutation.UpdatedAt(); !ok {
		v := theme.UpdateDefaultUpdatedAt()
		tu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *ThemeUpdate) check() error {
	if v, ok := tu.mutation.Name(); ok {
		if err := theme.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := tu.mutation.Thumbnail(); ok {
		if err := theme.ThumbnailValidator(v); err != nil {
			return &ValidationError{Name: "thumbnail", err: fmt.Errorf("ent: validator failed for field \"thumbnail\": %w", err)}
		}
	}
	return nil
}

func (tu *ThemeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   theme.Table,
			Columns: theme.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: theme.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: theme.FieldUpdatedAt,
		})
	}
	if value, ok := tu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: theme.FieldDeletedAt,
		})
	}
	if tu.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: theme.FieldDeletedAt,
		})
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: theme.FieldName,
		})
	}
	if value, ok := tu.mutation.Thumbnail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: theme.FieldThumbnail,
		})
	}
	if value, ok := tu.mutation.Publish(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: theme.FieldPublish,
		})
	}
	if tu.mutation.ThemeTemplatesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   theme.ThemeTemplatesTable,
			Columns: []string{theme.ThemeTemplatesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: themetemplate.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedThemeTemplatesIDs(); len(nodes) > 0 && !tu.mutation.ThemeTemplatesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   theme.ThemeTemplatesTable,
			Columns: []string{theme.ThemeTemplatesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: themetemplate.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.ThemeTemplatesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   theme.ThemeTemplatesTable,
			Columns: []string{theme.ThemeTemplatesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: themetemplate.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.ShopCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   theme.ShopTable,
			Columns: []string{theme.ShopColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: shop.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.ShopIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   theme.ShopTable,
			Columns: []string{theme.ShopColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: shop.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{theme.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ThemeUpdateOne is the builder for updating a single Theme entity.
type ThemeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ThemeMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (tuo *ThemeUpdateOne) SetUpdatedAt(t time.Time) *ThemeUpdateOne {
	tuo.mutation.SetUpdatedAt(t)
	return tuo
}

// SetDeletedAt sets the "deleted_at" field.
func (tuo *ThemeUpdateOne) SetDeletedAt(t time.Time) *ThemeUpdateOne {
	tuo.mutation.SetDeletedAt(t)
	return tuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tuo *ThemeUpdateOne) SetNillableDeletedAt(t *time.Time) *ThemeUpdateOne {
	if t != nil {
		tuo.SetDeletedAt(*t)
	}
	return tuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (tuo *ThemeUpdateOne) ClearDeletedAt() *ThemeUpdateOne {
	tuo.mutation.ClearDeletedAt()
	return tuo
}

// SetShopID sets the "shop_id" field.
func (tuo *ThemeUpdateOne) SetShopID(u uint64) *ThemeUpdateOne {
	tuo.mutation.SetShopID(u)
	return tuo
}

// SetNillableShopID sets the "shop_id" field if the given value is not nil.
func (tuo *ThemeUpdateOne) SetNillableShopID(u *uint64) *ThemeUpdateOne {
	if u != nil {
		tuo.SetShopID(*u)
	}
	return tuo
}

// ClearShopID clears the value of the "shop_id" field.
func (tuo *ThemeUpdateOne) ClearShopID() *ThemeUpdateOne {
	tuo.mutation.ClearShopID()
	return tuo
}

// SetName sets the "name" field.
func (tuo *ThemeUpdateOne) SetName(s string) *ThemeUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetThumbnail sets the "thumbnail" field.
func (tuo *ThemeUpdateOne) SetThumbnail(s string) *ThemeUpdateOne {
	tuo.mutation.SetThumbnail(s)
	return tuo
}

// SetPublish sets the "publish" field.
func (tuo *ThemeUpdateOne) SetPublish(b bool) *ThemeUpdateOne {
	tuo.mutation.SetPublish(b)
	return tuo
}

// SetNillablePublish sets the "publish" field if the given value is not nil.
func (tuo *ThemeUpdateOne) SetNillablePublish(b *bool) *ThemeUpdateOne {
	if b != nil {
		tuo.SetPublish(*b)
	}
	return tuo
}

// AddThemeTemplateIDs adds the "themeTemplates" edge to the ThemeTemplate entity by IDs.
func (tuo *ThemeUpdateOne) AddThemeTemplateIDs(ids ...uint64) *ThemeUpdateOne {
	tuo.mutation.AddThemeTemplateIDs(ids...)
	return tuo
}

// AddThemeTemplates adds the "themeTemplates" edges to the ThemeTemplate entity.
func (tuo *ThemeUpdateOne) AddThemeTemplates(t ...*ThemeTemplate) *ThemeUpdateOne {
	ids := make([]uint64, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.AddThemeTemplateIDs(ids...)
}

// SetShop sets the "shop" edge to the Shop entity.
func (tuo *ThemeUpdateOne) SetShop(s *Shop) *ThemeUpdateOne {
	return tuo.SetShopID(s.ID)
}

// Mutation returns the ThemeMutation object of the builder.
func (tuo *ThemeUpdateOne) Mutation() *ThemeMutation {
	return tuo.mutation
}

// ClearThemeTemplates clears all "themeTemplates" edges to the ThemeTemplate entity.
func (tuo *ThemeUpdateOne) ClearThemeTemplates() *ThemeUpdateOne {
	tuo.mutation.ClearThemeTemplates()
	return tuo
}

// RemoveThemeTemplateIDs removes the "themeTemplates" edge to ThemeTemplate entities by IDs.
func (tuo *ThemeUpdateOne) RemoveThemeTemplateIDs(ids ...uint64) *ThemeUpdateOne {
	tuo.mutation.RemoveThemeTemplateIDs(ids...)
	return tuo
}

// RemoveThemeTemplates removes "themeTemplates" edges to ThemeTemplate entities.
func (tuo *ThemeUpdateOne) RemoveThemeTemplates(t ...*ThemeTemplate) *ThemeUpdateOne {
	ids := make([]uint64, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.RemoveThemeTemplateIDs(ids...)
}

// ClearShop clears the "shop" edge to the Shop entity.
func (tuo *ThemeUpdateOne) ClearShop() *ThemeUpdateOne {
	tuo.mutation.ClearShop()
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *ThemeUpdateOne) Select(field string, fields ...string) *ThemeUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Theme entity.
func (tuo *ThemeUpdateOne) Save(ctx context.Context) (*Theme, error) {
	var (
		err  error
		node *Theme
	)
	tuo.defaults()
	if len(tuo.hooks) == 0 {
		if err = tuo.check(); err != nil {
			return nil, err
		}
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ThemeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuo.check(); err != nil {
				return nil, err
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *ThemeUpdateOne) SaveX(ctx context.Context) *Theme {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *ThemeUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *ThemeUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *ThemeUpdateOne) defaults() {
	if _, ok := tuo.mutation.UpdatedAt(); !ok {
		v := theme.UpdateDefaultUpdatedAt()
		tuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *ThemeUpdateOne) check() error {
	if v, ok := tuo.mutation.Name(); ok {
		if err := theme.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := tuo.mutation.Thumbnail(); ok {
		if err := theme.ThumbnailValidator(v); err != nil {
			return &ValidationError{Name: "thumbnail", err: fmt.Errorf("ent: validator failed for field \"thumbnail\": %w", err)}
		}
	}
	return nil
}

func (tuo *ThemeUpdateOne) sqlSave(ctx context.Context) (_node *Theme, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   theme.Table,
			Columns: theme.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: theme.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Theme.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, theme.FieldID)
		for _, f := range fields {
			if !theme.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != theme.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: theme.FieldUpdatedAt,
		})
	}
	if value, ok := tuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: theme.FieldDeletedAt,
		})
	}
	if tuo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: theme.FieldDeletedAt,
		})
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: theme.FieldName,
		})
	}
	if value, ok := tuo.mutation.Thumbnail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: theme.FieldThumbnail,
		})
	}
	if value, ok := tuo.mutation.Publish(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: theme.FieldPublish,
		})
	}
	if tuo.mutation.ThemeTemplatesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   theme.ThemeTemplatesTable,
			Columns: []string{theme.ThemeTemplatesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: themetemplate.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedThemeTemplatesIDs(); len(nodes) > 0 && !tuo.mutation.ThemeTemplatesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   theme.ThemeTemplatesTable,
			Columns: []string{theme.ThemeTemplatesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: themetemplate.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.ThemeTemplatesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   theme.ThemeTemplatesTable,
			Columns: []string{theme.ThemeTemplatesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: themetemplate.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.ShopCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   theme.ShopTable,
			Columns: []string{theme.ShopColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: shop.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.ShopIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   theme.ShopTable,
			Columns: []string{theme.ShopColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: shop.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Theme{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{theme.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
