// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated (@generated) by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/facebookincubator/symphony/graph/authz/models"
	"github.com/facebookincubator/symphony/graph/ent"
	"github.com/facebookincubator/symphony/graph/ent/user"
	"github.com/facebookincubator/symphony/graph/ent/usersgroup"
	"github.com/facebookincubator/symphony/pkg/actions/core"
)

type NamedNode interface {
	IsNamedNode()
}

type SystemPolicy interface {
	IsSystemPolicy()
}

type ActionsAction struct {
	ActionID    core.ActionID `json:"actionID"`
	Description string        `json:"description"`
	DataType    core.DataType `json:"dataType"`
}

type ActionsFilter struct {
	FilterID           string             `json:"filterID"`
	Description        string             `json:"description"`
	SupportedOperators []*ActionsOperator `json:"supportedOperators"`
}

type ActionsOperator struct {
	OperatorID  string        `json:"operatorID"`
	Description string        `json:"description"`
	DataType    core.DataType `json:"dataType"`
}

type ActionsRuleActionInput struct {
	ActionID core.ActionID `json:"actionID"`
	Data     string        `json:"data"`
}

type ActionsRuleFilterInput struct {
	FilterID   string `json:"filterID"`
	OperatorID string `json:"operatorID"`
	Data       string `json:"data"`
}

type ActionsRulesSearchResult struct {
	Results []*ent.ActionsRule `json:"results"`
	Count   int                `json:"count"`
}

type ActionsTrigger struct {
	ID               int              `json:"id"`
	TriggerID        core.TriggerID   `json:"triggerID"`
	Description      string           `json:"description"`
	SupportedActions []*ActionsAction `json:"supportedActions"`
	SupportedFilters []*ActionsFilter `json:"supportedFilters"`
}

type ActionsTriggersSearchResult struct {
	Results []*ActionsTrigger `json:"results"`
	Count   int               `json:"count"`
}

type AddActionsRuleInput struct {
	Name        string                    `json:"name"`
	TriggerID   core.TriggerID            `json:"triggerID"`
	RuleActions []*ActionsRuleActionInput `json:"ruleActions"`
	RuleFilters []*ActionsRuleFilterInput `json:"ruleFilters"`
}

type AddCustomerInput struct {
	Name       string  `json:"name"`
	ExternalID *string `json:"externalId"`
}

type AddEquipmentInput struct {
	Name               string           `json:"name"`
	Type               int              `json:"type"`
	Location           *int             `json:"location"`
	Parent             *int             `json:"parent"`
	PositionDefinition *int             `json:"positionDefinition"`
	Properties         []*PropertyInput `json:"properties"`
	WorkOrder          *int             `json:"workOrder"`
	ExternalID         *string          `json:"externalId"`
}

type AddEquipmentPortTypeInput struct {
	Name           string               `json:"name"`
	Properties     []*PropertyTypeInput `json:"properties"`
	LinkProperties []*PropertyTypeInput `json:"linkProperties"`
}

type AddEquipmentTypeInput struct {
	Name       string                    `json:"name"`
	Category   *string                   `json:"category"`
	Positions  []*EquipmentPositionInput `json:"positions"`
	Ports      []*EquipmentPortInput     `json:"ports"`
	Properties []*PropertyTypeInput      `json:"properties"`
}

type AddFloorPlanInput struct {
	Name             string         `json:"name"`
	LocationID       int            `json:"locationID"`
	Image            *AddImageInput `json:"image"`
	ReferenceX       int            `json:"referenceX"`
	ReferenceY       int            `json:"referenceY"`
	Latitude         float64        `json:"latitude"`
	Longitude        float64        `json:"longitude"`
	ReferencePoint1x int            `json:"referencePoint1X"`
	ReferencePoint1y int            `json:"referencePoint1Y"`
	ReferencePoint2x int            `json:"referencePoint2X"`
	ReferencePoint2y int            `json:"referencePoint2Y"`
	ScaleInMeters    float64        `json:"scaleInMeters"`
}

type AddHyperlinkInput struct {
	EntityType  ImageEntity `json:"entityType"`
	EntityID    int         `json:"entityId"`
	URL         string      `json:"url"`
	DisplayName *string     `json:"displayName"`
	Category    *string     `json:"category"`
}

type AddImageInput struct {
	EntityType  ImageEntity `json:"entityType"`
	EntityID    int         `json:"entityId"`
	ImgKey      string      `json:"imgKey"`
	FileName    string      `json:"fileName"`
	FileSize    int         `json:"fileSize"`
	Modified    time.Time   `json:"modified"`
	ContentType string      `json:"contentType"`
	Category    *string     `json:"category"`
}

type AddLinkInput struct {
	Sides      []*LinkSide      `json:"sides"`
	WorkOrder  *int             `json:"workOrder"`
	Properties []*PropertyInput `json:"properties"`
	ServiceIds []int            `json:"serviceIds"`
}

type AddLocationInput struct {
	Name       string           `json:"name"`
	Type       int              `json:"type"`
	Parent     *int             `json:"parent"`
	Latitude   *float64         `json:"latitude"`
	Longitude  *float64         `json:"longitude"`
	Properties []*PropertyInput `json:"properties"`
	ExternalID *string          `json:"externalID"`
}

type AddLocationTypeInput struct {
	Name                     string                         `json:"name"`
	MapType                  *string                        `json:"mapType"`
	MapZoomLevel             *int                           `json:"mapZoomLevel"`
	IsSite                   *bool                          `json:"isSite"`
	Properties               []*PropertyTypeInput           `json:"properties"`
	SurveyTemplateCategories []*SurveyTemplateCategoryInput `json:"surveyTemplateCategories"`
}

type AddPermissionsPolicyInput struct {
	Name           string                       `json:"name"`
	Description    *string                      `json:"description"`
	IsGlobal       *bool                        `json:"isGlobal"`
	InventoryInput *models.InventoryPolicyInput `json:"inventoryInput"`
	WorkforceInput *models.WorkforcePolicyInput `json:"workforceInput"`
}

type AddProjectInput struct {
	Name        string           `json:"name"`
	Description *string          `json:"description"`
	Creator     *string          `json:"creator"`
	CreatorID   *int             `json:"creatorId"`
	Type        int              `json:"type"`
	Location    *int             `json:"location"`
	Properties  []*PropertyInput `json:"properties"`
}

type AddProjectTypeInput struct {
	Name        string                      `json:"name"`
	Description *string                     `json:"description"`
	Properties  []*PropertyTypeInput        `json:"properties"`
	WorkOrders  []*WorkOrderDefinitionInput `json:"workOrders"`
}

type AddServiceEndpointInput struct {
	ID          int  `json:"id"`
	PortID      *int `json:"portId"`
	EquipmentID int  `json:"equipmentID"`
	Definition  int  `json:"definition"`
}

type AddUsersGroupInput struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type AddWorkOrderInput struct {
	Name                string                    `json:"name"`
	Description         *string                   `json:"description"`
	WorkOrderTypeID     int                       `json:"workOrderTypeId"`
	LocationID          *int                      `json:"locationId"`
	ProjectID           *int                      `json:"projectId"`
	Properties          []*PropertyInput          `json:"properties"`
	CheckList           []*CheckListItemInput     `json:"checkList"`
	OwnerName           *string                   `json:"ownerName"`
	OwnerID             *int                      `json:"ownerId"`
	CheckListCategories []*CheckListCategoryInput `json:"checkListCategories"`
	Assignee            *string                   `json:"assignee"`
	AssigneeID          *int                      `json:"assigneeId"`
	Index               *int                      `json:"index"`
	Status              *WorkOrderStatus          `json:"status"`
	Priority            *WorkOrderPriority        `json:"priority"`
}

type AddWorkOrderTypeInput struct {
	Name                string                      `json:"name"`
	Description         *string                     `json:"description"`
	Properties          []*PropertyTypeInput        `json:"properties"`
	CheckList           []*CheckListDefinitionInput `json:"checkList"`
	CheckListCategories []*CheckListCategoryInput   `json:"checkListCategories"`
}

