// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/testifysec/archivist/ent/dsse"
	"github.com/testifysec/archivist/ent/payloaddigest"
	"github.com/testifysec/archivist/ent/predicate"
)

// PayloadDigestUpdate is the builder for updating PayloadDigest entities.
type PayloadDigestUpdate struct {
	config
	hooks    []Hook
	mutation *PayloadDigestMutation
}

// Where appends a list predicates to the PayloadDigestUpdate builder.
func (pdu *PayloadDigestUpdate) Where(ps ...predicate.PayloadDigest) *PayloadDigestUpdate {
	pdu.mutation.Where(ps...)
	return pdu
}

// SetAlgorithm sets the "algorithm" field.
func (pdu *PayloadDigestUpdate) SetAlgorithm(s string) *PayloadDigestUpdate {
	pdu.mutation.SetAlgorithm(s)
	return pdu
}

// SetValue sets the "value" field.
func (pdu *PayloadDigestUpdate) SetValue(s string) *PayloadDigestUpdate {
	pdu.mutation.SetValue(s)
	return pdu
}

// SetDsseID sets the "dsse" edge to the Dsse entity by ID.
func (pdu *PayloadDigestUpdate) SetDsseID(id int) *PayloadDigestUpdate {
	pdu.mutation.SetDsseID(id)
	return pdu
}

// SetNillableDsseID sets the "dsse" edge to the Dsse entity by ID if the given value is not nil.
func (pdu *PayloadDigestUpdate) SetNillableDsseID(id *int) *PayloadDigestUpdate {
	if id != nil {
		pdu = pdu.SetDsseID(*id)
	}
	return pdu
}

// SetDsse sets the "dsse" edge to the Dsse entity.
func (pdu *PayloadDigestUpdate) SetDsse(d *Dsse) *PayloadDigestUpdate {
	return pdu.SetDsseID(d.ID)
}

// Mutation returns the PayloadDigestMutation object of the builder.
func (pdu *PayloadDigestUpdate) Mutation() *PayloadDigestMutation {
	return pdu.mutation
}

// ClearDsse clears the "dsse" edge to the Dsse entity.
func (pdu *PayloadDigestUpdate) ClearDsse() *PayloadDigestUpdate {
	pdu.mutation.ClearDsse()
	return pdu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pdu *PayloadDigestUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pdu.hooks) == 0 {
		if err = pdu.check(); err != nil {
			return 0, err
		}
		affected, err = pdu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PayloadDigestMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pdu.check(); err != nil {
				return 0, err
			}
			pdu.mutation = mutation
			affected, err = pdu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pdu.hooks) - 1; i >= 0; i-- {
			if pdu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pdu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pdu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pdu *PayloadDigestUpdate) SaveX(ctx context.Context) int {
	affected, err := pdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pdu *PayloadDigestUpdate) Exec(ctx context.Context) error {
	_, err := pdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pdu *PayloadDigestUpdate) ExecX(ctx context.Context) {
	if err := pdu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pdu *PayloadDigestUpdate) check() error {
	if v, ok := pdu.mutation.Algorithm(); ok {
		if err := payloaddigest.AlgorithmValidator(v); err != nil {
			return &ValidationError{Name: "algorithm", err: fmt.Errorf(`ent: validator failed for field "PayloadDigest.algorithm": %w`, err)}
		}
	}
	if v, ok := pdu.mutation.Value(); ok {
		if err := payloaddigest.ValueValidator(v); err != nil {
			return &ValidationError{Name: "value", err: fmt.Errorf(`ent: validator failed for field "PayloadDigest.value": %w`, err)}
		}
	}
	return nil
}

