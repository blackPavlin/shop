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
	"github.com/blackPavlin/shop/ent/image"
	"github.com/blackPavlin/shop/ent/predicate"
	"github.com/blackPavlin/shop/ent/productimage"
)

// ImageUpdate is the builder for updating Image entities.
type ImageUpdate struct {
	config
	hooks    []Hook
	mutation *ImageMutation
}

// Where appends a list predicates to the ImageUpdate builder.
func (iu *ImageUpdate) Where(ps ...predicate.Image) *ImageUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetUpdatedAt sets the "updated_at" field.
func (iu *ImageUpdate) SetUpdatedAt(t time.Time) *ImageUpdate {
	iu.mutation.SetUpdatedAt(t)
	return iu
}

// SetName sets the "name" field.
func (iu *ImageUpdate) SetName(s string) *ImageUpdate {
	iu.mutation.SetName(s)
	return iu
}

// SetOriginalName sets the "original_name" field.
func (iu *ImageUpdate) SetOriginalName(s string) *ImageUpdate {
	iu.mutation.SetOriginalName(s)
	return iu
}

// AddProductImageIDs adds the "product_images" edge to the ProductImage entity by IDs.
func (iu *ImageUpdate) AddProductImageIDs(ids ...int64) *ImageUpdate {
	iu.mutation.AddProductImageIDs(ids...)
	return iu
}

// AddProductImages adds the "product_images" edges to the ProductImage entity.
func (iu *ImageUpdate) AddProductImages(p ...*ProductImage) *ImageUpdate {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return iu.AddProductImageIDs(ids...)
}

// Mutation returns the ImageMutation object of the builder.
func (iu *ImageUpdate) Mutation() *ImageMutation {
	return iu.mutation
}

// ClearProductImages clears all "product_images" edges to the ProductImage entity.
func (iu *ImageUpdate) ClearProductImages() *ImageUpdate {
	iu.mutation.ClearProductImages()
	return iu
}

// RemoveProductImageIDs removes the "product_images" edge to ProductImage entities by IDs.
func (iu *ImageUpdate) RemoveProductImageIDs(ids ...int64) *ImageUpdate {
	iu.mutation.RemoveProductImageIDs(ids...)
	return iu
}

