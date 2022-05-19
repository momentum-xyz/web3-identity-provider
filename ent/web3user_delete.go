// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/OdysseyMomentumExperience/web3-identity-provider/ent/predicate"
	"github.com/OdysseyMomentumExperience/web3-identity-provider/ent/web3user"
)

// Web3UserDelete is the builder for deleting a Web3User entity.
type Web3UserDelete struct {
	config
	hooks    []Hook
	mutation *Web3UserMutation
}

// Where appends a list predicates to the Web3UserDelete builder.
func (wd *Web3UserDelete) Where(ps ...predicate.Web3User) *Web3UserDelete {
	wd.mutation.Where(ps...)
	return wd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wd *Web3UserDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(wd.hooks) == 0 {
		affected, err = wd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*Web3UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wd.mutation = mutation
			affected, err = wd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wd.hooks) - 1; i >= 0; i-- {
			if wd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (wd *Web3UserDelete) ExecX(ctx context.Context) int {
	n, err := wd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wd *Web3UserDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: web3user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: web3user.FieldID,
			},
		},
	}
	if ps := wd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, wd.driver, _spec)
}

// Web3UserDeleteOne is the builder for deleting a single Web3User entity.
type Web3UserDeleteOne struct {
	wd *Web3UserDelete
}

// Exec executes the deletion query.
func (wdo *Web3UserDeleteOne) Exec(ctx context.Context) error {
	n, err := wdo.wd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{web3user.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wdo *Web3UserDeleteOne) ExecX(ctx context.Context) {
	wdo.wd.ExecX(ctx)
}
