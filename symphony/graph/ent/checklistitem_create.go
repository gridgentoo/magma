// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/symphony/graph/ent/checklistitem"
	"github.com/facebookincubator/symphony/graph/ent/file"
	"github.com/facebookincubator/symphony/graph/ent/surveycellscan"
	"github.com/facebookincubator/symphony/graph/ent/surveywifiscan"
	"github.com/facebookincubator/symphony/graph/ent/workorder"
)

// CheckListItemCreate is the builder for creating a CheckListItem entity.
type CheckListItemCreate struct {
	config
	mutation *CheckListItemMutation
	hooks    []Hook
}

// SetTitle sets the title field.
func (clic *CheckListItemCreate) SetTitle(s string) *CheckListItemCreate {
	clic.mutation.SetTitle(s)
	return clic
}

// SetType sets the type field.
func (clic *CheckListItemCreate) SetType(s string) *CheckListItemCreate {
	clic.mutation.SetType(s)
	return clic
}

// SetIndex sets the index field.
func (clic *CheckListItemCreate) SetIndex(i int) *CheckListItemCreate {
	clic.mutation.SetIndex(i)
	return clic
}

// SetNillableIndex sets the index field if the given value is not nil.
func (clic *CheckListItemCreate) SetNillableIndex(i *int) *CheckListItemCreate {
	if i != nil {
		clic.SetIndex(*i)
	}
	return clic
}

// SetChecked sets the checked field.
func (clic *CheckListItemCreate) SetChecked(b bool) *CheckListItemCreate {
	clic.mutation.SetChecked(b)
	return clic
}

// SetNillableChecked sets the checked field if the given value is not nil.
func (clic *CheckListItemCreate) SetNillableChecked(b *bool) *CheckListItemCreate {
	if b != nil {
		clic.SetChecked(*b)
	}
	return clic
}

// SetStringVal sets the string_val field.
func (clic *CheckListItemCreate) SetStringVal(s string) *CheckListItemCreate {
	clic.mutation.SetStringVal(s)
	return clic
}

// SetNillableStringVal sets the string_val field if the given value is not nil.
func (clic *CheckListItemCreate) SetNillableStringVal(s *string) *CheckListItemCreate {
	if s != nil {
		clic.SetStringVal(*s)
	}
	return clic
}

// SetEnumValues sets the enum_values field.
func (clic *CheckListItemCreate) SetEnumValues(s string) *CheckListItemCreate {
	clic.mutation.SetEnumValues(s)
	return clic
}

// SetNillableEnumValues sets the enum_values field if the given value is not nil.
func (clic *CheckListItemCreate) SetNillableEnumValues(s *string) *CheckListItemCreate {
	if s != nil {
		clic.SetEnumValues(*s)
	}
	return clic
}

// SetEnumSelectionMode sets the enum_selection_mode field.
func (clic *CheckListItemCreate) SetEnumSelectionMode(s string) *CheckListItemCreate {
	clic.mutation.SetEnumSelectionMode(s)
	return clic
}

// SetNillableEnumSelectionMode sets the enum_selection_mode field if the given value is not nil.
func (clic *CheckListItemCreate) SetNillableEnumSelectionMode(s *string) *CheckListItemCreate {
	if s != nil {
		clic.SetEnumSelectionMode(*s)
	}
	return clic
}

// SetSelectedEnumValues sets the selected_enum_values field.
func (clic *CheckListItemCreate) SetSelectedEnumValues(s string) *CheckListItemCreate {
	clic.mutation.SetSelectedEnumValues(s)
	return clic
}

// SetNillableSelectedEnumValues sets the selected_enum_values field if the given value is not nil.
func (clic *CheckListItemCreate) SetNillableSelectedEnumValues(s *string) *CheckListItemCreate {
	if s != nil {
		clic.SetSelectedEnumValues(*s)
	}
	return clic
}

// SetYesNoVal sets the yes_no_val field.
func (clic *CheckListItemCreate) SetYesNoVal(cnv checklistitem.YesNoVal) *CheckListItemCreate {
	clic.mutation.SetYesNoVal(cnv)
	return clic
}

// SetNillableYesNoVal sets the yes_no_val field if the given value is not nil.
func (clic *CheckListItemCreate) SetNillableYesNoVal(cnv *checklistitem.YesNoVal) *CheckListItemCreate {
	if cnv != nil {
		clic.SetYesNoVal(*cnv)
	}
	return clic
}

// SetHelpText sets the help_text field.
func (clic *CheckListItemCreate) SetHelpText(s string) *CheckListItemCreate {
	clic.mutation.SetHelpText(s)
	return clic
}

// SetNillableHelpText sets the help_text field if the given value is not nil.
func (clic *CheckListItemCreate) SetNillableHelpText(s *string) *CheckListItemCreate {
	if s != nil {
		clic.SetHelpText(*s)
	}
	return clic
}

// AddFileIDs adds the files edge to File by ids.
func (clic *CheckListItemCreate) AddFileIDs(ids ...int) *CheckListItemCreate {
	clic.mutation.AddFileIDs(ids...)
	return clic
}

