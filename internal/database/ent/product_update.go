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
	"github.com/blackPavlin/shop/internal/database/ent/cart"
	"github.com/blackPavlin/shop/internal/database/ent/category"
	"github.com/blackPavlin/shop/internal/database/ent/orderproduct"
	"github.com/blackPavlin/shop/internal/database/ent/predicate"
	"github.com/blackPavlin/shop/internal/database/ent/product"
	"github.com/blackPavlin/shop/internal/database/ent/productimage"
)

// ProductUpdate is the builder for updating Product entities.
type ProductUpdate struct {
	config
	hooks    []Hook
	mutation *ProductMutation
}

// Where appends a list predicates to the ProductUpdate builder.
func (pu *ProductUpdate) Where(ps ...predicate.Product) *ProductUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *ProductUpdate) SetUpdatedAt(t time.Time) *ProductUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetCategoryID sets the "category_id" field.
func (pu *ProductUpdate) SetCategoryID(i int64) *ProductUpdate {
	pu.mutation.SetCategoryID(i)
	return pu
}

// SetNillableCategoryID sets the "category_id" field if the given value is not nil.
func (pu *ProductUpdate) SetNillableCategoryID(i *int64) *ProductUpdate {
	if i != nil {
		pu.SetCategoryID(*i)
	}
	return pu
}

// SetName sets the "name" field.
func (pu *ProductUpdate) SetName(s string) *ProductUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pu *ProductUpdate) SetNillableName(s *string) *ProductUpdate {
	if s != nil {
		pu.SetName(*s)
	}
	return pu
}

// SetDescription sets the "description" field.
func (pu *ProductUpdate) SetDescription(s string) *ProductUpdate {
	pu.mutation.SetDescription(s)
	return pu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pu *ProductUpdate) SetNillableDescription(s *string) *ProductUpdate {
	if s != nil {
		pu.SetDescription(*s)
	}
	return pu
}

// ClearDescription clears the value of the "description" field.
func (pu *ProductUpdate) ClearDescription() *ProductUpdate {
	pu.mutation.ClearDescription()
	return pu
}

