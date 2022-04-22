// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package resolver_test

import (
	"context"
	"testing"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/facebookincubator/symphony/graph/graphql/models"
	"github.com/facebookincubator/symphony/pkg/ent/user"
	"github.com/facebookincubator/symphony/pkg/viewer"
	"github.com/facebookincubator/symphony/pkg/viewer/viewertest"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func toUserStatusPointer(status user.Status) *user.Status {
	return &status
}

func TestEditUser(t *testing.T) {
	r := newTestResolver(t)
	defer r.Close()
	ctx := viewertest.NewContext(context.Background(), r.client)

	u := viewer.FromContext(ctx).(*viewer.UserViewer).User()
	require.Equal(t, user.StatusActive, u.Status)
	require.Empty(t, u.FirstName)

	mr := r.Mutation()
	u, err := mr.EditUser(ctx, models.EditUserInput{ID: u.ID, Status: toUserStatusPointer(user.StatusDeactivated), FirstName: pointer.ToString("John"), LastName: pointer.ToString("Doe")})
	require.NoError(t, err)
	require.Equal(t, user.StatusDeactivated, u.Status)
	require.Equal(t, "John", u.FirstName)
	require.Equal(t, "Doe", u.LastName)
}

func TestAddAndDeleteProfileImage(t *testing.T) {
	r := newTestResolver(t)
	defer r.Close()
	ctx := viewertest.NewContext(context.Background(), r.client)
	u := viewer.FromContext(ctx).(*viewer.UserViewer).User()

	mr := r.Mutation()
	file1, err := mr.AddImage(ctx, models.AddImageInput{
		EntityType:  models.ImageEntityUser,
		EntityID:    u.ID,
		ImgKey:      uuid.New().String(),
		FileName:    "profile_photo.png",
		FileSize:    123,
		Modified:    time.Now(),
		ContentType: "image/png",
	})
	require.NoError(t, err)
	file, err := u.ProfilePhoto(ctx)
	require.NoError(t, err)
	require.Equal(t, "profile_photo.png", file.Name)

	_, err = mr.DeleteImage(ctx, models.ImageEntityUser, u.ID, file1.ID)
	require.NoError(t, err)

	file, err = u.ProfilePhoto(ctx)
	require.NoError(t, err)
	require.Nil(t, file)
}
