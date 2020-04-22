// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/symphony/graph/ent/equipment"
	"github.com/facebookincubator/symphony/graph/ent/equipmentposition"
	"github.com/facebookincubator/symphony/graph/ent/equipmentpositiondefinition"
)

// EquipmentPosition is the model entity for the EquipmentPosition schema.
type EquipmentPosition struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EquipmentPositionQuery when eager-loading is set.
	Edges                         EquipmentPositionEdges `json:"edges"`
	equipment_positions           *int
	equipment_position_definition *int
}

// EquipmentPositionEdges holds the relations/edges for other nodes in the graph.
type EquipmentPositionEdges struct {
	// Definition holds the value of the definition edge.
	Definition *EquipmentPositionDefinition `gqlgen:"definition"`
	// Parent holds the value of the parent edge.
	Parent *Equipment `gqlgen:"parentEquipment"`
	// Attachment holds the value of the attachment edge.
	Attachment *Equipment `gqlgen:"attachedEquipment"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// DefinitionOrErr returns the Definition value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EquipmentPositionEdges) DefinitionOrErr() (*EquipmentPositionDefinition, error) {
	if e.loadedTypes[0] {
		if e.Definition == nil {
			// The edge definition was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: equipmentpositiondefinition.Label}
		}
		return e.Definition, nil
	}
	return nil, &NotLoadedError{edge: "definition"}
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EquipmentPositionEdges) ParentOrErr() (*Equipment, error) {
	if e.loadedTypes[1] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: equipment.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// AttachmentOrErr returns the Attachment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EquipmentPositionEdges) AttachmentOrErr() (*Equipment, error) {
	if e.loadedTypes[2] {
		if e.Attachment == nil {
			// The edge attachment was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: equipment.Label}
		}
		return e.Attachment, nil
	}
	return nil, &NotLoadedError{edge: "attachment"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EquipmentPosition) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
		&sql.NullTime{},  // create_time
		&sql.NullTime{},  // update_time
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*EquipmentPosition) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // equipment_positions
		&sql.NullInt64{}, // equipment_position_definition
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EquipmentPosition fields.
func (ep *EquipmentPosition) assignValues(values ...interface{}) error {
	if m, n := len(values), len(equipmentposition.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	ep.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field create_time", values[0])
	} else if value.Valid {
		ep.CreateTime = value.Time
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field update_time", values[1])
	} else if value.Valid {
		ep.UpdateTime = value.Time
	}
	values = values[2:]
	if len(values) == len(equipmentposition.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field equipment_positions", value)
		} else if value.Valid {
			ep.equipment_positions = new(int)
			*ep.equipment_positions = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field equipment_position_definition", value)
		} else if value.Valid {
			ep.equipment_position_definition = new(int)
			*ep.equipment_position_definition = int(value.Int64)
		}
	}
	return nil
}

// QueryDefinition queries the definition edge of the EquipmentPosition.
func (ep *EquipmentPosition) QueryDefinition() *EquipmentPositionDefinitionQuery {
	return (&EquipmentPositionClient{config: ep.config}).QueryDefinition(ep)
}

// QueryParent queries the parent edge of the EquipmentPosition.
func (ep *EquipmentPosition) QueryParent() *EquipmentQuery {
	return (&EquipmentPositionClient{config: ep.config}).QueryParent(ep)
}

// QueryAttachment queries the attachment edge of the EquipmentPosition.
func (ep *EquipmentPosition) QueryAttachment() *EquipmentQuery {
	return (&EquipmentPositionClient{config: ep.config}).QueryAttachment(ep)
}

// Update returns a builder for updating this EquipmentPosition.
// Note that, you need to call EquipmentPosition.Unwrap() before calling this method, if this EquipmentPosition
// was returned from a transaction, and the transaction was committed or rolled back.
func (ep *EquipmentPosition) Update() *EquipmentPositionUpdateOne {
	return (&EquipmentPositionClient{config: ep.config}).UpdateOne(ep)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (ep *EquipmentPosition) Unwrap() *EquipmentPosition {
	tx, ok := ep.config.driver.(*txDriver)
	if !ok {
		panic("ent: EquipmentPosition is not a transactional entity")
	}
	ep.config.driver = tx.drv
	return ep
}

// String implements the fmt.Stringer.
func (ep *EquipmentPosition) String() string {
	var builder strings.Builder
	builder.WriteString("EquipmentPosition(")
	builder.WriteString(fmt.Sprintf("id=%v", ep.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(ep.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(ep.UpdateTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// EquipmentPositions is a parsable slice of EquipmentPosition.
type EquipmentPositions []*EquipmentPosition

func (ep EquipmentPositions) config(cfg config) {
	for _i := range ep {
		ep[_i].config = cfg
	}
}
