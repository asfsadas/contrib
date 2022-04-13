// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/asfsadas/contrib/entproto/internal/entprototest/ent/explicitskippedmessage"
)

// ExplicitSkippedMessageCreate is the builder for creating a ExplicitSkippedMessage entity.
type ExplicitSkippedMessageCreate struct {
	config
	mutation *ExplicitSkippedMessageMutation
	hooks    []Hook
}

// Mutation returns the ExplicitSkippedMessageMutation object of the builder.
func (esmc *ExplicitSkippedMessageCreate) Mutation() *ExplicitSkippedMessageMutation {
	return esmc.mutation
}

// Save creates the ExplicitSkippedMessage in the database.
func (esmc *ExplicitSkippedMessageCreate) Save(ctx context.Context) (*ExplicitSkippedMessage, error) {
	var (
		err  error
		node *ExplicitSkippedMessage
	)
	if len(esmc.hooks) == 0 {
		if err = esmc.check(); err != nil {
			return nil, err
		}
		node, err = esmc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ExplicitSkippedMessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = esmc.check(); err != nil {
				return nil, err
			}
			esmc.mutation = mutation
			if node, err = esmc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(esmc.hooks) - 1; i >= 0; i-- {
			if esmc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = esmc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, esmc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (esmc *ExplicitSkippedMessageCreate) SaveX(ctx context.Context) *ExplicitSkippedMessage {
	v, err := esmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (esmc *ExplicitSkippedMessageCreate) Exec(ctx context.Context) error {
	_, err := esmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (esmc *ExplicitSkippedMessageCreate) ExecX(ctx context.Context) {
	if err := esmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (esmc *ExplicitSkippedMessageCreate) check() error {
	return nil
}

func (esmc *ExplicitSkippedMessageCreate) sqlSave(ctx context.Context) (*ExplicitSkippedMessage, error) {
	_node, _spec := esmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, esmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (esmc *ExplicitSkippedMessageCreate) createSpec() (*ExplicitSkippedMessage, *sqlgraph.CreateSpec) {
	var (
		_node = &ExplicitSkippedMessage{config: esmc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: explicitskippedmessage.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: explicitskippedmessage.FieldID,
			},
		}
	)
	return _node, _spec
}

// ExplicitSkippedMessageCreateBulk is the builder for creating many ExplicitSkippedMessage entities in bulk.
type ExplicitSkippedMessageCreateBulk struct {
	config
	builders []*ExplicitSkippedMessageCreate
}

// Save creates the ExplicitSkippedMessage entities in the database.
func (esmcb *ExplicitSkippedMessageCreateBulk) Save(ctx context.Context) ([]*ExplicitSkippedMessage, error) {
	specs := make([]*sqlgraph.CreateSpec, len(esmcb.builders))
	nodes := make([]*ExplicitSkippedMessage, len(esmcb.builders))
	mutators := make([]Mutator, len(esmcb.builders))
	for i := range esmcb.builders {
		func(i int, root context.Context) {
			builder := esmcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ExplicitSkippedMessageMutation)
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
					_, err = mutators[i+1].Mutate(root, esmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, esmcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, esmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (esmcb *ExplicitSkippedMessageCreateBulk) SaveX(ctx context.Context) []*ExplicitSkippedMessage {
	v, err := esmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (esmcb *ExplicitSkippedMessageCreateBulk) Exec(ctx context.Context) error {
	_, err := esmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (esmcb *ExplicitSkippedMessageCreateBulk) ExecX(ctx context.Context) {
	if err := esmcb.Exec(ctx); err != nil {
		panic(err)
	}
}
