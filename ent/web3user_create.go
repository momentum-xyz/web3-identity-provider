// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/OdysseyMomentumExperience/web3-identity-provider/ent/web3challenge"
	"github.com/OdysseyMomentumExperience/web3-identity-provider/ent/web3user"
	"github.com/google/uuid"
)

// Web3UserCreate is the builder for creating a Web3User entity.
type Web3UserCreate struct {
	config
	mutation *Web3UserMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetUUID sets the "uuid" field.
func (wc *Web3UserCreate) SetUUID(u uuid.UUID) *Web3UserCreate {
	wc.mutation.SetUUID(u)
	return wc
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (wc *Web3UserCreate) SetNillableUUID(u *uuid.UUID) *Web3UserCreate {
	if u != nil {
		wc.SetUUID(*u)
	}
	return wc
}

// SetAddress sets the "address" field.
func (wc *Web3UserCreate) SetAddress(s string) *Web3UserCreate {
	wc.mutation.SetAddress(s)
	return wc
}

// SetWalletType sets the "wallet_type" field.
func (wc *Web3UserCreate) SetWalletType(wt web3user.WalletType) *Web3UserCreate {
	wc.mutation.SetWalletType(wt)
	return wc
}

// SetNillableWalletType sets the "wallet_type" field if the given value is not nil.
func (wc *Web3UserCreate) SetNillableWalletType(wt *web3user.WalletType) *Web3UserCreate {
	if wt != nil {
		wc.SetWalletType(*wt)
	}
	return wc
}

// AddWeb3ChallengeIDs adds the "web3_challenges" edge to the Web3Challenge entity by IDs.
func (wc *Web3UserCreate) AddWeb3ChallengeIDs(ids ...int) *Web3UserCreate {
	wc.mutation.AddWeb3ChallengeIDs(ids...)
	return wc
}

// AddWeb3Challenges adds the "web3_challenges" edges to the Web3Challenge entity.
func (wc *Web3UserCreate) AddWeb3Challenges(w ...*Web3Challenge) *Web3UserCreate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wc.AddWeb3ChallengeIDs(ids...)
}

// Mutation returns the Web3UserMutation object of the builder.
func (wc *Web3UserCreate) Mutation() *Web3UserMutation {
	return wc.mutation
}

// Save creates the Web3User in the database.
func (wc *Web3UserCreate) Save(ctx context.Context) (*Web3User, error) {
	var (
		err  error
		node *Web3User
	)
	wc.defaults()
	if len(wc.hooks) == 0 {
		if err = wc.check(); err != nil {
			return nil, err
		}
		node, err = wc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*Web3UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wc.check(); err != nil {
				return nil, err
			}
			wc.mutation = mutation
			if node, err = wc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(wc.hooks) - 1; i >= 0; i-- {
			if wc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (wc *Web3UserCreate) SaveX(ctx context.Context) *Web3User {
	v, err := wc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wc *Web3UserCreate) Exec(ctx context.Context) error {
	_, err := wc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wc *Web3UserCreate) ExecX(ctx context.Context) {
	if err := wc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wc *Web3UserCreate) defaults() {
	if _, ok := wc.mutation.UUID(); !ok {
		v := web3user.DefaultUUID()
		wc.mutation.SetUUID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wc *Web3UserCreate) check() error {
	if _, ok := wc.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New(`ent: missing required field "Web3User.uuid"`)}
	}
	if _, ok := wc.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New(`ent: missing required field "Web3User.address"`)}
	}
	if v, ok := wc.mutation.WalletType(); ok {
		if err := web3user.WalletTypeValidator(v); err != nil {
			return &ValidationError{Name: "wallet_type", err: fmt.Errorf(`ent: validator failed for field "Web3User.wallet_type": %w`, err)}
		}
	}
	return nil
}

func (wc *Web3UserCreate) sqlSave(ctx context.Context) (*Web3User, error) {
	_node, _spec := wc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (wc *Web3UserCreate) createSpec() (*Web3User, *sqlgraph.CreateSpec) {
	var (
		_node = &Web3User{config: wc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: web3user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: web3user.FieldID,
			},
		}
	)
	_spec.OnConflict = wc.conflict
	if value, ok := wc.mutation.UUID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: web3user.FieldUUID,
		})
		_node.UUID = value
	}
	if value, ok := wc.mutation.Address(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: web3user.FieldAddress,
		})
		_node.Address = value
	}
	if value, ok := wc.mutation.WalletType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: web3user.FieldWalletType,
		})
		_node.WalletType = value
	}
	if nodes := wc.mutation.Web3ChallengesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   web3user.Web3ChallengesTable,
			Columns: []string{web3user.Web3ChallengesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: web3challenge.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Web3User.Create().
//		SetUUID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.Web3UserUpsert) {
//			SetUUID(v+v).
//		}).
//		Exec(ctx)
//
func (wc *Web3UserCreate) OnConflict(opts ...sql.ConflictOption) *Web3UserUpsertOne {
	wc.conflict = opts
	return &Web3UserUpsertOne{
		create: wc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Web3User.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (wc *Web3UserCreate) OnConflictColumns(columns ...string) *Web3UserUpsertOne {
	wc.conflict = append(wc.conflict, sql.ConflictColumns(columns...))
	return &Web3UserUpsertOne{
		create: wc,
	}
}

type (
	// Web3UserUpsertOne is the builder for "upsert"-ing
	//  one Web3User node.
	Web3UserUpsertOne struct {
		create *Web3UserCreate
	}

	// Web3UserUpsert is the "OnConflict" setter.
	Web3UserUpsert struct {
		*sql.UpdateSet
	}
)

// SetUUID sets the "uuid" field.
func (u *Web3UserUpsert) SetUUID(v uuid.UUID) *Web3UserUpsert {
	u.Set(web3user.FieldUUID, v)
	return u
}

// UpdateUUID sets the "uuid" field to the value that was provided on create.
func (u *Web3UserUpsert) UpdateUUID() *Web3UserUpsert {
	u.SetExcluded(web3user.FieldUUID)
	return u
}

// SetAddress sets the "address" field.
func (u *Web3UserUpsert) SetAddress(v string) *Web3UserUpsert {
	u.Set(web3user.FieldAddress, v)
	return u
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *Web3UserUpsert) UpdateAddress() *Web3UserUpsert {
	u.SetExcluded(web3user.FieldAddress)
	return u
}

// SetWalletType sets the "wallet_type" field.
func (u *Web3UserUpsert) SetWalletType(v web3user.WalletType) *Web3UserUpsert {
	u.Set(web3user.FieldWalletType, v)
	return u
}

// UpdateWalletType sets the "wallet_type" field to the value that was provided on create.
func (u *Web3UserUpsert) UpdateWalletType() *Web3UserUpsert {
	u.SetExcluded(web3user.FieldWalletType)
	return u
}

// ClearWalletType clears the value of the "wallet_type" field.
func (u *Web3UserUpsert) ClearWalletType() *Web3UserUpsert {
	u.SetNull(web3user.FieldWalletType)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Web3User.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *Web3UserUpsertOne) UpdateNewValues() *Web3UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Web3User.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *Web3UserUpsertOne) Ignore() *Web3UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *Web3UserUpsertOne) DoNothing() *Web3UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the Web3UserCreate.OnConflict
// documentation for more info.
func (u *Web3UserUpsertOne) Update(set func(*Web3UserUpsert)) *Web3UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&Web3UserUpsert{UpdateSet: update})
	}))
	return u
}

