// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"message-push/app/receiver/common/model/po/ent/businessclient"
	"message-push/app/receiver/common/model/po/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BusinessClientUpdate is the builder for updating BusinessClient entities.
type BusinessClientUpdate struct {
	config
	hooks     []Hook
	mutation  *BusinessClientMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the BusinessClientUpdate builder.
func (bcu *BusinessClientUpdate) Where(ps ...predicate.BusinessClient) *BusinessClientUpdate {
	bcu.mutation.Where(ps...)
	return bcu
}

// SetName sets the "name" field.
func (bcu *BusinessClientUpdate) SetName(s string) *BusinessClientUpdate {
	bcu.mutation.SetName(s)
	return bcu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (bcu *BusinessClientUpdate) SetNillableName(s *string) *BusinessClientUpdate {
	if s != nil {
		bcu.SetName(*s)
	}
	return bcu
}

// SetClientKey sets the "client_key" field.
func (bcu *BusinessClientUpdate) SetClientKey(s string) *BusinessClientUpdate {
	bcu.mutation.SetClientKey(s)
	return bcu
}

// SetNillableClientKey sets the "client_key" field if the given value is not nil.
func (bcu *BusinessClientUpdate) SetNillableClientKey(s *string) *BusinessClientUpdate {
	if s != nil {
		bcu.SetClientKey(*s)
	}
	return bcu
}

// SetAppID sets the "app_id" field.
func (bcu *BusinessClientUpdate) SetAppID(s string) *BusinessClientUpdate {
	bcu.mutation.SetAppID(s)
	return bcu
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (bcu *BusinessClientUpdate) SetNillableAppID(s *string) *BusinessClientUpdate {
	if s != nil {
		bcu.SetAppID(*s)
	}
	return bcu
}

// SetClientID sets the "client_id" field.
func (bcu *BusinessClientUpdate) SetClientID(s string) *BusinessClientUpdate {
	bcu.mutation.SetClientID(s)
	return bcu
}

// SetNillableClientID sets the "client_id" field if the given value is not nil.
func (bcu *BusinessClientUpdate) SetNillableClientID(s *string) *BusinessClientUpdate {
	if s != nil {
		bcu.SetClientID(*s)
	}
	return bcu
}

// SetClientType sets the "client_type" field.
func (bcu *BusinessClientUpdate) SetClientType(s string) *BusinessClientUpdate {
	bcu.mutation.SetClientType(s)
	return bcu
}

// SetNillableClientType sets the "client_type" field if the given value is not nil.
func (bcu *BusinessClientUpdate) SetNillableClientType(s *string) *BusinessClientUpdate {
	if s != nil {
		bcu.SetClientType(*s)
	}
	return bcu
}

// SetIsDel sets the "is_del" field.
func (bcu *BusinessClientUpdate) SetIsDel(i int) *BusinessClientUpdate {
	bcu.mutation.ResetIsDel()
	bcu.mutation.SetIsDel(i)
	return bcu
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (bcu *BusinessClientUpdate) SetNillableIsDel(i *int) *BusinessClientUpdate {
	if i != nil {
		bcu.SetIsDel(*i)
	}
	return bcu
}

// AddIsDel adds i to the "is_del" field.
func (bcu *BusinessClientUpdate) AddIsDel(i int) *BusinessClientUpdate {
	bcu.mutation.AddIsDel(i)
	return bcu
}

// SetUpdateTime sets the "update_time" field.
func (bcu *BusinessClientUpdate) SetUpdateTime(t time.Time) *BusinessClientUpdate {
	bcu.mutation.SetUpdateTime(t)
	return bcu
}

// Mutation returns the BusinessClientMutation object of the builder.
func (bcu *BusinessClientUpdate) Mutation() *BusinessClientMutation {
	return bcu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bcu *BusinessClientUpdate) Save(ctx context.Context) (int, error) {
	bcu.defaults()
	return withHooks(ctx, bcu.sqlSave, bcu.mutation, bcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bcu *BusinessClientUpdate) SaveX(ctx context.Context) int {
	affected, err := bcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bcu *BusinessClientUpdate) Exec(ctx context.Context) error {
	_, err := bcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcu *BusinessClientUpdate) ExecX(ctx context.Context) {
	if err := bcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bcu *BusinessClientUpdate) defaults() {
	if _, ok := bcu.mutation.UpdateTime(); !ok {
		v := businessclient.UpdateDefaultUpdateTime()
		bcu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bcu *BusinessClientUpdate) check() error {
	if v, ok := bcu.mutation.Name(); ok {
		if err := businessclient.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "BusinessClient.name": %w`, err)}
		}
	}
	if v, ok := bcu.mutation.ClientKey(); ok {
		if err := businessclient.ClientKeyValidator(v); err != nil {
			return &ValidationError{Name: "client_key", err: fmt.Errorf(`ent: validator failed for field "BusinessClient.client_key": %w`, err)}
		}
	}
	if v, ok := bcu.mutation.AppID(); ok {
		if err := businessclient.AppIDValidator(v); err != nil {
			return &ValidationError{Name: "app_id", err: fmt.Errorf(`ent: validator failed for field "BusinessClient.app_id": %w`, err)}
		}
	}
	if v, ok := bcu.mutation.ClientID(); ok {
		if err := businessclient.ClientIDValidator(v); err != nil {
			return &ValidationError{Name: "client_id", err: fmt.Errorf(`ent: validator failed for field "BusinessClient.client_id": %w`, err)}
		}
	}
	if v, ok := bcu.mutation.ClientType(); ok {
		if err := businessclient.ClientTypeValidator(v); err != nil {
			return &ValidationError{Name: "client_type", err: fmt.Errorf(`ent: validator failed for field "BusinessClient.client_type": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (bcu *BusinessClientUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *BusinessClientUpdate {
	bcu.modifiers = append(bcu.modifiers, modifiers...)
	return bcu
}

func (bcu *BusinessClientUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := bcu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(businessclient.Table, businessclient.Columns, sqlgraph.NewFieldSpec(businessclient.FieldID, field.TypeInt))
	if ps := bcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bcu.mutation.Name(); ok {
		_spec.SetField(businessclient.FieldName, field.TypeString, value)
	}
	if value, ok := bcu.mutation.ClientKey(); ok {
		_spec.SetField(businessclient.FieldClientKey, field.TypeString, value)
	}
	if value, ok := bcu.mutation.AppID(); ok {
		_spec.SetField(businessclient.FieldAppID, field.TypeString, value)
	}
	if value, ok := bcu.mutation.ClientID(); ok {
		_spec.SetField(businessclient.FieldClientID, field.TypeString, value)
	}
	if value, ok := bcu.mutation.ClientType(); ok {
		_spec.SetField(businessclient.FieldClientType, field.TypeString, value)
	}
	if value, ok := bcu.mutation.IsDel(); ok {
		_spec.SetField(businessclient.FieldIsDel, field.TypeInt, value)
	}
	if value, ok := bcu.mutation.AddedIsDel(); ok {
		_spec.AddField(businessclient.FieldIsDel, field.TypeInt, value)
	}
	if value, ok := bcu.mutation.UpdateTime(); ok {
		_spec.SetField(businessclient.FieldUpdateTime, field.TypeTime, value)
	}
	_spec.AddModifiers(bcu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, bcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{businessclient.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bcu.mutation.done = true
	return n, nil
}

// BusinessClientUpdateOne is the builder for updating a single BusinessClient entity.
type BusinessClientUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *BusinessClientMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetName sets the "name" field.
func (bcuo *BusinessClientUpdateOne) SetName(s string) *BusinessClientUpdateOne {
	bcuo.mutation.SetName(s)
	return bcuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (bcuo *BusinessClientUpdateOne) SetNillableName(s *string) *BusinessClientUpdateOne {
	if s != nil {
		bcuo.SetName(*s)
	}
	return bcuo
}

// SetClientKey sets the "client_key" field.
func (bcuo *BusinessClientUpdateOne) SetClientKey(s string) *BusinessClientUpdateOne {
	bcuo.mutation.SetClientKey(s)
	return bcuo
}

// SetNillableClientKey sets the "client_key" field if the given value is not nil.
func (bcuo *BusinessClientUpdateOne) SetNillableClientKey(s *string) *BusinessClientUpdateOne {
	if s != nil {
		bcuo.SetClientKey(*s)
	}
	return bcuo
}

// SetAppID sets the "app_id" field.
func (bcuo *BusinessClientUpdateOne) SetAppID(s string) *BusinessClientUpdateOne {
	bcuo.mutation.SetAppID(s)
	return bcuo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (bcuo *BusinessClientUpdateOne) SetNillableAppID(s *string) *BusinessClientUpdateOne {
	if s != nil {
		bcuo.SetAppID(*s)
	}
	return bcuo
}

// SetClientID sets the "client_id" field.
func (bcuo *BusinessClientUpdateOne) SetClientID(s string) *BusinessClientUpdateOne {
	bcuo.mutation.SetClientID(s)
	return bcuo
}

// SetNillableClientID sets the "client_id" field if the given value is not nil.
func (bcuo *BusinessClientUpdateOne) SetNillableClientID(s *string) *BusinessClientUpdateOne {
	if s != nil {
		bcuo.SetClientID(*s)
	}
	return bcuo
}

// SetClientType sets the "client_type" field.
func (bcuo *BusinessClientUpdateOne) SetClientType(s string) *BusinessClientUpdateOne {
	bcuo.mutation.SetClientType(s)
	return bcuo
}

// SetNillableClientType sets the "client_type" field if the given value is not nil.
func (bcuo *BusinessClientUpdateOne) SetNillableClientType(s *string) *BusinessClientUpdateOne {
	if s != nil {
		bcuo.SetClientType(*s)
	}
	return bcuo
}

// SetIsDel sets the "is_del" field.
func (bcuo *BusinessClientUpdateOne) SetIsDel(i int) *BusinessClientUpdateOne {
	bcuo.mutation.ResetIsDel()
	bcuo.mutation.SetIsDel(i)
	return bcuo
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (bcuo *BusinessClientUpdateOne) SetNillableIsDel(i *int) *BusinessClientUpdateOne {
	if i != nil {
		bcuo.SetIsDel(*i)
	}
	return bcuo
}

// AddIsDel adds i to the "is_del" field.
func (bcuo *BusinessClientUpdateOne) AddIsDel(i int) *BusinessClientUpdateOne {
	bcuo.mutation.AddIsDel(i)
	return bcuo
}

// SetUpdateTime sets the "update_time" field.
func (bcuo *BusinessClientUpdateOne) SetUpdateTime(t time.Time) *BusinessClientUpdateOne {
	bcuo.mutation.SetUpdateTime(t)
	return bcuo
}

// Mutation returns the BusinessClientMutation object of the builder.
func (bcuo *BusinessClientUpdateOne) Mutation() *BusinessClientMutation {
	return bcuo.mutation
}

// Where appends a list predicates to the BusinessClientUpdate builder.
func (bcuo *BusinessClientUpdateOne) Where(ps ...predicate.BusinessClient) *BusinessClientUpdateOne {
	bcuo.mutation.Where(ps...)
	return bcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bcuo *BusinessClientUpdateOne) Select(field string, fields ...string) *BusinessClientUpdateOne {
	bcuo.fields = append([]string{field}, fields...)
	return bcuo
}

// Save executes the query and returns the updated BusinessClient entity.
func (bcuo *BusinessClientUpdateOne) Save(ctx context.Context) (*BusinessClient, error) {
	bcuo.defaults()
	return withHooks(ctx, bcuo.sqlSave, bcuo.mutation, bcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bcuo *BusinessClientUpdateOne) SaveX(ctx context.Context) *BusinessClient {
	node, err := bcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bcuo *BusinessClientUpdateOne) Exec(ctx context.Context) error {
	_, err := bcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcuo *BusinessClientUpdateOne) ExecX(ctx context.Context) {
	if err := bcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bcuo *BusinessClientUpdateOne) defaults() {
	if _, ok := bcuo.mutation.UpdateTime(); !ok {
		v := businessclient.UpdateDefaultUpdateTime()
		bcuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bcuo *BusinessClientUpdateOne) check() error {
	if v, ok := bcuo.mutation.Name(); ok {
		if err := businessclient.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "BusinessClient.name": %w`, err)}
		}
	}
	if v, ok := bcuo.mutation.ClientKey(); ok {
		if err := businessclient.ClientKeyValidator(v); err != nil {
			return &ValidationError{Name: "client_key", err: fmt.Errorf(`ent: validator failed for field "BusinessClient.client_key": %w`, err)}
		}
	}
	if v, ok := bcuo.mutation.AppID(); ok {
		if err := businessclient.AppIDValidator(v); err != nil {
			return &ValidationError{Name: "app_id", err: fmt.Errorf(`ent: validator failed for field "BusinessClient.app_id": %w`, err)}
		}
	}
	if v, ok := bcuo.mutation.ClientID(); ok {
		if err := businessclient.ClientIDValidator(v); err != nil {
			return &ValidationError{Name: "client_id", err: fmt.Errorf(`ent: validator failed for field "BusinessClient.client_id": %w`, err)}
		}
	}
	if v, ok := bcuo.mutation.ClientType(); ok {
		if err := businessclient.ClientTypeValidator(v); err != nil {
			return &ValidationError{Name: "client_type", err: fmt.Errorf(`ent: validator failed for field "BusinessClient.client_type": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (bcuo *BusinessClientUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *BusinessClientUpdateOne {
	bcuo.modifiers = append(bcuo.modifiers, modifiers...)
	return bcuo
}

func (bcuo *BusinessClientUpdateOne) sqlSave(ctx context.Context) (_node *BusinessClient, err error) {
	if err := bcuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(businessclient.Table, businessclient.Columns, sqlgraph.NewFieldSpec(businessclient.FieldID, field.TypeInt))
	id, ok := bcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "BusinessClient.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := bcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, businessclient.FieldID)
		for _, f := range fields {
			if !businessclient.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != businessclient.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bcuo.mutation.Name(); ok {
		_spec.SetField(businessclient.FieldName, field.TypeString, value)
	}
	if value, ok := bcuo.mutation.ClientKey(); ok {
		_spec.SetField(businessclient.FieldClientKey, field.TypeString, value)
	}
	if value, ok := bcuo.mutation.AppID(); ok {
		_spec.SetField(businessclient.FieldAppID, field.TypeString, value)
	}
	if value, ok := bcuo.mutation.ClientID(); ok {
		_spec.SetField(businessclient.FieldClientID, field.TypeString, value)
	}
	if value, ok := bcuo.mutation.ClientType(); ok {
		_spec.SetField(businessclient.FieldClientType, field.TypeString, value)
	}
	if value, ok := bcuo.mutation.IsDel(); ok {
		_spec.SetField(businessclient.FieldIsDel, field.TypeInt, value)
	}
	if value, ok := bcuo.mutation.AddedIsDel(); ok {
		_spec.AddField(businessclient.FieldIsDel, field.TypeInt, value)
	}
	if value, ok := bcuo.mutation.UpdateTime(); ok {
		_spec.SetField(businessclient.FieldUpdateTime, field.TypeTime, value)
	}
	_spec.AddModifiers(bcuo.modifiers...)
	_node = &BusinessClient{config: bcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{businessclient.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	bcuo.mutation.done = true
	return _node, nil
}