type AdministrativePolicy struct {
	Access *BasicPermissionRule `json:"access"`
}

type BasicPermissionRule struct {
	IsAllowed models.PermissionValue `json:"isAllowed"`
}

type Cud struct {
	Create *BasicPermissionRule `json:"create"`
	Update *BasicPermissionRule `json:"update"`
	Delete *BasicPermissionRule `json:"delete"`
}

type CheckListCategoryInput struct {
	ID          *int                  `json:"id"`
	Title       string                `json:"title"`
	Description *string               `json:"description"`
	CheckList   []*CheckListItemInput `json:"checkList"`
}

type CheckListDefinitionInput struct {
	ID         *int              `json:"id"`
	Title      string            `json:"title"`
	Type       CheckListItemType `json:"type"`
	Index      *int              `json:"index"`
	EnumValues *string           `json:"enumValues"`
	HelpText   *string           `json:"helpText"`
}

type CheckListItemInput struct {
	ID                 *int                            `json:"id"`
	Title              string                          `json:"title"`
	Type               CheckListItemType               `json:"type"`
	Index              *int                            `json:"index"`
	HelpText           *string                         `json:"helpText"`
	EnumValues         *string                         `json:"enumValues"`
	EnumSelectionMode  *CheckListItemEnumSelectionMode `json:"enumSelectionMode"`
	SelectedEnumValues *string                         `json:"selectedEnumValues"`
	StringValue        *string                         `json:"stringValue"`
	Checked            *bool                           `json:"checked"`
	Files              []*FileInput                    `json:"files"`
	YesNoResponse      *YesNoResponse                  `json:"yesNoResponse"`
}

type CommentInput struct {
	EntityType CommentEntity `json:"entityType"`
	ID         int           `json:"id"`
	Text       string        `json:"text"`
}

type Device struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Up   *bool  `json:"up"`
}

type EditEquipmentInput struct {
	ID         int              `json:"id"`
	Name       string           `json:"name"`
	Properties []*PropertyInput `json:"properties"`
	DeviceID   *string          `json:"deviceID"`
	ExternalID *string          `json:"externalId"`
}

type EditEquipmentPortInput struct {
	Side       *LinkSide        `json:"side"`
	Properties []*PropertyInput `json:"properties"`
}

type EditEquipmentPortTypeInput struct {
	ID             int                  `json:"id"`
	Name           string               `json:"name"`
	Properties     []*PropertyTypeInput `json:"properties"`
	LinkProperties []*PropertyTypeInput `json:"linkProperties"`
}

type EditEquipmentTypeInput struct {
	ID         int                       `json:"id"`
	Name       string                    `json:"name"`
	Category   *string                   `json:"category"`
	Positions  []*EquipmentPositionInput `json:"positions"`
	Ports      []*EquipmentPortInput     `json:"ports"`
	Properties []*PropertyTypeInput      `json:"properties"`
}

type EditLinkInput struct {
	ID         int              `json:"id"`
	Properties []*PropertyInput `json:"properties"`
	ServiceIds []int            `json:"serviceIds"`
}

type EditLocationInput struct {
	ID         int              `json:"id"`
	Name       string           `json:"name"`
	Latitude   float64          `json:"latitude"`
	Longitude  float64          `json:"longitude"`
	Properties []*PropertyInput `json:"properties"`
	ExternalID *string          `json:"externalID"`
}

type EditLocationTypeInput struct {
	ID           int                  `json:"id"`
	Name         string               `json:"name"`
	MapType      *string              `json:"mapType"`
	MapZoomLevel *int                 `json:"mapZoomLevel"`
	IsSite       *bool                `json:"isSite"`
	Properties   []*PropertyTypeInput `json:"properties"`
}

type EditProjectInput struct {
	ID          int              `json:"id"`
	Name        string           `json:"name"`
	Description *string          `json:"description"`
	Creator     *string          `json:"creator"`
	CreatorID   *int             `json:"creatorId"`
	Type        int              `json:"type"`
	Location    *int             `json:"location"`
	Properties  []*PropertyInput `json:"properties"`
}

type EditProjectTypeInput struct {
	ID          int                         `json:"id"`
	Name        string                      `json:"name"`
	Description *string                     `json:"description"`
	Properties  []*PropertyTypeInput        `json:"properties"`
	WorkOrders  []*WorkOrderDefinitionInput `json:"workOrders"`
}

type EditReportFilterInput struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type EditUserInput struct {
	ID        int          `json:"id"`
	FirstName *string      `json:"firstName"`
	LastName  *string      `json:"lastName"`
	Status    *user.Status `json:"status"`
	Role      *user.Role   `json:"role"`
}

type EditUsersGroupInput struct {
	ID          int                `json:"id"`
	Name        *string            `json:"name"`
	Description *string            `json:"description"`
	Status      *usersgroup.Status `json:"status"`
}

type EditWorkOrderInput struct {
	ID                  int                       `json:"id"`
	Name                string                    `json:"name"`
	Description         *string                   `json:"description"`
	OwnerName           *string                   `json:"ownerName"`
	OwnerID             *int                      `json:"ownerId"`
	InstallDate         *time.Time                `json:"installDate"`
	Assignee            *string                   `json:"assignee"`
	AssigneeID          *int                      `json:"assigneeId"`
	Index               *int                      `json:"index"`
	Status              WorkOrderStatus           `json:"status"`
	Priority            WorkOrderPriority         `json:"priority"`
	ProjectID           *int                      `json:"projectId"`
	Properties          []*PropertyInput          `json:"properties"`
	CheckList           []*CheckListItemInput     `json:"checkList"`
	CheckListCategories []*CheckListCategoryInput `json:"checkListCategories"`
	LocationID          *int                      `json:"locationId"`
}

type EditWorkOrderTypeInput struct {
	ID                  int                         `json:"id"`
	Name                string                      `json:"name"`
	Description         *string                     `json:"description"`
	Properties          []*PropertyTypeInput        `json:"properties"`
	CheckList           []*CheckListDefinitionInput `json:"checkList"`
	CheckListCategories []*CheckListCategoryInput   `json:"checkListCategories"`
}

type EquipmentFilterInput struct {
	FilterType    EquipmentFilterType `json:"filterType"`
	Operator      FilterOperator      `json:"operator"`
	StringValue   *string             `json:"stringValue"`
	PropertyValue *PropertyTypeInput  `json:"propertyValue"`
	IDSet         []int               `json:"idSet"`
	StringSet     []string            `json:"stringSet"`
	MaxDepth      *int                `json:"maxDepth"`
}

type EquipmentPortInput struct {
	ID           *int    `json:"id"`
	Name         string  `json:"name"`
	Index        *int    `json:"index"`
	VisibleLabel *string `json:"visibleLabel"`
	PortTypeID   *int    `json:"portTypeID"`
	Bandwidth    *string `json:"bandwidth"`
}

type EquipmentPositionInput struct {
	ID           *int    `json:"id"`
	Name         string  `json:"name"`
	Index        *int    `json:"index"`
	VisibleLabel *string `json:"visibleLabel"`
}

type EquipmentSearchResult struct {
	Equipment []*ent.Equipment `json:"equipment"`
	Count     int              `json:"count"`
}

type FileInput struct {
	ID               *int      `json:"id"`
	FileName         string    `json:"fileName"`
	SizeInBytes      *int      `json:"sizeInBytes"`
	ModificationTime *int      `json:"modificationTime"`
	UploadTime       *int      `json:"uploadTime"`
	FileType         *FileType `json:"fileType"`
	MimeType         *string   `json:"mimeType"`
	StoreKey         string    `json:"storeKey"`
}

type GeneralFilter struct {
	FilterType    string            `json:"filterType"`
	Key           string            `json:"key"`
	Operator      FilterOperator    `json:"operator"`
	StringValue   *string           `json:"stringValue"`
	IDSet         []int             `json:"idSet"`
	StringSet     []string          `json:"stringSet"`
	BoolValue     *bool             `json:"boolValue"`
	PropertyValue *ent.PropertyType `json:"propertyValue"`
}

