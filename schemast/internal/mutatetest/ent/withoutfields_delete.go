// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/asfsadas/contrib/schemast/internal/mutatetest/ent/predicate"
	"github.com/asfsadas/contrib/schemast/internal/mutatetest/ent/withoutfields"
)

// WithoutFieldsDelete is the builder for deleting a WithoutFields entity.
type WithoutFieldsDelete struct {
	config
	hooks    []Hook
	mutation *WithoutFieldsMutation
}

// Where appends a list predicates to the WithoutFieldsDelete builder.
func (wfd *WithoutFieldsDelete) Where(ps ...predicate.WithoutFields) *WithoutFieldsDelete {
	wfd.mutation.Where(ps...)
	return wfd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wfd *WithoutFieldsDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(wfd.hooks) == 0 {
		affected, err = wfd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WithoutFieldsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wfd.mutation = mutation
			affected, err = wfd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wfd.hooks) - 1; i >= 0; i-- {
			if wfd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wfd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wfd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (wfd *WithoutFieldsDelete) ExecX(ctx context.Context) int {
	n, err := wfd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wfd *WithoutFieldsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: withoutfields.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: withoutfields.FieldID,
			},
		},
	}
	if ps := wfd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, wfd.driver, _spec)
}

// WithoutFieldsDeleteOne is the builder for deleting a single WithoutFields entity.
type WithoutFieldsDeleteOne struct {
	wfd *WithoutFieldsDelete
}

// Exec executes the deletion query.
func (wfdo *WithoutFieldsDeleteOne) Exec(ctx context.Context) error {
	n, err := wfdo.wfd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{withoutfields.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wfdo *WithoutFieldsDeleteOne) ExecX(ctx context.Context) {
	wfdo.wfd.ExecX(ctx)
}
