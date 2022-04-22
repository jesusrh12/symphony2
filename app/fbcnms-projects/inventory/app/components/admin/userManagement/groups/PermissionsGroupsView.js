/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */

import type {TableRowDataType} from '@symphony/design-system/components/Table/Table';
import type {UsersGroup} from '../data/UsersGroups';

import * as React from 'react';
import Table from '@symphony/design-system/components/Table/Table';
import fbt from 'fbt';
import withSuspense from '../../../../common/withSuspense';
import {GROUP_STATUSES} from '../utils/UserManagementUtils';
import {makeStyles} from '@material-ui/styles';
import {useMemo} from 'react';
import {useRouter} from '@fbcnms/ui/hooks';
import {useUsersGroups} from '../data/UsersGroups';

export const PERMISSION_GROUPS_VIEW_NAME = fbt(
  'Permission Groups',
  'Header for view showing system permissions groups settings',
);

export const PERMISSION_GROUPS_VIEW_SUBHEADER = fbt(
  "Create groups of users, choose which data they have access to, and what theyֿ're allowed to do with it.",
  'Subheader for view showing system permissions groups settings',
);

const useStyles = makeStyles(() => ({
  root: {
    flexGrow: 1,
  },
  narrowColumn: {
    width: '70%',
  },
  wideColumn: {
    width: '170%',
  },
}));

type GroupTableRow = TableRowDataType<UsersGroup>;

const group2GroupTableRow: (
  UsersGroup | GroupTableRow,
) => GroupTableRow = group => ({
  key: group.key || group.id,
  ...group,
});

function PermissionsGroupsView() {
  const classes = useStyles();
  const {history} = useRouter();
  const groups = useUsersGroups();
  const groupsTable = useMemo(() => groups.map(group2GroupTableRow), [groups]);

  const columns = [
    {
      key: 'name',
      title: (
        <fbt desc="Group Name column header in permission groups table">
          Group Name
        </fbt>
      ),
      getSortingValue: groupRow => groupRow.name,
      render: groupRow => groupRow.name,
      tooltip: groupRow => groupRow.name,
    },
    {
      key: 'description',
      title: (
        <fbt desc="Description column header in permission groups table">
          Description
        </fbt>
      ),
      getSortingValue: groupRow => groupRow.description,
      render: groupRow => groupRow.description || '',
      tooltip: groupRow => groupRow.description,
      titleClassName: classes.wideColumn,
      className: classes.wideColumn,
    },
    {
      key: 'members',
      title: (
        <fbt desc="Members column header in permission groups table">
          Members
        </fbt>
      ),
      getSortingValue: groupRow => groupRow.members.length,
      render: groupRow => groupRow.members.length,
      tooltip: groupRow => groupRow.members.length,
      titleClassName: classes.narrowColumn,
      className: classes.narrowColumn,
    },
    {
      key: 'status',
      title: (
        <fbt desc="Status column header in permission groups table">Status</fbt>
      ),
      getSortingValue: groupRow => GROUP_STATUSES[groupRow.status].value,
      render: groupRow => GROUP_STATUSES[groupRow.status].value,
      tooltip: groupRow => GROUP_STATUSES[groupRow.status].value,
      titleClassName: classes.narrowColumn,
      className: classes.narrowColumn,
    },
  ];

  return (
    <div className={classes.root}>
      <Table
        data={groupsTable}
        onActiveRowIdChanged={groupId => {
          if (groupId != null) {
            history.push(`group/${groupId}`);
          }
        }}
        columns={columns}
      />
    </div>
  );
}

export default withSuspense(PermissionsGroupsView);
