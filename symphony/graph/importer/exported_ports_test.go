// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package importer

import (
	"context"
	"strconv"
	"strings"
	"testing"

	"github.com/facebookincubator/symphony/graph/ent/propertytype"

	"github.com/facebookincubator/symphony/graph/ent/equipmentport"

	"github.com/facebookincubator/symphony/graph/ent/equipmentportdefinition"

	"github.com/AlekSi/pointer"
	"github.com/facebookincubator/symphony/graph/graphql/models"

	"github.com/facebookincubator/symphony/graph/viewer/viewertest"
	"github.com/stretchr/testify/require"
)

const (
	parentEquip  = "parentEquipmentName"
	parentEquip2 = "parentEquipmentName2"
	parentEquip3 = "parentEquipmentName3"

	currEquip   = "currEquipmentName"
	portName1   = "port1"
	portName2   = "port2"
	propNameStr = "propNameStr"
	posName     = "pos"

	propNameDate = "propNameDate"
	propNameBool = "propNameBool"
	propNameInt  = "propNameInt"
	locationL    = "locationL"
	locationM    = "locationM"
	locationS    = "locationS"
)

type portData struct {
	equipParentID   int
	equipParent2ID  int
	equipParent3ID  int
	equipChildID    int
	equipChild2ID   int
	portDef1        int
	parentPortInst1 int
	parentPortInst2 int
	parentPortInst3 int
	portDef2        int
	childPortInst1  int
	childPortInst2  int
	linkID          int
}

