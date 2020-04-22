// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/symphony/graph/ent/equipment"
	"github.com/facebookincubator/symphony/graph/ent/equipmentport"
	"github.com/facebookincubator/symphony/graph/ent/predicate"
	"github.com/facebookincubator/symphony/graph/ent/service"
	"github.com/facebookincubator/symphony/graph/ent/serviceendpoint"
	"github.com/facebookincubator/symphony/graph/ent/serviceendpointdefinition"
)

// ServiceEndpointUpdate is the builder for updating ServiceEndpoint entities.
type ServiceEndpointUpdate struct {
	config
	hooks      []Hook
	mutation   *ServiceEndpointMutation
	predicates []predicate.ServiceEndpoint
}

// Where adds a new predicate for the builder.
func (seu *ServiceEndpointUpdate) Where(ps ...predicate.ServiceEndpoint) *ServiceEndpointUpdate {
	seu.predicates = append(seu.predicates, ps...)
	return seu
}

// SetPortID sets the port edge to EquipmentPort by id.
func (seu *ServiceEndpointUpdate) SetPortID(id int) *ServiceEndpointUpdate {
	seu.mutation.SetPortID(id)
	return seu
}

// SetNillablePortID sets the port edge to EquipmentPort by id if the given value is not nil.
func (seu *ServiceEndpointUpdate) SetNillablePortID(id *int) *ServiceEndpointUpdate {
	if id != nil {
		seu = seu.SetPortID(*id)
	}
	return seu
}

// SetPort sets the port edge to EquipmentPort.
func (seu *ServiceEndpointUpdate) SetPort(e *EquipmentPort) *ServiceEndpointUpdate {
	return seu.SetPortID(e.ID)
}

// SetEquipmentID sets the equipment edge to Equipment by id.
func (seu *ServiceEndpointUpdate) SetEquipmentID(id int) *ServiceEndpointUpdate {
	seu.mutation.SetEquipmentID(id)
	return seu
}

// SetEquipment sets the equipment edge to Equipment.
func (seu *ServiceEndpointUpdate) SetEquipment(e *Equipment) *ServiceEndpointUpdate {
	return seu.SetEquipmentID(e.ID)
}

// SetServiceID sets the service edge to Service by id.
func (seu *ServiceEndpointUpdate) SetServiceID(id int) *ServiceEndpointUpdate {
	seu.mutation.SetServiceID(id)
	return seu
}

// SetService sets the service edge to Service.
func (seu *ServiceEndpointUpdate) SetService(s *Service) *ServiceEndpointUpdate {
	return seu.SetServiceID(s.ID)
}

// SetDefinitionID sets the definition edge to ServiceEndpointDefinition by id.
func (seu *ServiceEndpointUpdate) SetDefinitionID(id int) *ServiceEndpointUpdate {
	seu.mutation.SetDefinitionID(id)
	return seu
}

// SetNillableDefinitionID sets the definition edge to ServiceEndpointDefinition by id if the given value is not nil.
func (seu *ServiceEndpointUpdate) SetNillableDefinitionID(id *int) *ServiceEndpointUpdate {
	if id != nil {
		seu = seu.SetDefinitionID(*id)
	}
	return seu
}

// SetDefinition sets the definition edge to ServiceEndpointDefinition.
func (seu *ServiceEndpointUpdate) SetDefinition(s *ServiceEndpointDefinition) *ServiceEndpointUpdate {
	return seu.SetDefinitionID(s.ID)
}

// ClearPort clears the port edge to EquipmentPort.
func (seu *ServiceEndpointUpdate) ClearPort() *ServiceEndpointUpdate {
	seu.mutation.ClearPort()
	return seu
}

// ClearEquipment clears the equipment edge to Equipment.
func (seu *ServiceEndpointUpdate) ClearEquipment() *ServiceEndpointUpdate {
	seu.mutation.ClearEquipment()
	return seu
}

// ClearService clears the service edge to Service.
func (seu *ServiceEndpointUpdate) ClearService() *ServiceEndpointUpdate {
	seu.mutation.ClearService()
	return seu
}

