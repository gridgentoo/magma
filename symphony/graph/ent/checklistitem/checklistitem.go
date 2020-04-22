// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated (@generated) by entc, DO NOT EDIT.

package checklistitem

import (
	"fmt"
)

const (
	// Label holds the string label denoting the checklistitem type in the database.
	Label = "check_list_item"
	// FieldID holds the string denoting the id field in the database.
	FieldID                 = "id"                   // FieldTitle holds the string denoting the title vertex property in the database.
	FieldTitle              = "title"                // FieldType holds the string denoting the type vertex property in the database.
	FieldType               = "type"                 // FieldIndex holds the string denoting the index vertex property in the database.
	FieldIndex              = "index"                // FieldChecked holds the string denoting the checked vertex property in the database.
	FieldChecked            = "checked"              // FieldStringVal holds the string denoting the string_val vertex property in the database.
	FieldStringVal          = "string_val"           // FieldEnumValues holds the string denoting the enum_values vertex property in the database.
	FieldEnumValues         = "enum_values"          // FieldEnumSelectionMode holds the string denoting the enum_selection_mode vertex property in the database.
	FieldEnumSelectionMode  = "enum_selection_mode"  // FieldSelectedEnumValues holds the string denoting the selected_enum_values vertex property in the database.
	FieldSelectedEnumValues = "selected_enum_values" // FieldYesNoVal holds the string denoting the yes_no_val vertex property in the database.
	FieldYesNoVal           = "yes_no_val"           // FieldHelpText holds the string denoting the help_text vertex property in the database.
	FieldHelpText           = "help_text"

	// EdgeFiles holds the string denoting the files edge name in mutations.
	EdgeFiles = "files"
	// EdgeWifiScan holds the string denoting the wifi_scan edge name in mutations.
	EdgeWifiScan = "wifi_scan"
	// EdgeCellScan holds the string denoting the cell_scan edge name in mutations.
	EdgeCellScan = "cell_scan"
	// EdgeWorkOrder holds the string denoting the work_order edge name in mutations.
	EdgeWorkOrder = "work_order"

	// Table holds the table name of the checklistitem in the database.
	Table = "check_list_items"
	// FilesTable is the table the holds the files relation/edge.
	FilesTable = "files"
	// FilesInverseTable is the table name for the File entity.
	// It exists in this package in order to avoid circular dependency with the "file" package.
	FilesInverseTable = "files"
	// FilesColumn is the table column denoting the files relation/edge.
	FilesColumn = "check_list_item_files"
	// WifiScanTable is the table the holds the wifi_scan relation/edge.
	WifiScanTable = "survey_wi_fi_scans"
	// WifiScanInverseTable is the table name for the SurveyWiFiScan entity.
	// It exists in this package in order to avoid circular dependency with the "surveywifiscan" package.
	WifiScanInverseTable = "survey_wi_fi_scans"
	// WifiScanColumn is the table column denoting the wifi_scan relation/edge.
	WifiScanColumn = "survey_wi_fi_scan_checklist_item"
	// CellScanTable is the table the holds the cell_scan relation/edge.
	CellScanTable = "survey_cell_scans"
	// CellScanInverseTable is the table name for the SurveyCellScan entity.
	// It exists in this package in order to avoid circular dependency with the "surveycellscan" package.
	CellScanInverseTable = "survey_cell_scans"
	// CellScanColumn is the table column denoting the cell_scan relation/edge.
	CellScanColumn = "survey_cell_scan_checklist_item"
	// WorkOrderTable is the table the holds the work_order relation/edge.
	WorkOrderTable = "check_list_items"
	// WorkOrderInverseTable is the table name for the WorkOrder entity.
	// It exists in this package in order to avoid circular dependency with the "workorder" package.
	WorkOrderInverseTable = "work_orders"
	// WorkOrderColumn is the table column denoting the work_order relation/edge.
	WorkOrderColumn = "work_order_check_list_items"
)

// Columns holds all SQL columns for checklistitem fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldType,
	FieldIndex,
	FieldChecked,
	FieldStringVal,
	FieldEnumValues,
	FieldEnumSelectionMode,
	FieldSelectedEnumValues,
	FieldYesNoVal,
	FieldHelpText,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the CheckListItem type.
var ForeignKeys = []string{
	"check_list_category_check_list_items",
	"work_order_check_list_items",
}

// YesNoVal defines the type for the yes_no_val enum field.
type YesNoVal string

// YesNoVal values.
const (
	YesNoValYES YesNoVal = "YES"
	YesNoValNO  YesNoVal = "NO"
)

func (s YesNoVal) String() string {
	return string(s)
}

// YesNoValValidator is a validator for the "ynv" field enum values. It is called by the builders before save.
func YesNoValValidator(ynv YesNoVal) error {
	switch ynv {
	case YesNoValYES, YesNoValNO:
		return nil
	default:
		return fmt.Errorf("checklistitem: invalid enum value for yes_no_val field: %q", ynv)
	}
}