func preparePortTypeData(ctx context.Context, t *testing.T, r TestImporterResolver) portData {
	ids := prepareEquipmentTypeData(ctx, t, r)
	mr := r.importer.r.Mutation()

	ptyp, _ := mr.AddEquipmentPortType(ctx, models.AddEquipmentPortTypeInput{
		Name: "portType1",
		Properties: []*models.PropertyTypeInput{
			{
				Name:        propNameStr,
				Type:        models.PropertyKindString,
				StringValue: pointer.ToString("t1"),
			},
			{
				Name: propNameInt,
				Type: models.PropertyKindInt,
			},
		},
		LinkProperties: []*models.PropertyTypeInput{
			{
				Name: propNameInt,
				Type: models.PropertyKindInt,
			},
		},
	})
	port1 := models.EquipmentPortInput{
		Name:       portName1,
		PortTypeID: &ptyp.ID,
	}
	pos1 := models.EquipmentPositionInput{
		Name: posName,
	}
	etype, _ := r.client.EquipmentType.Get(ctx, ids.equipTypeID)
	etype, _ = mr.EditEquipmentType(ctx, models.EditEquipmentTypeInput{
		ID:        etype.ID,
		Name:      etype.Name,
		Ports:     []*models.EquipmentPortInput{&port1},
		Positions: []*models.EquipmentPositionInput{&pos1},
	})

	ptyp2, _ := mr.AddEquipmentPortType(ctx, models.AddEquipmentPortTypeInput{
		Name: "portType2",
		Properties: []*models.PropertyTypeInput{
			{
				Name:        propNameDate,
				Type:        models.PropertyKindDate,
				StringValue: pointer.ToString("1988-03-29"),
			},
			{
				Name: propNameBool,
				Type: models.PropertyKindBool,
			},
		},
		LinkProperties: []*models.PropertyTypeInput{
			{
				Name:        propNameDate,
				Type:        models.PropertyKindDate,
				StringValue: pointer.ToString("2020-01-01"),
			},
			{
				Name:         propNameBool,
				Type:         models.PropertyKindBool,
				BooleanValue: pointer.ToBool(true),
			},
		},
	})
	port2 := models.EquipmentPortInput{
		Name:       portName2,
		PortTypeID: &ptyp2.ID,
	}
	etype2, _ := r.client.EquipmentType.Get(ctx, ids.equipTypeID2)
	etype2, _ = mr.EditEquipmentType(ctx, models.EditEquipmentTypeInput{
		ID:    etype2.ID,
		Name:  etype2.Name,
		Ports: []*models.EquipmentPortInput{&port2},
	})
	gpLocation, _ := mr.AddLocation(ctx, models.AddLocationInput{
		Name: locationL,
		Type: ids.locTypeIDL,
	})
	pLocation, _ := mr.AddLocation(ctx, models.AddLocationInput{
		Name:   locationM,
		Type:   ids.locTypeIDM,
		Parent: &gpLocation.ID,
	})
	sLocation, _ := mr.AddLocation(ctx, models.AddLocationInput{
		Name:   locationS,
		Type:   ids.locTypeIDS,
		Parent: &pLocation.ID,
	})
	parentEquipment, _ := mr.AddEquipment(ctx, models.AddEquipmentInput{
		Name:     parentEquip,
		Type:     etype.ID,
		Location: &sLocation.ID,
	})
	portDef1 := etype.QueryPortDefinitions().OnlyX(ctx)
	posDef1 := etype.QueryPositionDefinitions().OnlyX(ctx)
	parentPortInst1 := parentEquipment.QueryPorts().Where(equipmentport.HasDefinitionWith(equipmentportdefinition.ID(portDef1.ID))).OnlyX(ctx)

	parentEquip2, _ := mr.AddEquipment(ctx, models.AddEquipmentInput{
		Name:     parentEquip2,
		Type:     etype.ID,
		Location: &sLocation.ID,
	})
	parentPortInst2 := parentEquip2.QueryPorts().Where(equipmentport.HasDefinitionWith(equipmentportdefinition.ID(portDef1.ID))).OnlyX(ctx)
	parentEquip3, _ := mr.AddEquipment(ctx, models.AddEquipmentInput{
		Name:     parentEquip3,
		Type:     etype.ID,
		Location: &sLocation.ID,
	})
	parentPortInst3 := parentEquip3.QueryPorts().Where(equipmentport.HasDefinitionWith(equipmentportdefinition.ID(portDef1.ID))).OnlyX(ctx)

	childEquip, _ := mr.AddEquipment(ctx, models.AddEquipmentInput{
		Name:               currEquip,
		Type:               etype2.ID,
		Parent:             &parentEquipment.ID,
		PositionDefinition: &posDef1.ID,
	})
	portDef2 := etype2.QueryPortDefinitions().OnlyX(ctx)

	childPortInst1 := childEquip.QueryPorts().Where(equipmentport.HasDefinitionWith(equipmentportdefinition.ID(portDef2.ID))).OnlyX(ctx)

	l, _ := mr.AddLink(ctx, models.AddLinkInput{
		Sides: []*models.LinkSide{
			{Equipment: childEquip.ID, Port: portDef2.ID},
			{Equipment: parentEquip2.ID, Port: portDef1.ID},
		},
	})

	childEquip2, _ := mr.AddEquipment(ctx, models.AddEquipmentInput{
		Name:               currEquip,
		Type:               etype2.ID,
		Parent:             &parentEquip2.ID,
		PositionDefinition: &posDef1.ID,
	})
	childPortDef2 := etype2.QueryPortDefinitions().OnlyX(ctx)

	childPortInst2 := childEquip.QueryPorts().Where(equipmentport.HasDefinitionWith(equipmentportdefinition.ID(childPortDef2.ID))).OnlyX(ctx)
	/* locL -> locM -> locS:
	parent1 (port1) -> child (port2[linked])
	parent2 (port1[linked]) -> child (port2)
	parent3 (port1)
	*/
	return portData{
		equipParentID:   parentEquipment.ID,
		equipParent2ID:  parentEquip2.ID,
		equipParent3ID:  parentEquip3.ID,
		equipChildID:    childEquip.ID,
		equipChild2ID:   childEquip2.ID,
		portDef1:        portDef1.ID,
		parentPortInst1: parentPortInst1.ID,
		parentPortInst2: parentPortInst2.ID,
		parentPortInst3: parentPortInst3.ID,

		portDef2:       portDef2.ID,
		childPortInst1: childPortInst1.ID,
		childPortInst2: childPortInst2.ID,
		linkID:         l.ID,
	}
}

