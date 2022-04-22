// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package authz_test

import (
	"context"
	"testing"

	"github.com/facebookincubator/symphony/pkg/authz"
	"github.com/facebookincubator/symphony/pkg/ent/reportfilter"
	"github.com/facebookincubator/symphony/pkg/ent/user"
	"github.com/facebookincubator/symphony/pkg/viewer/viewertest"

	"github.com/stretchr/testify/require"
)

func TestReportFilterCanAlwaysBeWritten(t *testing.T) {
	c := viewertest.NewTestClient(t)
	ctx := viewertest.NewContext(
		context.Background(),
		c,
		viewertest.WithRole(user.RoleUser),
		viewertest.WithPermissions(authz.EmptyPermissions()))
	reportFilter, err := c.ReportFilter.Create().
		SetName("ReportFilter").
		SetEntity(reportfilter.EntityWorkOrder).
		Save(ctx)
	require.NoError(t, err)
	err = c.ReportFilter.UpdateOne(reportFilter).
		SetName("NewReportFilter").
		Exec(ctx)
	require.NoError(t, err)
	err = c.ReportFilter.DeleteOne(reportFilter).
		Exec(ctx)
	require.NoError(t, err)
}
