// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/asfsadas/contrib/entproto/internal/entprototest/ent/twomethodservice"
)

// TwoMethodServiceCreate is the builder for creating a TwoMethodService entity.
type TwoMethodServiceCreate struct {
	config
	mutation *TwoMethodServiceMutation
	hooks    []Hook
}

// Mutation returns the TwoMethodServiceMutation object of the builder.
func (tmsc *TwoMethodServiceCreate) Mutation() *TwoMethodServiceMutation {
	return tmsc.mutation
}

// Save creates the TwoMethodService in the database.
func (tmsc *TwoMethodServiceCreate) Save(ctx context.Context) (*TwoMethodService, error) {
	var (
		err  error
		node *TwoMethodService
	)
	if len(tmsc.hooks) == 0 {
		if err = tmsc.check(); err != nil {
			return nil, err
		}
		node, err = tmsc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TwoMethodServiceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tmsc.check(); err != nil {
				return nil, err
			}
			tmsc.mutation = mutation
			if node, err = tmsc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tmsc.hooks) - 1; i >= 0; i-- {
			if tmsc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tmsc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tmsc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tmsc *TwoMethodServiceCreate) SaveX(ctx context.Context) *TwoMethodService {
	v, err := tmsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tmsc *TwoMethodServiceCreate) Exec(ctx context.Context) error {
	_, err := tmsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tmsc *TwoMethodServiceCreate) ExecX(ctx context.Context) {
	if err := tmsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tmsc *TwoMethodServiceCreate) check() error {
	return nil
}

func (tmsc *TwoMethodServiceCreate) sqlSave(ctx context.Context) (*TwoMethodService, error) {
	_node, _spec := tmsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tmsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (tmsc *TwoMethodServiceCreate) createSpec() (*TwoMethodService, *sqlgraph.CreateSpec) {
	var (
		_node = &TwoMethodService{config: tmsc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: twomethodservice.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: twomethodservice.FieldID,
			},
		}
	)
	return _node, _spec
}

// TwoMethodServiceCreateBulk is the builder for creating many TwoMethodService entities in bulk.
type TwoMethodServiceCreateBulk struct {
	config
	builders []*TwoMethodServiceCreate
}

// Save creates the TwoMethodService entities in the database.
func (tmscb *TwoMethodServiceCreateBulk) Save(ctx context.Context) ([]*TwoMethodService, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tmscb.builders))
	nodes := make([]*TwoMethodService, len(tmscb.builders))
	mutators := make([]Mutator, len(tmscb.builders))
	for i := range tmscb.builders {
		func(i int, root context.Context) {
			builder := tmscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TwoMethodServiceMutation)
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
					_, err = mutators[i+1].Mutate(root, tmscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tmscb.driver, spec); err != nil {
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
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, tmscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tmscb *TwoMethodServiceCreateBulk) SaveX(ctx context.Context) []*TwoMethodService {
	v, err := tmscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tmscb *TwoMethodServiceCreateBulk) Exec(ctx context.Context) error {
	_, err := tmscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tmscb *TwoMethodServiceCreateBulk) ExecX(ctx context.Context) {
	if err := tmscb.Exec(ctx); err != nil {
		panic(err)
	}
}
