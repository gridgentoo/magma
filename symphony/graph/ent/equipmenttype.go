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
	"github.com/facebookincubator/symphony/graph/ent/equipmentcategory"
	"github.com/facebookincubator/symphony/graph/ent/equipmenttype"
)

// EquipmentType is the model entity for the EquipmentType schema.
type EquipmentType struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EquipmentTypeQuery when eager-loading is set.
	Edges                   EquipmentTypeEdges `json:"edges"`
	equipment_type_category *int
}

// EquipmentTypeEdges holds the relations/edges for other nodes in the graph.
type EquipmentTypeEdges struct {
	// PortDefinitions holds the value of the port_definitions edge.
	PortDefinitions []*EquipmentPortDefinition `gqlgen:"portDefinitions"`
	// PositionDefinitions holds the value of the position_definitions edge.
	PositionDefinitions []*EquipmentPositionDefinition `gqlgen:"positionDefinitions"`
	// PropertyTypes holds the value of the property_types edge.
	PropertyTypes []*PropertyType `gqlgen:"propertyTypes"`
	// Equipment holds the value of the equipment edge.
	Equipment []*Equipment `gqlgen:"equipments"`
	// Category holds the value of the category edge.
	Category *EquipmentCategory `gqlgen:"category"`
	// ServiceEndpointDefinitions holds the value of the service_endpoint_definitions edge.
	ServiceEndpointDefinitions []*ServiceEndpointDefinition `gqlgen:"serviceEndpointDefinitions"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// PortDefinitionsOrErr returns the PortDefinitions value or an error if the edge
// was not loaded in eager-loading.
func (e EquipmentTypeEdges) PortDefinitionsOrErr() ([]*EquipmentPortDefinition, error) {
	if e.loadedTypes[0] {
		return e.PortDefinitions, nil
	}
	return nil, &NotLoadedError{edge: "port_definitions"}
}

// PositionDefinitionsOrErr returns the PositionDefinitions value or an error if the edge
// was not loaded in eager-loading.
func (e EquipmentTypeEdges) PositionDefinitionsOrErr() ([]*EquipmentPositionDefinition, error) {
	if e.loadedTypes[1] {
		return e.PositionDefinitions, nil
	}
	return nil, &NotLoadedError{edge: "position_definitions"}
}

// PropertyTypesOrErr returns the PropertyTypes value or an error if the edge
// was not loaded in eager-loading.
func (e EquipmentTypeEdges) PropertyTypesOrErr() ([]*PropertyType, error) {
	if e.loadedTypes[2] {
		return e.PropertyTypes, nil
	}
	return nil, &NotLoadedError{edge: "property_types"}
}

// EquipmentOrErr returns the Equipment value or an error if the edge
// was not loaded in eager-loading.
func (e EquipmentTypeEdges) EquipmentOrErr() ([]*Equipment, error) {
	if e.loadedTypes[3] {
		return e.Equipment, nil
	}
	return nil, &NotLoadedError{edge: "equipment"}
}

// CategoryOrErr returns the Category value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EquipmentTypeEdges) CategoryOrErr() (*EquipmentCategory, error) {
	if e.loadedTypes[4] {
		if e.Category == nil {
			// The edge category was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: equipmentcategory.Label}
		}
		return e.Category, nil
	}
	return nil, &NotLoadedError{edge: "category"}
}

// ServiceEndpointDefinitionsOrErr returns the ServiceEndpointDefinitions value or an error if the edge
// was not loaded in eager-loading.
func (e EquipmentTypeEdges) ServiceEndpointDefinitionsOrErr() ([]*ServiceEndpointDefinition, error) {
	if e.loadedTypes[5] {
		return e.ServiceEndpointDefinitions, nil
	}
	return nil, &NotLoadedError{edge: "service_endpoint_definitions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EquipmentType) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullTime{},   // create_time
		&sql.NullTime{},   // update_time
		&sql.NullString{}, // name
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*EquipmentType) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // equipment_type_category
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EquipmentType fields.
func (et *EquipmentType) assignValues(values ...interface{}) error {
	if m, n := len(values), len(equipmenttype.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	et.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field create_time", values[0])
	} else if value.Valid {
		et.CreateTime = value.Time
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field update_time", values[1])
	} else if value.Valid {
		et.UpdateTime = value.Time
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[2])
	} else if value.Valid {
		et.Name = value.String
	}
	values = values[3:]
	if len(values) == len(equipmenttype.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field equipment_type_category", value)
		} else if value.Valid {
			et.equipment_type_category = new(int)
			*et.equipment_type_category = int(value.Int64)
		}
	}
	return nil
}

// QueryPortDefinitions queries the port_definitions edge of the EquipmentType.
func (et *EquipmentType) QueryPortDefinitions() *EquipmentPortDefinitionQuery {
	return (&EquipmentTypeClient{config: et.config}).QueryPortDefinitions(et)
}

// QueryPositionDefinitions queries the position_definitions edge of the EquipmentType.
func (et *EquipmentType) QueryPositionDefinitions() *EquipmentPositionDefinitionQuery {
	return (&EquipmentTypeClient{config: et.config}).QueryPositionDefinitions(et)
}

// QueryPropertyTypes queries the property_types edge of the EquipmentType.
func (et *EquipmentType) QueryPropertyTypes() *PropertyTypeQuery {
	return (&EquipmentTypeClient{config: et.config}).QueryPropertyTypes(et)
}

// QueryEquipment queries the equipment edge of the EquipmentType.
func (et *EquipmentType) QueryEquipment() *EquipmentQuery {
	return (&EquipmentTypeClient{config: et.config}).QueryEquipment(et)
}

// QueryCategory queries the category edge of the EquipmentType.
func (et *EquipmentType) QueryCategory() *EquipmentCategoryQuery {
	return (&EquipmentTypeClient{config: et.config}).QueryCategory(et)
}

// QueryServiceEndpointDefinitions queries the service_endpoint_definitions edge of the EquipmentType.
func (et *EquipmentType) QueryServiceEndpointDefinitions() *ServiceEndpointDefinitionQuery {
	return (&EquipmentTypeClient{config: et.config}).QueryServiceEndpointDefinitions(et)
}

// Update returns a builder for updating this EquipmentType.
// Note that, you need to call EquipmentType.Unwrap() before calling this method, if this EquipmentType
// was returned from a transaction, and the transaction was committed or rolled back.
func (et *EquipmentType) Update() *EquipmentTypeUpdateOne {
	return (&EquipmentTypeClient{config: et.config}).UpdateOne(et)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (et *EquipmentType) Unwrap() *EquipmentType {
	tx, ok := et.config.driver.(*txDriver)
	if !ok {
		panic("ent: EquipmentType is not a transactional entity")
	}
	et.config.driver = tx.drv
	return et
}

// String implements the fmt.Stringer.
func (et *EquipmentType) String() string {
	var builder strings.Builder
	builder.WriteString("EquipmentType(")
	builder.WriteString(fmt.Sprintf("id=%v", et.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(et.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(et.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", name=")
	builder.WriteString(et.Name)
	builder.WriteByte(')')
	return builder.String()
}

// EquipmentTypes is a parsable slice of EquipmentType.
type EquipmentTypes []*EquipmentType

func (et EquipmentTypes) config(cfg config) {
	for _i := range et {
		et[_i].config = cfg
	}
}
