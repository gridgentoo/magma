// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package resolver

import (
	"context"
	"fmt"

	"github.com/facebookincubator/symphony/graph/ent"
	"github.com/facebookincubator/symphony/graph/graphql/models"
)

type propertyTypeResolver struct{}

func (propertyTypeResolver) Type(_ context.Context, obj *ent.PropertyType) (models.PropertyKind, error) {
	return models.PropertyKind(obj.Type), nil
}

type propertyResolver struct{}

func (propertyResolver) PropertyType(ctx context.Context, obj *ent.Property) (*ent.PropertyType, error) {
	return obj.QueryType().Only(ctx)
}

func (propertyResolver) NodeValue(ctx context.Context, property *ent.Property) (models.NamedNode, error) {
	propertyType, err := property.QueryType().Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("querying property type %w", err)
	}
	switch propertyType.NodeType {
	case "location":
		l, err := property.QueryLocationValue().Only(ctx)
		return l, ent.MaskNotFound(err)
	case "equipment":
		e, err := property.QueryEquipmentValue().Only(ctx)
		return e, ent.MaskNotFound(err)
	case "service":
		s, err := property.QueryServiceValue().Only(ctx)
		return s, ent.MaskNotFound(err)
	case "work_order":
		s, err := property.QueryWorkOrderValue().Only(ctx)
		return s, ent.MaskNotFound(err)
	default:
		return nil, nil
	}
}