type GeneralFilterInput struct {
	FilterType    string             `json:"filterType"`
	Key           string             `json:"key"`
	Operator      FilterOperator     `json:"operator"`
	StringValue   *string            `json:"stringValue"`
	IDSet         []int              `json:"idSet"`
	StringSet     []string           `json:"stringSet"`
	BoolValue     *bool              `json:"boolValue"`
	PropertyValue *PropertyTypeInput `json:"propertyValue"`
}

type InventoryPolicy struct {
	Read          *BasicPermissionRule `json:"read"`
	Location      *Cud                 `json:"location"`
	Equipment     *Cud                 `json:"equipment"`
	EquipmentType *Cud                 `json:"equipmentType"`
	LocationType  *Cud                 `json:"locationType"`
	PortType      *Cud                 `json:"portType"`
	ServiceType   *Cud                 `json:"serviceType"`
}

func (InventoryPolicy) IsSystemPolicy() {}

type LatestPythonPackageResult struct {
	LastPythonPackage         *PythonPackage `json:"lastPythonPackage"`
	LastBreakingPythonPackage *PythonPackage `json:"lastBreakingPythonPackage"`
}

type LinkFilterInput struct {
	FilterType    LinkFilterType     `json:"filterType"`
	Operator      FilterOperator     `json:"operator"`
	StringValue   *string            `json:"stringValue"`
	PropertyValue *PropertyTypeInput `json:"propertyValue"`
	IDSet         []int              `json:"idSet"`
	StringSet     []string           `json:"stringSet"`
	MaxDepth      *int               `json:"maxDepth"`
}

type LinkSearchResult struct {
	Links []*ent.Link `json:"links"`
	Count int         `json:"count"`
}

type LinkSide struct {
	Equipment int `json:"equipment"`
	Port      int `json:"port"`
}

type LocationFilterInput struct {
	FilterType    LocationFilterType `json:"filterType"`
	Operator      FilterOperator     `json:"operator"`
	BoolValue     *bool              `json:"boolValue"`
	StringValue   *string            `json:"stringValue"`
	PropertyValue *PropertyTypeInput `json:"propertyValue"`
	IDSet         []int              `json:"idSet"`
	StringSet     []string           `json:"stringSet"`
	MaxDepth      *int               `json:"maxDepth"`
}

type LocationSearchResult struct {
	Locations []*ent.Location `json:"locations"`
	Count     int             `json:"count"`
}

type LocationTypeIndex struct {
	LocationTypeID int `json:"locationTypeID"`
	Index          int `json:"index"`
}

type NetworkTopology struct {
	Nodes []ent.Noder     `json:"nodes"`
	Links []*TopologyLink `json:"links"`
}

type PermissionSettings struct {
	CanWrite            bool                  `json:"canWrite"`
	AdminPolicy         *AdministrativePolicy `json:"adminPolicy"`
	InventoryPolicy     *InventoryPolicy      `json:"inventoryPolicy"`
	WorkforcePermission *WorkforcePolicy      `json:"workforcePermission"`
}

type PortFilterInput struct {
	FilterType    PortFilterType     `json:"filterType"`
	Operator      FilterOperator     `json:"operator"`
	BoolValue     *bool              `json:"boolValue"`
	StringValue   *string            `json:"stringValue"`
	PropertyValue *PropertyTypeInput `json:"propertyValue"`
	IDSet         []int              `json:"idSet"`
	StringSet     []string           `json:"stringSet"`
	MaxDepth      *int               `json:"maxDepth"`
}

type PortSearchResult struct {
	Ports []*ent.EquipmentPort `json:"ports"`
	Count int                  `json:"count"`
}

type ProjectFilterInput struct {
	FilterType  ProjectFilterType `json:"filterType"`
	Operator    FilterOperator    `json:"operator"`
	StringValue *string           `json:"stringValue"`
}

type PropertyInput struct {
	ID                 *int     `json:"id"`
	PropertyTypeID     int      `json:"propertyTypeID"`
	StringValue        *string  `json:"stringValue"`
	IntValue           *int     `json:"intValue"`
	BooleanValue       *bool    `json:"booleanValue"`
	FloatValue         *float64 `json:"floatValue"`
	LatitudeValue      *float64 `json:"latitudeValue"`
	LongitudeValue     *float64 `json:"longitudeValue"`
	RangeFromValue     *float64 `json:"rangeFromValue"`
	RangeToValue       *float64 `json:"rangeToValue"`
	NodeIDValue        *int     `json:"nodeIDValue"`
	IsEditable         *bool    `json:"isEditable"`
	IsInstanceProperty *bool    `json:"isInstanceProperty"`
}

type PropertyTypeInput struct {
	ID                 *int         `json:"id"`
	ExternalID         *string      `json:"externalId"`
	Name               string       `json:"name"`
	Type               PropertyKind `json:"type"`
	NodeType           *string      `json:"nodeType"`
	Index              *int         `json:"index"`
	Category           *string      `json:"category"`
	StringValue        *string      `json:"stringValue"`
	IntValue           *int         `json:"intValue"`
	BooleanValue       *bool        `json:"booleanValue"`
	FloatValue         *float64     `json:"floatValue"`
	LatitudeValue      *float64     `json:"latitudeValue"`
	LongitudeValue     *float64     `json:"longitudeValue"`
	RangeFromValue     *float64     `json:"rangeFromValue"`
	RangeToValue       *float64     `json:"rangeToValue"`
	IsEditable         *bool        `json:"isEditable"`
	IsInstanceProperty *bool        `json:"isInstanceProperty"`
	IsMandatory        *bool        `json:"isMandatory"`
	IsDeleted          *bool        `json:"isDeleted"`
}

type PythonPackage struct {
	Version           string    `json:"version"`
	WhlFileKey        string    `json:"whlFileKey"`
	UploadTime        time.Time `json:"uploadTime"`
	HasBreakingChange bool      `json:"hasBreakingChange"`
}

type ReportFilterInput struct {
	Name    string                `json:"name"`
	Entity  FilterEntity          `json:"entity"`
	Filters []*GeneralFilterInput `json:"filters"`
}

// A connection to a list of search entries.
type SearchEntriesConnection struct {
	// A list of search entry edges.
	Edges []*SearchEntryEdge `json:"edges"`
	// Information to aid in pagination.
	PageInfo *ent.PageInfo `json:"pageInfo"`
}

type SearchEntry struct {
	EntityID   int     `json:"entityId"`
	EntityType string  `json:"entityType"`
	Name       string  `json:"name"`
	Type       string  `json:"type"`
	ExternalID *string `json:"externalId"`
}

// A search entry edge in a connection.
type SearchEntryEdge struct {
	// The search entry at the end of the edge.
	Node *SearchEntry `json:"node"`
	// A cursor for use in pagination.
	Cursor ent.Cursor `json:"cursor"`
}

// A search entry edge in a connection.
type SearchNodeEdge struct {
	// The search entry at the end of the edge.
	Node ent.Noder `json:"node"`
	// A cursor for use in pagination.
	Cursor ent.Cursor `json:"cursor"`
}

// A connection to a list of search entries.
type SearchNodesConnection struct {
	// A list of search entry edges.
	Edges []*SearchNodeEdge `json:"edges"`
	// Information to aid in pagination.
	PageInfo *ent.PageInfo `json:"pageInfo"`
}

type ServiceCreateData struct {
	Name               string           `json:"name"`
	ExternalID         *string          `json:"externalId"`
	Status             *ServiceStatus   `json:"status"`
	ServiceTypeID      int              `json:"serviceTypeId"`
	CustomerID         *int             `json:"customerId"`
	UpstreamServiceIds []int            `json:"upstreamServiceIds"`
	Properties         []*PropertyInput `json:"properties"`
}

type ServiceEditData struct {
	ID                 int              `json:"id"`
	Name               *string          `json:"name"`
	ExternalID         *string          `json:"externalId"`
	Status             *ServiceStatus   `json:"status"`
	CustomerID         *int             `json:"customerId"`
	UpstreamServiceIds []int            `json:"upstreamServiceIds"`
	Properties         []*PropertyInput `json:"properties"`
}

