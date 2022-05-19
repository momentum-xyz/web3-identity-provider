// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/OdysseyMomentumExperience/web3-identity-provider/ent/predicate"
	"github.com/OdysseyMomentumExperience/web3-identity-provider/ent/web3challenge"
	"github.com/OdysseyMomentumExperience/web3-identity-provider/ent/web3user"
	"github.com/google/uuid"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeWeb3Challenge = "Web3Challenge"
	TypeWeb3User      = "Web3User"
)

// Web3ChallengeMutation represents an operation that mutates the Web3Challenge nodes in the graph.
type Web3ChallengeMutation struct {
	config
	op               Op
	typ              string
	id               *int
	uuid             *uuid.UUID
	login_challenge  *string
	web3_challenge   *string
	clearedFields    map[string]struct{}
	web3_user        *int
	clearedweb3_user bool
	done             bool
	oldValue         func(context.Context) (*Web3Challenge, error)
	predicates       []predicate.Web3Challenge
}

var _ ent.Mutation = (*Web3ChallengeMutation)(nil)

// web3challengeOption allows management of the mutation configuration using functional options.
type web3challengeOption func(*Web3ChallengeMutation)

// newWeb3ChallengeMutation creates new mutation for the Web3Challenge entity.
func newWeb3ChallengeMutation(c config, op Op, opts ...web3challengeOption) *Web3ChallengeMutation {
	m := &Web3ChallengeMutation{
		config:        c,
		op:            op,
		typ:           TypeWeb3Challenge,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withWeb3ChallengeID sets the ID field of the mutation.
func withWeb3ChallengeID(id int) web3challengeOption {
	return func(m *Web3ChallengeMutation) {
		var (
			err   error
			once  sync.Once
			value *Web3Challenge
		)
		m.oldValue = func(ctx context.Context) (*Web3Challenge, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Web3Challenge.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withWeb3Challenge sets the old Web3Challenge of the mutation.
func withWeb3Challenge(node *Web3Challenge) web3challengeOption {
	return func(m *Web3ChallengeMutation) {
		m.oldValue = func(context.Context) (*Web3Challenge, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m Web3ChallengeMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m Web3ChallengeMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *Web3ChallengeMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *Web3ChallengeMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Web3Challenge.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetUUID sets the "uuid" field.
func (m *Web3ChallengeMutation) SetUUID(u uuid.UUID) {
	m.uuid = &u
}

// UUID returns the value of the "uuid" field in the mutation.
func (m *Web3ChallengeMutation) UUID() (r uuid.UUID, exists bool) {
	v := m.uuid
	if v == nil {
		return
	}
	return *v, true
}

// OldUUID returns the old "uuid" field's value of the Web3Challenge entity.
// If the Web3Challenge object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *Web3ChallengeMutation) OldUUID(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUUID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUUID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUUID: %w", err)
	}
	return oldValue.UUID, nil
}

// ResetUUID resets all changes to the "uuid" field.
func (m *Web3ChallengeMutation) ResetUUID() {
	m.uuid = nil
}

// SetWeb3UserID sets the "web3_user_id" field.
func (m *Web3ChallengeMutation) SetWeb3UserID(i int) {
	m.web3_user = &i
}

// Web3UserID returns the value of the "web3_user_id" field in the mutation.
func (m *Web3ChallengeMutation) Web3UserID() (r int, exists bool) {
	v := m.web3_user
	if v == nil {
		return
	}
	return *v, true
}

// OldWeb3UserID returns the old "web3_user_id" field's value of the Web3Challenge entity.
// If the Web3Challenge object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *Web3ChallengeMutation) OldWeb3UserID(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldWeb3UserID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldWeb3UserID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldWeb3UserID: %w", err)
	}
	return oldValue.Web3UserID, nil
}

// ResetWeb3UserID resets all changes to the "web3_user_id" field.
func (m *Web3ChallengeMutation) ResetWeb3UserID() {
	m.web3_user = nil
}

// SetLoginChallenge sets the "login_challenge" field.
func (m *Web3ChallengeMutation) SetLoginChallenge(s string) {
	m.login_challenge = &s
}

// LoginChallenge returns the value of the "login_challenge" field in the mutation.
func (m *Web3ChallengeMutation) LoginChallenge() (r string, exists bool) {
	v := m.login_challenge
	if v == nil {
		return
	}
	return *v, true
}

// OldLoginChallenge returns the old "login_challenge" field's value of the Web3Challenge entity.
// If the Web3Challenge object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *Web3ChallengeMutation) OldLoginChallenge(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLoginChallenge is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLoginChallenge requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLoginChallenge: %w", err)
	}
	return oldValue.LoginChallenge, nil
}

// ResetLoginChallenge resets all changes to the "login_challenge" field.
func (m *Web3ChallengeMutation) ResetLoginChallenge() {
	m.login_challenge = nil
}

// SetWeb3Challenge sets the "web3_challenge" field.
func (m *Web3ChallengeMutation) SetWeb3Challenge(s string) {
	m.web3_challenge = &s
}

// Web3Challenge returns the value of the "web3_challenge" field in the mutation.
func (m *Web3ChallengeMutation) Web3Challenge() (r string, exists bool) {
	v := m.web3_challenge
	if v == nil {
		return
	}
	return *v, true
}

// OldWeb3Challenge returns the old "web3_challenge" field's value of the Web3Challenge entity.
// If the Web3Challenge object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *Web3ChallengeMutation) OldWeb3Challenge(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldWeb3Challenge is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldWeb3Challenge requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldWeb3Challenge: %w", err)
	}
	return oldValue.Web3Challenge, nil
}

