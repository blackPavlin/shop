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
	"github.com/blackPavlin/shop/ent/predicate"
	"github.com/blackPavlin/shop/ent/product"
	"github.com/blackPavlin/shop/ent/user"
)

// CartUpdate is the builder for updating Cart entities.
type CartUpdate struct {
	config
	hooks    []Hook
	mutation *CartMutation
}

// Where appends a list predicates to the CartUpdate builder.
func (cu *CartUpdate) Where(ps ...predicate.Cart) *CartUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CartUpdate) SetUpdatedAt(t time.Time) *CartUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetUserID sets the "user_id" field.
func (cu *CartUpdate) SetUserID(i int64) *CartUpdate {
	cu.mutation.SetUserID(i)
	return cu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (cu *CartUpdate) SetNillableUserID(i *int64) *CartUpdate {
	if i != nil {
		cu.SetUserID(*i)
	}
	return cu
}

// SetProductID sets the "product_id" field.
func (cu *CartUpdate) SetProductID(i int64) *CartUpdate {
	cu.mutation.SetProductID(i)
	return cu
}

// SetNillableProductID sets the "product_id" field if the given value is not nil.
func (cu *CartUpdate) SetNillableProductID(i *int64) *CartUpdate {
	if i != nil {
		cu.SetProductID(*i)
	}
	return cu
}

// SetAmount sets the "amount" field.
func (cu *CartUpdate) SetAmount(i int64) *CartUpdate {
	cu.mutation.ResetAmount()
	cu.mutation.SetAmount(i)
	return cu
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (cu *CartUpdate) SetNillableAmount(i *int64) *CartUpdate {
	if i != nil {
		cu.SetAmount(*i)
	}
	return cu
}

// AddAmount adds i to the "amount" field.
func (cu *CartUpdate) AddAmount(i int64) *CartUpdate {
	cu.mutation.AddAmount(i)
	return cu
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (cu *CartUpdate) SetUsersID(id int64) *CartUpdate {
	cu.mutation.SetUsersID(id)
	return cu
}

// SetUsers sets the "users" edge to the User entity.
func (cu *CartUpdate) SetUsers(u *User) *CartUpdate {
	return cu.SetUsersID(u.ID)
}

// SetProductsID sets the "products" edge to the Product entity by ID.
func (cu *CartUpdate) SetProductsID(id int64) *CartUpdate {
	cu.mutation.SetProductsID(id)
	return cu
}

// SetProducts sets the "products" edge to the Product entity.
func (cu *CartUpdate) SetProducts(p *Product) *CartUpdate {
	return cu.SetProductsID(p.ID)
}

// Mutation returns the CartMutation object of the builder.
func (cu *CartUpdate) Mutation() *CartMutation {
	return cu.mutation
}

// ClearUsers clears the "users" edge to the User entity.
func (cu *CartUpdate) ClearUsers() *CartUpdate {
	cu.mutation.ClearUsers()
	return cu
}

// ClearProducts clears the "products" edge to the Product entity.
func (cu *CartUpdate) ClearProducts() *CartUpdate {
	cu.mutation.ClearProducts()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CartUpdate) Save(ctx context.Context) (int, error) {
	cu.defaults()
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CartUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CartUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CartUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CartUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := cart.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CartUpdate) check() error {
	if v, ok := cu.mutation.Amount(); ok {
		if err := cart.AmountValidator(v); err != nil {
			return &ValidationError{Name: "amount", err: fmt.Errorf(`ent: validator failed for field "Cart.amount": %w`, err)}
		}
	}
	if _, ok := cu.mutation.UsersID(); cu.mutation.UsersCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Cart.users"`)
	}
	if _, ok := cu.mutation.ProductsID(); cu.mutation.ProductsCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Cart.products"`)
	}
	return nil
}

func (cu *CartUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(cart.Table, cart.Columns, sqlgraph.NewFieldSpec(cart.FieldID, field.TypeInt64))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(cart.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.Amount(); ok {
		_spec.SetField(cart.FieldAmount, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.AddedAmount(); ok {
		_spec.AddField(cart.FieldAmount, field.TypeInt64, value)
	}
	if cu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cart.UsersTable,
			Columns: []string{cart.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cart.UsersTable,
			Columns: []string{cart.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cart.ProductsTable,
			Columns: []string{cart.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cart.ProductsTable,
			Columns: []string{cart.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cart.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CartUpdateOne is the builder for updating a single Cart entity.
type CartUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CartMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CartUpdateOne) SetUpdatedAt(t time.Time) *CartUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetUserID sets the "user_id" field.
func (cuo *CartUpdateOne) SetUserID(i int64) *CartUpdateOne {
	cuo.mutation.SetUserID(i)
	return cuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (cuo *CartUpdateOne) SetNillableUserID(i *int64) *CartUpdateOne {
	if i != nil {
		cuo.SetUserID(*i)
	}
	return cuo
}

// SetProductID sets the "product_id" field.
func (cuo *CartUpdateOne) SetProductID(i int64) *CartUpdateOne {
	cuo.mutation.SetProductID(i)
	return cuo
}

// SetNillableProductID sets the "product_id" field if the given value is not nil.
func (cuo *CartUpdateOne) SetNillableProductID(i *int64) *CartUpdateOne {
	if i != nil {
		cuo.SetProductID(*i)
	}
	return cuo
}

// SetAmount sets the "amount" field.
func (cuo *CartUpdateOne) SetAmount(i int64) *CartUpdateOne {
	cuo.mutation.ResetAmount()
	cuo.mutation.SetAmount(i)
	return cuo
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (cuo *CartUpdateOne) SetNillableAmount(i *int64) *CartUpdateOne {
	if i != nil {
		cuo.SetAmount(*i)
	}
	return cuo
}

// AddAmount adds i to the "amount" field.
func (cuo *CartUpdateOne) AddAmount(i int64) *CartUpdateOne {
	cuo.mutation.AddAmount(i)
	return cuo
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (cuo *CartUpdateOne) SetUsersID(id int64) *CartUpdateOne {
	cuo.mutation.SetUsersID(id)
	return cuo
}

// SetUsers sets the "users" edge to the User entity.
func (cuo *CartUpdateOne) SetUsers(u *User) *CartUpdateOne {
	return cuo.SetUsersID(u.ID)
}

// SetProductsID sets the "products" edge to the Product entity by ID.
func (cuo *CartUpdateOne) SetProductsID(id int64) *CartUpdateOne {
	cuo.mutation.SetProductsID(id)
	return cuo
}

// SetProducts sets the "products" edge to the Product entity.
func (cuo *CartUpdateOne) SetProducts(p *Product) *CartUpdateOne {
	return cuo.SetProductsID(p.ID)
}

// Mutation returns the CartMutation object of the builder.
func (cuo *CartUpdateOne) Mutation() *CartMutation {
	return cuo.mutation
}

// ClearUsers clears the "users" edge to the User entity.
func (cuo *CartUpdateOne) ClearUsers() *CartUpdateOne {
	cuo.mutation.ClearUsers()
	return cuo
}

// ClearProducts clears the "products" edge to the Product entity.
func (cuo *CartUpdateOne) ClearProducts() *CartUpdateOne {
	cuo.mutation.ClearProducts()
	return cuo
}

// Where appends a list predicates to the CartUpdate builder.
func (cuo *CartUpdateOne) Where(ps ...predicate.Cart) *CartUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CartUpdateOne) Select(field string, fields ...string) *CartUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Cart entity.
func (cuo *CartUpdateOne) Save(ctx context.Context) (*Cart, error) {
	cuo.defaults()
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CartUpdateOne) SaveX(ctx context.Context) *Cart {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CartUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CartUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CartUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := cart.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CartUpdateOne) check() error {
	if v, ok := cuo.mutation.Amount(); ok {
		if err := cart.AmountValidator(v); err != nil {
			return &ValidationError{Name: "amount", err: fmt.Errorf(`ent: validator failed for field "Cart.amount": %w`, err)}
		}
	}
	if _, ok := cuo.mutation.UsersID(); cuo.mutation.UsersCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Cart.users"`)
	}
	if _, ok := cuo.mutation.ProductsID(); cuo.mutation.ProductsCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Cart.products"`)
	}
	return nil
}

func (cuo *CartUpdateOne) sqlSave(ctx context.Context) (_node *Cart, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(cart.Table, cart.Columns, sqlgraph.NewFieldSpec(cart.FieldID, field.TypeInt64))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Cart.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cart.FieldID)
		for _, f := range fields {
			if !cart.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != cart.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(cart.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.Amount(); ok {
		_spec.SetField(cart.FieldAmount, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.AddedAmount(); ok {
		_spec.AddField(cart.FieldAmount, field.TypeInt64, value)
	}
	if cuo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cart.UsersTable,
			Columns: []string{cart.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cart.UsersTable,
			Columns: []string{cart.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cart.ProductsTable,
			Columns: []string{cart.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cart.ProductsTable,
			Columns: []string{cart.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Cart{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cart.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