func TestPortTitleInputValidation(t *testing.T) {
	r := newImporterTestResolver(t)
	importer := r.importer
	defer r.drv.Close()

	ctx := newImportContext(viewertest.NewContext(context.Background(), r.client))
	var (
		portDataHeader = [...]string{"Port ID", "Port Name", "Port Type", "Equipment Name", "Equipment Type"}
		parentsHeader  = [...]string{"Parent Equipment (3)", "Parent Equipment (2)", "Parent Equipment", "Equipment Position"}
		linkDataHeader = [...]string{"Linked Port ID", "Linked Port Name", "Linked Equipment ID", "Linked Equipment"}
		servicesHeader = [...]string{"Service Names"}
	)
	prepareBasicData(ctx, t, *r)
	header, _ := NewImportHeader([]string{"aa"}, ImportEntityPort)
	err := importer.inputValidationsPorts(ctx, header)
	require.Error(t, err)
	header, _ = NewImportHeader(portDataHeader[:], ImportEntityPort)
	err = importer.inputValidationsPorts(ctx, header)
	require.Error(t, err)
	header, _ = NewImportHeader(linkDataHeader[:], ImportEntityPort)
	err = importer.inputValidationsPorts(ctx, header)
	require.Error(t, err)

	locationTypeNotInOrder := append(append(append(append(portDataHeader[:], []string{locTypeNameS, locTypeNameM, locTypeNameL}...), parentsHeader[:]...), linkDataHeader[:]...), servicesHeader[:]...)
	header, _ = NewImportHeader(locationTypeNotInOrder, ImportEntityPort)
	err = importer.inputValidationsPorts(ctx, header)
	require.Error(t, err)

	locationTypeInOrder := append(append(append(append(portDataHeader[:], []string{locTypeNameL, locTypeNameM, locTypeNameS}...), parentsHeader[:]...), linkDataHeader[:]...), servicesHeader[:]...)
	header, _ = NewImportHeader(locationTypeInOrder, ImportEntityPort)
	err = importer.inputValidationsPorts(ctx, header)
	require.NoError(t, err)
}

