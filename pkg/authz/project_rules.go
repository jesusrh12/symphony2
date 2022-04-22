// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package authz

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/symphony/pkg/ent/workorderdefinition"

	"github.com/facebookincubator/symphony/pkg/authz/models"
	"github.com/facebookincubator/symphony/pkg/ent"
	"github.com/facebookincubator/symphony/pkg/ent/predicate"
	"github.com/facebookincubator/symphony/pkg/ent/privacy"
	"github.com/facebookincubator/symphony/pkg/ent/project"
	"github.com/facebookincubator/symphony/pkg/ent/projecttype"
	"github.com/facebookincubator/symphony/pkg/ent/user"
	"github.com/facebookincubator/symphony/pkg/viewer"
)

func getProjectType(ctx context.Context, m *ent.ProjectMutation) (*int, error) {
	id, exists := m.ID()
	if !exists {
		return nil, nil
	}
	projectTypeID, err := m.Client().ProjectType.Query().
		Where(projecttype.HasProjectsWith(project.ID(id))).
		OnlyID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch project type id: %w", err)
	}
	return &projectTypeID, nil
}

func projectCudBasedCheck(ctx context.Context, cud *models.WorkforceCud, m *ent.ProjectMutation) (bool, error) {
	if m.Op().Is(ent.OpCreate) {
		typeID, exists := m.TypeID()
		if !exists {
			return false, errors.New("creating project with no type")
		}
		return checkWorkforce(ctx, cud.Create, nil, &typeID, nil), nil
	}
	projectTypeID, err := getProjectType(ctx, m)
	if err != nil {
		return false, err
	}
	if m.Op().Is(ent.OpUpdateOne) {
		return checkWorkforce(ctx, cud.Update, nil, projectTypeID, nil), nil
	}
	return checkWorkforce(ctx, cud.Delete, nil, projectTypeID, nil), nil
}

func projectReadPredicate(ctx context.Context) predicate.Project {
	var predicates []predicate.Project
	rule := FromContext(ctx).WorkforcePolicy.Read
	switch rule.IsAllowed {
	case models.PermissionValueYes:
		return nil
	case models.PermissionValueByCondition:
		predicates = append(predicates,
			project.HasTypeWith(projecttype.IDIn(rule.ProjectTypeIds...)))
	}
	if v, exists := viewer.FromContext(ctx).(*viewer.UserViewer); exists {
		predicates = append(predicates,
			project.HasCreatorWith(user.ID(v.User().ID)),
		)
	}
	if woPredicate := workOrderReadPredicate(ctx); woPredicate != nil {
		predicates = append(predicates,
			project.HasWorkOrdersWith(woPredicate))
	}
	return project.Or(predicates...)
}

func isUserProjectCreator(ctx context.Context, userID int, project *ent.Project) (bool, error) {
	creatorID, err := project.QueryCreator().OnlyID(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return false, fmt.Errorf("failed to fetch project creator: %w", err)
		}
		return false, nil
	}
	return creatorID == userID, nil
}

func isCreatorChanged(ctx context.Context, m *ent.ProjectMutation) (bool, error) {
	var currCreatorID *int
	creatorIDToSet, created := m.CreatorID()
	creatorCleared := m.CreatorCleared()
	if !created && !creatorCleared {
		return false, nil
	}
	projectID, exists := m.ID()
	if !exists {
		return created, nil
	}
	creatorID, err := m.Client().User.Query().
		Where(user.HasCreatedProjectsWith(project.ID(projectID))).
		OnlyID(ctx)
	if err == nil {
		currCreatorID = &creatorID
	}
	if err != nil && !ent.IsNotFound(err) {
		return false, privacy.Denyf("failed to fetch creator: %w", err)
	}
	switch {
	case currCreatorID == nil && created:
		return true, nil
	case currCreatorID != nil && created && *currCreatorID != creatorIDToSet:
		return true, nil
	case currCreatorID != nil && creatorCleared:
		return true, nil
	}
	return false, nil
}