type ServiceEndpointDefinitionInput struct {
	ID              *int    `json:"id"`
	Name            string  `json:"name"`
	Role            *string `json:"role"`
	Index           int     `json:"index"`
	EquipmentTypeID int     `json:"equipmentTypeID"`
}

type ServiceFilterInput struct {
	FilterType    ServiceFilterType  `json:"filterType"`
	Operator      FilterOperator     `json:"operator"`
	StringValue   *string            `json:"stringValue"`
	PropertyValue *PropertyTypeInput `json:"propertyValue"`
	IDSet         []int              `json:"idSet"`
	StringSet     []string           `json:"stringSet"`
	MaxDepth      *int               `json:"maxDepth"`
}

type ServiceSearchResult struct {
	Services []*ent.Service `json:"services"`
	Count    int            `json:"count"`
}

type ServiceTypeCreateData struct {
	Name        string                            `json:"name"`
	HasCustomer bool                              `json:"hasCustomer"`
	Properties  []*PropertyTypeInput              `json:"properties"`
	Endpoints   []*ServiceEndpointDefinitionInput `json:"endpoints"`
}

type ServiceTypeEditData struct {
	ID          int                               `json:"id"`
	Name        string                            `json:"name"`
	HasCustomer bool                              `json:"hasCustomer"`
	Properties  []*PropertyTypeInput              `json:"properties"`
	Endpoints   []*ServiceEndpointDefinitionInput `json:"endpoints"`
}

type SurveyCellScanData struct {
	NetworkType           CellularNetworkType `json:"networkType"`
	SignalStrength        int                 `json:"signalStrength"`
	Timestamp             *int                `json:"timestamp"`
	BaseStationID         *string             `json:"baseStationID"`
	NetworkID             *string             `json:"networkID"`
	SystemID              *string             `json:"systemID"`
	CellID                *string             `json:"cellID"`
	LocationAreaCode      *string             `json:"locationAreaCode"`
	MobileCountryCode     *string             `json:"mobileCountryCode"`
	MobileNetworkCode     *string             `json:"mobileNetworkCode"`
	PrimaryScramblingCode *string             `json:"primaryScramblingCode"`
	Operator              *string             `json:"operator"`
	Arfcn                 *int                `json:"arfcn"`
	PhysicalCellID        *string             `json:"physicalCellID"`
	TrackingAreaCode      *string             `json:"trackingAreaCode"`
	TimingAdvance         *int                `json:"timingAdvance"`
	Earfcn                *int                `json:"earfcn"`
	Uarfcn                *int                `json:"uarfcn"`
	Latitude              *float64            `json:"latitude"`
	Longitude             *float64            `json:"longitude"`
}

type SurveyCreateData struct {
	Name                string                    `json:"name"`
	OwnerName           *string                   `json:"ownerName"`
	CreationTimestamp   *int                      `json:"creationTimestamp"`
	CompletionTimestamp int                       `json:"completionTimestamp"`
	Status              *SurveyStatus             `json:"status"`
	LocationID          int                       `json:"locationID"`
	SurveyResponses     []*SurveyQuestionResponse `json:"surveyResponses"`
}

type SurveyQuestionResponse struct {
	FormName         *string               `json:"formName"`
	FormDescription  *string               `json:"formDescription"`
	FormIndex        int                   `json:"formIndex"`
	QuestionFormat   *SurveyQuestionType   `json:"questionFormat"`
	QuestionText     string                `json:"questionText"`
	QuestionIndex    int                   `json:"questionIndex"`
	BoolData         *bool                 `json:"boolData"`
	EmailData        *string               `json:"emailData"`
	Latitude         *float64              `json:"latitude"`
	Longitude        *float64              `json:"longitude"`
	LocationAccuracy *float64              `json:"locationAccuracy"`
	Altitude         *float64              `json:"altitude"`
	PhoneData        *string               `json:"phoneData"`
	TextData         *string               `json:"textData"`
	FloatData        *float64              `json:"floatData"`
	IntData          *int                  `json:"intData"`
	DateData         *int                  `json:"dateData"`
	PhotoData        *FileInput            `json:"photoData"`
	WifiData         []*SurveyWiFiScanData `json:"wifiData"`
	CellData         []*SurveyCellScanData `json:"cellData"`
	ImagesData       []*FileInput          `json:"imagesData"`
}

type SurveyTemplateCategoryInput struct {
	ID                      *int                           `json:"id"`
	CategoryTitle           string                         `json:"categoryTitle"`
	CategoryDescription     string                         `json:"categoryDescription"`
	SurveyTemplateQuestions []*SurveyTemplateQuestionInput `json:"surveyTemplateQuestions"`
}

type SurveyTemplateQuestionInput struct {
	ID                  *int               `json:"id"`
	QuestionTitle       string             `json:"questionTitle"`
	QuestionDescription string             `json:"questionDescription"`
	QuestionType        SurveyQuestionType `json:"questionType"`
	Index               int                `json:"index"`
}

type SurveyWiFiScanData struct {
	Timestamp    int      `json:"timestamp"`
	Frequency    int      `json:"frequency"`
	Channel      int      `json:"channel"`
	Bssid        string   `json:"bssid"`
	Strength     int      `json:"strength"`
	Ssid         *string  `json:"ssid"`
	Band         *string  `json:"band"`
	ChannelWidth *int     `json:"channelWidth"`
	Capabilities *string  `json:"capabilities"`
	Latitude     *float64 `json:"latitude"`
	Longitude    *float64 `json:"longitude"`
}

type TechnicianCheckListItemInput struct {
	ID                 int                   `json:"id"`
	SelectedEnumValues *string               `json:"selectedEnumValues"`
	StringValue        *string               `json:"stringValue"`
	Checked            *bool                 `json:"checked"`
	YesNoResponse      *YesNoResponse        `json:"yesNoResponse"`
	WifiData           []*SurveyWiFiScanData `json:"wifiData"`
	CellData           []*SurveyCellScanData `json:"cellData"`
}

type TechnicianInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type TechnicianWorkOrderUploadInput struct {
	WorkOrderID int                             `json:"workOrderId"`
	Checklist   []*TechnicianCheckListItemInput `json:"checklist"`
}

type TopologyLink struct {
	Type   TopologyLinkType `json:"type"`
	Source ent.Noder        `json:"source"`
	Target ent.Noder        `json:"target"`
}

type UpdateUserGroupsInput struct {
	ID             int   `json:"id"`
	AddGroupIds    []int `json:"addGroupIds"`
	RemoveGroupIds []int `json:"removeGroupIds"`
}

type UpdateUsersGroupMembersInput struct {
	ID            int   `json:"id"`
	AddUserIds    []int `json:"addUserIds"`
	RemoveUserIds []int `json:"removeUserIds"`
}

type UserFilterInput struct {
	FilterType    UserFilterType     `json:"filterType"`
	Operator      FilterOperator     `json:"operator"`
	StringValue   *string            `json:"stringValue"`
	PropertyValue *PropertyTypeInput `json:"propertyValue"`
	IDSet         []int              `json:"idSet"`
	StringSet     []string           `json:"stringSet"`
	MaxDepth      *int               `json:"maxDepth"`
}

type UserSearchResult struct {
	Users []*ent.User `json:"users"`
	Count int         `json:"count"`
}

type WorkOrderDefinitionInput struct {
	ID    *int `json:"id"`
	Index *int `json:"index"`
	Type  int  `json:"type"`
}

type WorkOrderExecutionResult struct {
	ID               int              `json:"id"`
	Name             string           `json:"name"`
	EquipmentAdded   []*ent.Equipment `json:"equipmentAdded"`
	EquipmentRemoved []int            `json:"equipmentRemoved"`
	LinkAdded        []*ent.Link      `json:"linkAdded"`
	LinkRemoved      []int            `json:"linkRemoved"`
}

type WorkOrderFilterInput struct {
	FilterType    WorkOrderFilterType `json:"filterType"`
	Operator      FilterOperator      `json:"operator"`
	StringValue   *string             `json:"stringValue"`
	IDSet         []int               `json:"idSet"`
	StringSet     []string            `json:"stringSet"`
	PropertyValue *PropertyTypeInput  `json:"propertyValue"`
	MaxDepth      *int                `json:"maxDepth"`
}

