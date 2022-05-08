// Code generated by entc, DO NOT EDIT.

package ent

import (
	"app-api/ent/bktemplatesection"
	"app-api/ent/predicate"
	"app-api/ent/templatesection"
	"app-api/ent/templatesectionversion"
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BkTemplateSectionQuery is the builder for querying BkTemplateSection entities.
type BkTemplateSectionQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.BkTemplateSection
	// eager-loading edges.
	withTemplateSection *TemplateSectionQuery
	withVersion         *TemplateSectionVersionQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BkTemplateSectionQuery builder.
func (btsq *BkTemplateSectionQuery) Where(ps ...predicate.BkTemplateSection) *BkTemplateSectionQuery {
	btsq.predicates = append(btsq.predicates, ps...)
	return btsq
}

// Limit adds a limit step to the query.
func (btsq *BkTemplateSectionQuery) Limit(limit int) *BkTemplateSectionQuery {
	btsq.limit = &limit
	return btsq
}

// Offset adds an offset step to the query.
func (btsq *BkTemplateSectionQuery) Offset(offset int) *BkTemplateSectionQuery {
	btsq.offset = &offset
	return btsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (btsq *BkTemplateSectionQuery) Unique(unique bool) *BkTemplateSectionQuery {
	btsq.unique = &unique
	return btsq
}

// Order adds an order step to the query.
func (btsq *BkTemplateSectionQuery) Order(o ...OrderFunc) *BkTemplateSectionQuery {
	btsq.order = append(btsq.order, o...)
	return btsq
}

// QueryTemplateSection chains the current query on the "templateSection" edge.
func (btsq *BkTemplateSectionQuery) QueryTemplateSection() *TemplateSectionQuery {
	query := &TemplateSectionQuery{config: btsq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := btsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := btsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(bktemplatesection.Table, bktemplatesection.FieldID, selector),
			sqlgraph.To(templatesection.Table, templatesection.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, bktemplatesection.TemplateSectionTable, bktemplatesection.TemplateSectionColumn),
		)
		fromU = sqlgraph.SetNeighbors(btsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryVersion chains the current query on the "version" edge.
func (btsq *BkTemplateSectionQuery) QueryVersion() *TemplateSectionVersionQuery {
	query := &TemplateSectionVersionQuery{config: btsq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := btsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := btsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(bktemplatesection.Table, bktemplatesection.FieldID, selector),
			sqlgraph.To(templatesectionversion.Table, templatesectionversion.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, bktemplatesection.VersionTable, bktemplatesection.VersionColumn),
		)
		fromU = sqlgraph.SetNeighbors(btsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first BkTemplateSection entity from the query.
// Returns a *NotFoundError when no BkTemplateSection was found.
func (btsq *BkTemplateSectionQuery) First(ctx context.Context) (*BkTemplateSection, error) {
	nodes, err := btsq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{bktemplatesection.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (btsq *BkTemplateSectionQuery) FirstX(ctx context.Context) *BkTemplateSection {
	node, err := btsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BkTemplateSection ID from the query.
// Returns a *NotFoundError when no BkTemplateSection ID was found.
func (btsq *BkTemplateSectionQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = btsq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{bktemplatesection.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (btsq *BkTemplateSectionQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := btsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BkTemplateSection entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one BkTemplateSection entity is not found.
// Returns a *NotFoundError when no BkTemplateSection entities are found.
func (btsq *BkTemplateSectionQuery) Only(ctx context.Context) (*BkTemplateSection, error) {
	nodes, err := btsq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{bktemplatesection.Label}
	default:
		return nil, &NotSingularError{bktemplatesection.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (btsq *BkTemplateSectionQuery) OnlyX(ctx context.Context) *BkTemplateSection {
	node, err := btsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BkTemplateSection ID in the query.
// Returns a *NotSingularError when exactly one BkTemplateSection ID is not found.
// Returns a *NotFoundError when no entities are found.
func (btsq *BkTemplateSectionQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = btsq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{bktemplatesection.Label}
	default:
		err = &NotSingularError{bktemplatesection.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (btsq *BkTemplateSectionQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := btsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BkTemplateSections.
func (btsq *BkTemplateSectionQuery) All(ctx context.Context) ([]*BkTemplateSection, error) {
	if err := btsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return btsq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (btsq *BkTemplateSectionQuery) AllX(ctx context.Context) []*BkTemplateSection {
	nodes, err := btsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BkTemplateSection IDs.
func (btsq *BkTemplateSectionQuery) IDs(ctx context.Context) ([]uint64, error) {
	var ids []uint64
	if err := btsq.Select(bktemplatesection.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (btsq *BkTemplateSectionQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := btsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (btsq *BkTemplateSectionQuery) Count(ctx context.Context) (int, error) {
	if err := btsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return btsq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (btsq *BkTemplateSectionQuery) CountX(ctx context.Context) int {
	count, err := btsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (btsq *BkTemplateSectionQuery) Exist(ctx context.Context) (bool, error) {
	if err := btsq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return btsq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (btsq *BkTemplateSectionQuery) ExistX(ctx context.Context) bool {
	exist, err := btsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BkTemplateSectionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (btsq *BkTemplateSectionQuery) Clone() *BkTemplateSectionQuery {
	if btsq == nil {
		return nil
	}
	return &BkTemplateSectionQuery{
		config:              btsq.config,
		limit:               btsq.limit,
		offset:              btsq.offset,
		order:               append([]OrderFunc{}, btsq.order...),
		predicates:          append([]predicate.BkTemplateSection{}, btsq.predicates...),
		withTemplateSection: btsq.withTemplateSection.Clone(),
		withVersion:         btsq.withVersion.Clone(),
		// clone intermediate query.
		sql:  btsq.sql.Clone(),
		path: btsq.path,
	}
}

// WithTemplateSection tells the query-builder to eager-load the nodes that are connected to
// the "templateSection" edge. The optional arguments are used to configure the query builder of the edge.
func (btsq *BkTemplateSectionQuery) WithTemplateSection(opts ...func(*TemplateSectionQuery)) *BkTemplateSectionQuery {
	query := &TemplateSectionQuery{config: btsq.config}
	for _, opt := range opts {
		opt(query)
	}
	btsq.withTemplateSection = query
	return btsq
}

// WithVersion tells the query-builder to eager-load the nodes that are connected to
// the "version" edge. The optional arguments are used to configure the query builder of the edge.
func (btsq *BkTemplateSectionQuery) WithVersion(opts ...func(*TemplateSectionVersionQuery)) *BkTemplateSectionQuery {
	query := &TemplateSectionVersionQuery{config: btsq.config}
	for _, opt := range opts {
		opt(query)
	}
	btsq.withVersion = query
	return btsq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"createdAt"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BkTemplateSection.Query().
//		GroupBy(bktemplatesection.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (btsq *BkTemplateSectionQuery) GroupBy(field string, fields ...string) *BkTemplateSectionGroupBy {
	group := &BkTemplateSectionGroupBy{config: btsq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := btsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return btsq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"createdAt"`
//	}
//
//	client.BkTemplateSection.Query().
//		Select(bktemplatesection.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (btsq *BkTemplateSectionQuery) Select(fields ...string) *BkTemplateSectionSelect {
	btsq.fields = append(btsq.fields, fields...)
	return &BkTemplateSectionSelect{BkTemplateSectionQuery: btsq}
}

func (btsq *BkTemplateSectionQuery) prepareQuery(ctx context.Context) error {
	for _, f := range btsq.fields {
		if !bktemplatesection.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if btsq.path != nil {
		prev, err := btsq.path(ctx)
		if err != nil {
			return err
		}
		btsq.sql = prev
	}
	return nil
}

func (btsq *BkTemplateSectionQuery) sqlAll(ctx context.Context) ([]*BkTemplateSection, error) {
	var (
		nodes       = []*BkTemplateSection{}
		_spec       = btsq.querySpec()
		loadedTypes = [2]bool{
			btsq.withTemplateSection != nil,
			btsq.withVersion != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &BkTemplateSection{config: btsq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, btsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := btsq.withTemplateSection; query != nil {
		ids := make([]uint64, 0, len(nodes))
		nodeids := make(map[uint64][]*BkTemplateSection)
		for i := range nodes {
			fk := nodes[i].TemplateSectionID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(templatesection.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "template_section_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.TemplateSection = n
			}
		}
	}

	if query := btsq.withVersion; query != nil {
		ids := make([]uint64, 0, len(nodes))
		nodeids := make(map[uint64][]*BkTemplateSection)
		for i := range nodes {
			fk := nodes[i].VersionID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(templatesectionversion.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "version_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Version = n
			}
		}
	}

	return nodes, nil
}

func (btsq *BkTemplateSectionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := btsq.querySpec()
	return sqlgraph.CountNodes(ctx, btsq.driver, _spec)
}

func (btsq *BkTemplateSectionQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := btsq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (btsq *BkTemplateSectionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   bktemplatesection.Table,
			Columns: bktemplatesection.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: bktemplatesection.FieldID,
			},
		},
		From:   btsq.sql,
		Unique: true,
	}
	if unique := btsq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := btsq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bktemplatesection.FieldID)
		for i := range fields {
			if fields[i] != bktemplatesection.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := btsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := btsq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := btsq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := btsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (btsq *BkTemplateSectionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(btsq.driver.Dialect())
	t1 := builder.Table(bktemplatesection.Table)
	columns := btsq.fields
	if len(columns) == 0 {
		columns = bktemplatesection.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if btsq.sql != nil {
		selector = btsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range btsq.predicates {
		p(selector)
	}
	for _, p := range btsq.order {
		p(selector)
	}
	if offset := btsq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := btsq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BkTemplateSectionGroupBy is the group-by builder for BkTemplateSection entities.
type BkTemplateSectionGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (btsgb *BkTemplateSectionGroupBy) Aggregate(fns ...AggregateFunc) *BkTemplateSectionGroupBy {
	btsgb.fns = append(btsgb.fns, fns...)
	return btsgb
}

// Scan applies the group-by query and scans the result into the given value.
func (btsgb *BkTemplateSectionGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := btsgb.path(ctx)
	if err != nil {
		return err
	}
	btsgb.sql = query
	return btsgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (btsgb *BkTemplateSectionGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := btsgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (btsgb *BkTemplateSectionGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(btsgb.fields) > 1 {
		return nil, errors.New("ent: BkTemplateSectionGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := btsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (btsgb *BkTemplateSectionGroupBy) StringsX(ctx context.Context) []string {
	v, err := btsgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (btsgb *BkTemplateSectionGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = btsgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bktemplatesection.Label}
	default:
		err = fmt.Errorf("ent: BkTemplateSectionGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (btsgb *BkTemplateSectionGroupBy) StringX(ctx context.Context) string {
	v, err := btsgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (btsgb *BkTemplateSectionGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(btsgb.fields) > 1 {
		return nil, errors.New("ent: BkTemplateSectionGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := btsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (btsgb *BkTemplateSectionGroupBy) IntsX(ctx context.Context) []int {
	v, err := btsgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (btsgb *BkTemplateSectionGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = btsgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bktemplatesection.Label}
	default:
		err = fmt.Errorf("ent: BkTemplateSectionGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (btsgb *BkTemplateSectionGroupBy) IntX(ctx context.Context) int {
	v, err := btsgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (btsgb *BkTemplateSectionGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(btsgb.fields) > 1 {
		return nil, errors.New("ent: BkTemplateSectionGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := btsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (btsgb *BkTemplateSectionGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := btsgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (btsgb *BkTemplateSectionGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = btsgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bktemplatesection.Label}
	default:
		err = fmt.Errorf("ent: BkTemplateSectionGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (btsgb *BkTemplateSectionGroupBy) Float64X(ctx context.Context) float64 {
	v, err := btsgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (btsgb *BkTemplateSectionGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(btsgb.fields) > 1 {
		return nil, errors.New("ent: BkTemplateSectionGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := btsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (btsgb *BkTemplateSectionGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := btsgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (btsgb *BkTemplateSectionGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = btsgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bktemplatesection.Label}
	default:
		err = fmt.Errorf("ent: BkTemplateSectionGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (btsgb *BkTemplateSectionGroupBy) BoolX(ctx context.Context) bool {
	v, err := btsgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (btsgb *BkTemplateSectionGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range btsgb.fields {
		if !bktemplatesection.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := btsgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := btsgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (btsgb *BkTemplateSectionGroupBy) sqlQuery() *sql.Selector {
	selector := btsgb.sql.Select()
	aggregation := make([]string, 0, len(btsgb.fns))
	for _, fn := range btsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(btsgb.fields)+len(btsgb.fns))
		for _, f := range btsgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(btsgb.fields...)...)
}

// BkTemplateSectionSelect is the builder for selecting fields of BkTemplateSection entities.
type BkTemplateSectionSelect struct {
	*BkTemplateSectionQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (btss *BkTemplateSectionSelect) Scan(ctx context.Context, v interface{}) error {
	if err := btss.prepareQuery(ctx); err != nil {
		return err
	}
	btss.sql = btss.BkTemplateSectionQuery.sqlQuery(ctx)
	return btss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (btss *BkTemplateSectionSelect) ScanX(ctx context.Context, v interface{}) {
	if err := btss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (btss *BkTemplateSectionSelect) Strings(ctx context.Context) ([]string, error) {
	if len(btss.fields) > 1 {
		return nil, errors.New("ent: BkTemplateSectionSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := btss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (btss *BkTemplateSectionSelect) StringsX(ctx context.Context) []string {
	v, err := btss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (btss *BkTemplateSectionSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = btss.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bktemplatesection.Label}
	default:
		err = fmt.Errorf("ent: BkTemplateSectionSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (btss *BkTemplateSectionSelect) StringX(ctx context.Context) string {
	v, err := btss.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (btss *BkTemplateSectionSelect) Ints(ctx context.Context) ([]int, error) {
	if len(btss.fields) > 1 {
		return nil, errors.New("ent: BkTemplateSectionSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := btss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (btss *BkTemplateSectionSelect) IntsX(ctx context.Context) []int {
	v, err := btss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (btss *BkTemplateSectionSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = btss.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bktemplatesection.Label}
	default:
		err = fmt.Errorf("ent: BkTemplateSectionSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (btss *BkTemplateSectionSelect) IntX(ctx context.Context) int {
	v, err := btss.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (btss *BkTemplateSectionSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(btss.fields) > 1 {
		return nil, errors.New("ent: BkTemplateSectionSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := btss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (btss *BkTemplateSectionSelect) Float64sX(ctx context.Context) []float64 {
	v, err := btss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (btss *BkTemplateSectionSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = btss.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bktemplatesection.Label}
	default:
		err = fmt.Errorf("ent: BkTemplateSectionSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (btss *BkTemplateSectionSelect) Float64X(ctx context.Context) float64 {
	v, err := btss.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (btss *BkTemplateSectionSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(btss.fields) > 1 {
		return nil, errors.New("ent: BkTemplateSectionSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := btss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (btss *BkTemplateSectionSelect) BoolsX(ctx context.Context) []bool {
	v, err := btss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (btss *BkTemplateSectionSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = btss.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bktemplatesection.Label}
	default:
		err = fmt.Errorf("ent: BkTemplateSectionSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (btss *BkTemplateSectionSelect) BoolX(ctx context.Context) bool {
	v, err := btss.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (btss *BkTemplateSectionSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := btss.sql.Query()
	if err := btss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}