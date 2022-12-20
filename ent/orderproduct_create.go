// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/blackPavlin/shop/ent/orderproduct"
)

// OrderProductCreate is the builder for creating a OrderProduct entity.
type OrderProductCreate struct {
	config
	mutation *OrderProductMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (opc *OrderProductCreate) SetCreatedAt(t time.Time) *OrderProductCreate {
	opc.mutation.SetCreatedAt(t)
	return opc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (opc *OrderProductCreate) SetNillableCreatedAt(t *time.Time) *OrderProductCreate {
	if t != nil {
		opc.SetCreatedAt(*t)
	}
	return opc
}

// SetUpdatedAt sets the "updated_at" field.
func (opc *OrderProductCreate) SetUpdatedAt(t time.Time) *OrderProductCreate {
	opc.mutation.SetUpdatedAt(t)
	return opc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (opc *OrderProductCreate) SetNillableUpdatedAt(t *time.Time) *OrderProductCreate {
	if t != nil {
		opc.SetUpdatedAt(*t)
	}
	return opc
}

// Mutation returns the OrderProductMutation object of the builder.
func (opc *OrderProductCreate) Mutation() *OrderProductMutation {
	return opc.mutation
}

// Save creates the OrderProduct in the database.
func (opc *OrderProductCreate) Save(ctx context.Context) (*OrderProduct, error) {
	var (
		err  error
		node *OrderProduct
	)
	opc.defaults()
	if len(opc.hooks) == 0 {
		if err = opc.check(); err != nil {
			return nil, err
		}
		node, err = opc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderProductMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = opc.check(); err != nil {
				return nil, err
			}
			opc.mutation = mutation
			if node, err = opc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(opc.hooks) - 1; i >= 0; i-- {
			if opc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = opc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, opc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*OrderProduct)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from OrderProductMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (opc *OrderProductCreate) SaveX(ctx context.Context) *OrderProduct {
	v, err := opc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (opc *OrderProductCreate) Exec(ctx context.Context) error {
	_, err := opc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (opc *OrderProductCreate) ExecX(ctx context.Context) {
	if err := opc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (opc *OrderProductCreate) defaults() {
	if _, ok := opc.mutation.CreatedAt(); !ok {
		v := orderproduct.DefaultCreatedAt()
		opc.mutation.SetCreatedAt(v)
	}
	if _, ok := opc.mutation.UpdatedAt(); !ok {
		v := orderproduct.DefaultUpdatedAt()
		opc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (opc *OrderProductCreate) check() error {
	if _, ok := opc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "OrderProduct.created_at"`)}
	}
	if _, ok := opc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "OrderProduct.updated_at"`)}
	}
	return nil
}

func (opc *OrderProductCreate) sqlSave(ctx context.Context) (*OrderProduct, error) {
	_node, _spec := opc.createSpec()
	if err := sqlgraph.CreateNode(ctx, opc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int64(id)
	return _node, nil
}

func (opc *OrderProductCreate) createSpec() (*OrderProduct, *sqlgraph.CreateSpec) {
	var (
		_node = &OrderProduct{config: opc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: orderproduct.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: orderproduct.FieldID,
			},
		}
	)
	if value, ok := opc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: orderproduct.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := opc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: orderproduct.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// OrderProductCreateBulk is the builder for creating many OrderProduct entities in bulk.
type OrderProductCreateBulk struct {
	config
	builders []*OrderProductCreate
}

// Save creates the OrderProduct entities in the database.
func (opcb *OrderProductCreateBulk) Save(ctx context.Context) ([]*OrderProduct, error) {
	specs := make([]*sqlgraph.CreateSpec, len(opcb.builders))
	nodes := make([]*OrderProduct, len(opcb.builders))
	mutators := make([]Mutator, len(opcb.builders))
	for i := range opcb.builders {
		func(i int, root context.Context) {
			builder := opcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderProductMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, opcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, opcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, opcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (opcb *OrderProductCreateBulk) SaveX(ctx context.Context) []*OrderProduct {
	v, err := opcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (opcb *OrderProductCreateBulk) Exec(ctx context.Context) error {
	_, err := opcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (opcb *OrderProductCreateBulk) ExecX(ctx context.Context) {
	if err := opcb.Exec(ctx); err != nil {
		panic(err)
	}
}
