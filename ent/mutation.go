// Code generated by entc, DO NOT EDIT.

package ent

import (
	"betterdsc/ent/predicate"
	"betterdsc/ent/serveremote"
	"context"
	"errors"
	"fmt"
	"sync"

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
	TypeServerEmote = "ServerEmote"
)

// ServerEmoteMutation represents an operation that mutates the ServerEmote nodes in the graph.
type ServerEmoteMutation struct {
	config
	op            Op
	typ           string
	id            *int
	server_id     *string
	emote_id      *string
	code          *string
	image_type    *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*ServerEmote, error)
	predicates    []predicate.ServerEmote
}

var _ ent.Mutation = (*ServerEmoteMutation)(nil)

// serveremoteOption allows management of the mutation configuration using functional options.
type serveremoteOption func(*ServerEmoteMutation)

// newServerEmoteMutation creates new mutation for the ServerEmote entity.
func newServerEmoteMutation(c config, op Op, opts ...serveremoteOption) *ServerEmoteMutation {
	m := &ServerEmoteMutation{
		config:        c,
		op:            op,
		typ:           TypeServerEmote,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withServerEmoteID sets the ID field of the mutation.
func withServerEmoteID(id int) serveremoteOption {
	return func(m *ServerEmoteMutation) {
		var (
			err   error
			once  sync.Once
			value *ServerEmote
		)
		m.oldValue = func(ctx context.Context) (*ServerEmote, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().ServerEmote.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withServerEmote sets the old ServerEmote of the mutation.
func withServerEmote(node *ServerEmote) serveremoteOption {
	return func(m *ServerEmoteMutation) {
		m.oldValue = func(context.Context) (*ServerEmote, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ServerEmoteMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ServerEmoteMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ServerEmoteMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *ServerEmoteMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().ServerEmote.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetServerID sets the "server_id" field.
func (m *ServerEmoteMutation) SetServerID(s string) {
	m.server_id = &s
}

// ServerID returns the value of the "server_id" field in the mutation.
func (m *ServerEmoteMutation) ServerID() (r string, exists bool) {
	v := m.server_id
	if v == nil {
		return
	}
	return *v, true
}

// OldServerID returns the old "server_id" field's value of the ServerEmote entity.
// If the ServerEmote object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ServerEmoteMutation) OldServerID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldServerID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldServerID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldServerID: %w", err)
	}
	return oldValue.ServerID, nil
}

// ResetServerID resets all changes to the "server_id" field.
func (m *ServerEmoteMutation) ResetServerID() {
	m.server_id = nil
}

// SetEmoteID sets the "emote_id" field.
func (m *ServerEmoteMutation) SetEmoteID(s string) {
	m.emote_id = &s
}

// EmoteID returns the value of the "emote_id" field in the mutation.
func (m *ServerEmoteMutation) EmoteID() (r string, exists bool) {
	v := m.emote_id
	if v == nil {
		return
	}
	return *v, true
}

// OldEmoteID returns the old "emote_id" field's value of the ServerEmote entity.
// If the ServerEmote object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ServerEmoteMutation) OldEmoteID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEmoteID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEmoteID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmoteID: %w", err)
	}
	return oldValue.EmoteID, nil
}

// ResetEmoteID resets all changes to the "emote_id" field.
func (m *ServerEmoteMutation) ResetEmoteID() {
	m.emote_id = nil
}

// SetCode sets the "code" field.
func (m *ServerEmoteMutation) SetCode(s string) {
	m.code = &s
}

// Code returns the value of the "code" field in the mutation.
func (m *ServerEmoteMutation) Code() (r string, exists bool) {
	v := m.code
	if v == nil {
		return
	}
	return *v, true
}

// OldCode returns the old "code" field's value of the ServerEmote entity.
// If the ServerEmote object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ServerEmoteMutation) OldCode(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCode is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCode requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCode: %w", err)
	}
	return oldValue.Code, nil
}

// ResetCode resets all changes to the "code" field.
func (m *ServerEmoteMutation) ResetCode() {
	m.code = nil
}

// SetImageType sets the "image_type" field.
func (m *ServerEmoteMutation) SetImageType(s string) {
	m.image_type = &s
}

// ImageType returns the value of the "image_type" field in the mutation.
func (m *ServerEmoteMutation) ImageType() (r string, exists bool) {
	v := m.image_type
	if v == nil {
		return
	}
	return *v, true
}

// OldImageType returns the old "image_type" field's value of the ServerEmote entity.
// If the ServerEmote object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ServerEmoteMutation) OldImageType(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldImageType is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldImageType requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldImageType: %w", err)
	}
	return oldValue.ImageType, nil
}

// ResetImageType resets all changes to the "image_type" field.
func (m *ServerEmoteMutation) ResetImageType() {
	m.image_type = nil
}

// Where appends a list predicates to the ServerEmoteMutation builder.
func (m *ServerEmoteMutation) Where(ps ...predicate.ServerEmote) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *ServerEmoteMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (ServerEmote).
func (m *ServerEmoteMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ServerEmoteMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.server_id != nil {
		fields = append(fields, serveremote.FieldServerID)
	}
	if m.emote_id != nil {
		fields = append(fields, serveremote.FieldEmoteID)
	}
	if m.code != nil {
		fields = append(fields, serveremote.FieldCode)
	}
	if m.image_type != nil {
		fields = append(fields, serveremote.FieldImageType)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ServerEmoteMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case serveremote.FieldServerID:
		return m.ServerID()
	case serveremote.FieldEmoteID:
		return m.EmoteID()
	case serveremote.FieldCode:
		return m.Code()
	case serveremote.FieldImageType:
		return m.ImageType()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ServerEmoteMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case serveremote.FieldServerID:
		return m.OldServerID(ctx)
	case serveremote.FieldEmoteID:
		return m.OldEmoteID(ctx)
	case serveremote.FieldCode:
		return m.OldCode(ctx)
	case serveremote.FieldImageType:
		return m.OldImageType(ctx)
	}
	return nil, fmt.Errorf("unknown ServerEmote field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ServerEmoteMutation) SetField(name string, value ent.Value) error {
	switch name {
	case serveremote.FieldServerID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetServerID(v)
		return nil
	case serveremote.FieldEmoteID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmoteID(v)
		return nil
	case serveremote.FieldCode:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCode(v)
		return nil
	case serveremote.FieldImageType:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetImageType(v)
		return nil
	}
	return fmt.Errorf("unknown ServerEmote field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ServerEmoteMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ServerEmoteMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ServerEmoteMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown ServerEmote numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ServerEmoteMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ServerEmoteMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ServerEmoteMutation) ClearField(name string) error {
	return fmt.Errorf("unknown ServerEmote nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ServerEmoteMutation) ResetField(name string) error {
	switch name {
	case serveremote.FieldServerID:
		m.ResetServerID()
		return nil
	case serveremote.FieldEmoteID:
		m.ResetEmoteID()
		return nil
	case serveremote.FieldCode:
		m.ResetCode()
		return nil
	case serveremote.FieldImageType:
		m.ResetImageType()
		return nil
	}
	return fmt.Errorf("unknown ServerEmote field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ServerEmoteMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ServerEmoteMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ServerEmoteMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ServerEmoteMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ServerEmoteMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ServerEmoteMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ServerEmoteMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown ServerEmote unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ServerEmoteMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown ServerEmote edge %s", name)
}
