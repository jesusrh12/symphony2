// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package authz

import (
	"context"

	"github.com/facebookincubator/symphony/pkg/authz/models"
	"github.com/facebookincubator/symphony/pkg/ent"
	"github.com/facebookincubator/symphony/pkg/ent/privacy"
	"github.com/facebookincubator/symphony/pkg/ent/user"
	"github.com/facebookincubator/symphony/pkg/viewer"
)

func cudBasedCheck(cud *models.Cud, m ent.Mutation) bool {
	var permission *models.BasicPermissionRule
	switch {
	case m.Op().Is(ent.OpCreate):
		permission = cud.Create
	case m.Op().Is(ent.OpUpdateOne | ent.OpUpdate):
		permission = cud.Update
	case m.Op().Is(ent.OpDeleteOne | ent.OpDelete):
		permission = cud.Delete
	default:
		return false
	}
	return permission.IsAllowed == models.PermissionValueYes
}

func allowOrSkip(r *models.BasicPermissionRule) error {
	if r.IsAllowed == models.PermissionValueYes {
		return privacy.Allow
	}
	return privacy.Skip
}

func allowOrSkipLocations(r *models.LocationPermissionRule, locationTypeID int) error {
	switch r.IsAllowed {
	case models.PermissionValueYes:
		return privacy.Allow
	case models.PermissionValueByCondition:
		for _, typeID := range r.LocationTypeIds {
			if typeID == locationTypeID {
				return privacy.Allow
			}
		}
	}
	return privacy.Skip
}

func privacyDecision(allowed bool) error {
	if allowed {
		return privacy.Allow
	}
	return privacy.Skip
}

func checkWorkforce(ctx context.Context, r *models.WorkforcePermissionRule, workOrderTypeID *int, projectTypeID *int, woOrganizationID *int) bool {
	switch r.IsAllowed {
	case models.PermissionValueYes:
		if woOrganizationID != nil {
			userViewer, ok := viewer.FromContext(ctx).(*viewer.UserViewer)
			if !ok {
				return false
			}

			//  uOrg, err := userViewer.User().Organization(ctx)
			uOrg, err := userViewer.User().QueryOrganization().OnlyID(ctx)
			if err != nil || uOrg == 0 {
				return false
			}

			if uOrg == *woOrganizationID {
				return true
			}
			// return false
		} else {
			return true
		}
	case models.PermissionValueByCondition:
		if workOrderTypeID != nil {
			for _, typeID := range r.WorkOrderTypeIds {
				if typeID == *workOrderTypeID {
					if woOrganizationID != nil {
						for _, typeInterID := range r.OrganizationIds {
							if typeInterID == *woOrganizationID {
								return true
							}
						}
					} else {
						return true
					}
				}
			}
		}
		if projectTypeID != nil {
			for _, typeID := range r.ProjectTypeIds {
				if typeID == *projectTypeID {
					return true
				}
			}
		}
	}
	return false
}

func cudBasedRule(cud *models.Cud, m ent.Mutation) error {
	return privacyDecision(cudBasedCheck(cud, m))
}

func userHasFullPermissions(v viewer.Viewer) bool {
	return v.Role() == user.RoleOwner || v.Role() == user.RoleAdmin
}

func allowWritePermissionsRule() privacy.MutationRule {
	return privacy.MutationRuleFunc(func(ctx context.Context, _ ent.Mutation) error {
		v := viewer.FromContext(ctx)
		return privacyDecision(userHasFullPermissions(v))
	})
}

func allowReadPermissionsRule() privacy.QueryRule {
	return privacy.QueryRuleFunc(func(ctx context.Context, _ ent.Query) error {
		v := viewer.FromContext(ctx)
		return privacyDecision(userHasFullPermissions(v))
	})
}

func denyIfNoPermissionSettingsRule() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		if FromContext(ctx) == nil {
			return privacy.Deny
		}
		return privacy.Skip
	})
}

func allowOrSkipWorkOrder(ctx context.Context, p *models.PermissionSettings, wo *ent.WorkOrder) error {
	switch allowed, err := isViewerWorkOrderOwnerOrAssignee(ctx, wo); {
	case err != nil:
		return privacy.Denyf("cannot check work order viewer relation: %w", err)
	case allowed:
		return privacy.Allow
	}
	workOrderTypeID, err := wo.QueryType().OnlyID(ctx)
	if err != nil {
		return privacy.Denyf("cannot fetch work order type id: %w", err)
	}
	organizationID, err := wo.QueryOrganization().OnlyID(ctx)
	if err != nil || organizationID == 0 {
		return privacyDecision(
			checkWorkforce(
				ctx, p.WorkforcePolicy.Data.Update, &workOrderTypeID, nil, nil,
			),
		)
	}
	return privacyDecision(
		checkWorkforce(
			ctx, p.WorkforcePolicy.Data.Update, &workOrderTypeID, nil, &organizationID,
		),
	)
}

func allowOrSkipProject(ctx context.Context, p *models.PermissionSettings, proj *ent.Project) error {
	if userViewer, ok := viewer.FromContext(ctx).(*viewer.UserViewer); ok {
		switch isCreator, err := isUserProjectCreator(ctx, userViewer.User().ID, proj); {
		case err != nil:
			return privacy.Denyf("cannot check project viewer relation: %w", err)
		case isCreator:
			return privacy.Allow
		}
	}

	projectTypeID, err := proj.QueryType().OnlyID(ctx)
	if err != nil {
		return privacy.Denyf("cannot fetch project type id: %w", err)
	}
	return privacyDecision(
		checkWorkforce(
			ctx, p.WorkforcePolicy.Data.Update, nil, &projectTypeID, nil,
		),
	)
}

func denyBulkEditOrDeleteRule() privacy.MutationRule {
	return privacy.DenyMutationOperationRule(ent.OpUpdate | ent.OpDelete)
}