func allowOrSkipWorkOrderDefinition(ctx context.Context, client *ent.Client, workOrderDefinitionID int) error {
	workOrderDefinition, err := client.WorkOrderDefinition.Query().
		Where(workorderdefinition.ID(workOrderDefinitionID)).
		WithProjectTemplate().
		WithProjectType().
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return privacy.Denyf("failed to query work order definition: %w", err)
		}
		return privacy.Skip
	}
	switch {
	case workOrderDefinition.Edges.ProjectTemplate != nil:
		return privacy.Allow
	case workOrderDefinition.Edges.ProjectType != nil:
		return allowOrSkip(FromContext(ctx).WorkforcePolicy.Templates.Update)
	}
	return privacy.Skip
}

// ProjectWritePolicyRule grants write permission to project based on policy.
func ProjectWritePolicyRule() privacy.MutationRule {
	return privacy.ProjectMutationRuleFunc(func(ctx context.Context, m *ent.ProjectMutation) error {
		cud := FromContext(ctx).WorkforcePolicy.Data
		allowed, err := projectCudBasedCheck(ctx, cud, m)
		if err != nil {
			return privacy.Denyf(err.Error())
		}
		if !m.Op().Is(ent.OpCreate) {
			creatorChanged, err := isCreatorChanged(ctx, m)
			if err != nil {
				return privacy.Denyf(err.Error())
			}
			if creatorChanged {
				projectTypeID, err := getProjectType(ctx, m)
				if err != nil {
					return err
				}
				allowed = allowed && checkWorkforce(ctx, cud.TransferOwnership, nil, projectTypeID, nil)
			}
		}
		if allowed {
			return privacy.Allow
		}
		return privacy.Skip
	})
}

// AllowProjectCreatorWrite grants write permission if user is creator of project
func AllowProjectCreatorWrite() privacy.MutationRule {
	return privacy.ProjectMutationRuleFunc(func(ctx context.Context, m *ent.ProjectMutation) error {
		if m.Op().Is(ent.OpDeleteOne) {
			return privacy.Skip
		}
		projectID, exists := m.ID()
		if !exists {
			return privacy.Skip
		}
		userViewer, ok := viewer.FromContext(ctx).(*viewer.UserViewer)
		if !ok {
			return privacy.Skip
		}
		proj, err := m.Client().Project.Get(ctx, projectID)
		if err != nil {
			if !ent.IsNotFound(err) {
				return privacy.Denyf("failed to fetch project: %w", err)
			}
			return privacy.Skip
		}

		isCreator, err := isUserProjectCreator(ctx, userViewer.User().ID, proj)
		if err != nil {
			return privacy.Denyf(err.Error())
		}
		if isCreator {
			return privacy.Allow
		}
		return privacy.Skip
	})
}

// ProjectReadPolicyRule grants read permission to project based on policy.
func ProjectReadPolicyRule() privacy.QueryRule {
	return privacy.ProjectQueryRuleFunc(func(ctx context.Context, q *ent.ProjectQuery) error {
		projectPredicate := projectReadPredicate(ctx)
		if projectPredicate != nil {
			q.Where(projectPredicate)
		}
		return privacy.Skip
	})
}

// ProjectTypeWritePolicyRule grants write permission to project type based on policy.
func ProjectTypeWritePolicyRule() privacy.MutationRule {
	return privacy.MutationRuleFunc(func(ctx context.Context, m ent.Mutation) error {
		return cudBasedRule(FromContext(ctx).WorkforcePolicy.Templates, m)
	})
}

// WorkOrderDefinitionWritePolicyRule grants write permission to work order definition based on policy.
func WorkOrderDefinitionWritePolicyRule() privacy.MutationRule {
	return privacy.WorkOrderDefinitionMutationRuleFunc(func(ctx context.Context, m *ent.WorkOrderDefinitionMutation) error {
		if m.Op().Is(ent.OpCreate) {
			if _, exists := m.ProjectTemplateID(); exists {
				return privacy.Allow
			}
			if _, exists := m.ProjectTypeID(); exists {
				return allowOrSkip(FromContext(ctx).WorkforcePolicy.Templates.Update)
			}
		} else {
			workOrderDefinitionID, exists := m.ID()
			if !exists {
				return privacy.Skip
			}
			return allowOrSkipWorkOrderDefinition(ctx, m.Client(), workOrderDefinitionID)
		}
		return privacy.Skip
	})
}