type WorkOrderSearchResult struct {
	WorkOrders []*ent.WorkOrder `json:"workOrders"`
	Count      int              `json:"count"`
}

type WorkforceCud struct {
	Create            *BasicPermissionRule `json:"create"`
	Update            *BasicPermissionRule `json:"update"`
	Delete            *BasicPermissionRule `json:"delete"`
	Assign            *BasicPermissionRule `json:"assign"`
	TransferOwnership *BasicPermissionRule `json:"transferOwnership"`
}

type WorkforcePolicy struct {
	Read      *BasicPermissionRule `json:"read"`
	Data      *WorkforceCud        `json:"data"`
	Templates *Cud                 `json:"templates"`
}

func (WorkforcePolicy) IsSystemPolicy() {}

type CellularNetworkType string

const (
	CellularNetworkTypeCdma  CellularNetworkType = "CDMA"
	CellularNetworkTypeGsm   CellularNetworkType = "GSM"
	CellularNetworkTypeLte   CellularNetworkType = "LTE"
	CellularNetworkTypeWcdma CellularNetworkType = "WCDMA"
)

var AllCellularNetworkType = []CellularNetworkType{
	CellularNetworkTypeCdma,
	CellularNetworkTypeGsm,
	CellularNetworkTypeLte,
	CellularNetworkTypeWcdma,
}

func (e CellularNetworkType) IsValid() bool {
	switch e {
	case CellularNetworkTypeCdma, CellularNetworkTypeGsm, CellularNetworkTypeLte, CellularNetworkTypeWcdma:
		return true
	}
	return false
}

func (e CellularNetworkType) String() string {
	return string(e)
}

func (e *CellularNetworkType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CellularNetworkType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CellularNetworkType", str)
	}
	return nil
}

func (e CellularNetworkType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type CheckListItemEnumSelectionMode string

const (
	CheckListItemEnumSelectionModeSingle   CheckListItemEnumSelectionMode = "single"
	CheckListItemEnumSelectionModeMultiple CheckListItemEnumSelectionMode = "multiple"
)

var AllCheckListItemEnumSelectionMode = []CheckListItemEnumSelectionMode{
	CheckListItemEnumSelectionModeSingle,
	CheckListItemEnumSelectionModeMultiple,
}

func (e CheckListItemEnumSelectionMode) IsValid() bool {
	switch e {
	case CheckListItemEnumSelectionModeSingle, CheckListItemEnumSelectionModeMultiple:
		return true
	}
	return false
}

func (e CheckListItemEnumSelectionMode) String() string {
	return string(e)
}

func (e *CheckListItemEnumSelectionMode) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CheckListItemEnumSelectionMode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CheckListItemEnumSelectionMode", str)
	}
	return nil
}

func (e CheckListItemEnumSelectionMode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type CheckListItemType string

const (
	CheckListItemTypeSimple   CheckListItemType = "simple"
	CheckListItemTypeString   CheckListItemType = "string"
	CheckListItemTypeEnum     CheckListItemType = "enum"
	CheckListItemTypeFiles    CheckListItemType = "files"
	CheckListItemTypeYesNo    CheckListItemType = "yes_no"
	CheckListItemTypeCellScan CheckListItemType = "cell_scan"
	CheckListItemTypeWifiScan CheckListItemType = "wifi_scan"
)

var AllCheckListItemType = []CheckListItemType{
	CheckListItemTypeSimple,
	CheckListItemTypeString,
	CheckListItemTypeEnum,
	CheckListItemTypeFiles,
	CheckListItemTypeYesNo,
	CheckListItemTypeCellScan,
	CheckListItemTypeWifiScan,
}

func (e CheckListItemType) IsValid() bool {
	switch e {
	case CheckListItemTypeSimple, CheckListItemTypeString, CheckListItemTypeEnum, CheckListItemTypeFiles, CheckListItemTypeYesNo, CheckListItemTypeCellScan, CheckListItemTypeWifiScan:
		return true
	}
	return false
}

func (e CheckListItemType) String() string {
	return string(e)
}

func (e *CheckListItemType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CheckListItemType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CheckListItemType", str)
	}
	return nil
}

func (e CheckListItemType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type CommentEntity string

const (
	CommentEntityWorkOrder CommentEntity = "WORK_ORDER"
	CommentEntityProject   CommentEntity = "PROJECT"
)

var AllCommentEntity = []CommentEntity{
	CommentEntityWorkOrder,
	CommentEntityProject,
}

func (e CommentEntity) IsValid() bool {
	switch e {
	case CommentEntityWorkOrder, CommentEntityProject:
		return true
	}
	return false
}

func (e CommentEntity) String() string {
	return string(e)
}

func (e *CommentEntity) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CommentEntity(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CommentEntity", str)
	}
	return nil
}

func (e CommentEntity) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// what type of equipment we filter about
type EquipmentFilterType string

const (
	EquipmentFilterTypeEquipInstName       EquipmentFilterType = "EQUIP_INST_NAME"
	EquipmentFilterTypeEquipInstExternalID EquipmentFilterType = "EQUIP_INST_EXTERNAL_ID"
	EquipmentFilterTypeProperty            EquipmentFilterType = "PROPERTY"
	EquipmentFilterTypeLocationInst        EquipmentFilterType = "LOCATION_INST"
	EquipmentFilterTypeEquipmentType       EquipmentFilterType = "EQUIPMENT_TYPE"
)

var AllEquipmentFilterType = []EquipmentFilterType{
	EquipmentFilterTypeEquipInstName,
	EquipmentFilterTypeEquipInstExternalID,
	EquipmentFilterTypeProperty,
	EquipmentFilterTypeLocationInst,
	EquipmentFilterTypeEquipmentType,
}

func (e EquipmentFilterType) IsValid() bool {
	switch e {
	case EquipmentFilterTypeEquipInstName, EquipmentFilterTypeEquipInstExternalID, EquipmentFilterTypeProperty, EquipmentFilterTypeLocationInst, EquipmentFilterTypeEquipmentType:
		return true
	}
	return false
}

func (e EquipmentFilterType) String() string {
	return string(e)
}

func (e *EquipmentFilterType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EquipmentFilterType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EquipmentFilterType", str)
	}
	return nil
}

func (e EquipmentFilterType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type FileType string

const (
	FileTypeImage FileType = "IMAGE"
	FileTypeFile  FileType = "FILE"
)

var AllFileType = []FileType{
	FileTypeImage,
	FileTypeFile,
}

func (e FileType) IsValid() bool {
	switch e {
	case FileTypeImage, FileTypeFile:
		return true
	}
	return false
}

func (e FileType) String() string {
	return string(e)
}

func (e *FileType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FileType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FileType", str)
	}
	return nil
}

func (e FileType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type FilterEntity string

const (
	FilterEntityWorkOrder FilterEntity = "WORK_ORDER"
	FilterEntityPort      FilterEntity = "PORT"
	FilterEntityEquipment FilterEntity = "EQUIPMENT"
	FilterEntityLink      FilterEntity = "LINK"
	FilterEntityLocation  FilterEntity = "LOCATION"
	FilterEntityService   FilterEntity = "SERVICE"
)

var AllFilterEntity = []FilterEntity{
	FilterEntityWorkOrder,
	FilterEntityPort,
	FilterEntityEquipment,
	FilterEntityLink,
	FilterEntityLocation,
	FilterEntityService,
}

func (e FilterEntity) IsValid() bool {
	switch e {
	case FilterEntityWorkOrder, FilterEntityPort, FilterEntityEquipment, FilterEntityLink, FilterEntityLocation, FilterEntityService:
		return true
	}
	return false
}

func (e FilterEntity) String() string {
	return string(e)
}

func (e *FilterEntity) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FilterEntity(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FilterEntity", str)
	}
	return nil
}

func (e FilterEntity) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// operators to filter search by
type FilterOperator string

