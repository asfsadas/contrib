// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/asfsadas/contrib/schemast/internal/mutatetest/ent/predicate"
	"github.com/asfsadas/contrib/schemast/internal/mutatetest/ent/withoutfields"
)

// WithoutFieldsUpdate is the builder for updating WithoutFields entities.
type WithoutFieldsUpdate struct {
	config
	hooks    []Hook
	mutation *WithoutFieldsMutation
}

// Where appends a list predicates to the WithoutFieldsUpdate builder.
func (wfu *WithoutFieldsUpdate) Where(ps ...predicate.WithoutFields) *WithoutFieldsUpdate {
	wfu.mutation.Where(ps...)
	return wfu
}

// Mutation returns the WithoutFieldsMutation object of the builder.
func (wfu *WithoutFieldsUpdate) Mutation() *WithoutFieldsMutation {
	return wfu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wfu *WithoutFieldsUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(wfu.hooks) == 0 {
		affected, err = wfu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WithoutFieldsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wfu.mutation = mutation
			affected, err = wfu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wfu.hooks) - 1; i >= 0; i-- {
			if wfu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wfu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wfu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (wfu *WithoutFieldsUpdate) SaveX(ctx context.Context) int {
	affected, err := wfu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wfu *WithoutFieldsUpdate) Exec(ctx context.Context) error {
	_, err := wfu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wfu *WithoutFieldsUpdate) ExecX(ctx context.Context) {
	if err := wfu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wfu *WithoutFieldsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   withoutfields.Table,
			Columns: withoutfields.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: withoutfields.FieldID,
			},
		},
	}
	if ps := wfu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wfu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{withoutfields.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// WithoutFieldsUpdateOne is the builder for updating a single WithoutFields entity.
type WithoutFieldsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WithoutFieldsMutation
}

// Mutation returns the WithoutFieldsMutation object of the builder.
func (wfuo *WithoutFieldsUpdateOne) Mutation() *WithoutFieldsMutation {
	return wfuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wfuo *WithoutFieldsUpdateOne) Select(field string, fields ...string) *WithoutFieldsUpdateOne {
	wfuo.fields = append([]string{field}, fields...)
	return wfuo
}

// Save executes the query and returns the updated WithoutFields entity.
func (wfuo *WithoutFieldsUpdateOne) Save(ctx context.Context) (*WithoutFields, error) {
	var (
		err  error
		node *WithoutFields
	)
	if len(wfuo.hooks) == 0 {
		node, err = wfuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WithoutFieldsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wfuo.mutation = mutation
			node, err = wfuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wfuo.hooks) - 1; i >= 0; i-- {
			if wfuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wfuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wfuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (wfuo *WithoutFieldsUpdateOne) SaveX(ctx context.Context) *WithoutFields {
	node, err := wfuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wfuo *WithoutFieldsUpdateOne) Exec(ctx context.Context) error {
	_, err := wfuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wfuo *WithoutFieldsUpdateOne) ExecX(ctx context.Context) {
	if err := wfuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wfuo *WithoutFieldsUpdateOne) sqlSave(ctx context.Context) (_node *WithoutFields, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   withoutfields.Table,
			Columns: withoutfields.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: withoutfields.FieldID,
			},
		},
	}
	id, ok := wfuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "WithoutFields.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wfuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, withoutfields.FieldID)
		for _, f := range fields {
			if !withoutfields.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != withoutfields.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wfuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_node = &WithoutFields{config: wfuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wfuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{withoutfields.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