// SetUUID sets the "uuid" field.
func (u *Web3UserUpsertOne) SetUUID(v uuid.UUID) *Web3UserUpsertOne {
	return u.Update(func(s *Web3UserUpsert) {
		s.SetUUID(v)
	})
}

// UpdateUUID sets the "uuid" field to the value that was provided on create.
func (u *Web3UserUpsertOne) UpdateUUID() *Web3UserUpsertOne {
	return u.Update(func(s *Web3UserUpsert) {
		s.UpdateUUID()
	})
}

// SetAddress sets the "address" field.
func (u *Web3UserUpsertOne) SetAddress(v string) *Web3UserUpsertOne {
	return u.Update(func(s *Web3UserUpsert) {
		s.SetAddress(v)
	})
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *Web3UserUpsertOne) UpdateAddress() *Web3UserUpsertOne {
	return u.Update(func(s *Web3UserUpsert) {
		s.UpdateAddress()
	})
}

// SetWalletType sets the "wallet_type" field.
func (u *Web3UserUpsertOne) SetWalletType(v web3user.WalletType) *Web3UserUpsertOne {
	return u.Update(func(s *Web3UserUpsert) {
		s.SetWalletType(v)
	})
}

// UpdateWalletType sets the "wallet_type" field to the value that was provided on create.
func (u *Web3UserUpsertOne) UpdateWalletType() *Web3UserUpsertOne {
	return u.Update(func(s *Web3UserUpsert) {
		s.UpdateWalletType()
	})
}

