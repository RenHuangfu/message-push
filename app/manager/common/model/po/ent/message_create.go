// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"message-push/app/receiver/common/model/po/ent/message"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MessageCreate is the builder for creating a Message entity.
type MessageCreate struct {
	config
	mutation *MessageMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetAppID sets the "app_id" field.
func (mc *MessageCreate) SetAppID(i int) *MessageCreate {
	mc.mutation.SetAppID(i)
	return mc
}

// SetClientIds sets the "client_ids" field.
func (mc *MessageCreate) SetClientIds(i []int) *MessageCreate {
	mc.mutation.SetClientIds(i)
	return mc
}

// SetContent sets the "content" field.
func (mc *MessageCreate) SetContent(s string) *MessageCreate {
	mc.mutation.SetContent(s)
	return mc
}

// SetDelay sets the "delay" field.
func (mc *MessageCreate) SetDelay(i int) *MessageCreate {
	mc.mutation.SetDelay(i)
	return mc
}

// SetSendTime sets the "send_time" field.
func (mc *MessageCreate) SetSendTime(t time.Time) *MessageCreate {
	mc.mutation.SetSendTime(t)
	return mc
}

// SetSendCount sets the "send_count" field.
func (mc *MessageCreate) SetSendCount(i int) *MessageCreate {
	mc.mutation.SetSendCount(i)
	return mc
}

// SetStatus sets the "status" field.
func (mc *MessageCreate) SetStatus(i int) *MessageCreate {
	mc.mutation.SetStatus(i)
	return mc
}

// SetIsDel sets the "is_del" field.
func (mc *MessageCreate) SetIsDel(i int) *MessageCreate {
	mc.mutation.SetIsDel(i)
	return mc
}

// SetCreateTime sets the "create_time" field.
func (mc *MessageCreate) SetCreateTime(t time.Time) *MessageCreate {
	mc.mutation.SetCreateTime(t)
	return mc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (mc *MessageCreate) SetNillableCreateTime(t *time.Time) *MessageCreate {
	if t != nil {
		mc.SetCreateTime(*t)
	}
	return mc
}

// SetUpdateTime sets the "update_time" field.
func (mc *MessageCreate) SetUpdateTime(t time.Time) *MessageCreate {
	mc.mutation.SetUpdateTime(t)
	return mc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (mc *MessageCreate) SetNillableUpdateTime(t *time.Time) *MessageCreate {
	if t != nil {
		mc.SetUpdateTime(*t)
	}
	return mc
}

// SetID sets the "id" field.
func (mc *MessageCreate) SetID(i int) *MessageCreate {
	mc.mutation.SetID(i)
	return mc
}

// Mutation returns the MessageMutation object of the builder.
func (mc *MessageCreate) Mutation() *MessageMutation {
	return mc.mutation
}

// Save creates the Message in the database.
func (mc *MessageCreate) Save(ctx context.Context) (*Message, error) {
	mc.defaults()
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MessageCreate) SaveX(ctx context.Context) *Message {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MessageCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MessageCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MessageCreate) defaults() {
	if _, ok := mc.mutation.CreateTime(); !ok {
		v := message.DefaultCreateTime()
		mc.mutation.SetCreateTime(v)
	}
	if _, ok := mc.mutation.UpdateTime(); !ok {
		v := message.DefaultUpdateTime()
		mc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MessageCreate) check() error {
	if _, ok := mc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "Message.app_id"`)}
	}
	if _, ok := mc.mutation.ClientIds(); !ok {
		return &ValidationError{Name: "client_ids", err: errors.New(`ent: missing required field "Message.client_ids"`)}
	}
	if _, ok := mc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "Message.content"`)}
	}
	if _, ok := mc.mutation.Delay(); !ok {
		return &ValidationError{Name: "delay", err: errors.New(`ent: missing required field "Message.delay"`)}
	}
	if _, ok := mc.mutation.SendTime(); !ok {
		return &ValidationError{Name: "send_time", err: errors.New(`ent: missing required field "Message.send_time"`)}
	}
	if _, ok := mc.mutation.SendCount(); !ok {
		return &ValidationError{Name: "send_count", err: errors.New(`ent: missing required field "Message.send_count"`)}
	}
	if _, ok := mc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Message.status"`)}
	}
	if _, ok := mc.mutation.IsDel(); !ok {
		return &ValidationError{Name: "is_del", err: errors.New(`ent: missing required field "Message.is_del"`)}
	}
	if _, ok := mc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Message.create_time"`)}
	}
	if _, ok := mc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Message.update_time"`)}
	}
	return nil
}