const (
	FilterOperatorIs              FilterOperator = "IS"
	FilterOperatorContains        FilterOperator = "CONTAINS"
	FilterOperatorIsOneOf         FilterOperator = "IS_ONE_OF"
	FilterOperatorIsNotOneOf      FilterOperator = "IS_NOT_ONE_OF"
	FilterOperatorDateGreaterThan FilterOperator = "DATE_GREATER_THAN"
	FilterOperatorDateLessThan    FilterOperator = "DATE_LESS_THAN"
)

var AllFilterOperator = []FilterOperator{
	FilterOperatorIs,
	FilterOperatorContains,
	FilterOperatorIsOneOf,
	FilterOperatorIsNotOneOf,
	FilterOperatorDateGreaterThan,
	FilterOperatorDateLessThan,
}

func (e FilterOperator) IsValid() bool {
	switch e {
	case FilterOperatorIs, FilterOperatorContains, FilterOperatorIsOneOf, FilterOperatorIsNotOneOf, FilterOperatorDateGreaterThan, FilterOperatorDateLessThan:
		return true
	}
	return false
}

func (e FilterOperator) String() string {
	return string(e)
}

func (e *FilterOperator) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FilterOperator(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FilterOperator", str)
	}
	return nil
}

func (e FilterOperator) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Equipment planned status
type FutureState string

const (
	FutureStateInstall FutureState = "INSTALL"
	FutureStateRemove  FutureState = "REMOVE"
)

var AllFutureState = []FutureState{
	FutureStateInstall,
	FutureStateRemove,
}

func (e FutureState) IsValid() bool {
	switch e {
	case FutureStateInstall, FutureStateRemove:
		return true
	}
	return false
}

func (e FutureState) String() string {
	return string(e)
}

func (e *FutureState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FutureState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FutureState", str)
	}
	return nil
}

func (e FutureState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ImageEntity string

const (
	ImageEntityLocation      ImageEntity = "LOCATION"
	ImageEntityWorkOrder     ImageEntity = "WORK_ORDER"
	ImageEntitySiteSurvey    ImageEntity = "SITE_SURVEY"
	ImageEntityEquipment     ImageEntity = "EQUIPMENT"
	ImageEntityUser          ImageEntity = "USER"
	ImageEntityChecklistItem ImageEntity = "CHECKLIST_ITEM"
)

var AllImageEntity = []ImageEntity{
	ImageEntityLocation,
	ImageEntityWorkOrder,
	ImageEntitySiteSurvey,
	ImageEntityEquipment,
	ImageEntityUser,
	ImageEntityChecklistItem,
}

func (e ImageEntity) IsValid() bool {
	switch e {
	case ImageEntityLocation, ImageEntityWorkOrder, ImageEntitySiteSurvey, ImageEntityEquipment, ImageEntityUser, ImageEntityChecklistItem:
		return true
	}
	return false
}

func (e ImageEntity) String() string {
	return string(e)
}

func (e *ImageEntity) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ImageEntity(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ImageEntity", str)
	}
	return nil
}

func (e ImageEntity) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// what filters should we apply on links
type LinkFilterType string

const (
	LinkFilterTypeLinkFutureStatus LinkFilterType = "LINK_FUTURE_STATUS"
	LinkFilterTypeEquipmentType    LinkFilterType = "EQUIPMENT_TYPE"
	LinkFilterTypeLocationInst     LinkFilterType = "LOCATION_INST"
	LinkFilterTypeProperty         LinkFilterType = "PROPERTY"
	LinkFilterTypeServiceInst      LinkFilterType = "SERVICE_INST"
	LinkFilterTypeEquipmentInst    LinkFilterType = "EQUIPMENT_INST"
)

var AllLinkFilterType = []LinkFilterType{
	LinkFilterTypeLinkFutureStatus,
	LinkFilterTypeEquipmentType,
	LinkFilterTypeLocationInst,
	LinkFilterTypeProperty,
	LinkFilterTypeServiceInst,
	LinkFilterTypeEquipmentInst,
}

func (e LinkFilterType) IsValid() bool {
	switch e {
	case LinkFilterTypeLinkFutureStatus, LinkFilterTypeEquipmentType, LinkFilterTypeLocationInst, LinkFilterTypeProperty, LinkFilterTypeServiceInst, LinkFilterTypeEquipmentInst:
		return true
	}
	return false
}

func (e LinkFilterType) String() string {
	return string(e)
}

func (e *LinkFilterType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = LinkFilterType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid LinkFilterType", str)
	}
	return nil
}

func (e LinkFilterType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// what filters should we apply on locations
type LocationFilterType string

const (
	LocationFilterTypeLocationInst             LocationFilterType = "LOCATION_INST"
	LocationFilterTypeLocationInstName         LocationFilterType = "LOCATION_INST_NAME"
	LocationFilterTypeLocationInstExternalID   LocationFilterType = "LOCATION_INST_EXTERNAL_ID"
	LocationFilterTypeLocationType             LocationFilterType = "LOCATION_TYPE"
	LocationFilterTypeLocationInstHasEquipment LocationFilterType = "LOCATION_INST_HAS_EQUIPMENT"
	LocationFilterTypeProperty                 LocationFilterType = "PROPERTY"
)

var AllLocationFilterType = []LocationFilterType{
	LocationFilterTypeLocationInst,
	LocationFilterTypeLocationInstName,
	LocationFilterTypeLocationInstExternalID,
	LocationFilterTypeLocationType,
	LocationFilterTypeLocationInstHasEquipment,
	LocationFilterTypeProperty,
}

func (e LocationFilterType) IsValid() bool {
	switch e {
	case LocationFilterTypeLocationInst, LocationFilterTypeLocationInstName, LocationFilterTypeLocationInstExternalID, LocationFilterTypeLocationType, LocationFilterTypeLocationInstHasEquipment, LocationFilterTypeProperty:
		return true
	}
	return false
}

func (e LocationFilterType) String() string {
	return string(e)
}

func (e *LocationFilterType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = LocationFilterType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid LocationFilterType", str)
	}
	return nil
}

func (e LocationFilterType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// what filters should we apply on ports
type PortFilterType string

const (
	PortFilterTypePortDef           PortFilterType = "PORT_DEF"
	PortFilterTypePortInstHasLink   PortFilterType = "PORT_INST_HAS_LINK"
	PortFilterTypePortInstEquipment PortFilterType = "PORT_INST_EQUIPMENT"
	PortFilterTypeLocationInst      PortFilterType = "LOCATION_INST"
	PortFilterTypeProperty          PortFilterType = "PROPERTY"
	PortFilterTypeServiceInst       PortFilterType = "SERVICE_INST"
)

var AllPortFilterType = []PortFilterType{
	PortFilterTypePortDef,
	PortFilterTypePortInstHasLink,
	PortFilterTypePortInstEquipment,
	PortFilterTypeLocationInst,
	PortFilterTypeProperty,
	PortFilterTypeServiceInst,
}

func (e PortFilterType) IsValid() bool {
	switch e {
	case PortFilterTypePortDef, PortFilterTypePortInstHasLink, PortFilterTypePortInstEquipment, PortFilterTypeLocationInst, PortFilterTypeProperty, PortFilterTypeServiceInst:
		return true
	}
	return false
}

func (e PortFilterType) String() string {
	return string(e)
}

func (e *PortFilterType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PortFilterType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PortFilterType", str)
	}
	return nil
}

func (e PortFilterType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ProjectFilterType string

const (
	ProjectFilterTypeProjectName ProjectFilterType = "PROJECT_NAME"
)

var AllProjectFilterType = []ProjectFilterType{
	ProjectFilterTypeProjectName,
}

func (e ProjectFilterType) IsValid() bool {
	switch e {
	case ProjectFilterTypeProjectName:
		return true
	}
	return false
}

func (e ProjectFilterType) String() string {
	return string(e)
}

func (e *ProjectFilterType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ProjectFilterType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ProjectFilterType", str)
	}
	return nil
}

