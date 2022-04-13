// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/asfsadas/contrib/entproto/internal/entprototest/ent/allmethodsservice"
	"github.com/asfsadas/contrib/entproto/internal/entprototest/ent/predicate"
)

// AllMethodsServiceDelete is the builder for deleting a AllMethodsService entity.
type AllMethodsServiceDelete struct {
	config
	hooks    []Hook
	mutation *AllMethodsServiceMutation
}

// Where appends a list predicates to the AllMethodsServiceDelete builder.
func (amsd *AllMethodsServiceDelete) Where(ps ...predicate.AllMethodsService) *AllMethodsServiceDelete {
	amsd.mutation.Where(ps...)
	return amsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (amsd *AllMethodsServiceDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(amsd.hooks) == 0 {
		affected, err = amsd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AllMethodsServiceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			amsd.mutation = mutation
			affected, err = amsd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(amsd.hooks) - 1; i >= 0; i-- {
			if amsd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = amsd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, amsd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (amsd *AllMethodsServiceDelete) ExecX(ctx context.Context) int {
	n, err := amsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (amsd *AllMethodsServiceDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: allmethodsservice.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: allmethodsservice.FieldID,
			},
		},
	}
	if ps := amsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, amsd.driver, _spec)
}

// AllMethodsServiceDeleteOne is the builder for deleting a single AllMethodsService entity.
type AllMethodsServiceDeleteOne struct {
	amsd *AllMethodsServiceDelete
}

// Exec executes the deletion query.
func (amsdo *AllMethodsServiceDeleteOne) Exec(ctx context.Context) error {
	n, err := amsdo.amsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{allmethodsservice.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (amsdo *AllMethodsServiceDeleteOne) ExecX(ctx context.Context) {
	amsdo.amsd.ExecX(ctx)
}