func (mc *MessageCreate) sqlSave(ctx context.Context) (*Message, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MessageCreate) createSpec() (*Message, *sqlgraph.CreateSpec) {
	var (
		_node = &Message{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(message.Table, sqlgraph.NewFieldSpec(message.FieldID, field.TypeInt))
	)
	_spec.OnConflict = mc.conflict
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mc.mutation.AppID(); ok {
		_spec.SetField(message.FieldAppID, field.TypeInt, value)
		_node.AppID = value
	}
	if value, ok := mc.mutation.ClientIds(); ok {
		_spec.SetField(message.FieldClientIds, field.TypeJSON, value)
		_node.ClientIds = value
	}
	if value, ok := mc.mutation.Content(); ok {
		_spec.SetField(message.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if value, ok := mc.mutation.Delay(); ok {
		_spec.SetField(message.FieldDelay, field.TypeInt, value)
		_node.Delay = value
	}
	if value, ok := mc.mutation.SendTime(); ok {
		_spec.SetField(message.FieldSendTime, field.TypeTime, value)
		_node.SendTime = value
	}
	if value, ok := mc.mutation.SendCount(); ok {
		_spec.SetField(message.FieldSendCount, field.TypeInt, value)
		_node.SendCount = value
	}
	if value, ok := mc.mutation.Status(); ok {
		_spec.SetField(message.FieldStatus, field.TypeInt, value)
		_node.Status = value
	}
	if value, ok := mc.mutation.IsDel(); ok {
		_spec.SetField(message.FieldIsDel, field.TypeInt, value)
		_node.IsDel = value
	}
	if value, ok := mc.mutation.CreateTime(); ok {
		_spec.SetField(message.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := mc.mutation.UpdateTime(); ok {
		_spec.SetField(message.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Message.Create().
//		SetAppID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MessageUpsert) {
//			SetAppID(v+v).
//		}).
//		Exec(ctx)
func (mc *MessageCreate) OnConflict(opts ...sql.ConflictOption) *MessageUpsertOne {
	mc.conflict = opts
	return &MessageUpsertOne{
		create: mc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Message.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mc *MessageCreate) OnConflictColumns(columns ...string) *MessageUpsertOne {
	mc.conflict = append(mc.conflict, sql.ConflictColumns(columns...))
	return &MessageUpsertOne{
		create: mc,
	}
}

type (
	// MessageUpsertOne is the builder for "upsert"-ing
	//  one Message node.
	MessageUpsertOne struct {
		create *MessageCreate
	}

	// MessageUpsert is the "OnConflict" setter.
	MessageUpsert struct {
		*sql.UpdateSet
	}
)

// SetAppID sets the "app_id" field.
func (u *MessageUpsert) SetAppID(v int) *MessageUpsert {
	u.Set(message.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *MessageUpsert) UpdateAppID() *MessageUpsert {
	u.SetExcluded(message.FieldAppID)
	return u
}

// AddAppID adds v to the "app_id" field.
func (u *MessageUpsert) AddAppID(v int) *MessageUpsert {
	u.Add(message.FieldAppID, v)
	return u
}

// SetClientIds sets the "client_ids" field.
func (u *MessageUpsert) SetClientIds(v []int) *MessageUpsert {
	u.Set(message.FieldClientIds, v)
	return u
}

// UpdateClientIds sets the "client_ids" field to the value that was provided on create.
func (u *MessageUpsert) UpdateClientIds() *MessageUpsert {
	u.SetExcluded(message.FieldClientIds)
	return u
}

// SetContent sets the "content" field.
func (u *MessageUpsert) SetContent(v string) *MessageUpsert {
	u.Set(message.FieldContent, v)
	return u
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *MessageUpsert) UpdateContent() *MessageUpsert {
	u.SetExcluded(message.FieldContent)
	return u
}

// SetDelay sets the "delay" field.
func (u *MessageUpsert) SetDelay(v int) *MessageUpsert {
	u.Set(message.FieldDelay, v)
	return u
}

// UpdateDelay sets the "delay" field to the value that was provided on create.
func (u *MessageUpsert) UpdateDelay() *MessageUpsert {
	u.SetExcluded(message.FieldDelay)
	return u
}

// AddDelay adds v to the "delay" field.
func (u *MessageUpsert) AddDelay(v int) *MessageUpsert {
	u.Add(message.FieldDelay, v)
	return u
}

// SetSendTime sets the "send_time" field.
func (u *MessageUpsert) SetSendTime(v time.Time) *MessageUpsert {
	u.Set(message.FieldSendTime, v)
	return u
}

// UpdateSendTime sets the "send_time" field to the value that was provided on create.
func (u *MessageUpsert) UpdateSendTime() *MessageUpsert {
	u.SetExcluded(message.FieldSendTime)
	return u
}

// SetSendCount sets the "send_count" field.
func (u *MessageUpsert) SetSendCount(v int) *MessageUpsert {
	u.Set(message.FieldSendCount, v)
	return u
}

// UpdateSendCount sets the "send_count" field to the value that was provided on create.
func (u *MessageUpsert) UpdateSendCount() *MessageUpsert {
	u.SetExcluded(message.FieldSendCount)
	return u
}

// AddSendCount adds v to the "send_count" field.
func (u *MessageUpsert) AddSendCount(v int) *MessageUpsert {
	u.Add(message.FieldSendCount, v)
	return u
}

// SetStatus sets the "status" field.
func (u *MessageUpsert) SetStatus(v int) *MessageUpsert {
	u.Set(message.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *MessageUpsert) UpdateStatus() *MessageUpsert {
	u.SetExcluded(message.FieldStatus)
	return u
}

// AddStatus adds v to the "status" field.
func (u *MessageUpsert) AddStatus(v int) *MessageUpsert {
	u.Add(message.FieldStatus, v)
	return u
}

// SetIsDel sets the "is_del" field.
func (u *MessageUpsert) SetIsDel(v int) *MessageUpsert {
	u.Set(message.FieldIsDel, v)
	return u
}

// UpdateIsDel sets the "is_del" field to the value that was provided on create.
func (u *MessageUpsert) UpdateIsDel() *MessageUpsert {
	u.SetExcluded(message.FieldIsDel)
	return u
}

// AddIsDel adds v to the "is_del" field.
func (u *MessageUpsert) AddIsDel(v int) *MessageUpsert {
	u.Add(message.FieldIsDel, v)
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *MessageUpsert) SetUpdateTime(v time.Time) *MessageUpsert {
	u.Set(message.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *MessageUpsert) UpdateUpdateTime() *MessageUpsert {
	u.SetExcluded(message.FieldUpdateTime)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Message.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(message.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MessageUpsertOne) UpdateNewValues() *MessageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(message.FieldID)
		}
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(message.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Message.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *MessageUpsertOne) Ignore() *MessageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MessageUpsertOne) DoNothing() *MessageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MessageCreate.OnConflict
// documentation for more info.
func (u *MessageUpsertOne) Update(set func(*MessageUpsert)) *MessageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MessageUpsert{UpdateSet: update})
	}))
	return u
}

// SetAppID sets the "app_id" field.
func (u *MessageUpsertOne) SetAppID(v int) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetAppID(v)
	})
}

// AddAppID adds v to the "app_id" field.
func (u *MessageUpsertOne) AddAppID(v int) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.AddAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateAppID() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateAppID()
	})
}