func (e ProjectFilterType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PropertyEntity string

const (
	PropertyEntityEquipment PropertyEntity = "EQUIPMENT"
	PropertyEntityService   PropertyEntity = "SERVICE"
	PropertyEntityLink      PropertyEntity = "LINK"
	PropertyEntityPort      PropertyEntity = "PORT"
	PropertyEntityLocation  PropertyEntity = "LOCATION"
	PropertyEntityWorkOrder PropertyEntity = "WORK_ORDER"
	PropertyEntityProject   PropertyEntity = "PROJECT"
)

var AllPropertyEntity = []PropertyEntity{
	PropertyEntityEquipment,
	PropertyEntityService,
	PropertyEntityLink,
	PropertyEntityPort,
	PropertyEntityLocation,
	PropertyEntityWorkOrder,
	PropertyEntityProject,
}

func (e PropertyEntity) IsValid() bool {
	switch e {
	case PropertyEntityEquipment, PropertyEntityService, PropertyEntityLink, PropertyEntityPort, PropertyEntityLocation, PropertyEntityWorkOrder, PropertyEntityProject:
		return true
	}
	return false
}

func (e PropertyEntity) String() string {
	return string(e)
}

func (e *PropertyEntity) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PropertyEntity(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PropertyEntity", str)
	}
	return nil
}

func (e PropertyEntity) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PropertyKind string

const (
	PropertyKindString        PropertyKind = "string"
	PropertyKindInt           PropertyKind = "int"
	PropertyKindBool          PropertyKind = "bool"
	PropertyKindFloat         PropertyKind = "float"
	PropertyKindDate          PropertyKind = "date"
	PropertyKindEnum          PropertyKind = "enum"
	PropertyKindRange         PropertyKind = "range"
	PropertyKindEmail         PropertyKind = "email"
	PropertyKindGpsLocation   PropertyKind = "gps_location"
	PropertyKindDatetimeLocal PropertyKind = "datetime_local"
	PropertyKindNode          PropertyKind = "node"
)

var AllPropertyKind = []PropertyKind{
	PropertyKindString,
	PropertyKindInt,
	PropertyKindBool,
	PropertyKindFloat,
	PropertyKindDate,
	PropertyKindEnum,
	PropertyKindRange,
	PropertyKindEmail,
	PropertyKindGpsLocation,
	PropertyKindDatetimeLocal,
	PropertyKindNode,
}

func (e PropertyKind) IsValid() bool {
	switch e {
	case PropertyKindString, PropertyKindInt, PropertyKindBool, PropertyKindFloat, PropertyKindDate, PropertyKindEnum, PropertyKindRange, PropertyKindEmail, PropertyKindGpsLocation, PropertyKindDatetimeLocal, PropertyKindNode:
		return true
	}
	return false
}

func (e PropertyKind) String() string {
	return string(e)
}

func (e *PropertyKind) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PropertyKind(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PropertyKind", str)
	}
	return nil
}

func (e PropertyKind) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// what filters should we apply on services
type ServiceFilterType string

const (
	ServiceFilterTypeServiceInstName         ServiceFilterType = "SERVICE_INST_NAME"
	ServiceFilterTypeServiceStatus           ServiceFilterType = "SERVICE_STATUS"
	ServiceFilterTypeServiceType             ServiceFilterType = "SERVICE_TYPE"
	ServiceFilterTypeServiceInstExternalID   ServiceFilterType = "SERVICE_INST_EXTERNAL_ID"
	ServiceFilterTypeServiceInstCustomerName ServiceFilterType = "SERVICE_INST_CUSTOMER_NAME"
	ServiceFilterTypeProperty                ServiceFilterType = "PROPERTY"
	ServiceFilterTypeLocationInst            ServiceFilterType = "LOCATION_INST"
	ServiceFilterTypeEquipmentInService      ServiceFilterType = "EQUIPMENT_IN_SERVICE"
)

var AllServiceFilterType = []ServiceFilterType{
	ServiceFilterTypeServiceInstName,
	ServiceFilterTypeServiceStatus,
	ServiceFilterTypeServiceType,
	ServiceFilterTypeServiceInstExternalID,
	ServiceFilterTypeServiceInstCustomerName,
	ServiceFilterTypeProperty,
	ServiceFilterTypeLocationInst,
	ServiceFilterTypeEquipmentInService,
}

func (e ServiceFilterType) IsValid() bool {
	switch e {
	case ServiceFilterTypeServiceInstName, ServiceFilterTypeServiceStatus, ServiceFilterTypeServiceType, ServiceFilterTypeServiceInstExternalID, ServiceFilterTypeServiceInstCustomerName, ServiceFilterTypeProperty, ServiceFilterTypeLocationInst, ServiceFilterTypeEquipmentInService:
		return true
	}
	return false
}

func (e ServiceFilterType) String() string {
	return string(e)
}

func (e *ServiceFilterType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ServiceFilterType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ServiceFilterType", str)
	}
	return nil
}

func (e ServiceFilterType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ServiceStatus string

const (
	ServiceStatusPending      ServiceStatus = "PENDING"
	ServiceStatusInService    ServiceStatus = "IN_SERVICE"
	ServiceStatusMaintenance  ServiceStatus = "MAINTENANCE"
	ServiceStatusDisconnected ServiceStatus = "DISCONNECTED"
)

var AllServiceStatus = []ServiceStatus{
	ServiceStatusPending,
	ServiceStatusInService,
	ServiceStatusMaintenance,
	ServiceStatusDisconnected,
}

func (e ServiceStatus) IsValid() bool {
	switch e {
	case ServiceStatusPending, ServiceStatusInService, ServiceStatusMaintenance, ServiceStatusDisconnected:
		return true
	}
	return false
}

func (e ServiceStatus) String() string {
	return string(e)
}

func (e *ServiceStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ServiceStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ServiceStatus", str)
	}
	return nil
}

func (e ServiceStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SurveyQuestionType string

const (
	SurveyQuestionTypeBool     SurveyQuestionType = "BOOL"
	SurveyQuestionTypeEmail    SurveyQuestionType = "EMAIL"
	SurveyQuestionTypeCoords   SurveyQuestionType = "COORDS"
	SurveyQuestionTypePhone    SurveyQuestionType = "PHONE"
	SurveyQuestionTypeText     SurveyQuestionType = "TEXT"
	SurveyQuestionTypeTextarea SurveyQuestionType = "TEXTAREA"
	SurveyQuestionTypePhoto    SurveyQuestionType = "PHOTO"
	SurveyQuestionTypeWifi     SurveyQuestionType = "WIFI"
	SurveyQuestionTypeCellular SurveyQuestionType = "CELLULAR"
	SurveyQuestionTypeFloat    SurveyQuestionType = "FLOAT"
	SurveyQuestionTypeInteger  SurveyQuestionType = "INTEGER"
	SurveyQuestionTypeDate     SurveyQuestionType = "DATE"
)

var AllSurveyQuestionType = []SurveyQuestionType{
	SurveyQuestionTypeBool,
	SurveyQuestionTypeEmail,
	SurveyQuestionTypeCoords,
	SurveyQuestionTypePhone,
	SurveyQuestionTypeText,
	SurveyQuestionTypeTextarea,
	SurveyQuestionTypePhoto,
	SurveyQuestionTypeWifi,
	SurveyQuestionTypeCellular,
	SurveyQuestionTypeFloat,
	SurveyQuestionTypeInteger,
	SurveyQuestionTypeDate,
}

func (e SurveyQuestionType) IsValid() bool {
	switch e {
	case SurveyQuestionTypeBool, SurveyQuestionTypeEmail, SurveyQuestionTypeCoords, SurveyQuestionTypePhone, SurveyQuestionTypeText, SurveyQuestionTypeTextarea, SurveyQuestionTypePhoto, SurveyQuestionTypeWifi, SurveyQuestionTypeCellular, SurveyQuestionTypeFloat, SurveyQuestionTypeInteger, SurveyQuestionTypeDate:
		return true
	}
	return false
}

func (e SurveyQuestionType) String() string {
	return string(e)
}

func (e *SurveyQuestionType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SurveyQuestionType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SurveyQuestionType", str)
	}
	return nil
}

