// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package resolver

import (
	"context"

	"github.com/facebookincubator/symphony/graph/ent"
	"github.com/facebookincubator/symphony/graph/graphql/models"
)

type checkListCategoryResolver struct{}

func (checkListCategoryResolver) CheckList(ctx context.Context, obj *ent.CheckListCategory) ([]*ent.CheckListItem, error) {
	return obj.QueryCheckListItems().All(ctx)
}

type checkListItemResolver struct{}

func (checkListItemResolver) Type(ctx context.Context, obj *ent.CheckListItem) (models.CheckListItemType, error) {
	return models.CheckListItemType(obj.Type), nil
}

func (checkListItemResolver) EnumSelectionMode(_ context.Context, item *ent.CheckListItem) (*models.CheckListItemEnumSelectionMode, error) {
	selectionMode := models.CheckListItemEnumSelectionMode(item.EnumSelectionMode)
	return &selectionMode, nil
}

func (checkListItemResolver) Files(ctx context.Context, item *ent.CheckListItem) ([]*ent.File, error) {
	return item.QueryFiles().All(ctx)
}

func (checkListItemResolver) YesNoResponse(ctx context.Context, item *ent.CheckListItem) (*models.YesNoResponse, error) {
	yesNoResponse := models.YesNoResponse(item.YesNoVal)
	if yesNoResponse.IsValid() {
		return &yesNoResponse, nil
	}
	return nil, nil
}

func (checkListItemResolver) WifiData(ctx context.Context, item *ent.CheckListItem) ([]*ent.SurveyWiFiScan, error) {
	return item.QueryWifiScan().All(ctx)
}

func (checkListItemResolver) CellData(ctx context.Context, item *ent.CheckListItem) ([]*ent.SurveyCellScan, error) {
	return item.QueryCellScan().All(ctx)
}

type checkListItemDefinitionResolver struct{}

func (checkListItemDefinitionResolver) Type(ctx context.Context, obj *ent.CheckListItemDefinition) (models.CheckListItemType, error) {
	return models.CheckListItemType(obj.Type), nil
}
