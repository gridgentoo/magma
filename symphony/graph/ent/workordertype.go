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
	"github.com/facebookincubator/symphony/graph/ent/workordertype"
)

// WorkOrderType is the model entity for the WorkOrderType schema.
type WorkOrderType struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WorkOrderTypeQuery when eager-loading is set.
	Edges WorkOrderTypeEdges `json:"edges"`
}

// WorkOrderTypeEdges holds the relations/edges for other nodes in the graph.
type WorkOrderTypeEdges struct {
	// WorkOrders holds the value of the work_orders edge.
	WorkOrders []*WorkOrder
	// PropertyTypes holds the value of the property_types edge.
	PropertyTypes []*PropertyType
	// Definitions holds the value of the definitions edge.
	Definitions []*WorkOrderDefinition
	// CheckListCategories holds the value of the check_list_categories edge.
	CheckListCategories []*CheckListCategory
	// CheckListDefinitions holds the value of the check_list_definitions edge.
	CheckListDefinitions []*CheckListItemDefinition
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// WorkOrdersOrErr returns the WorkOrders value or an error if the edge
// was not loaded in eager-loading.
func (e WorkOrderTypeEdges) WorkOrdersOrErr() ([]*WorkOrder, error) {
	if e.loadedTypes[0] {
		return e.WorkOrders, nil
	}
	return nil, &NotLoadedError{edge: "work_orders"}
}

// PropertyTypesOrErr returns the PropertyTypes value or an error if the edge
// was not loaded in eager-loading.
func (e WorkOrderTypeEdges) PropertyTypesOrErr() ([]*PropertyType, error) {
	if e.loadedTypes[1] {
		return e.PropertyTypes, nil
	}
	return nil, &NotLoadedError{edge: "property_types"}
}

// DefinitionsOrErr returns the Definitions value or an error if the edge
// was not loaded in eager-loading.
func (e WorkOrderTypeEdges) DefinitionsOrErr() ([]*WorkOrderDefinition, error) {
	if e.loadedTypes[2] {
		return e.Definitions, nil
	}
	return nil, &NotLoadedError{edge: "definitions"}
}

// CheckListCategoriesOrErr returns the CheckListCategories value or an error if the edge
// was not loaded in eager-loading.
func (e WorkOrderTypeEdges) CheckListCategoriesOrErr() ([]*CheckListCategory, error) {
	if e.loadedTypes[3] {
		return e.CheckListCategories, nil
	}
	return nil, &NotLoadedError{edge: "check_list_categories"}
}

// CheckListDefinitionsOrErr returns the CheckListDefinitions value or an error if the edge
// was not loaded in eager-loading.
func (e WorkOrderTypeEdges) CheckListDefinitionsOrErr() ([]*CheckListItemDefinition, error) {
	if e.loadedTypes[4] {
		return e.CheckListDefinitions, nil
	}
	return nil, &NotLoadedError{edge: "check_list_definitions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*WorkOrderType) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullTime{},   // create_time
		&sql.NullTime{},   // update_time
		&sql.NullString{}, // name
		&sql.NullString{}, // description
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WorkOrderType fields.
func (wot *WorkOrderType) assignValues(values ...interface{}) error {
	if m, n := len(values), len(workordertype.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	wot.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field create_time", values[0])
	} else if value.Valid {
		wot.CreateTime = value.Time
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field update_time", values[1])
	} else if value.Valid {
		wot.UpdateTime = value.Time
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[2])
	} else if value.Valid {
		wot.Name = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field description", values[3])
	} else if value.Valid {
		wot.Description = value.String
	}
	return nil
}

// QueryWorkOrders queries the work_orders edge of the WorkOrderType.
func (wot *WorkOrderType) QueryWorkOrders() *WorkOrderQuery {
	return (&WorkOrderTypeClient{config: wot.config}).QueryWorkOrders(wot)
}

// QueryPropertyTypes queries the property_types edge of the WorkOrderType.
func (wot *WorkOrderType) QueryPropertyTypes() *PropertyTypeQuery {
	return (&WorkOrderTypeClient{config: wot.config}).QueryPropertyTypes(wot)
}

// QueryDefinitions queries the definitions edge of the WorkOrderType.
func (wot *WorkOrderType) QueryDefinitions() *WorkOrderDefinitionQuery {
	return (&WorkOrderTypeClient{config: wot.config}).QueryDefinitions(wot)
}

// QueryCheckListCategories queries the check_list_categories edge of the WorkOrderType.
func (wot *WorkOrderType) QueryCheckListCategories() *CheckListCategoryQuery {
	return (&WorkOrderTypeClient{config: wot.config}).QueryCheckListCategories(wot)
}

// QueryCheckListDefinitions queries the check_list_definitions edge of the WorkOrderType.
func (wot *WorkOrderType) QueryCheckListDefinitions() *CheckListItemDefinitionQuery {
	return (&WorkOrderTypeClient{config: wot.config}).QueryCheckListDefinitions(wot)
}

// Update returns a builder for updating this WorkOrderType.
// Note that, you need to call WorkOrderType.Unwrap() before calling this method, if this WorkOrderType
// was returned from a transaction, and the transaction was committed or rolled back.
func (wot *WorkOrderType) Update() *WorkOrderTypeUpdateOne {
	return (&WorkOrderTypeClient{config: wot.config}).UpdateOne(wot)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (wot *WorkOrderType) Unwrap() *WorkOrderType {
	tx, ok := wot.config.driver.(*txDriver)
	if !ok {
		panic("ent: WorkOrderType is not a transactional entity")
	}
	wot.config.driver = tx.drv
	return wot
}

// String implements the fmt.Stringer.
func (wot *WorkOrderType) String() string {
	var builder strings.Builder
	builder.WriteString("WorkOrderType(")
	builder.WriteString(fmt.Sprintf("id=%v", wot.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(wot.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(wot.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", name=")
	builder.WriteString(wot.Name)
	builder.WriteString(", description=")
	builder.WriteString(wot.Description)
	builder.WriteByte(')')
	return builder.String()
}

// WorkOrderTypes is a parsable slice of WorkOrderType.
type WorkOrderTypes []*WorkOrderType

func (wot WorkOrderTypes) config(cfg config) {
	for _i := range wot {
		wot[_i].config = cfg
	}
}
