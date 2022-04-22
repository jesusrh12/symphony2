// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package resolver_test

import (
	"testing"

	"github.com/facebookincubator/symphony/pkg/viewer/viewertest"
	"github.com/stretchr/testify/require"
)

func TestQueryMe(t *testing.T) {
	resolver := newTestResolver(t)
	defer resolver.Close()
	c := resolver.GraphClient()

	var rsp struct {
		Me struct {
			Tenant string
			User   struct {
				AuthID string
			}
		}
	}
	c.MustPost("query { me { tenant, user { authID } } }", &rsp)
	require.Equal(t, viewertest.DefaultTenant, rsp.Me.Tenant)
	require.Equal(t, viewertest.DefaultUser, rsp.Me.User.AuthID)
}