// SetAmount sets the "amount" field.
func (pu *ProductUpdate) SetAmount(i int64) *ProductUpdate {
	pu.mutation.ResetAmount()
	pu.mutation.SetAmount(i)
	return pu
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (pu *ProductUpdate) SetNillableAmount(i *int64) *ProductUpdate {
	if i != nil {
		pu.SetAmount(*i)
	}
	return pu
}

// AddAmount adds i to the "amount" field.
func (pu *ProductUpdate) AddAmount(i int64) *ProductUpdate {
	pu.mutation.AddAmount(i)
	return pu
}

// SetPrice sets the "price" field.
func (pu *ProductUpdate) SetPrice(i int64) *ProductUpdate {
	pu.mutation.ResetPrice()
	pu.mutation.SetPrice(i)
	return pu
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (pu *ProductUpdate) SetNillablePrice(i *int64) *ProductUpdate {
	if i != nil {
		pu.SetPrice(*i)
	}
	return pu
}

// AddPrice adds i to the "price" field.
func (pu *ProductUpdate) AddPrice(i int64) *ProductUpdate {
	pu.mutation.AddPrice(i)
	return pu
}

// SetCategoriesID sets the "categories" edge to the Category entity by ID.
func (pu *ProductUpdate) SetCategoriesID(id int64) *ProductUpdate {
	pu.mutation.SetCategoriesID(id)
	return pu
}

// SetCategories sets the "categories" edge to the Category entity.
func (pu *ProductUpdate) SetCategories(c *Category) *ProductUpdate {
	return pu.SetCategoriesID(c.ID)
}

// AddCartIDs adds the "carts" edge to the Cart entity by IDs.
func (pu *ProductUpdate) AddCartIDs(ids ...int64) *ProductUpdate {
	pu.mutation.AddCartIDs(ids...)
	return pu
}

// AddCarts adds the "carts" edges to the Cart entity.
func (pu *ProductUpdate) AddCarts(c ...*Cart) *ProductUpdate {
	ids := make([]int64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pu.AddCartIDs(ids...)
}

// AddProductImageIDs adds the "product_images" edge to the ProductImage entity by IDs.
func (pu *ProductUpdate) AddProductImageIDs(ids ...int64) *ProductUpdate {
	pu.mutation.AddProductImageIDs(ids...)
	return pu
}

// AddProductImages adds the "product_images" edges to the ProductImage entity.
func (pu *ProductUpdate) AddProductImages(p ...*ProductImage) *ProductUpdate {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddProductImageIDs(ids...)
}

// AddOrderProductIDs adds the "order_products" edge to the OrderProduct entity by IDs.
func (pu *ProductUpdate) AddOrderProductIDs(ids ...int64) *ProductUpdate {
	pu.mutation.AddOrderProductIDs(ids...)
	return pu
}

// AddOrderProducts adds the "order_products" edges to the OrderProduct entity.
func (pu *ProductUpdate) AddOrderProducts(o ...*OrderProduct) *ProductUpdate {
	ids := make([]int64, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return pu.AddOrderProductIDs(ids...)
}

// Mutation returns the ProductMutation object of the builder.
func (pu *ProductUpdate) Mutation() *ProductMutation {
	return pu.mutation
}

// ClearCategories clears the "categories" edge to the Category entity.
func (pu *ProductUpdate) ClearCategories() *ProductUpdate {
	pu.mutation.ClearCategories()
	return pu
}

// ClearCarts clears all "carts" edges to the Cart entity.
func (pu *ProductUpdate) ClearCarts() *ProductUpdate {
	pu.mutation.ClearCarts()
	return pu
}

// RemoveCartIDs removes the "carts" edge to Cart entities by IDs.
func (pu *ProductUpdate) RemoveCartIDs(ids ...int64) *ProductUpdate {
	pu.mutation.RemoveCartIDs(ids...)
	return pu
}

// RemoveCarts removes "carts" edges to Cart entities.
func (pu *ProductUpdate) RemoveCarts(c ...*Cart) *ProductUpdate {
	ids := make([]int64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pu.RemoveCartIDs(ids...)
}

// ClearProductImages clears all "product_images" edges to the ProductImage entity.
func (pu *ProductUpdate) ClearProductImages() *ProductUpdate {
	pu.mutation.ClearProductImages()
	return pu
}

// RemoveProductImageIDs removes the "product_images" edge to ProductImage entities by IDs.
func (pu *ProductUpdate) RemoveProductImageIDs(ids ...int64) *ProductUpdate {
	pu.mutation.RemoveProductImageIDs(ids...)
	return pu
}

// RemoveProductImages removes "product_images" edges to ProductImage entities.
func (pu *ProductUpdate) RemoveProductImages(p ...*ProductImage) *ProductUpdate {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemoveProductImageIDs(ids...)
}

// ClearOrderProducts clears all "order_products" edges to the OrderProduct entity.
func (pu *ProductUpdate) ClearOrderProducts() *ProductUpdate {
	pu.mutation.ClearOrderProducts()
	return pu
}

// RemoveOrderProductIDs removes the "order_products" edge to OrderProduct entities by IDs.
func (pu *ProductUpdate) RemoveOrderProductIDs(ids ...int64) *ProductUpdate {
	pu.mutation.RemoveOrderProductIDs(ids...)
	return pu
}

// RemoveOrderProducts removes "order_products" edges to OrderProduct entities.
func (pu *ProductUpdate) RemoveOrderProducts(o ...*OrderProduct) *ProductUpdate {
	ids := make([]int64, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return pu.RemoveOrderProductIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProductUpdate) Save(ctx context.Context) (int, error) {
	pu.defaults()
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProductUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProductUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProductUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *ProductUpdate) defaults() {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		v := product.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *ProductUpdate) check() error {
	if v, ok := pu.mutation.Name(); ok {
		if err := product.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Product.name": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Amount(); ok {
		if err := product.AmountValidator(v); err != nil {
			return &ValidationError{Name: "amount", err: fmt.Errorf(`ent: validator failed for field "Product.amount": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Price(); ok {
		if err := product.PriceValidator(v); err != nil {
			return &ValidationError{Name: "price", err: fmt.Errorf(`ent: validator failed for field "Product.price": %w`, err)}
		}
	}
	if pu.mutation.CategoriesCleared() && len(pu.mutation.CategoriesIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Product.categories"`)
	}
	return nil
}

func (pu *ProductUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(product.Table, product.Columns, sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt64))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(product.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(product.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.Description(); ok {
		_spec.SetField(product.FieldDescription, field.TypeString, value)
	}
	if pu.mutation.DescriptionCleared() {
		_spec.ClearField(product.FieldDescription, field.TypeString)
	}
	if value, ok := pu.mutation.Amount(); ok {
		_spec.SetField(product.FieldAmount, field.TypeInt64, value)
	}
	if value, ok := pu.mutation.AddedAmount(); ok {
		_spec.AddField(product.FieldAmount, field.TypeInt64, value)
	}
	if value, ok := pu.mutation.Price(); ok {
		_spec.SetField(product.FieldPrice, field.TypeInt64, value)
	}
	if value, ok := pu.mutation.AddedPrice(); ok {
		_spec.AddField(product.FieldPrice, field.TypeInt64, value)
	}
	if pu.mutation.CategoriesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.CategoriesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.CartsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedCartsIDs(); len(nodes) > 0 && !pu.mutation.CartsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.CartsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.ProductImagesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedProductImagesIDs(); len(nodes) > 0 && !pu.mutation.ProductImagesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ProductImagesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.OrderProductsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedOrderProductsIDs(); len(nodes) > 0 && !pu.mutation.OrderProductsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.OrderProductsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{product.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// ProductUpdateOne is the builder for updating a single Product entity.
type ProductUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProductMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *ProductUpdateOne) SetUpdatedAt(t time.Time) *ProductUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetCategoryID sets the "category_id" field.
func (puo *ProductUpdateOne) SetCategoryID(i int64) *ProductUpdateOne {
	puo.mutation.SetCategoryID(i)
	return puo
}

// SetNillableCategoryID sets the "category_id" field if the given value is not nil.
func (puo *ProductUpdateOne) SetNillableCategoryID(i *int64) *ProductUpdateOne {
	if i != nil {
		puo.SetCategoryID(*i)
	}
	return puo
}

// SetName sets the "name" field.
func (puo *ProductUpdateOne) SetName(s string) *ProductUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (puo *ProductUpdateOne) SetNillableName(s *string) *ProductUpdateOne {
	if s != nil {
		puo.SetName(*s)
	}
	return puo
}

// SetDescription sets the "description" field.
func (puo *ProductUpdateOne) SetDescription(s string) *ProductUpdateOne {
	puo.mutation.SetDescription(s)
	return puo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (puo *ProductUpdateOne) SetNillableDescription(s *string) *ProductUpdateOne {
	if s != nil {
		puo.SetDescription(*s)
	}
	return puo
}

// ClearDescription clears the value of the "description" field.
func (puo *ProductUpdateOne) ClearDescription() *ProductUpdateOne {
	puo.mutation.ClearDescription()
	return puo
}

// SetAmount sets the "amount" field.
func (puo *ProductUpdateOne) SetAmount(i int64) *ProductUpdateOne {
	puo.mutation.ResetAmount()
	puo.mutation.SetAmount(i)
	return puo
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (puo *ProductUpdateOne) SetNillableAmount(i *int64) *ProductUpdateOne {
	if i != nil {
		puo.SetAmount(*i)
	}
	return puo
}

// AddAmount adds i to the "amount" field.
func (puo *ProductUpdateOne) AddAmount(i int64) *ProductUpdateOne {
	puo.mutation.AddAmount(i)
	return puo
}

// SetPrice sets the "price" field.
func (puo *ProductUpdateOne) SetPrice(i int64) *ProductUpdateOne {
	puo.mutation.ResetPrice()
	puo.mutation.SetPrice(i)
	return puo
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (puo *ProductUpdateOne) SetNillablePrice(i *int64) *ProductUpdateOne {
	if i != nil {
		puo.SetPrice(*i)
	}
	return puo
}

// AddPrice adds i to the "price" field.
func (puo *ProductUpdateOne) AddPrice(i int64) *ProductUpdateOne {
	puo.mutation.AddPrice(i)
	return puo
}

// SetCategoriesID sets the "categories" edge to the Category entity by ID.
func (puo *ProductUpdateOne) SetCategoriesID(id int64) *ProductUpdateOne {
	puo.mutation.SetCategoriesID(id)
	return puo
}

// SetCategories sets the "categories" edge to the Category entity.
func (puo *ProductUpdateOne) SetCategories(c *Category) *ProductUpdateOne {
	return puo.SetCategoriesID(c.ID)
}

// AddCartIDs adds the "carts" edge to the Cart entity by IDs.
func (puo *ProductUpdateOne) AddCartIDs(ids ...int64) *ProductUpdateOne {
	puo.mutation.AddCartIDs(ids...)
	return puo
}

// AddCarts adds the "carts" edges to the Cart entity.
func (puo *ProductUpdateOne) AddCarts(c ...*Cart) *ProductUpdateOne {
	ids := make([]int64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return puo.AddCartIDs(ids...)
}

// AddProductImageIDs adds the "product_images" edge to the ProductImage entity by IDs.
func (puo *ProductUpdateOne) AddProductImageIDs(ids ...int64) *ProductUpdateOne {
	puo.mutation.AddProductImageIDs(ids...)
	return puo
}

// AddProductImages adds the "product_images" edges to the ProductImage entity.
func (puo *ProductUpdateOne) AddProductImages(p ...*ProductImage) *ProductUpdateOne {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddProductImageIDs(ids...)
}

// AddOrderProductIDs adds the "order_products" edge to the OrderProduct entity by IDs.
func (puo *ProductUpdateOne) AddOrderProductIDs(ids ...int64) *ProductUpdateOne {
	puo.mutation.AddOrderProductIDs(ids...)
	return puo
}

// AddOrderProducts adds the "order_products" edges to the OrderProduct entity.
func (puo *ProductUpdateOne) AddOrderProducts(o ...*OrderProduct) *ProductUpdateOne {
	ids := make([]int64, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return puo.AddOrderProductIDs(ids...)
}

// Mutation returns the ProductMutation object of the builder.
func (puo *ProductUpdateOne) Mutation() *ProductMutation {
	return puo.mutation
}

// ClearCategories clears the "categories" edge to the Category entity.
func (puo *ProductUpdateOne) ClearCategories() *ProductUpdateOne {
	puo.mutation.ClearCategories()
	return puo
}

// ClearCarts clears all "carts" edges to the Cart entity.
func (puo *ProductUpdateOne) ClearCarts() *ProductUpdateOne {
	puo.mutation.ClearCarts()
	return puo
}

// RemoveCartIDs removes the "carts" edge to Cart entities by IDs.
func (puo *ProductUpdateOne) RemoveCartIDs(ids ...int64) *ProductUpdateOne {
	puo.mutation.RemoveCartIDs(ids...)
	return puo
}

// RemoveCarts removes "carts" edges to Cart entities.
func (puo *ProductUpdateOne) RemoveCarts(c ...*Cart) *ProductUpdateOne {
	ids := make([]int64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return puo.RemoveCartIDs(ids...)
}

// ClearProductImages clears all "product_images" edges to the ProductImage entity.
func (puo *ProductUpdateOne) ClearProductImages() *ProductUpdateOne {
	puo.mutation.ClearProductImages()
	return puo
}

// RemoveProductImageIDs removes the "product_images" edge to ProductImage entities by IDs.
func (puo *ProductUpdateOne) RemoveProductImageIDs(ids ...int64) *ProductUpdateOne {
	puo.mutation.RemoveProductImageIDs(ids...)
	return puo
}

// RemoveProductImages removes "product_images" edges to ProductImage entities.
func (puo *ProductUpdateOne) RemoveProductImages(p ...*ProductImage) *ProductUpdateOne {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemoveProductImageIDs(ids...)
}

// ClearOrderProducts clears all "order_products" edges to the OrderProduct entity.
func (puo *ProductUpdateOne) ClearOrderProducts() *ProductUpdateOne {
	puo.mutation.ClearOrderProducts()
	return puo
}

// RemoveOrderProductIDs removes the "order_products" edge to OrderProduct entities by IDs.
func (puo *ProductUpdateOne) RemoveOrderProductIDs(ids ...int64) *ProductUpdateOne {
	puo.mutation.RemoveOrderProductIDs(ids...)
	return puo
}

// RemoveOrderProducts removes "order_products" edges to OrderProduct entities.
func (puo *ProductUpdateOne) RemoveOrderProducts(o ...*OrderProduct) *ProductUpdateOne {
	ids := make([]int64, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return puo.RemoveOrderProductIDs(ids...)
}

// Where appends a list predicates to the ProductUpdate builder.
func (puo *ProductUpdateOne) Where(ps ...predicate.Product) *ProductUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProductUpdateOne) Select(field string, fields ...string) *ProductUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Product entity.
func (puo *ProductUpdateOne) Save(ctx context.Context) (*Product, error) {
	puo.defaults()
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProductUpdateOne) SaveX(ctx context.Context) *Product {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProductUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProductUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *ProductUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		v := product.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *ProductUpdateOne) check() error {
	if v, ok := puo.mutation.Name(); ok {
		if err := product.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Product.name": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Amount(); ok {
		if err := product.AmountValidator(v); err != nil {
			return &ValidationError{Name: "amount", err: fmt.Errorf(`ent: validator failed for field "Product.amount": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Price(); ok {
		if err := product.PriceValidator(v); err != nil {
			return &ValidationError{Name: "price", err: fmt.Errorf(`ent: validator failed for field "Product.price": %w`, err)}
		}
	}
	if puo.mutation.CategoriesCleared() && len(puo.mutation.CategoriesIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Product.categories"`)
	}
	return nil
}

func (puo *ProductUpdateOne) sqlSave(ctx context.Context) (_node *Product, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(product.Table, product.Columns, sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt64))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Product.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, product.FieldID)
		for _, f := range fields {
			if !product.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != product.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(product.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(product.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.Description(); ok {
		_spec.SetField(product.FieldDescription, field.TypeString, value)
	}
	if puo.mutation.DescriptionCleared() {
		_spec.ClearField(product.FieldDescription, field.TypeString)
	}
	if value, ok := puo.mutation.Amount(); ok {
		_spec.SetField(product.FieldAmount, field.TypeInt64, value)
	}
	if value, ok := puo.mutation.AddedAmount(); ok {
		_spec.AddField(product.FieldAmount, field.TypeInt64, value)
	}
	if value, ok := puo.mutation.Price(); ok {
		_spec.SetField(product.FieldPrice, field.TypeInt64, value)
	}
	if value, ok := puo.mutation.AddedPrice(); ok {
		_spec.AddField(product.FieldPrice, field.TypeInt64, value)
	}
	if puo.mutation.CategoriesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.CategoriesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.CartsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedCartsIDs(); len(nodes) > 0 && !puo.mutation.CartsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.CartsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.ProductImagesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedProductImagesIDs(); len(nodes) > 0 && !puo.mutation.ProductImagesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ProductImagesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.OrderProductsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedOrderProductsIDs(); len(nodes) > 0 && !puo.mutation.OrderProductsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.OrderProductsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Product{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{product.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