// ClearDefinition clears the definition edge to ServiceEndpointDefinition.
func (seu *ServiceEndpointUpdate) ClearDefinition() *ServiceEndpointUpdate {
	seu.mutation.ClearDefinition()
	return seu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (seu *ServiceEndpointUpdate) Save(ctx context.Context) (int, error) {
	if _, ok := seu.mutation.UpdateTime(); !ok {
		v := serviceendpoint.UpdateDefaultUpdateTime()
		seu.mutation.SetUpdateTime(v)
	}

	if _, ok := seu.mutation.EquipmentID(); seu.mutation.EquipmentCleared() && !ok {
		return 0, errors.New("ent: clearing a unique edge \"equipment\"")
	}

	if _, ok := seu.mutation.ServiceID(); seu.mutation.ServiceCleared() && !ok {
		return 0, errors.New("ent: clearing a unique edge \"service\"")
	}

	var (
		err      error
		affected int
	)
	if len(seu.hooks) == 0 {
		affected, err = seu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ServiceEndpointMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			seu.mutation = mutation
			affected, err = seu.sqlSave(ctx)
			return affected, err
		})
		for i := len(seu.hooks) - 1; i >= 0; i-- {
			mut = seu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, seu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (seu *ServiceEndpointUpdate) SaveX(ctx context.Context) int {
	affected, err := seu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (seu *ServiceEndpointUpdate) Exec(ctx context.Context) error {
	_, err := seu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (seu *ServiceEndpointUpdate) ExecX(ctx context.Context) {
	if err := seu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (seu *ServiceEndpointUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   serviceendpoint.Table,
			Columns: serviceendpoint.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: serviceendpoint.FieldID,
			},
		},
	}
	if ps := seu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := seu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: serviceendpoint.FieldUpdateTime,
		})
	}
	if seu.mutation.PortCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   serviceendpoint.PortTable,
			Columns: []string{serviceendpoint.PortColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: equipmentport.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := seu.mutation.PortIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   serviceendpoint.PortTable,
			Columns: []string{serviceendpoint.PortColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: equipmentport.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if seu.mutation.EquipmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   serviceendpoint.EquipmentTable,
			Columns: []string{serviceendpoint.EquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: equipment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := seu.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   serviceendpoint.EquipmentTable,
			Columns: []string{serviceendpoint.EquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: equipment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if seu.mutation.ServiceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   serviceendpoint.ServiceTable,
			Columns: []string{serviceendpoint.ServiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: service.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := seu.mutation.ServiceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   serviceendpoint.ServiceTable,
			Columns: []string{serviceendpoint.ServiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: service.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if seu.mutation.DefinitionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   serviceendpoint.DefinitionTable,
			Columns: []string{serviceendpoint.DefinitionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: serviceendpointdefinition.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := seu.mutation.DefinitionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   serviceendpoint.DefinitionTable,
			Columns: []string{serviceendpoint.DefinitionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: serviceendpointdefinition.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, seu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{serviceendpoint.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ServiceEndpointUpdateOne is the builder for updating a single ServiceEndpoint entity.
type ServiceEndpointUpdateOne struct {
	config
	hooks    []Hook
	mutation *ServiceEndpointMutation
}

// SetPortID sets the port edge to EquipmentPort by id.
func (seuo *ServiceEndpointUpdateOne) SetPortID(id int) *ServiceEndpointUpdateOne {
	seuo.mutation.SetPortID(id)
	return seuo
}

// SetNillablePortID sets the port edge to EquipmentPort by id if the given value is not nil.
func (seuo *ServiceEndpointUpdateOne) SetNillablePortID(id *int) *ServiceEndpointUpdateOne {
	if id != nil {
		seuo = seuo.SetPortID(*id)
	}
	return seuo
}

// SetPort sets the port edge to EquipmentPort.
func (seuo *ServiceEndpointUpdateOne) SetPort(e *EquipmentPort) *ServiceEndpointUpdateOne {
	return seuo.SetPortID(e.ID)
}

// SetEquipmentID sets the equipment edge to Equipment by id.
func (seuo *ServiceEndpointUpdateOne) SetEquipmentID(id int) *ServiceEndpointUpdateOne {
	seuo.mutation.SetEquipmentID(id)
	return seuo
}

// SetEquipment sets the equipment edge to Equipment.
func (seuo *ServiceEndpointUpdateOne) SetEquipment(e *Equipment) *ServiceEndpointUpdateOne {
	return seuo.SetEquipmentID(e.ID)
}

// SetServiceID sets the service edge to Service by id.
func (seuo *ServiceEndpointUpdateOne) SetServiceID(id int) *ServiceEndpointUpdateOne {
	seuo.mutation.SetServiceID(id)
	return seuo
}

// SetService sets the service edge to Service.
func (seuo *ServiceEndpointUpdateOne) SetService(s *Service) *ServiceEndpointUpdateOne {
	return seuo.SetServiceID(s.ID)
}

// SetDefinitionID sets the definition edge to ServiceEndpointDefinition by id.
func (seuo *ServiceEndpointUpdateOne) SetDefinitionID(id int) *ServiceEndpointUpdateOne {
	seuo.mutation.SetDefinitionID(id)
	return seuo
}

// SetNillableDefinitionID sets the definition edge to ServiceEndpointDefinition by id if the given value is not nil.
func (seuo *ServiceEndpointUpdateOne) SetNillableDefinitionID(id *int) *ServiceEndpointUpdateOne {
	if id != nil {
		seuo = seuo.SetDefinitionID(*id)
	}
	return seuo
}

// SetDefinition sets the definition edge to ServiceEndpointDefinition.
func (seuo *ServiceEndpointUpdateOne) SetDefinition(s *ServiceEndpointDefinition) *ServiceEndpointUpdateOne {
	return seuo.SetDefinitionID(s.ID)
}

// ClearPort clears the port edge to EquipmentPort.
func (seuo *ServiceEndpointUpdateOne) ClearPort() *ServiceEndpointUpdateOne {
	seuo.mutation.ClearPort()
	return seuo
}

// ClearEquipment clears the equipment edge to Equipment.
func (seuo *ServiceEndpointUpdateOne) ClearEquipment() *ServiceEndpointUpdateOne {
	seuo.mutation.ClearEquipment()
	return seuo
}

// ClearService clears the service edge to Service.
func (seuo *ServiceEndpointUpdateOne) ClearService() *ServiceEndpointUpdateOne {
	seuo.mutation.ClearService()
	return seuo
}

// ClearDefinition clears the definition edge to ServiceEndpointDefinition.
func (seuo *ServiceEndpointUpdateOne) ClearDefinition() *ServiceEndpointUpdateOne {
	seuo.mutation.ClearDefinition()
	return seuo
}

// Save executes the query and returns the updated entity.
func (seuo *ServiceEndpointUpdateOne) Save(ctx context.Context) (*ServiceEndpoint, error) {
	if _, ok := seuo.mutation.UpdateTime(); !ok {
		v := serviceendpoint.UpdateDefaultUpdateTime()
		seuo.mutation.SetUpdateTime(v)
	}

	if _, ok := seuo.mutation.EquipmentID(); seuo.mutation.EquipmentCleared() && !ok {
		return nil, errors.New("ent: clearing a unique edge \"equipment\"")
	}

	if _, ok := seuo.mutation.ServiceID(); seuo.mutation.ServiceCleared() && !ok {
		return nil, errors.New("ent: clearing a unique edge \"service\"")
	}

	var (
		err  error
		node *ServiceEndpoint
	)
	if len(seuo.hooks) == 0 {
		node, err = seuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ServiceEndpointMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			seuo.mutation = mutation
			node, err = seuo.sqlSave(ctx)
			return node, err
		})
		for i := len(seuo.hooks) - 1; i >= 0; i-- {
			mut = seuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, seuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (seuo *ServiceEndpointUpdateOne) SaveX(ctx context.Context) *ServiceEndpoint {
	se, err := seuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return se
}

// Exec executes the query on the entity.
func (seuo *ServiceEndpointUpdateOne) Exec(ctx context.Context) error {
	_, err := seuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (seuo *ServiceEndpointUpdateOne) ExecX(ctx context.Context) {
	if err := seuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (seuo *ServiceEndpointUpdateOne) sqlSave(ctx context.Context) (se *ServiceEndpoint, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   serviceendpoint.Table,
			Columns: serviceendpoint.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: serviceendpoint.FieldID,
			},
		},
	}
	id, ok := seuo.mutation.ID()
	if !ok {
		return nil, fmt.Errorf("missing ServiceEndpoint.ID for update")
	}
	_spec.Node.ID.Value = id
	if value, ok := seuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: serviceendpoint.FieldUpdateTime,
		})
	}
	if seuo.mutation.PortCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   serviceendpoint.PortTable,
			Columns: []string{serviceendpoint.PortColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: equipmentport.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := seuo.mutation.PortIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   serviceendpoint.PortTable,
			Columns: []string{serviceendpoint.PortColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: equipmentport.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if seuo.mutation.EquipmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   serviceendpoint.EquipmentTable,
			Columns: []string{serviceendpoint.EquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: equipment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := seuo.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   serviceendpoint.EquipmentTable,
			Columns: []string{serviceendpoint.EquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: equipment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if seuo.mutation.ServiceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   serviceendpoint.ServiceTable,
			Columns: []string{serviceendpoint.ServiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: service.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := seuo.mutation.ServiceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   serviceendpoint.ServiceTable,
			Columns: []string{serviceendpoint.ServiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: service.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if seuo.mutation.DefinitionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   serviceendpoint.DefinitionTable,
			Columns: []string{serviceendpoint.DefinitionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: serviceendpointdefinition.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := seuo.mutation.DefinitionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   serviceendpoint.DefinitionTable,
			Columns: []string{serviceendpoint.DefinitionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: serviceendpointdefinition.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	se = &ServiceEndpoint{config: seuo.config}
	_spec.Assign = se.assignValues
	_spec.ScanValues = se.scanValues()
	if err = sqlgraph.UpdateNode(ctx, seuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{serviceendpoint.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return se, nil
}
