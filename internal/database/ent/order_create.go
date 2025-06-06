// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/blackPavlin/shop/internal/database/ent/order"
	"github.com/blackPavlin/shop/internal/database/ent/orderproduct"
	"github.com/blackPavlin/shop/internal/database/ent/user"
)

// OrderCreate is the builder for creating a Order entity.
type OrderCreate struct {
	config
	mutation *OrderMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (oc *OrderCreate) SetCreatedAt(t time.Time) *OrderCreate {
	oc.mutation.SetCreatedAt(t)
	return oc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (oc *OrderCreate) SetNillableCreatedAt(t *time.Time) *OrderCreate {
	if t != nil {
		oc.SetCreatedAt(*t)
	}
	return oc
}

// SetUpdatedAt sets the "updated_at" field.
func (oc *OrderCreate) SetUpdatedAt(t time.Time) *OrderCreate {
	oc.mutation.SetUpdatedAt(t)
	return oc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (oc *OrderCreate) SetNillableUpdatedAt(t *time.Time) *OrderCreate {
	if t != nil {
		oc.SetUpdatedAt(*t)
	}
	return oc
}

// SetUserID sets the "user_id" field.
func (oc *OrderCreate) SetUserID(i int64) *OrderCreate {
	oc.mutation.SetUserID(i)
	return oc
}

// SetStatus sets the "status" field.
func (oc *OrderCreate) SetStatus(o order.Status) *OrderCreate {
	oc.mutation.SetStatus(o)
	return oc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (oc *OrderCreate) SetNillableStatus(o *order.Status) *OrderCreate {
	if o != nil {
		oc.SetStatus(*o)
	}
	return oc
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (oc *OrderCreate) SetUsersID(id int64) *OrderCreate {
	oc.mutation.SetUsersID(id)
	return oc
}

// SetUsers sets the "users" edge to the User entity.
func (oc *OrderCreate) SetUsers(u *User) *OrderCreate {
	return oc.SetUsersID(u.ID)
}

// AddOrderProductIDs adds the "order_products" edge to the OrderProduct entity by IDs.
func (oc *OrderCreate) AddOrderProductIDs(ids ...int64) *OrderCreate {
	oc.mutation.AddOrderProductIDs(ids...)
	return oc
}

// AddOrderProducts adds the "order_products" edges to the OrderProduct entity.
func (oc *OrderCreate) AddOrderProducts(o ...*OrderProduct) *OrderCreate {
	ids := make([]int64, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oc.AddOrderProductIDs(ids...)
}

// Mutation returns the OrderMutation object of the builder.
func (oc *OrderCreate) Mutation() *OrderMutation {
	return oc.mutation
}

// Save creates the Order in the database.
func (oc *OrderCreate) Save(ctx context.Context) (*Order, error) {
	oc.defaults()
	return withHooks(ctx, oc.sqlSave, oc.mutation, oc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OrderCreate) SaveX(ctx context.Context) *Order {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oc *OrderCreate) Exec(ctx context.Context) error {
	_, err := oc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oc *OrderCreate) ExecX(ctx context.Context) {
	if err := oc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oc *OrderCreate) defaults() {
	if _, ok := oc.mutation.CreatedAt(); !ok {
		v := order.DefaultCreatedAt()
		oc.mutation.SetCreatedAt(v)
	}
	if _, ok := oc.mutation.UpdatedAt(); !ok {
		v := order.DefaultUpdatedAt()
		oc.mutation.SetUpdatedAt(v)
	}
	if _, ok := oc.mutation.Status(); !ok {
		v := order.DefaultStatus
		oc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oc *OrderCreate) check() error {
	if _, ok := oc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Order.created_at"`)}
	}
	if _, ok := oc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Order.updated_at"`)}
	}
	if _, ok := oc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Order.user_id"`)}
	}
	if _, ok := oc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Order.status"`)}
	}
	if v, ok := oc.mutation.Status(); ok {
		if err := order.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Order.status": %w`, err)}
		}
	}
	if len(oc.mutation.UsersIDs()) == 0 {
		return &ValidationError{Name: "users", err: errors.New(`ent: missing required edge "Order.users"`)}
	}
	return nil
}

func (oc *OrderCreate) sqlSave(ctx context.Context) (*Order, error) {
	if err := oc.check(); err != nil {
		return nil, err
	}
	_node, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int64(id)
	oc.mutation.id = &_node.ID
	oc.mutation.done = true
	return _node, nil
}

func (oc *OrderCreate) createSpec() (*Order, *sqlgraph.CreateSpec) {
	var (
		_node = &Order{config: oc.config}
		_spec = sqlgraph.NewCreateSpec(order.Table, sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = oc.conflict
	if value, ok := oc.mutation.CreatedAt(); ok {
		_spec.SetField(order.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := oc.mutation.UpdatedAt(); ok {
		_spec.SetField(order.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := oc.mutation.Status(); ok {
		_spec.SetField(order.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if nodes := oc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.UsersTable,
			Columns: []string{order.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.OrderProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   order.OrderProductsTable,
			Columns: []string{order.OrderProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orderproduct.FieldID, field.TypeInt64),
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
//	client.Order.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OrderUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (oc *OrderCreate) OnConflict(opts ...sql.ConflictOption) *OrderUpsertOne {
	oc.conflict = opts
	return &OrderUpsertOne{
		create: oc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Order.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (oc *OrderCreate) OnConflictColumns(columns ...string) *OrderUpsertOne {
	oc.conflict = append(oc.conflict, sql.ConflictColumns(columns...))
	return &OrderUpsertOne{
		create: oc,
	}
}

type (
	// OrderUpsertOne is the builder for "upsert"-ing
	//  one Order node.
	OrderUpsertOne struct {
		create *OrderCreate
	}

	// OrderUpsert is the "OnConflict" setter.
	OrderUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *OrderUpsert) SetUpdatedAt(v time.Time) *OrderUpsert {
	u.Set(order.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *OrderUpsert) UpdateUpdatedAt() *OrderUpsert {
	u.SetExcluded(order.FieldUpdatedAt)
	return u
}

// SetUserID sets the "user_id" field.
func (u *OrderUpsert) SetUserID(v int64) *OrderUpsert {
	u.Set(order.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *OrderUpsert) UpdateUserID() *OrderUpsert {
	u.SetExcluded(order.FieldUserID)
	return u
}

// SetStatus sets the "status" field.
func (u *OrderUpsert) SetStatus(v order.Status) *OrderUpsert {
	u.Set(order.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *OrderUpsert) UpdateStatus() *OrderUpsert {
	u.SetExcluded(order.FieldStatus)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Order.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *OrderUpsertOne) UpdateNewValues() *OrderUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(order.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Order.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *OrderUpsertOne) Ignore() *OrderUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OrderUpsertOne) DoNothing() *OrderUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OrderCreate.OnConflict
// documentation for more info.
func (u *OrderUpsertOne) Update(set func(*OrderUpsert)) *OrderUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OrderUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *OrderUpsertOne) SetUpdatedAt(v time.Time) *OrderUpsertOne {
	return u.Update(func(s *OrderUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *OrderUpsertOne) UpdateUpdatedAt() *OrderUpsertOne {
	return u.Update(func(s *OrderUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetUserID sets the "user_id" field.
func (u *OrderUpsertOne) SetUserID(v int64) *OrderUpsertOne {
	return u.Update(func(s *OrderUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *OrderUpsertOne) UpdateUserID() *OrderUpsertOne {
	return u.Update(func(s *OrderUpsert) {
		s.UpdateUserID()
	})
}

// SetStatus sets the "status" field.
func (u *OrderUpsertOne) SetStatus(v order.Status) *OrderUpsertOne {
	return u.Update(func(s *OrderUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *OrderUpsertOne) UpdateStatus() *OrderUpsertOne {
	return u.Update(func(s *OrderUpsert) {
		s.UpdateStatus()
	})
}

// Exec executes the query.
func (u *OrderUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for OrderCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OrderUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *OrderUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *OrderUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// OrderCreateBulk is the builder for creating many Order entities in bulk.
type OrderCreateBulk struct {
	config
	err      error
	builders []*OrderCreate
	conflict []sql.ConflictOption
}

// Save creates the Order entities in the database.
func (ocb *OrderCreateBulk) Save(ctx context.Context) ([]*Order, error) {
	if ocb.err != nil {
		return nil, ocb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*Order, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderMutation)
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
					_, err = mutators[i+1].Mutate(root, ocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ocb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
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
		if _, err := mutators[0].Mutate(ctx, ocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ocb *OrderCreateBulk) SaveX(ctx context.Context) []*Order {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ocb *OrderCreateBulk) Exec(ctx context.Context) error {
	_, err := ocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocb *OrderCreateBulk) ExecX(ctx context.Context) {
	if err := ocb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Order.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OrderUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (ocb *OrderCreateBulk) OnConflict(opts ...sql.ConflictOption) *OrderUpsertBulk {
	ocb.conflict = opts
	return &OrderUpsertBulk{
		create: ocb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Order.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ocb *OrderCreateBulk) OnConflictColumns(columns ...string) *OrderUpsertBulk {
	ocb.conflict = append(ocb.conflict, sql.ConflictColumns(columns...))
	return &OrderUpsertBulk{
		create: ocb,
	}
}

// OrderUpsertBulk is the builder for "upsert"-ing
// a bulk of Order nodes.
type OrderUpsertBulk struct {
	create *OrderCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Order.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *OrderUpsertBulk) UpdateNewValues() *OrderUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(order.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Order.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *OrderUpsertBulk) Ignore() *OrderUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OrderUpsertBulk) DoNothing() *OrderUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OrderCreateBulk.OnConflict
// documentation for more info.
func (u *OrderUpsertBulk) Update(set func(*OrderUpsert)) *OrderUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OrderUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *OrderUpsertBulk) SetUpdatedAt(v time.Time) *OrderUpsertBulk {
	return u.Update(func(s *OrderUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *OrderUpsertBulk) UpdateUpdatedAt() *OrderUpsertBulk {
	return u.Update(func(s *OrderUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetUserID sets the "user_id" field.
func (u *OrderUpsertBulk) SetUserID(v int64) *OrderUpsertBulk {
	return u.Update(func(s *OrderUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *OrderUpsertBulk) UpdateUserID() *OrderUpsertBulk {
	return u.Update(func(s *OrderUpsert) {
		s.UpdateUserID()
	})
}

// SetStatus sets the "status" field.
func (u *OrderUpsertBulk) SetStatus(v order.Status) *OrderUpsertBulk {
	return u.Update(func(s *OrderUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *OrderUpsertBulk) UpdateStatus() *OrderUpsertBulk {
	return u.Update(func(s *OrderUpsert) {
		s.UpdateStatus()
	})
}

// Exec executes the query.
func (u *OrderUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the OrderCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for OrderCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OrderUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
