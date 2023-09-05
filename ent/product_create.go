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
	"github.com/blackPavlin/shop/ent/cart"
	"github.com/blackPavlin/shop/ent/category"
	"github.com/blackPavlin/shop/ent/orderproduct"
	"github.com/blackPavlin/shop/ent/product"
	"github.com/blackPavlin/shop/ent/productimage"
)

// ProductCreate is the builder for creating a Product entity.
type ProductCreate struct {
	config
	mutation *ProductMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (pc *ProductCreate) SetCreatedAt(t time.Time) *ProductCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *ProductCreate) SetNillableCreatedAt(t *time.Time) *ProductCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *ProductCreate) SetUpdatedAt(t time.Time) *ProductCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *ProductCreate) SetNillableUpdatedAt(t *time.Time) *ProductCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetCategoryID sets the "category_id" field.
func (pc *ProductCreate) SetCategoryID(i int64) *ProductCreate {
	pc.mutation.SetCategoryID(i)
	return pc
}

// SetName sets the "name" field.
func (pc *ProductCreate) SetName(s string) *ProductCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetDescription sets the "description" field.
func (pc *ProductCreate) SetDescription(s string) *ProductCreate {
	pc.mutation.SetDescription(s)
	return pc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pc *ProductCreate) SetNillableDescription(s *string) *ProductCreate {
	if s != nil {
		pc.SetDescription(*s)
	}
	return pc
}

