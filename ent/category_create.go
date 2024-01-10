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
	"github.com/blackPavlin/shop/ent/category"
	"github.com/blackPavlin/shop/ent/product"
)

// CategoryCreate is the builder for creating a Category entity.
type CategoryCreate struct {
	config
	mutation *CategoryMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (cc *CategoryCreate) SetCreatedAt(t time.Time) *CategoryCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CategoryCreate) SetNillableCreatedAt(t *time.Time) *CategoryCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CategoryCreate) SetUpdatedAt(t time.Time) *CategoryCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CategoryCreate) SetNillableUpdatedAt(t *time.Time) *CategoryCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetName sets the "name" field.
func (cc *CategoryCreate) SetName(s string) *CategoryCreate {
	cc.mutation.SetName(s)
	return cc
}

// AddProductIDs adds the "products" edge to the Product entity by IDs.
func (cc *CategoryCreate) AddProductIDs(ids ...int64) *CategoryCreate {
	cc.mutation.AddProductIDs(ids...)
	return cc
}

// AddProducts adds the "products" edges to the Product entity.
func (cc *CategoryCreate) AddProducts(p ...*Product) *CategoryCreate {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cc.AddProductIDs(ids...)
}

// Mutation returns the CategoryMutation object of the builder.
func (cc *CategoryCreate) Mutation() *CategoryMutation {
	return cc.mutation
}

// Save creates the Category in the database.
func (cc *CategoryCreate) Save(ctx context.Context) (*Category, error) {
	cc.defaults()
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CategoryCreate) SaveX(ctx context.Context) *Category {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CategoryCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CategoryCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CategoryCreate) defaults() {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := category.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := category.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CategoryCreate) check() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Category.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Category.updated_at"`)}
	}
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Category.name"`)}
	}
	if v, ok := cc.mutation.Name(); ok {
		if err := category.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Category.name": %w`, err)}
		}
	}
	return nil
}

func (cc *CategoryCreate) sqlSave(ctx context.Context) (*Category, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int64(id)
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CategoryCreate) createSpec() (*Category, *sqlgraph.CreateSpec) {
	var (
		_node = &Category{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(category.Table, sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = cc.conflict
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(category.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(category.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(category.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if nodes := cc.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.ProductsTable,
			Columns: []string{category.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt64),
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
//	client.Category.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CategoryUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (cc *CategoryCreate) OnConflict(opts ...sql.ConflictOption) *CategoryUpsertOne {
	cc.conflict = opts
	return &CategoryUpsertOne{
		create: cc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Category.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cc *CategoryCreate) OnConflictColumns(columns ...string) *CategoryUpsertOne {
	cc.conflict = append(cc.conflict, sql.ConflictColumns(columns...))
	return &CategoryUpsertOne{
		create: cc,
	}
}

type (
	// CategoryUpsertOne is the builder for "upsert"-ing
	//  one Category node.
	CategoryUpsertOne struct {
		create *CategoryCreate
	}

	// CategoryUpsert is the "OnConflict" setter.
	CategoryUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *CategoryUpsert) SetUpdatedAt(v time.Time) *CategoryUpsert {
	u.Set(category.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CategoryUpsert) UpdateUpdatedAt() *CategoryUpsert {
	u.SetExcluded(category.FieldUpdatedAt)
	return u
}

// SetName sets the "name" field.
func (u *CategoryUpsert) SetName(v string) *CategoryUpsert {
	u.Set(category.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CategoryUpsert) UpdateName() *CategoryUpsert {
	u.SetExcluded(category.FieldName)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Category.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *CategoryUpsertOne) UpdateNewValues() *CategoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(category.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Category.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *CategoryUpsertOne) Ignore() *CategoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CategoryUpsertOne) DoNothing() *CategoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CategoryCreate.OnConflict
// documentation for more info.
func (u *CategoryUpsertOne) Update(set func(*CategoryUpsert)) *CategoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CategoryUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CategoryUpsertOne) SetUpdatedAt(v time.Time) *CategoryUpsertOne {
	return u.Update(func(s *CategoryUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CategoryUpsertOne) UpdateUpdatedAt() *CategoryUpsertOne {
	return u.Update(func(s *CategoryUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetName sets the "name" field.
func (u *CategoryUpsertOne) SetName(v string) *CategoryUpsertOne {
	return u.Update(func(s *CategoryUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CategoryUpsertOne) UpdateName() *CategoryUpsertOne {
	return u.Update(func(s *CategoryUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *CategoryUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CategoryCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CategoryUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CategoryUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CategoryUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CategoryCreateBulk is the builder for creating many Category entities in bulk.
type CategoryCreateBulk struct {
	config
	err      error
	builders []*CategoryCreate
	conflict []sql.ConflictOption
}

// Save creates the Category entities in the database.
func (ccb *CategoryCreateBulk) Save(ctx context.Context) ([]*Category, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Category, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CategoryMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CategoryCreateBulk) SaveX(ctx context.Context) []*Category {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CategoryCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CategoryCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Category.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CategoryUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (ccb *CategoryCreateBulk) OnConflict(opts ...sql.ConflictOption) *CategoryUpsertBulk {
	ccb.conflict = opts
	return &CategoryUpsertBulk{
		create: ccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Category.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ccb *CategoryCreateBulk) OnConflictColumns(columns ...string) *CategoryUpsertBulk {
	ccb.conflict = append(ccb.conflict, sql.ConflictColumns(columns...))
	return &CategoryUpsertBulk{
		create: ccb,
	}
}

// CategoryUpsertBulk is the builder for "upsert"-ing
// a bulk of Category nodes.
type CategoryUpsertBulk struct {
	create *CategoryCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Category.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *CategoryUpsertBulk) UpdateNewValues() *CategoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(category.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Category.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *CategoryUpsertBulk) Ignore() *CategoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CategoryUpsertBulk) DoNothing() *CategoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CategoryCreateBulk.OnConflict
// documentation for more info.
func (u *CategoryUpsertBulk) Update(set func(*CategoryUpsert)) *CategoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CategoryUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CategoryUpsertBulk) SetUpdatedAt(v time.Time) *CategoryUpsertBulk {
	return u.Update(func(s *CategoryUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CategoryUpsertBulk) UpdateUpdatedAt() *CategoryUpsertBulk {
	return u.Update(func(s *CategoryUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetName sets the "name" field.
func (u *CategoryUpsertBulk) SetName(v string) *CategoryUpsertBulk {
	return u.Update(func(s *CategoryUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CategoryUpsertBulk) UpdateName() *CategoryUpsertBulk {
	return u.Update(func(s *CategoryUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *CategoryUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the CategoryCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CategoryCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CategoryUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
