// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package resolverutil

import (
	"github.com/facebookincubator/symphony/graph/graphql/models"
	"github.com/facebookincubator/symphony/pkg/ent"
	"github.com/facebookincubator/symphony/pkg/ent/schema/enum"
	"github.com/facebookincubator/symphony/pkg/ent/usersgroup"

	"github.com/pkg/errors"
)

func handleUsersGroupFilter(q *ent.UsersGroupQuery, filter *models.UsersGroupFilterInput) (*ent.UsersGroupQuery, error) {
	if filter.FilterType == models.UsersGroupFilterTypeGroupName {
		return usersGroupFilter(q, filter)
	}
	return nil, errors.Errorf("filter type is not supported: %s", filter.FilterType)
}

func usersGroupFilter(q *ent.UsersGroupQuery, filter *models.UsersGroupFilterInput) (*ent.UsersGroupQuery, error) {
	switch {
	case filter.Operator == enum.FilterOperatorIs && filter.StringValue != nil:
		return q.Where(usersgroup.NameEqualFold(*filter.StringValue)), nil
	case filter.Operator == enum.FilterOperatorContains && filter.StringValue != nil:
		return q.Where(usersgroup.NameContainsFold(*filter.StringValue)), nil
	}
	return nil, errors.Errorf("operation %s is not supported with value of %#v", filter.Operator, filter.StringValue)
}