// ResetWeb3Challenge resets all changes to the "web3_challenge" field.
func (m *Web3ChallengeMutation) ResetWeb3Challenge() {
	m.web3_challenge = nil
}

// ClearWeb3User clears the "web3_user" edge to the Web3User entity.
func (m *Web3ChallengeMutation) ClearWeb3User() {
	m.clearedweb3_user = true
}

// Web3UserCleared reports if the "web3_user" edge to the Web3User entity was cleared.
func (m *Web3ChallengeMutation) Web3UserCleared() bool {
	return m.clearedweb3_user
}

// Web3UserIDs returns the "web3_user" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// Web3UserID instead. It exists only for internal usage by the builders.
func (m *Web3ChallengeMutation) Web3UserIDs() (ids []int) {
	if id := m.web3_user; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetWeb3User resets all changes to the "web3_user" edge.
func (m *Web3ChallengeMutation) ResetWeb3User() {
	m.web3_user = nil
	m.clearedweb3_user = false
}

// Where appends a list predicates to the Web3ChallengeMutation builder.
func (m *Web3ChallengeMutation) Where(ps ...predicate.Web3Challenge) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *Web3ChallengeMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Web3Challenge).
func (m *Web3ChallengeMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *Web3ChallengeMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.uuid != nil {
		fields = append(fields, web3challenge.FieldUUID)
	}
	if m.web3_user != nil {
		fields = append(fields, web3challenge.FieldWeb3UserID)
	}
	if m.login_challenge != nil {
		fields = append(fields, web3challenge.FieldLoginChallenge)
	}
	if m.web3_challenge != nil {
		fields = append(fields, web3challenge.FieldWeb3Challenge)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *Web3ChallengeMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case web3challenge.FieldUUID:
		return m.UUID()
	case web3challenge.FieldWeb3UserID:
		return m.Web3UserID()
	case web3challenge.FieldLoginChallenge:
		return m.LoginChallenge()
	case web3challenge.FieldWeb3Challenge:
		return m.Web3Challenge()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *Web3ChallengeMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case web3challenge.FieldUUID:
		return m.OldUUID(ctx)
	case web3challenge.FieldWeb3UserID:
		return m.OldWeb3UserID(ctx)
	case web3challenge.FieldLoginChallenge:
		return m.OldLoginChallenge(ctx)
	case web3challenge.FieldWeb3Challenge:
		return m.OldWeb3Challenge(ctx)
	}
	return nil, fmt.Errorf("unknown Web3Challenge field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *Web3ChallengeMutation) SetField(name string, value ent.Value) error {
	switch name {
	case web3challenge.FieldUUID:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUUID(v)
		return nil
	case web3challenge.FieldWeb3UserID:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetWeb3UserID(v)
		return nil
	case web3challenge.FieldLoginChallenge:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLoginChallenge(v)
		return nil
	case web3challenge.FieldWeb3Challenge:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetWeb3Challenge(v)
		return nil
	}
	return fmt.Errorf("unknown Web3Challenge field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *Web3ChallengeMutation) AddedFields() []string {
	var fields []string
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *Web3ChallengeMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *Web3ChallengeMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Web3Challenge numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *Web3ChallengeMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *Web3ChallengeMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *Web3ChallengeMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Web3Challenge nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *Web3ChallengeMutation) ResetField(name string) error {
	switch name {
	case web3challenge.FieldUUID:
		m.ResetUUID()
		return nil
	case web3challenge.FieldWeb3UserID:
		m.ResetWeb3UserID()
		return nil
	case web3challenge.FieldLoginChallenge:
		m.ResetLoginChallenge()
		return nil
	case web3challenge.FieldWeb3Challenge:
		m.ResetWeb3Challenge()
		return nil
	}
	return fmt.Errorf("unknown Web3Challenge field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *Web3ChallengeMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.web3_user != nil {
		edges = append(edges, web3challenge.EdgeWeb3User)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *Web3ChallengeMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case web3challenge.EdgeWeb3User:
		if id := m.web3_user; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *Web3ChallengeMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *Web3ChallengeMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *Web3ChallengeMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedweb3_user {
		edges = append(edges, web3challenge.EdgeWeb3User)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *Web3ChallengeMutation) EdgeCleared(name string) bool {
	switch name {
	case web3challenge.EdgeWeb3User:
		return m.clearedweb3_user
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *Web3ChallengeMutation) ClearEdge(name string) error {
	switch name {
	case web3challenge.EdgeWeb3User:
		m.ClearWeb3User()
		return nil
	}
	return fmt.Errorf("unknown Web3Challenge unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *Web3ChallengeMutation) ResetEdge(name string) error {
	switch name {
	case web3challenge.EdgeWeb3User:
		m.ResetWeb3User()
		return nil
	}
	return fmt.Errorf("unknown Web3Challenge edge %s", name)
}

// Web3UserMutation represents an operation that mutates the Web3User nodes in the graph.
type Web3UserMutation struct {
	config
	op                     Op
	typ                    string
	id                     *int
	uuid                   *uuid.UUID
	address                *string
	wallet_type            *web3user.WalletType
	clearedFields          map[string]struct{}
	web3_challenges        map[int]struct{}
	removedweb3_challenges map[int]struct{}
	clearedweb3_challenges bool
	done                   bool
	oldValue               func(context.Context) (*Web3User, error)
	predicates             []predicate.Web3User
}

var _ ent.Mutation = (*Web3UserMutation)(nil)

// web3userOption allows management of the mutation configuration using functional options.
type web3userOption func(*Web3UserMutation)

// newWeb3UserMutation creates new mutation for the Web3User entity.
func newWeb3UserMutation(c config, op Op, opts ...web3userOption) *Web3UserMutation {
	m := &Web3UserMutation{
		config:        c,
		op:            op,
		typ:           TypeWeb3User,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withWeb3UserID sets the ID field of the mutation.
func withWeb3UserID(id int) web3userOption {
	return func(m *Web3UserMutation) {
		var (
			err   error
			once  sync.Once
			value *Web3User
		)
		m.oldValue = func(ctx context.Context) (*Web3User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Web3User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withWeb3User sets the old Web3User of the mutation.
func withWeb3User(node *Web3User) web3userOption {
	return func(m *Web3UserMutation) {
		m.oldValue = func(context.Context) (*Web3User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m Web3UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m Web3UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *Web3UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *Web3UserMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Web3User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetUUID sets the "uuid" field.
func (m *Web3UserMutation) SetUUID(u uuid.UUID) {
	m.uuid = &u
}

// UUID returns the value of the "uuid" field in the mutation.
func (m *Web3UserMutation) UUID() (r uuid.UUID, exists bool) {
	v := m.uuid
	if v == nil {
		return
	}
	return *v, true
}

// OldUUID returns the old "uuid" field's value of the Web3User entity.
// If the Web3User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *Web3UserMutation) OldUUID(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUUID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUUID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUUID: %w", err)
	}
	return oldValue.UUID, nil
}

// ResetUUID resets all changes to the "uuid" field.
func (m *Web3UserMutation) ResetUUID() {
	m.uuid = nil
}

// SetAddress sets the "address" field.
func (m *Web3UserMutation) SetAddress(s string) {
	m.address = &s
}

// Address returns the value of the "address" field in the mutation.
func (m *Web3UserMutation) Address() (r string, exists bool) {
	v := m.address
	if v == nil {
		return
	}
	return *v, true
}

// OldAddress returns the old "address" field's value of the Web3User entity.
// If the Web3User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *Web3UserMutation) OldAddress(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAddress is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAddress requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAddress: %w", err)
	}
	return oldValue.Address, nil
}

// ResetAddress resets all changes to the "address" field.
func (m *Web3UserMutation) ResetAddress() {
	m.address = nil
}

// SetWalletType sets the "wallet_type" field.
func (m *Web3UserMutation) SetWalletType(wt web3user.WalletType) {
	m.wallet_type = &wt
}

// WalletType returns the value of the "wallet_type" field in the mutation.
func (m *Web3UserMutation) WalletType() (r web3user.WalletType, exists bool) {
	v := m.wallet_type
	if v == nil {
		return
	}
	return *v, true
}

// OldWalletType returns the old "wallet_type" field's value of the Web3User entity.
// If the Web3User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *Web3UserMutation) OldWalletType(ctx context.Context) (v web3user.WalletType, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldWalletType is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldWalletType requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldWalletType: %w", err)
	}
	return oldValue.WalletType, nil
}

// ClearWalletType clears the value of the "wallet_type" field.
func (m *Web3UserMutation) ClearWalletType() {
	m.wallet_type = nil
	m.clearedFields[web3user.FieldWalletType] = struct{}{}
}

// WalletTypeCleared returns if the "wallet_type" field was cleared in this mutation.
func (m *Web3UserMutation) WalletTypeCleared() bool {
	_, ok := m.clearedFields[web3user.FieldWalletType]
	return ok
}

// ResetWalletType resets all changes to the "wallet_type" field.
func (m *Web3UserMutation) ResetWalletType() {
	m.wallet_type = nil
	delete(m.clearedFields, web3user.FieldWalletType)
}

// AddWeb3ChallengeIDs adds the "web3_challenges" edge to the Web3Challenge entity by ids.
func (m *Web3UserMutation) AddWeb3ChallengeIDs(ids ...int) {
	if m.web3_challenges == nil {
		m.web3_challenges = make(map[int]struct{})
	}
	for i := range ids {
		m.web3_challenges[ids[i]] = struct{}{}
	}
}

// ClearWeb3Challenges clears the "web3_challenges" edge to the Web3Challenge entity.
func (m *Web3UserMutation) ClearWeb3Challenges() {
	m.clearedweb3_challenges = true
}

// Web3ChallengesCleared reports if the "web3_challenges" edge to the Web3Challenge entity was cleared.
func (m *Web3UserMutation) Web3ChallengesCleared() bool {
	return m.clearedweb3_challenges
}

// RemoveWeb3ChallengeIDs removes the "web3_challenges" edge to the Web3Challenge entity by IDs.
func (m *Web3UserMutation) RemoveWeb3ChallengeIDs(ids ...int) {
	if m.removedweb3_challenges == nil {
		m.removedweb3_challenges = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.web3_challenges, ids[i])
		m.removedweb3_challenges[ids[i]] = struct{}{}
	}
}

// RemovedWeb3Challenges returns the removed IDs of the "web3_challenges" edge to the Web3Challenge entity.
func (m *Web3UserMutation) RemovedWeb3ChallengesIDs() (ids []int) {
	for id := range m.removedweb3_challenges {
		ids = append(ids, id)
	}
	return
}

// Web3ChallengesIDs returns the "web3_challenges" edge IDs in the mutation.
func (m *Web3UserMutation) Web3ChallengesIDs() (ids []int) {
	for id := range m.web3_challenges {
		ids = append(ids, id)
	}
	return
}

// ResetWeb3Challenges resets all changes to the "web3_challenges" edge.
func (m *Web3UserMutation) ResetWeb3Challenges() {
	m.web3_challenges = nil
	m.clearedweb3_challenges = false
	m.removedweb3_challenges = nil
}

// Where appends a list predicates to the Web3UserMutation builder.
func (m *Web3UserMutation) Where(ps ...predicate.Web3User) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *Web3UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Web3User).
func (m *Web3UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *Web3UserMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.uuid != nil {
		fields = append(fields, web3user.FieldUUID)
	}
	if m.address != nil {
		fields = append(fields, web3user.FieldAddress)
	}
	if m.wallet_type != nil {
		fields = append(fields, web3user.FieldWalletType)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *Web3UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case web3user.FieldUUID:
		return m.UUID()
	case web3user.FieldAddress:
		return m.Address()
	case web3user.FieldWalletType:
		return m.WalletType()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *Web3UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case web3user.FieldUUID:
		return m.OldUUID(ctx)
	case web3user.FieldAddress:
		return m.OldAddress(ctx)
	case web3user.FieldWalletType:
		return m.OldWalletType(ctx)
	}
	return nil, fmt.Errorf("unknown Web3User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *Web3UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case web3user.FieldUUID:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUUID(v)
		return nil
	case web3user.FieldAddress:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAddress(v)
		return nil
	case web3user.FieldWalletType:
		v, ok := value.(web3user.WalletType)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetWalletType(v)
		return nil
	}
	return fmt.Errorf("unknown Web3User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *Web3UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *Web3UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *Web3UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Web3User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *Web3UserMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(web3user.FieldWalletType) {
		fields = append(fields, web3user.FieldWalletType)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *Web3UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *Web3UserMutation) ClearField(name string) error {
	switch name {
	case web3user.FieldWalletType:
		m.ClearWalletType()
		return nil
	}
	return fmt.Errorf("unknown Web3User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *Web3UserMutation) ResetField(name string) error {
	switch name {
	case web3user.FieldUUID:
		m.ResetUUID()
		return nil
	case web3user.FieldAddress:
		m.ResetAddress()
		return nil
	case web3user.FieldWalletType:
		m.ResetWalletType()
		return nil
	}
	return fmt.Errorf("unknown Web3User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *Web3UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.web3_challenges != nil {
		edges = append(edges, web3user.EdgeWeb3Challenges)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *Web3UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case web3user.EdgeWeb3Challenges:
		ids := make([]ent.Value, 0, len(m.web3_challenges))
		for id := range m.web3_challenges {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *Web3UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedweb3_challenges != nil {
		edges = append(edges, web3user.EdgeWeb3Challenges)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *Web3UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case web3user.EdgeWeb3Challenges:
		ids := make([]ent.Value, 0, len(m.removedweb3_challenges))
		for id := range m.removedweb3_challenges {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *Web3UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedweb3_challenges {
		edges = append(edges, web3user.EdgeWeb3Challenges)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *Web3UserMutation) EdgeCleared(name string) bool {
	switch name {
	case web3user.EdgeWeb3Challenges:
		return m.clearedweb3_challenges
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *Web3UserMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Web3User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *Web3UserMutation) ResetEdge(name string) error {
	switch name {
	case web3user.EdgeWeb3Challenges:
		m.ResetWeb3Challenges()
		return nil
	}
	return fmt.Errorf("unknown Web3User edge %s", name)
}
