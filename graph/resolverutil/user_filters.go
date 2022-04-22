// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package resolverutil

import (
	"strings"

	"github.com/facebookincubator/symphony/graph/graphql/models"
	"github.com/facebookincubator/symphony/pkg/ent"
	"github.com/facebookincubator/symphony/pkg/ent/organization"
	"github.com/facebookincubator/symphony/pkg/ent/predicate"
	"github.com/facebookincubator/symphony/pkg/ent/schema/enum"
	"github.com/facebookincubator/symphony/pkg/ent/user"
	"github.com/pkg/errors"
)

func handleUserFilter(q *ent.UserQuery, filter *models.UserFilterInput) (*ent.UserQuery, error) {
	switch filter.FilterType {
	case models.UserFilterTypeUserName:
		return userNameFilter(q, filter)
	case models.UserFilterTypeUserStatus:
		return userStatusFilter(q, filter)
	case models.UserFilterTypeUserOrganization:
		return userOrganizationFilter(q, filter)
	default:
		return nil, errors.Errorf("filter type is not supported: %s", filter.FilterType)
	}
}

func userOrganizationFilter(q *ent.UserQuery, filter *models.UserFilterInput) (*ent.UserQuery, error) {
	if filter.Operator == enum.FilterOperatorIsOneOf && filter.IDSet != nil {
		return q.Where(user.HasOrganizationWith(organization.IDIn(filter.IDSet...))), nil
	}
	return nil, errors.Errorf("operation is not supported: %s", filter.Operator)
}

func userStatusFilter(q *ent.UserQuery, filter *models.UserFilterInput) (*ent.UserQuery, error) {
	if filter.Operator == enum.FilterOperatorIs {
		return q.Where(user.StatusEQ(*filter.StatusValue)), nil
	}
	return nil, errors.Errorf("operation %q not supported", filter.Operator)
}

func userNameFilter(q *ent.UserQuery, filter *models.UserFilterInput) (*ent.UserQuery, error) {
	if filter.Operator == enum.FilterOperatorContains {
		terms := strings.Split(*filter.StringValue, " ")
		qp := user.And()
		for _, s := range terms {
			qp = user.And(qp, userStringPredicate(s))
		}
		return q.Where(qp), nil
	}
	return nil, errors.Errorf("operation %q not supported", filter.Operator)
}

func userStringPredicate(s string) predicate.User {
	return user.Or(user.EmailContainsFold(s),
		user.FirstNameContainsFold(s),
		user.LastNameContainsFold(s))
}