// AddFiles adds the files edges to File.
func (clic *CheckListItemCreate) AddFiles(f ...*File) *CheckListItemCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return clic.AddFileIDs(ids...)
}

// AddWifiScanIDs adds the wifi_scan edge to SurveyWiFiScan by ids.
func (clic *CheckListItemCreate) AddWifiScanIDs(ids ...int) *CheckListItemCreate {
	clic.mutation.AddWifiScanIDs(ids...)
	return clic
}

// AddWifiScan adds the wifi_scan edges to SurveyWiFiScan.
func (clic *CheckListItemCreate) AddWifiScan(s ...*SurveyWiFiScan) *CheckListItemCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return clic.AddWifiScanIDs(ids...)
}

// AddCellScanIDs adds the cell_scan edge to SurveyCellScan by ids.
func (clic *CheckListItemCreate) AddCellScanIDs(ids ...int) *CheckListItemCreate {
	clic.mutation.AddCellScanIDs(ids...)
	return clic
}

// AddCellScan adds the cell_scan edges to SurveyCellScan.
func (clic *CheckListItemCreate) AddCellScan(s ...*SurveyCellScan) *CheckListItemCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return clic.AddCellScanIDs(ids...)
}

// SetWorkOrderID sets the work_order edge to WorkOrder by id.
func (clic *CheckListItemCreate) SetWorkOrderID(id int) *CheckListItemCreate {
	clic.mutation.SetWorkOrderID(id)
	return clic
}

// SetNillableWorkOrderID sets the work_order edge to WorkOrder by id if the given value is not nil.
func (clic *CheckListItemCreate) SetNillableWorkOrderID(id *int) *CheckListItemCreate {
	if id != nil {
		clic = clic.SetWorkOrderID(*id)
	}
	return clic
}

// SetWorkOrder sets the work_order edge to WorkOrder.
func (clic *CheckListItemCreate) SetWorkOrder(w *WorkOrder) *CheckListItemCreate {
	return clic.SetWorkOrderID(w.ID)
}

// Save creates the CheckListItem in the database.
func (clic *CheckListItemCreate) Save(ctx context.Context) (*CheckListItem, error) {
	if _, ok := clic.mutation.Title(); !ok {
		return nil, errors.New("ent: missing required field \"title\"")
	}
	if _, ok := clic.mutation.GetType(); !ok {
		return nil, errors.New("ent: missing required field \"type\"")
	}
	if v, ok := clic.mutation.YesNoVal(); ok {
		if err := checklistitem.YesNoValValidator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"yes_no_val\": %v", err)
		}
	}
	var (
		err  error
		node *CheckListItem
	)
	if len(clic.hooks) == 0 {
		node, err = clic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CheckListItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			clic.mutation = mutation
			node, err = clic.sqlSave(ctx)
			return node, err
		})
		for i := len(clic.hooks) - 1; i >= 0; i-- {
			mut = clic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, clic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (clic *CheckListItemCreate) SaveX(ctx context.Context) *CheckListItem {
	v, err := clic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (clic *CheckListItemCreate) sqlSave(ctx context.Context) (*CheckListItem, error) {
	var (
		cli   = &CheckListItem{config: clic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: checklistitem.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: checklistitem.FieldID,
			},
		}
	)
	if value, ok := clic.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: checklistitem.FieldTitle,
		})
		cli.Title = value
	}
	if value, ok := clic.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: checklistitem.FieldType,
		})
		cli.Type = value
	}
	if value, ok := clic.mutation.Index(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: checklistitem.FieldIndex,
		})
		cli.Index = value
	}
	if value, ok := clic.mutation.Checked(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: checklistitem.FieldChecked,
		})
		cli.Checked = value
	}
	if value, ok := clic.mutation.StringVal(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: checklistitem.FieldStringVal,
		})
		cli.StringVal = value
	}
	if value, ok := clic.mutation.EnumValues(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: checklistitem.FieldEnumValues,
		})
		cli.EnumValues = value
	}
	if value, ok := clic.mutation.EnumSelectionMode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: checklistitem.FieldEnumSelectionMode,
		})
		cli.EnumSelectionMode = value
	}
	if value, ok := clic.mutation.SelectedEnumValues(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: checklistitem.FieldSelectedEnumValues,
		})
		cli.SelectedEnumValues = value
	}
	if value, ok := clic.mutation.YesNoVal(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: checklistitem.FieldYesNoVal,
		})
		cli.YesNoVal = value
	}
	if value, ok := clic.mutation.HelpText(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: checklistitem.FieldHelpText,
		})
		cli.HelpText = &value
	}
	if nodes := clic.mutation.FilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   checklistitem.FilesTable,
			Columns: []string{checklistitem.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: file.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := clic.mutation.WifiScanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   checklistitem.WifiScanTable,
			Columns: []string{checklistitem.WifiScanColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: surveywifiscan.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := clic.mutation.CellScanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   checklistitem.CellScanTable,
			Columns: []string{checklistitem.CellScanColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: surveycellscan.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := clic.mutation.WorkOrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   checklistitem.WorkOrderTable,
			Columns: []string{checklistitem.WorkOrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: workorder.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, clic.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	cli.ID = int(id)
	return cli, nil
}