// SetClientIds sets the "client_ids" field.
func (u *MessageUpsertOne) SetClientIds(v []int) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetClientIds(v)
	})
}

// UpdateClientIds sets the "client_ids" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateClientIds() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateClientIds()
	})
}

// SetContent sets the "content" field.
func (u *MessageUpsertOne) SetContent(v string) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateContent() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateContent()
	})
}

// SetDelay sets the "delay" field.
func (u *MessageUpsertOne) SetDelay(v int) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetDelay(v)
	})
}

// AddDelay adds v to the "delay" field.
func (u *MessageUpsertOne) AddDelay(v int) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.AddDelay(v)
	})
}

// UpdateDelay sets the "delay" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateDelay() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateDelay()
	})
}

// SetSendTime sets the "send_time" field.
func (u *MessageUpsertOne) SetSendTime(v time.Time) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetSendTime(v)
	})
}

// UpdateSendTime sets the "send_time" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateSendTime() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateSendTime()
	})
}

// SetSendCount sets the "send_count" field.
func (u *MessageUpsertOne) SetSendCount(v int) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetSendCount(v)
	})
}

// AddSendCount adds v to the "send_count" field.
func (u *MessageUpsertOne) AddSendCount(v int) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.AddSendCount(v)
	})
}

// UpdateSendCount sets the "send_count" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateSendCount() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateSendCount()
	})
}

// SetStatus sets the "status" field.
func (u *MessageUpsertOne) SetStatus(v int) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetStatus(v)
	})
}

// AddStatus adds v to the "status" field.
func (u *MessageUpsertOne) AddStatus(v int) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.AddStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateStatus() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateStatus()
	})
}

// SetIsDel sets the "is_del" field.
func (u *MessageUpsertOne) SetIsDel(v int) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetIsDel(v)
	})
}

// AddIsDel adds v to the "is_del" field.
func (u *MessageUpsertOne) AddIsDel(v int) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.AddIsDel(v)
	})
}

// UpdateIsDel sets the "is_del" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateIsDel() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateIsDel()
	})
}

// SetUpdateTime sets the "update_time" field.
func (u *MessageUpsertOne) SetUpdateTime(v time.Time) *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *MessageUpsertOne) UpdateUpdateTime() *MessageUpsertOne {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateUpdateTime()
	})
}