func TestGeneralPortsImport(t *testing.T) {
	r := newImporterTestResolver(t)
	importer := r.importer
	defer r.drv.Close()

	ctx := newImportContext(viewertest.NewContext(context.Background(), r.client))
	ids := preparePortTypeData(ctx, t, *r)
	prepareSvcData(ctx, t, *r)

	def1 := r.client.EquipmentPortDefinition.GetX(ctx, ids.portDef1)
	typ1 := def1.QueryEquipmentPortType().OnlyX(ctx)
	equip1 := r.client.Equipment.GetX(ctx, ids.equipParentID)
	equip2 := r.client.Equipment.GetX(ctx, ids.equipParent2ID)
	etyp1 := equip1.QueryType().OnlyX(ctx)

	def2 := r.client.EquipmentPortDefinition.GetX(ctx, ids.portDef2)
	typ2 := def2.QueryEquipmentPortType().OnlyX(ctx)
	childEquip := r.client.Equipment.GetX(ctx, ids.equipChildID)
	etyp2 := childEquip.QueryType().OnlyX(ctx)
	var (
		portDataHeader = [...]string{"Port ID", "Port Name", "Port Type", "Equipment Name", "Equipment Type"}
		parentsHeader  = [...]string{"Parent Equipment (3)", "Parent Equipment (2)", "Parent Equipment", "Equipment Position"}
		linkDataHeader = [...]string{"Linked Port ID", "Linked Port Name", "Linked Equipment ID", "Linked Equipment"}
		servicesHeader = [...]string{"Service Names"}
		row1           = []string{strconv.Itoa(ids.parentPortInst1), def1.Name, typ1.Name, equip1.Name, etyp1.Name, locationL, locationM, locationS, "", "", "", "", "", "", "", "", strings.Join([]string{svcName, svc2Name}, ";"), "updateVal", "54"}
		row2           = []string{strconv.Itoa(ids.parentPortInst2), def1.Name, typ1.Name, equip2.Name, etyp1.Name, locationL, locationM, locationS, "", "", "", "", strconv.Itoa(ids.childPortInst1), def2.Name, strconv.Itoa(childEquip.ID), childEquip.Name,
			strings.Join([]string{svcName, svc2Name}, ";"), "updateVal2", "55", "", ""}
		row3 = []string{strconv.Itoa(ids.childPortInst1), def2.Name, typ2.Name, childEquip.Name, etyp2.Name, locationL, locationM, locationS, "", "", equip1.Name, posName, strconv.Itoa(ids.parentPortInst2), def1.Name, strconv.Itoa(equip2.ID), equip2.Name,
			strings.Join([]string{svcName, svc2Name}, ";"), "", "", "1988-01-01", "true"}
	)

	locationTypeInOrder := append(append(append(append(portDataHeader[:], []string{locTypeNameL, locTypeNameM, locTypeNameS}...), parentsHeader[:]...), linkDataHeader[:]...), servicesHeader[:]...)
	titleWithProperties := append(locationTypeInOrder, propNameStr, propNameInt, propNameDate, propNameBool)

	fl, _ := NewImportHeader(titleWithProperties, ImportEntityPort)
	err := importer.inputValidationsPorts(ctx, fl)
	require.NoError(t, err)

	r1, _ := NewImportRecord(row1, fl)
	port1, err := importer.validateLineForExistingPort(ctx, ids.parentPortInst1, r1)
	require.NoError(t, err)
	ptypes, err := importer.validatePropertiesForPortType(ctx, r1, port1.QueryDefinition().QueryEquipmentPortType().OnlyX(ctx), ImportEntityPort)
	require.NoError(t, err)
	require.Len(t, ptypes, 2)
	require.NotEqual(t, ptypes[0].PropertyTypeID, ptypes[1].PropertyTypeID)
	for _, value := range ptypes {
		ptyp := etyp1.QueryPortDefinitions().QueryEquipmentPortType().QueryPropertyTypes().Where(propertytype.ID(value.PropertyTypeID)).OnlyX(ctx)
		switch ptyp.Name {
		case propNameStr:
			require.Equal(t, *value.StringValue, "updateVal")
			require.Equal(t, ptyp.Type, "string")
		case propNameInt:
			require.Equal(t, *value.IntValue, 54)
			require.Equal(t, ptyp.Type, "int")
		default:
			require.Fail(t, "property type name should be one of the two")
		}
	}

	r2, _ := NewImportRecord(row2, fl)

	port2, err := importer.validateLineForExistingPort(ctx, ids.parentPortInst2, r2)
	require.NoError(t, err)
	ptypes2, err := importer.validatePropertiesForPortType(ctx, r2, port2.QueryDefinition().QueryEquipmentPortType().OnlyX(ctx), ImportEntityPort)
	require.NoError(t, err)
	require.Len(t, ptypes2, 2)
	require.NotEqual(t, ptypes2[0].PropertyTypeID, ptypes2[1].PropertyTypeID)
	for _, value := range ptypes2 {
		ptyp := etyp1.QueryPortDefinitions().QueryEquipmentPortType().QueryPropertyTypes().Where(propertytype.ID(value.PropertyTypeID)).OnlyX(ctx)
		switch ptyp.Name {
		case propNameStr:
			require.Equal(t, *value.StringValue, "updateVal2")
			require.Equal(t, ptyp.Type, "string")
		case propNameInt:
			require.Equal(t, *value.IntValue, 55)
			require.Equal(t, ptyp.Type, "int")
		default:
			require.Fail(t, "property type name should be one of the two")
		}
	}

	r3, _ := NewImportRecord(row3, fl)

	port3, err := importer.validateLineForExistingPort(ctx, ids.childPortInst1, r3)
	require.NoError(t, err)
	ptypes3, err := importer.validatePropertiesForPortType(ctx, r3, port3.QueryDefinition().QueryEquipmentPortType().OnlyX(ctx), ImportEntityPort)
	require.NoError(t, err)
	require.Len(t, ptypes3, 2)
	require.NotEqual(t, ptypes3[0].PropertyTypeID, ptypes3[1].PropertyTypeID)
	for _, value := range ptypes3 {
		ptyp := etyp2.QueryPortDefinitions().QueryEquipmentPortType().QueryPropertyTypes().Where(propertytype.ID(value.PropertyTypeID)).OnlyX(ctx)
		switch ptyp.Name {
		case propNameDate:
			require.Equal(t, *value.StringValue, "1988-01-01")
			require.Equal(t, ptyp.Type, "date")
		case propNameBool:
			require.Equal(t, *value.BooleanValue, true)
			require.Equal(t, ptyp.Type, "bool")
		default:
			require.Fail(t, "property type name should be one of the two")
		}
	}
}
