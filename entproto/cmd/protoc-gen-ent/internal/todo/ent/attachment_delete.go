// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/asfsadas/contrib/entproto/cmd/protoc-gen-ent/internal/todo/ent/attachment"
	"github.com/asfsadas/contrib/entproto/cmd/protoc-gen-ent/internal/todo/ent/predicate"
)

// AttachmentDelete is the builder for deleting a Attachment entity.
type AttachmentDelete struct {
	config
	hooks    []Hook
	mutation *AttachmentMutation
}

// Where appends a list predicates to the AttachmentDelete builder.
func (ad *AttachmentDelete) Where(ps ...predicate.Attachment) *AttachmentDelete {
	ad.mutation.Where(ps...)
	return ad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ad *AttachmentDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ad.hooks) == 0 {
		affected, err = ad.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AttachmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ad.mutation = mutation
			affected, err = ad.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ad.hooks) - 1; i >= 0; i-- {
			if ad.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ad.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ad.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ad *AttachmentDelete) ExecX(ctx context.Context) int {
	n, err := ad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ad *AttachmentDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: attachment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: attachment.FieldID,
			},
		},
	}
	if ps := ad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ad.driver, _spec)
}

// AttachmentDeleteOne is the builder for deleting a single Attachment entity.
type AttachmentDeleteOne struct {
	ad *AttachmentDelete
}

// Exec executes the deletion query.
func (ado *AttachmentDeleteOne) Exec(ctx context.Context) error {
	n, err := ado.ad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{attachment.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ado *AttachmentDeleteOne) ExecX(ctx context.Context) {
	ado.ad.ExecX(ctx)
}