// ClearWalletType clears the value of the "wallet_type" field.
func (u *Web3UserUpsertOne) ClearWalletType() *Web3UserUpsertOne {
	return u.Update(func(s *Web3UserUpsert) {
		s.ClearWalletType()
	})
}

// Exec executes the query.
func (u *Web3UserUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for Web3UserCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *Web3UserUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *Web3UserUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *Web3UserUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// Web3UserCreateBulk is the builder for creating many Web3User entities in bulk.
type Web3UserCreateBulk struct {
	config
	builders []*Web3UserCreate
	conflict []sql.ConflictOption
}

// Save creates the Web3User entities in the database.
func (wcb *Web3UserCreateBulk) Save(ctx context.Context) ([]*Web3User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wcb.builders))
	nodes := make([]*Web3User, len(wcb.builders))
	mutators := make([]Mutator, len(wcb.builders))
	for i := range wcb.builders {
		func(i int, root context.Context) {
			builder := wcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*Web3UserMutation)
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
					_, err = mutators[i+1].Mutate(root, wcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = wcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, wcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wcb *Web3UserCreateBulk) SaveX(ctx context.Context) []*Web3User {
	v, err := wcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wcb *Web3UserCreateBulk) Exec(ctx context.Context) error {
	_, err := wcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcb *Web3UserCreateBulk) ExecX(ctx context.Context) {
	if err := wcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Web3User.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.Web3UserUpsert) {
//			SetUUID(v+v).
//		}).
//		Exec(ctx)
//
func (wcb *Web3UserCreateBulk) OnConflict(opts ...sql.ConflictOption) *Web3UserUpsertBulk {
	wcb.conflict = opts
	return &Web3UserUpsertBulk{
		create: wcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Web3User.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (wcb *Web3UserCreateBulk) OnConflictColumns(columns ...string) *Web3UserUpsertBulk {
	wcb.conflict = append(wcb.conflict, sql.ConflictColumns(columns...))
	return &Web3UserUpsertBulk{
		create: wcb,
	}
}

// Web3UserUpsertBulk is the builder for "upsert"-ing
// a bulk of Web3User nodes.
type Web3UserUpsertBulk struct {
	create *Web3UserCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Web3User.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *Web3UserUpsertBulk) UpdateNewValues() *Web3UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Web3User.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *Web3UserUpsertBulk) Ignore() *Web3UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *Web3UserUpsertBulk) DoNothing() *Web3UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the Web3UserCreateBulk.OnConflict
// documentation for more info.
func (u *Web3UserUpsertBulk) Update(set func(*Web3UserUpsert)) *Web3UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&Web3UserUpsert{UpdateSet: update})
	}))
	return u
}

// SetUUID sets the "uuid" field.
func (u *Web3UserUpsertBulk) SetUUID(v uuid.UUID) *Web3UserUpsertBulk {
	return u.Update(func(s *Web3UserUpsert) {
		s.SetUUID(v)
	})
}

// UpdateUUID sets the "uuid" field to the value that was provided on create.
func (u *Web3UserUpsertBulk) UpdateUUID() *Web3UserUpsertBulk {
	return u.Update(func(s *Web3UserUpsert) {
		s.UpdateUUID()
	})
}

// SetAddress sets the "address" field.
func (u *Web3UserUpsertBulk) SetAddress(v string) *Web3UserUpsertBulk {
	return u.Update(func(s *Web3UserUpsert) {
		s.SetAddress(v)
	})
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *Web3UserUpsertBulk) UpdateAddress() *Web3UserUpsertBulk {
	return u.Update(func(s *Web3UserUpsert) {
		s.UpdateAddress()
	})
}

// SetWalletType sets the "wallet_type" field.
func (u *Web3UserUpsertBulk) SetWalletType(v web3user.WalletType) *Web3UserUpsertBulk {
	return u.Update(func(s *Web3UserUpsert) {
		s.SetWalletType(v)
	})
}

// UpdateWalletType sets the "wallet_type" field to the value that was provided on create.
func (u *Web3UserUpsertBulk) UpdateWalletType() *Web3UserUpsertBulk {
	return u.Update(func(s *Web3UserUpsert) {
		s.UpdateWalletType()
	})
}

// ClearWalletType clears the value of the "wallet_type" field.
func (u *Web3UserUpsertBulk) ClearWalletType() *Web3UserUpsertBulk {
	return u.Update(func(s *Web3UserUpsert) {
		s.ClearWalletType()
	})
}

// Exec executes the query.
func (u *Web3UserUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the Web3UserCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for Web3UserCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *Web3UserUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}