// Exec executes the query.
func (u *MessageUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MessageCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MessageUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *MessageUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *MessageUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// MessageCreateBulk is the builder for creating many Message entities in bulk.
type MessageCreateBulk struct {
	config
	err      error
	builders []*MessageCreate
	conflict []sql.ConflictOption
}

// Save creates the Message entities in the database.
func (mcb *MessageCreateBulk) Save(ctx context.Context) ([]*Message, error) {
	if mcb.err != nil {
		return nil, mcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Message, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MessageMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = mcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MessageCreateBulk) SaveX(ctx context.Context) []*Message {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MessageCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MessageCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Message.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MessageUpsert) {
//			SetAppID(v+v).
//		}).
//		Exec(ctx)
func (mcb *MessageCreateBulk) OnConflict(opts ...sql.ConflictOption) *MessageUpsertBulk {
	mcb.conflict = opts
	return &MessageUpsertBulk{
		create: mcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Message.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mcb *MessageCreateBulk) OnConflictColumns(columns ...string) *MessageUpsertBulk {
	mcb.conflict = append(mcb.conflict, sql.ConflictColumns(columns...))
	return &MessageUpsertBulk{
		create: mcb,
	}
}

// MessageUpsertBulk is the builder for "upsert"-ing
// a bulk of Message nodes.
type MessageUpsertBulk struct {
	create *MessageCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Message.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(message.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MessageUpsertBulk) UpdateNewValues() *MessageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(message.FieldID)
			}
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(message.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Message.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *MessageUpsertBulk) Ignore() *MessageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MessageUpsertBulk) DoNothing() *MessageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MessageCreateBulk.OnConflict
// documentation for more info.
func (u *MessageUpsertBulk) Update(set func(*MessageUpsert)) *MessageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MessageUpsert{UpdateSet: update})
	}))
	return u
}

// SetAppID sets the "app_id" field.
func (u *MessageUpsertBulk) SetAppID(v int) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetAppID(v)
	})
}

// AddAppID adds v to the "app_id" field.
func (u *MessageUpsertBulk) AddAppID(v int) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.AddAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateAppID() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateAppID()
	})
}

// SetClientIds sets the "client_ids" field.
func (u *MessageUpsertBulk) SetClientIds(v []int) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetClientIds(v)
	})
}

// UpdateClientIds sets the "client_ids" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateClientIds() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateClientIds()
	})
}

// SetContent sets the "content" field.
func (u *MessageUpsertBulk) SetContent(v string) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateContent() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateContent()
	})
}

// SetDelay sets the "delay" field.
func (u *MessageUpsertBulk) SetDelay(v int) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetDelay(v)
	})
}

// AddDelay adds v to the "delay" field.
func (u *MessageUpsertBulk) AddDelay(v int) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.AddDelay(v)
	})
}

// UpdateDelay sets the "delay" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateDelay() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateDelay()
	})
}

// SetSendTime sets the "send_time" field.
func (u *MessageUpsertBulk) SetSendTime(v time.Time) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetSendTime(v)
	})
}

// UpdateSendTime sets the "send_time" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateSendTime() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateSendTime()
	})
}

// SetSendCount sets the "send_count" field.
func (u *MessageUpsertBulk) SetSendCount(v int) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetSendCount(v)
	})
}

// AddSendCount adds v to the "send_count" field.
func (u *MessageUpsertBulk) AddSendCount(v int) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.AddSendCount(v)
	})
}

// UpdateSendCount sets the "send_count" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateSendCount() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateSendCount()
	})
}

// SetStatus sets the "status" field.
func (u *MessageUpsertBulk) SetStatus(v int) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetStatus(v)
	})
}

// AddStatus adds v to the "status" field.
func (u *MessageUpsertBulk) AddStatus(v int) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.AddStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateStatus() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateStatus()
	})
}

// SetIsDel sets the "is_del" field.
func (u *MessageUpsertBulk) SetIsDel(v int) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetIsDel(v)
	})
}

// AddIsDel adds v to the "is_del" field.
func (u *MessageUpsertBulk) AddIsDel(v int) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.AddIsDel(v)
	})
}

// UpdateIsDel sets the "is_del" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateIsDel() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateIsDel()
	})
}

// SetUpdateTime sets the "update_time" field.
func (u *MessageUpsertBulk) SetUpdateTime(v time.Time) *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *MessageUpsertBulk) UpdateUpdateTime() *MessageUpsertBulk {
	return u.Update(func(s *MessageUpsert) {
		s.UpdateUpdateTime()
	})
}

// Exec executes the query.
func (u *MessageUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the MessageCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MessageCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MessageUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