// RemoveProductImages removes "product_images" edges to ProductImage entities.
func (iu *ImageUpdate) RemoveProductImages(p ...*ProductImage) *ImageUpdate {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return iu.RemoveProductImageIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *ImageUpdate) Save(ctx context.Context) (int, error) {
	iu.defaults()
	return withHooks(ctx, iu.sqlSave, iu.mutation, iu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iu *ImageUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *ImageUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *ImageUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iu *ImageUpdate) defaults() {
	if _, ok := iu.mutation.UpdatedAt(); !ok {
		v := image.UpdateDefaultUpdatedAt()
		iu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iu *ImageUpdate) check() error {
	if v, ok := iu.mutation.Name(); ok {
		if err := image.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Image.name": %w`, err)}
		}
	}
	if v, ok := iu.mutation.OriginalName(); ok {
		if err := image.OriginalNameValidator(v); err != nil {
			return &ValidationError{Name: "original_name", err: fmt.Errorf(`ent: validator failed for field "Image.original_name": %w`, err)}
		}
	}
	return nil
}

func (iu *ImageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := iu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(image.Table, image.Columns, sqlgraph.NewFieldSpec(image.FieldID, field.TypeInt64))
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.UpdatedAt(); ok {
		_spec.SetField(image.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := iu.mutation.Name(); ok {
		_spec.SetField(image.FieldName, field.TypeString, value)
	}
	if value, ok := iu.mutation.OriginalName(); ok {
		_spec.SetField(image.FieldOriginalName, field.TypeString, value)
	}
	if iu.mutation.ProductImagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   image.ProductImagesTable,
			Columns: []string{image.ProductImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(productimage.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.RemovedProductImagesIDs(); len(nodes) > 0 && !iu.mutation.ProductImagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   image.ProductImagesTable,
			Columns: []string{image.ProductImagesColumn},
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
	if nodes := iu.mutation.ProductImagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   image.ProductImagesTable,
			Columns: []string{image.ProductImagesColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{image.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	iu.mutation.done = true
	return n, nil
}

// ImageUpdateOne is the builder for updating a single Image entity.
type ImageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ImageMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (iuo *ImageUpdateOne) SetUpdatedAt(t time.Time) *ImageUpdateOne {
	iuo.mutation.SetUpdatedAt(t)
	return iuo
}

// SetName sets the "name" field.
func (iuo *ImageUpdateOne) SetName(s string) *ImageUpdateOne {
	iuo.mutation.SetName(s)
	return iuo
}

// SetOriginalName sets the "original_name" field.
func (iuo *ImageUpdateOne) SetOriginalName(s string) *ImageUpdateOne {
	iuo.mutation.SetOriginalName(s)
	return iuo
}

// AddProductImageIDs adds the "product_images" edge to the ProductImage entity by IDs.
func (iuo *ImageUpdateOne) AddProductImageIDs(ids ...int64) *ImageUpdateOne {
	iuo.mutation.AddProductImageIDs(ids...)
	return iuo
}

// AddProductImages adds the "product_images" edges to the ProductImage entity.
func (iuo *ImageUpdateOne) AddProductImages(p ...*ProductImage) *ImageUpdateOne {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return iuo.AddProductImageIDs(ids...)
}

// Mutation returns the ImageMutation object of the builder.
func (iuo *ImageUpdateOne) Mutation() *ImageMutation {
	return iuo.mutation
}

// ClearProductImages clears all "product_images" edges to the ProductImage entity.
func (iuo *ImageUpdateOne) ClearProductImages() *ImageUpdateOne {
	iuo.mutation.ClearProductImages()
	return iuo
}

// RemoveProductImageIDs removes the "product_images" edge to ProductImage entities by IDs.
func (iuo *ImageUpdateOne) RemoveProductImageIDs(ids ...int64) *ImageUpdateOne {
	iuo.mutation.RemoveProductImageIDs(ids...)
	return iuo
}

// RemoveProductImages removes "product_images" edges to ProductImage entities.
func (iuo *ImageUpdateOne) RemoveProductImages(p ...*ProductImage) *ImageUpdateOne {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return iuo.RemoveProductImageIDs(ids...)
}

// Where appends a list predicates to the ImageUpdate builder.
func (iuo *ImageUpdateOne) Where(ps ...predicate.Image) *ImageUpdateOne {
	iuo.mutation.Where(ps...)
	return iuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *ImageUpdateOne) Select(field string, fields ...string) *ImageUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Image entity.
func (iuo *ImageUpdateOne) Save(ctx context.Context) (*Image, error) {
	iuo.defaults()
	return withHooks(ctx, iuo.sqlSave, iuo.mutation, iuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *ImageUpdateOne) SaveX(ctx context.Context) *Image {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *ImageUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *ImageUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iuo *ImageUpdateOne) defaults() {
	if _, ok := iuo.mutation.UpdatedAt(); !ok {
		v := image.UpdateDefaultUpdatedAt()
		iuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iuo *ImageUpdateOne) check() error {
	if v, ok := iuo.mutation.Name(); ok {
		if err := image.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Image.name": %w`, err)}
		}
	}
	if v, ok := iuo.mutation.OriginalName(); ok {
		if err := image.OriginalNameValidator(v); err != nil {
			return &ValidationError{Name: "original_name", err: fmt.Errorf(`ent: validator failed for field "Image.original_name": %w`, err)}
		}
	}
	return nil
}

func (iuo *ImageUpdateOne) sqlSave(ctx context.Context) (_node *Image, err error) {
	if err := iuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(image.Table, image.Columns, sqlgraph.NewFieldSpec(image.FieldID, field.TypeInt64))
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Image.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, image.FieldID)
		for _, f := range fields {
			if !image.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != image.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.UpdatedAt(); ok {
		_spec.SetField(image.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := iuo.mutation.Name(); ok {
		_spec.SetField(image.FieldName, field.TypeString, value)
	}
	if value, ok := iuo.mutation.OriginalName(); ok {
		_spec.SetField(image.FieldOriginalName, field.TypeString, value)
	}
	if iuo.mutation.ProductImagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   image.ProductImagesTable,
			Columns: []string{image.ProductImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(productimage.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.RemovedProductImagesIDs(); len(nodes) > 0 && !iuo.mutation.ProductImagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   image.ProductImagesTable,
			Columns: []string{image.ProductImagesColumn},
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
	if nodes := iuo.mutation.ProductImagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   image.ProductImagesTable,
			Columns: []string{image.ProductImagesColumn},
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
	_node = &Image{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{image.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	iuo.mutation.done = true
	return _node, nil
}