func (pdu *PayloadDigestUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   payloaddigest.Table,
			Columns: payloaddigest.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: payloaddigest.FieldID,
			},
		},
	}
	if ps := pdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pdu.mutation.Algorithm(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: payloaddigest.FieldAlgorithm,
		})
	}
	if value, ok := pdu.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: payloaddigest.FieldValue,
		})
	}
	if pdu.mutation.DsseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   payloaddigest.DsseTable,
			Columns: []string{payloaddigest.DsseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dsse.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pdu.mutation.DsseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   payloaddigest.DsseTable,
			Columns: []string{payloaddigest.DsseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dsse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{payloaddigest.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// PayloadDigestUpdateOne is the builder for updating a single PayloadDigest entity.
type PayloadDigestUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PayloadDigestMutation
}

// SetAlgorithm sets the "algorithm" field.
func (pduo *PayloadDigestUpdateOne) SetAlgorithm(s string) *PayloadDigestUpdateOne {
	pduo.mutation.SetAlgorithm(s)
	return pduo
}

// SetValue sets the "value" field.
func (pduo *PayloadDigestUpdateOne) SetValue(s string) *PayloadDigestUpdateOne {
	pduo.mutation.SetValue(s)
	return pduo
}

// SetDsseID sets the "dsse" edge to the Dsse entity by ID.
func (pduo *PayloadDigestUpdateOne) SetDsseID(id int) *PayloadDigestUpdateOne {
	pduo.mutation.SetDsseID(id)
	return pduo
}

// SetNillableDsseID sets the "dsse" edge to the Dsse entity by ID if the given value is not nil.
func (pduo *PayloadDigestUpdateOne) SetNillableDsseID(id *int) *PayloadDigestUpdateOne {
	if id != nil {
		pduo = pduo.SetDsseID(*id)
	}
	return pduo
}

// SetDsse sets the "dsse" edge to the Dsse entity.
func (pduo *PayloadDigestUpdateOne) SetDsse(d *Dsse) *PayloadDigestUpdateOne {
	return pduo.SetDsseID(d.ID)
}

// Mutation returns the PayloadDigestMutation object of the builder.
func (pduo *PayloadDigestUpdateOne) Mutation() *PayloadDigestMutation {
	return pduo.mutation
}

// ClearDsse clears the "dsse" edge to the Dsse entity.
func (pduo *PayloadDigestUpdateOne) ClearDsse() *PayloadDigestUpdateOne {
	pduo.mutation.ClearDsse()
	return pduo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pduo *PayloadDigestUpdateOne) Select(field string, fields ...string) *PayloadDigestUpdateOne {
	pduo.fields = append([]string{field}, fields...)
	return pduo
}

// Save executes the query and returns the updated PayloadDigest entity.
func (pduo *PayloadDigestUpdateOne) Save(ctx context.Context) (*PayloadDigest, error) {
	var (
		err  error
		node *PayloadDigest
	)
	if len(pduo.hooks) == 0 {
		if err = pduo.check(); err != nil {
			return nil, err
		}
		node, err = pduo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PayloadDigestMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pduo.check(); err != nil {
				return nil, err
			}
			pduo.mutation = mutation
			node, err = pduo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pduo.hooks) - 1; i >= 0; i-- {
			if pduo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pduo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, pduo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*PayloadDigest)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PayloadDigestMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (pduo *PayloadDigestUpdateOne) SaveX(ctx context.Context) *PayloadDigest {
	node, err := pduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pduo *PayloadDigestUpdateOne) Exec(ctx context.Context) error {
	_, err := pduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pduo *PayloadDigestUpdateOne) ExecX(ctx context.Context) {
	if err := pduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pduo *PayloadDigestUpdateOne) check() error {
	if v, ok := pduo.mutation.Algorithm(); ok {
		if err := payloaddigest.AlgorithmValidator(v); err != nil {
			return &ValidationError{Name: "algorithm", err: fmt.Errorf(`ent: validator failed for field "PayloadDigest.algorithm": %w`, err)}
		}
	}
	if v, ok := pduo.mutation.Value(); ok {
		if err := payloaddigest.ValueValidator(v); err != nil {
			return &ValidationError{Name: "value", err: fmt.Errorf(`ent: validator failed for field "PayloadDigest.value": %w`, err)}
		}
	}
	return nil
}

func (pduo *PayloadDigestUpdateOne) sqlSave(ctx context.Context) (_node *PayloadDigest, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   payloaddigest.Table,
			Columns: payloaddigest.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: payloaddigest.FieldID,
			},
		},
	}
	id, ok := pduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PayloadDigest.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, payloaddigest.FieldID)
		for _, f := range fields {
			if !payloaddigest.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != payloaddigest.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pduo.mutation.Algorithm(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: payloaddigest.FieldAlgorithm,
		})
	}
	if value, ok := pduo.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: payloaddigest.FieldValue,
		})
	}
	if pduo.mutation.DsseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   payloaddigest.DsseTable,
			Columns: []string{payloaddigest.DsseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dsse.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pduo.mutation.DsseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   payloaddigest.DsseTable,
			Columns: []string{payloaddigest.DsseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dsse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PayloadDigest{config: pduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{payloaddigest.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