func (e SurveyQuestionType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SurveyStatus string

const (
	SurveyStatusPlanned    SurveyStatus = "PLANNED"
	SurveyStatusInprogress SurveyStatus = "INPROGRESS"
	SurveyStatusCompleted  SurveyStatus = "COMPLETED"
)

var AllSurveyStatus = []SurveyStatus{
	SurveyStatusPlanned,
	SurveyStatusInprogress,
	SurveyStatusCompleted,
}

func (e SurveyStatus) IsValid() bool {
	switch e {
	case SurveyStatusPlanned, SurveyStatusInprogress, SurveyStatusCompleted:
		return true
	}
	return false
}

func (e SurveyStatus) String() string {
	return string(e)
}

func (e *SurveyStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SurveyStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SurveyStatus", str)
	}
	return nil
}

func (e SurveyStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TopologyLinkType string

const (
	TopologyLinkTypePhysical TopologyLinkType = "PHYSICAL"
)

var AllTopologyLinkType = []TopologyLinkType{
	TopologyLinkTypePhysical,
}

func (e TopologyLinkType) IsValid() bool {
	switch e {
	case TopologyLinkTypePhysical:
		return true
	}
	return false
}

func (e TopologyLinkType) String() string {
	return string(e)
}

func (e *TopologyLinkType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TopologyLinkType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TopologyLinkType", str)
	}
	return nil
}

func (e TopologyLinkType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// what filters should we apply on users
type UserFilterType string

const (
	UserFilterTypeUserName UserFilterType = "USER_NAME"
)

var AllUserFilterType = []UserFilterType{
	UserFilterTypeUserName,
}

func (e UserFilterType) IsValid() bool {
	switch e {
	case UserFilterTypeUserName:
		return true
	}
	return false
}

func (e UserFilterType) String() string {
	return string(e)
}

func (e *UserFilterType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserFilterType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserFilterType", str)
	}
	return nil
}

func (e UserFilterType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// what type of work order we filter about
type WorkOrderFilterType string

const (
	WorkOrderFilterTypeWorkOrderName         WorkOrderFilterType = "WORK_ORDER_NAME"
	WorkOrderFilterTypeWorkOrderStatus       WorkOrderFilterType = "WORK_ORDER_STATUS"
	WorkOrderFilterTypeWorkOrderOwner        WorkOrderFilterType = "WORK_ORDER_OWNER"
	WorkOrderFilterTypeWorkOrderOwnedBy      WorkOrderFilterType = "WORK_ORDER_OWNED_BY"
	WorkOrderFilterTypeWorkOrderType         WorkOrderFilterType = "WORK_ORDER_TYPE"
	WorkOrderFilterTypeWorkOrderCreationDate WorkOrderFilterType = "WORK_ORDER_CREATION_DATE"
	WorkOrderFilterTypeWorkOrderInstallDate  WorkOrderFilterType = "WORK_ORDER_INSTALL_DATE"
	WorkOrderFilterTypeWorkOrderAssignee     WorkOrderFilterType = "WORK_ORDER_ASSIGNEE"
	WorkOrderFilterTypeWorkOrderAssignedTo   WorkOrderFilterType = "WORK_ORDER_ASSIGNED_TO"
	WorkOrderFilterTypeWorkOrderLocationInst WorkOrderFilterType = "WORK_ORDER_LOCATION_INST"
	WorkOrderFilterTypeWorkOrderPriority     WorkOrderFilterType = "WORK_ORDER_PRIORITY"
	WorkOrderFilterTypeLocationInst          WorkOrderFilterType = "LOCATION_INST"
)

var AllWorkOrderFilterType = []WorkOrderFilterType{
	WorkOrderFilterTypeWorkOrderName,
	WorkOrderFilterTypeWorkOrderStatus,
	WorkOrderFilterTypeWorkOrderOwner,
	WorkOrderFilterTypeWorkOrderOwnedBy,
	WorkOrderFilterTypeWorkOrderType,
	WorkOrderFilterTypeWorkOrderCreationDate,
	WorkOrderFilterTypeWorkOrderInstallDate,
	WorkOrderFilterTypeWorkOrderAssignee,
	WorkOrderFilterTypeWorkOrderAssignedTo,
	WorkOrderFilterTypeWorkOrderLocationInst,
	WorkOrderFilterTypeWorkOrderPriority,
	WorkOrderFilterTypeLocationInst,
}

func (e WorkOrderFilterType) IsValid() bool {
	switch e {
	case WorkOrderFilterTypeWorkOrderName, WorkOrderFilterTypeWorkOrderStatus, WorkOrderFilterTypeWorkOrderOwner, WorkOrderFilterTypeWorkOrderOwnedBy, WorkOrderFilterTypeWorkOrderType, WorkOrderFilterTypeWorkOrderCreationDate, WorkOrderFilterTypeWorkOrderInstallDate, WorkOrderFilterTypeWorkOrderAssignee, WorkOrderFilterTypeWorkOrderAssignedTo, WorkOrderFilterTypeWorkOrderLocationInst, WorkOrderFilterTypeWorkOrderPriority, WorkOrderFilterTypeLocationInst:
		return true
	}
	return false
}

func (e WorkOrderFilterType) String() string {
	return string(e)
}

func (e *WorkOrderFilterType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = WorkOrderFilterType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid WorkOrderFilterType", str)
	}
	return nil
}

func (e WorkOrderFilterType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Work Order priority
type WorkOrderPriority string

const (
	WorkOrderPriorityUrgent WorkOrderPriority = "URGENT"
	WorkOrderPriorityHigh   WorkOrderPriority = "HIGH"
	WorkOrderPriorityMedium WorkOrderPriority = "MEDIUM"
	WorkOrderPriorityLow    WorkOrderPriority = "LOW"
	WorkOrderPriorityNone   WorkOrderPriority = "NONE"
)

var AllWorkOrderPriority = []WorkOrderPriority{
	WorkOrderPriorityUrgent,
	WorkOrderPriorityHigh,
	WorkOrderPriorityMedium,
	WorkOrderPriorityLow,
	WorkOrderPriorityNone,
}

func (e WorkOrderPriority) IsValid() bool {
	switch e {
	case WorkOrderPriorityUrgent, WorkOrderPriorityHigh, WorkOrderPriorityMedium, WorkOrderPriorityLow, WorkOrderPriorityNone:
		return true
	}
	return false
}

func (e WorkOrderPriority) String() string {
	return string(e)
}

func (e *WorkOrderPriority) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = WorkOrderPriority(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid WorkOrderPriority", str)
	}
	return nil
}

func (e WorkOrderPriority) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Work Order status
type WorkOrderStatus string

const (
	WorkOrderStatusPending WorkOrderStatus = "PENDING"
	WorkOrderStatusPlanned WorkOrderStatus = "PLANNED"
	WorkOrderStatusDone    WorkOrderStatus = "DONE"
)

var AllWorkOrderStatus = []WorkOrderStatus{
	WorkOrderStatusPending,
	WorkOrderStatusPlanned,
	WorkOrderStatusDone,
}

func (e WorkOrderStatus) IsValid() bool {
	switch e {
	case WorkOrderStatusPending, WorkOrderStatusPlanned, WorkOrderStatusDone:
		return true
	}
	return false
}

func (e WorkOrderStatus) String() string {
	return string(e)
}

func (e *WorkOrderStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = WorkOrderStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid WorkOrderStatus", str)
	}
	return nil
}

func (e WorkOrderStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type YesNoResponse string

const (
	YesNoResponseYes YesNoResponse = "YES"
	YesNoResponseNo  YesNoResponse = "NO"
)

var AllYesNoResponse = []YesNoResponse{
	YesNoResponseYes,
	YesNoResponseNo,
}

func (e YesNoResponse) IsValid() bool {
	switch e {
	case YesNoResponseYes, YesNoResponseNo:
		return true
	}
	return false
}

func (e YesNoResponse) String() string {
	return string(e)
}

func (e *YesNoResponse) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = YesNoResponse(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid YesNoResponse", str)
	}
	return nil
}

func (e YesNoResponse) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
