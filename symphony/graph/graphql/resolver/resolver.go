// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package resolver

import (
	"context"
	"net/http"

	"github.com/facebookincubator/symphony/graph/ent"
	"github.com/facebookincubator/symphony/graph/event"
	"github.com/facebookincubator/symphony/graph/graphql/generated"
	"github.com/facebookincubator/symphony/pkg/log"
)

type (
	// Config configures resolver.
	Config struct {
		Client     *ent.Client
		Logger     log.Logger
		Subscriber event.Subscriber
	}

	// Option allows for managing resolver configuration using functional options.
	Option func(*resolver)

	resolver struct {
		logger   log.Logger
		event    struct{ event.Subscriber }
		mutation struct{ transactional bool }
		orc8r    struct{ client *http.Client }
	}
)

// New creates a graphql resolver.
func New(cfg Config, opts ...Option) generated.ResolverRoot {
	r := &resolver{logger: cfg.Logger}
	r.event.Subscriber = cfg.Subscriber
	r.mutation.transactional = true
	for _, opt := range opts {
		opt(r)
	}
	return r
}

// WithTransaction if set to true, will wraps the mutation with transaction.
func WithTransaction(b bool) Option {
	return func(r *resolver) {
		r.mutation.transactional = b
	}
}

// WithOrc8rClient is used to provide orchestrator http client.
func WithOrc8rClient(client *http.Client) Option {
	return func(r *resolver) {
		r.orc8r.client = client
	}
}

func (resolver) ClientFrom(ctx context.Context) *ent.Client {
	client := ent.FromContext(ctx)
	if client == nil {
		panic("no ClientFrom attached to context")
	}
	return client
}

func (r resolver) Equipment() generated.EquipmentResolver {
	return equipmentResolver{r}
}

func (resolver) EquipmentPort() generated.EquipmentPortResolver {
	return equipmentPortResolver{}
}

func (resolver) EquipmentPosition() generated.EquipmentPositionResolver {
	return equipmentPositionResolver{}
}

func (resolver) EquipmentPortDefinition() generated.EquipmentPortDefinitionResolver {
	return equipmentPortDefinitionResolver{}
}

func (resolver) EquipmentPortType() generated.EquipmentPortTypeResolver {
	return equipmentPortTypeResolver{}
}

func (resolver) EquipmentType() generated.EquipmentTypeResolver {
	return equipmentTypeResolver{}
}

func (resolver) File() generated.FileResolver {
	return fileResolver{}
}

func (resolver) User() generated.UserResolver {
	return userResolver{}
}

func (resolver) UsersGroup() generated.UsersGroupResolver {
	return usersGroupResolver{}
}

func (resolver) Link() generated.LinkResolver {
	return linkResolver{}
}

func (resolver) Location() generated.LocationResolver {
	return locationResolver{}
}

func (resolver) LocationType() generated.LocationTypeResolver {
	return locationTypeResolver{}
}

func (resolver) FloorPlan() generated.FloorPlanResolver {
	return floorPlanResolver{}
}

func (r resolver) Mutation() (mr generated.MutationResolver) {
	mr = mutationResolver{r}
	if r.mutation.transactional {
		mr = txResolver{mr}
	}
	return mr
}

func (r resolver) Query() generated.QueryResolver {
	return queryResolver{r}
}

func (r resolver) Subscription() generated.SubscriptionResolver {
	return subscriptionResolver{r}
}

func (resolver) WorkOrder() generated.WorkOrderResolver {
	return workOrderResolver{}
}

func (resolver) WorkOrderType() generated.WorkOrderTypeResolver {
	return workOrderTypeResolver{}
}

func (resolver) WorkOrderDefinition() generated.WorkOrderDefinitionResolver {
	return workOrderDefinitionResolver{}
}

func (r resolver) Survey() generated.SurveyResolver {
	return surveyResolver{}
}

func (r resolver) SurveyQuestion() generated.SurveyQuestionResolver {
	return surveyQuestionResolver{}
}

func (r resolver) SurveyTemplateQuestion() generated.SurveyTemplateQuestionResolver {
	return surveyTemplateQuestionResolver{}
}

func (r resolver) SurveyTemplateCategory() generated.SurveyTemplateCategoryResolver {
	return surveyTemplateCategoryResolver{}
}

func (r resolver) SurveyCellScan() generated.SurveyCellScanResolver {
	return surveyCellScanResolver{}
}

func (r resolver) SurveyWiFiScan() generated.SurveyWiFiScanResolver {
	return surveyWiFiScanResolver{}
}

func (r resolver) PropertyType() generated.PropertyTypeResolver {
	return propertyTypeResolver{}
}

func (r resolver) Property() generated.PropertyResolver {
	return propertyResolver{}
}

func (resolver) Service() generated.ServiceResolver {
	return serviceResolver{}
}

func (r resolver) ServiceType() generated.ServiceTypeResolver {
	return serviceTypeResolver{}
}

func (r resolver) ServiceEndpoint() generated.ServiceEndpointResolver {
	return serviceEndpointResolver{}
}

func (resolver) Project() generated.ProjectResolver {
	return projectResolver{}
}

func (resolver) ProjectType() generated.ProjectTypeResolver {
	return projectTypeResolver{}
}

func (resolver) CheckListCategory() generated.CheckListCategoryResolver {
	return checkListCategoryResolver{}
}

func (resolver) CheckListItem() generated.CheckListItemResolver {
	return checkListItemResolver{}
}

func (resolver) CheckListItemDefinition() generated.CheckListItemDefinitionResolver {
	return checkListItemDefinitionResolver{}
}

func (resolver) ActionsRule() generated.ActionsRuleResolver {
	return actionsRuleResolver{}
}

func (resolver) ActionsRuleAction() generated.ActionsRuleActionResolver {
	return actionsRuleActionResolver{}
}

func (resolver) ActionsRuleFilter() generated.ActionsRuleFilterResolver {
	return actionsRuleFilterResolver{}
}

func (resolver) ActionsTrigger() generated.ActionsTriggerResolver {
	return actionsTriggerResolver{}
}

func (resolver) Viewer() generated.ViewerResolver {
	return viewerResolver{}
}

func (r resolver) ReportFilter() generated.ReportFilterResolver {
	return reportFilterResolver{}
}

func (r resolver) Comment() generated.CommentResolver {
	return commentResolver{}
}

func (r resolver) ServiceEndpointDefinition() generated.ServiceEndpointDefinitionResolver {
	return serviceEndpointTypeResolver{}
}

func (r resolver) PermissionsPolicy() generated.PermissionsPolicyResolver {
	return permissionsPolicyResolver{}
}