// SetAmount sets the "amount" field.
func (pc *ProductCreate) SetAmount(i int64) *ProductCreate {
	pc.mutation.SetAmount(i)
	return pc
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (pc *ProductCreate) SetNillableAmount(i *int64) *ProductCreate {
	if i != nil {
		pc.SetAmount(*i)
	}
	return pc
}

// SetPrice sets the "price" field.
func (pc *ProductCreate) SetPrice(i int64) *ProductCreate {
	pc.mutation.SetPrice(i)
	return pc
}

// SetCategoriesID sets the "categories" edge to the Category entity by ID.
func (pc *ProductCreate) SetCategoriesID(id int64) *ProductCreate {
	pc.mutation.SetCategoriesID(id)
	return pc
}

// SetCategories sets the "categories" edge to the Category entity.
func (pc *ProductCreate) SetCategories(c *Category) *ProductCreate {
	return pc.SetCategoriesID(c.ID)
}

// AddCartIDs adds the "carts" edge to the Cart entity by IDs.
func (pc *ProductCreate) AddCartIDs(ids ...int64) *ProductCreate {
	pc.mutation.AddCartIDs(ids...)
	return pc
}

// AddCarts adds the "carts" edges to the Cart entity.
func (pc *ProductCreate) AddCarts(c ...*Cart) *ProductCreate {
	ids := make([]int64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pc.AddCartIDs(ids...)
}

// AddProductImageIDs adds the "product_images" edge to the ProductImage entity by IDs.
func (pc *ProductCreate) AddProductImageIDs(ids ...int64) *ProductCreate {
	pc.mutation.AddProductImageIDs(ids...)
	return pc
}

// AddProductImages adds the "product_images" edges to the ProductImage entity.
func (pc *ProductCreate) AddProductImages(p ...*ProductImage) *ProductCreate {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddProductImageIDs(ids...)
}

// AddOrderProductIDs adds the "order_products" edge to the OrderProduct entity by IDs.
func (pc *ProductCreate) AddOrderProductIDs(ids ...int64) *ProductCreate {
	pc.mutation.AddOrderProductIDs(ids...)
	return pc
}

// AddOrderProducts adds the "order_products" edges to the OrderProduct entity.
func (pc *ProductCreate) AddOrderProducts(o ...*OrderProduct) *ProductCreate {
	ids := make([]int64, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return pc.AddOrderProductIDs(ids...)
}

// Mutation returns the ProductMutation object of the builder.
func (pc *ProductCreate) Mutation() *ProductMutation {
	return pc.mutation
}

// Save creates the Product in the database.
func (pc *ProductCreate) Save(ctx context.Context) (*Product, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProductCreate) SaveX(ctx context.Context) *Product {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ProductCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ProductCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *ProductCreate) defaults() {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := product.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := product.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pc.mutation.Amount(); !ok {
		v := product.DefaultAmount
		pc.mutation.SetAmount(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProductCreate) check() error {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Product.created_at"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Product.updated_at"`)}
	}
	if _, ok := pc.mutation.CategoryID(); !ok {
		return &ValidationError{Name: "category_id", err: errors.New(`ent: missing required field "Product.category_id"`)}
	}
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Product.name"`)}
	}
	if v, ok := pc.mutation.Name(); ok {
		if err := product.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Product.name": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "Product.amount"`)}
	}
	if v, ok := pc.mutation.Amount(); ok {
		if err := product.AmountValidator(v); err != nil {
			return &ValidationError{Name: "amount", err: fmt.Errorf(`ent: validator failed for field "Product.amount": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`ent: missing required field "Product.price"`)}
	}
	if v, ok := pc.mutation.Price(); ok {
		if err := product.PriceValidator(v); err != nil {
			return &ValidationError{Name: "price", err: fmt.Errorf(`ent: validator failed for field "Product.price": %w`, err)}
		}
	}
	if _, ok := pc.mutation.CategoriesID(); !ok {
		return &ValidationError{Name: "categories", err: errors.New(`ent: missing required edge "Product.categories"`)}
	}
	return nil
}

func (pc *ProductCreate) sqlSave(ctx context.Context) (*Product, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int64(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *ProductCreate) createSpec() (*Product, *sqlgraph.CreateSpec) {
	var (
		_node = &Product{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(product.Table, sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = pc.conflict
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(product.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(product.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(product.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.Description(); ok {
		_spec.SetField(product.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := pc.mutation.Amount(); ok {
		_spec.SetField(product.FieldAmount, field.TypeInt64, value)
		_node.Amount = value
	}
	if value, ok := pc.mutation.Price(); ok {
		_spec.SetField(product.FieldPrice, field.TypeInt64, value)
		_node.Price = value
	}
	if nodes := pc.mutation.CategoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.CategoriesTable,
			Columns: []string{product.CategoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CategoryID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.CartsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.CartsTable,
			Columns: []string{product.CartsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cart.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ProductImagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.ProductImagesTable,
			Columns: []string{product.ProductImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(productimage.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.OrderProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.OrderProductsTable,
			Columns: []string{product.OrderProductsColumn},
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
//	client.Product.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ProductUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (pc *ProductCreate) OnConflict(opts ...sql.ConflictOption) *ProductUpsertOne {
	pc.conflict = opts
	return &ProductUpsertOne{
		create: pc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Product.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pc *ProductCreate) OnConflictColumns(columns ...string) *ProductUpsertOne {
	pc.conflict = append(pc.conflict, sql.ConflictColumns(columns...))
	return &ProductUpsertOne{
		create: pc,
	}
}

type (
	// ProductUpsertOne is the builder for "upsert"-ing
	//  one Product node.
	ProductUpsertOne struct {
		create *ProductCreate
	}

	// ProductUpsert is the "OnConflict" setter.
	ProductUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *ProductUpsert) SetUpdatedAt(v time.Time) *ProductUpsert {
	u.Set(product.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ProductUpsert) UpdateUpdatedAt() *ProductUpsert {
	u.SetExcluded(product.FieldUpdatedAt)
	return u
}

// SetCategoryID sets the "category_id" field.
func (u *ProductUpsert) SetCategoryID(v int64) *ProductUpsert {
	u.Set(product.FieldCategoryID, v)
	return u
}

// UpdateCategoryID sets the "category_id" field to the value that was provided on create.
func (u *ProductUpsert) UpdateCategoryID() *ProductUpsert {
	u.SetExcluded(product.FieldCategoryID)
	return u
}

// SetName sets the "name" field.
func (u *ProductUpsert) SetName(v string) *ProductUpsert {
	u.Set(product.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ProductUpsert) UpdateName() *ProductUpsert {
	u.SetExcluded(product.FieldName)
	return u
}

// SetDescription sets the "description" field.
func (u *ProductUpsert) SetDescription(v string) *ProductUpsert {
	u.Set(product.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *ProductUpsert) UpdateDescription() *ProductUpsert {
	u.SetExcluded(product.FieldDescription)
	return u
}

// ClearDescription clears the value of the "description" field.
func (u *ProductUpsert) ClearDescription() *ProductUpsert {
	u.SetNull(product.FieldDescription)
	return u
}

// SetAmount sets the "amount" field.
func (u *ProductUpsert) SetAmount(v int64) *ProductUpsert {
	u.Set(product.FieldAmount, v)
	return u
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *ProductUpsert) UpdateAmount() *ProductUpsert {
	u.SetExcluded(product.FieldAmount)
	return u
}

// AddAmount adds v to the "amount" field.
func (u *ProductUpsert) AddAmount(v int64) *ProductUpsert {
	u.Add(product.FieldAmount, v)
	return u
}

// SetPrice sets the "price" field.
func (u *ProductUpsert) SetPrice(v int64) *ProductUpsert {
	u.Set(product.FieldPrice, v)
	return u
}

// UpdatePrice sets the "price" field to the value that was provided on create.
func (u *ProductUpsert) UpdatePrice() *ProductUpsert {
	u.SetExcluded(product.FieldPrice)
	return u
}

// AddPrice adds v to the "price" field.
func (u *ProductUpsert) AddPrice(v int64) *ProductUpsert {
	u.Add(product.FieldPrice, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Product.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ProductUpsertOne) UpdateNewValues() *ProductUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(product.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Product.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ProductUpsertOne) Ignore() *ProductUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ProductUpsertOne) DoNothing() *ProductUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ProductCreate.OnConflict
// documentation for more info.
func (u *ProductUpsertOne) Update(set func(*ProductUpsert)) *ProductUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ProductUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ProductUpsertOne) SetUpdatedAt(v time.Time) *ProductUpsertOne {
	return u.Update(func(s *ProductUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ProductUpsertOne) UpdateUpdatedAt() *ProductUpsertOne {
	return u.Update(func(s *ProductUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetCategoryID sets the "category_id" field.
func (u *ProductUpsertOne) SetCategoryID(v int64) *ProductUpsertOne {
	return u.Update(func(s *ProductUpsert) {
		s.SetCategoryID(v)
	})
}

// UpdateCategoryID sets the "category_id" field to the value that was provided on create.
func (u *ProductUpsertOne) UpdateCategoryID() *ProductUpsertOne {
	return u.Update(func(s *ProductUpsert) {
		s.UpdateCategoryID()
	})
}

// SetName sets the "name" field.
func (u *ProductUpsertOne) SetName(v string) *ProductUpsertOne {
	return u.Update(func(s *ProductUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ProductUpsertOne) UpdateName() *ProductUpsertOne {
	return u.Update(func(s *ProductUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *ProductUpsertOne) SetDescription(v string) *ProductUpsertOne {
	return u.Update(func(s *ProductUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *ProductUpsertOne) UpdateDescription() *ProductUpsertOne {
	return u.Update(func(s *ProductUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *ProductUpsertOne) ClearDescription() *ProductUpsertOne {
	return u.Update(func(s *ProductUpsert) {
		s.ClearDescription()
	})
}

// SetAmount sets the "amount" field.
func (u *ProductUpsertOne) SetAmount(v int64) *ProductUpsertOne {
	return u.Update(func(s *ProductUpsert) {
		s.SetAmount(v)
	})
}

// AddAmount adds v to the "amount" field.
func (u *ProductUpsertOne) AddAmount(v int64) *ProductUpsertOne {
	return u.Update(func(s *ProductUpsert) {
		s.AddAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *ProductUpsertOne) UpdateAmount() *ProductUpsertOne {
	return u.Update(func(s *ProductUpsert) {
		s.UpdateAmount()
	})
}

// SetPrice sets the "price" field.
func (u *ProductUpsertOne) SetPrice(v int64) *ProductUpsertOne {
	return u.Update(func(s *ProductUpsert) {
		s.SetPrice(v)
	})
}

// AddPrice adds v to the "price" field.
func (u *ProductUpsertOne) AddPrice(v int64) *ProductUpsertOne {
	return u.Update(func(s *ProductUpsert) {
		s.AddPrice(v)
	})
}

// UpdatePrice sets the "price" field to the value that was provided on create.
func (u *ProductUpsertOne) UpdatePrice() *ProductUpsertOne {
	return u.Update(func(s *ProductUpsert) {
		s.UpdatePrice()
	})
}

// Exec executes the query.
func (u *ProductUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ProductCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ProductUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ProductUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ProductUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ProductCreateBulk is the builder for creating many Product entities in bulk.
type ProductCreateBulk struct {
	config
	builders []*ProductCreate
	conflict []sql.ConflictOption
}

// Save creates the Product entities in the database.
func (pcb *ProductCreateBulk) Save(ctx context.Context) ([]*Product, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Product, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProductMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = pcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProductCreateBulk) SaveX(ctx context.Context) []*Product {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ProductCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ProductCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Product.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ProductUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (pcb *ProductCreateBulk) OnConflict(opts ...sql.ConflictOption) *ProductUpsertBulk {
	pcb.conflict = opts
	return &ProductUpsertBulk{
		create: pcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Product.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pcb *ProductCreateBulk) OnConflictColumns(columns ...string) *ProductUpsertBulk {
	pcb.conflict = append(pcb.conflict, sql.ConflictColumns(columns...))
	return &ProductUpsertBulk{
		create: pcb,
	}
}

// ProductUpsertBulk is the builder for "upsert"-ing
// a bulk of Product nodes.
type ProductUpsertBulk struct {
	create *ProductCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Product.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ProductUpsertBulk) UpdateNewValues() *ProductUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(product.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Product.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ProductUpsertBulk) Ignore() *ProductUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ProductUpsertBulk) DoNothing() *ProductUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ProductCreateBulk.OnConflict
// documentation for more info.
func (u *ProductUpsertBulk) Update(set func(*ProductUpsert)) *ProductUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ProductUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ProductUpsertBulk) SetUpdatedAt(v time.Time) *ProductUpsertBulk {
	return u.Update(func(s *ProductUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ProductUpsertBulk) UpdateUpdatedAt() *ProductUpsertBulk {
	return u.Update(func(s *ProductUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetCategoryID sets the "category_id" field.
func (u *ProductUpsertBulk) SetCategoryID(v int64) *ProductUpsertBulk {
	return u.Update(func(s *ProductUpsert) {
		s.SetCategoryID(v)
	})
}

// UpdateCategoryID sets the "category_id" field to the value that was provided on create.
func (u *ProductUpsertBulk) UpdateCategoryID() *ProductUpsertBulk {
	return u.Update(func(s *ProductUpsert) {
		s.UpdateCategoryID()
	})
}

// SetName sets the "name" field.
func (u *ProductUpsertBulk) SetName(v string) *ProductUpsertBulk {
	return u.Update(func(s *ProductUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ProductUpsertBulk) UpdateName() *ProductUpsertBulk {
	return u.Update(func(s *ProductUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *ProductUpsertBulk) SetDescription(v string) *ProductUpsertBulk {
	return u.Update(func(s *ProductUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *ProductUpsertBulk) UpdateDescription() *ProductUpsertBulk {
	return u.Update(func(s *ProductUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *ProductUpsertBulk) ClearDescription() *ProductUpsertBulk {
	return u.Update(func(s *ProductUpsert) {
		s.ClearDescription()
	})
}

// SetAmount sets the "amount" field.
func (u *ProductUpsertBulk) SetAmount(v int64) *ProductUpsertBulk {
	return u.Update(func(s *ProductUpsert) {
		s.SetAmount(v)
	})
}

// AddAmount adds v to the "amount" field.
func (u *ProductUpsertBulk) AddAmount(v int64) *ProductUpsertBulk {
	return u.Update(func(s *ProductUpsert) {
		s.AddAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *ProductUpsertBulk) UpdateAmount() *ProductUpsertBulk {
	return u.Update(func(s *ProductUpsert) {
		s.UpdateAmount()
	})
}

// SetPrice sets the "price" field.
func (u *ProductUpsertBulk) SetPrice(v int64) *ProductUpsertBulk {
	return u.Update(func(s *ProductUpsert) {
		s.SetPrice(v)
	})
}

// AddPrice adds v to the "price" field.
func (u *ProductUpsertBulk) AddPrice(v int64) *ProductUpsertBulk {
	return u.Update(func(s *ProductUpsert) {
		s.AddPrice(v)
	})
}

// UpdatePrice sets the "price" field to the value that was provided on create.
func (u *ProductUpsertBulk) UpdatePrice() *ProductUpsertBulk {
	return u.Update(func(s *ProductUpsert) {
		s.UpdatePrice()
	})
}

// Exec executes the query.
func (u *ProductUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ProductCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ProductCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ProductUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
