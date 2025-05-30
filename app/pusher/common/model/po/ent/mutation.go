// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"message-push/app/receiver/common/model/po/ent/demo"
	"message-push/app/receiver/common/model/po/ent/predicate"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeDemo = "Demo"
)

// DemoMutation represents an operation that mutates the Demo nodes in the graph.
type DemoMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	create_time   *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Demo, error)
	predicates    []predicate.Demo
}

var _ ent.Mutation = (*DemoMutation)(nil)

// demoOption allows management of the mutation configuration using functional options.
type demoOption func(*DemoMutation)

// newDemoMutation creates new mutation for the Demo entity.
func newDemoMutation(c config, op Op, opts ...demoOption) *DemoMutation {
	m := &DemoMutation{
		config:        c,
		op:            op,
		typ:           TypeDemo,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withDemoID sets the ID field of the mutation.
func withDemoID(id int) demoOption {
	return func(m *DemoMutation) {
		var (
			err   error
			once  sync.Once
			value *Demo
		)
		m.oldValue = func(ctx context.Context) (*Demo, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Demo.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withDemo sets the old Demo of the mutation.
func withDemo(node *Demo) demoOption {
	return func(m *DemoMutation) {
		m.oldValue = func(context.Context) (*Demo, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m DemoMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m DemoMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Demo entities.
func (m *DemoMutation) SetID(id int) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *DemoMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *DemoMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Demo.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *DemoMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *DemoMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Demo entity.
// If the Demo object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DemoMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ClearName clears the value of the "name" field.
func (m *DemoMutation) ClearName() {
	m.name = nil
	m.clearedFields[demo.FieldName] = struct{}{}
}

// NameCleared returns if the "name" field was cleared in this mutation.
func (m *DemoMutation) NameCleared() bool {
	_, ok := m.clearedFields[demo.FieldName]
	return ok
}

// ResetName resets all changes to the "name" field.
func (m *DemoMutation) ResetName() {
	m.name = nil
	delete(m.clearedFields, demo.FieldName)
}

// SetCreateTime sets the "create_time" field.
func (m *DemoMutation) SetCreateTime(t time.Time) {
	m.create_time = &t
}

// CreateTime returns the value of the "create_time" field in the mutation.
func (m *DemoMutation) CreateTime() (r time.Time, exists bool) {
	v := m.create_time
	if v == nil {
		return
	}
	return *v, true
}

// OldCreateTime returns the old "create_time" field's value of the Demo entity.
// If the Demo object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DemoMutation) OldCreateTime(ctx context.Context) (v *time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreateTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreateTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreateTime: %w", err)
	}
	return oldValue.CreateTime, nil
}

// ClearCreateTime clears the value of the "create_time" field.
func (m *DemoMutation) ClearCreateTime() {
	m.create_time = nil
	m.clearedFields[demo.FieldCreateTime] = struct{}{}
}

// CreateTimeCleared returns if the "create_time" field was cleared in this mutation.
func (m *DemoMutation) CreateTimeCleared() bool {
	_, ok := m.clearedFields[demo.FieldCreateTime]
	return ok
}

// ResetCreateTime resets all changes to the "create_time" field.
func (m *DemoMutation) ResetCreateTime() {
	m.create_time = nil
	delete(m.clearedFields, demo.FieldCreateTime)
}

// Where appends a list predicates to the DemoMutation builder.
func (m *DemoMutation) Where(ps ...predicate.Demo) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the DemoMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *DemoMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Demo, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *DemoMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *DemoMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Demo).
func (m *DemoMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *DemoMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.name != nil {
		fields = append(fields, demo.FieldName)
	}
	if m.create_time != nil {
		fields = append(fields, demo.FieldCreateTime)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *DemoMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case demo.FieldName:
		return m.Name()
	case demo.FieldCreateTime:
		return m.CreateTime()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *DemoMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case demo.FieldName:
		return m.OldName(ctx)
	case demo.FieldCreateTime:
		return m.OldCreateTime(ctx)
	}
	return nil, fmt.Errorf("unknown Demo field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *DemoMutation) SetField(name string, value ent.Value) error {
	switch name {
	case demo.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case demo.FieldCreateTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreateTime(v)
		return nil
	}
	return fmt.Errorf("unknown Demo field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *DemoMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *DemoMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *DemoMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Demo numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *DemoMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(demo.FieldName) {
		fields = append(fields, demo.FieldName)
	}
	if m.FieldCleared(demo.FieldCreateTime) {
		fields = append(fields, demo.FieldCreateTime)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *DemoMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *DemoMutation) ClearField(name string) error {
	switch name {
	case demo.FieldName:
		m.ClearName()
		return nil
	case demo.FieldCreateTime:
		m.ClearCreateTime()
		return nil
	}
	return fmt.Errorf("unknown Demo nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *DemoMutation) ResetField(name string) error {
	switch name {
	case demo.FieldName:
		m.ResetName()
		return nil
	case demo.FieldCreateTime:
		m.ResetCreateTime()
		return nil
	}
	return fmt.Errorf("unknown Demo field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *DemoMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *DemoMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *DemoMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *DemoMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *DemoMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *DemoMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *DemoMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Demo unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *DemoMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Demo edge %s", name)
}
