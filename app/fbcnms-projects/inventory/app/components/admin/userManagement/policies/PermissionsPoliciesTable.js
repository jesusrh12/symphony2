/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */

import type {PermissionsPolicy} from '../data/PermissionsPolicies';
import type {
  TableDesignProps,
  TableRowDataType,
  TableSelectionProps,
} from '@symphony/design-system/components/Table/Table';

import * as React from 'react';
import LockIcon from '@symphony/design-system/icons/Indications/LockIcon';
import Table from '@symphony/design-system/components/Table/Table';
import Text from '@symphony/design-system/components/Text';
import classNames from 'classnames';
import fbt from 'fbt';
import symphony from '@symphony/design-system/theme/symphony';
import {POLICY_TYPES} from '../utils/UserManagementUtils';
import {makeStyles} from '@material-ui/styles';
import {useMemo} from 'react';

export const SYSTEM_DEFAULT_POLICY_PREFIX = fbt('Global Policy', '');

const ALL_USERS = `${fbt('All Users', '')}`;

const useStyles = makeStyles(() => ({
  narrowColumn: {
    width: '70%',
  },
  wideColumn: {
    width: '170%',
  },
  nameCell: {
    display: 'flex',
    alignItems: 'center',
    '&:not($disabled)': {
      fill: symphony.palette.D700,
    },
    '&>:not(:first-child)': {
      marginLeft: '8px',
    },
  },
  defaultPolicyPrefix: {
    textDecoration: 'underline',
    marginRight: '4px',
  },
  disabled: {},
}));

type PolicyTableRow = TableRowDataType<PermissionsPolicy>;

const policy2PolicyTableRow: (
  PermissionsPolicy | PolicyTableRow,
) => PolicyTableRow = policy => ({
  key: policy.id,
  ...policy,
  alwaysShowOnTop: policy.isSystemDefault,
});

const getPolicyUsersCount = (PolicyRow: PolicyTableRow) =>
  PolicyRow.isGlobal ? ALL_USERS : PolicyRow.groups.length;

const getPolicyType = (PolicyRow: PolicyTableRow) => {
  switch (PolicyRow.type) {
    case POLICY_TYPES.InventoryPolicy.key:
      return POLICY_TYPES.InventoryPolicy.value;
    case POLICY_TYPES.WorkforcePolicy.key:
      return POLICY_TYPES.WorkforcePolicy.value;
    default:
      return null;
  }
};

type Props = $ReadOnly<{|
  policies: $ReadOnlyArray<PermissionsPolicy> | $ReadOnlyArray<PolicyTableRow>,
  onPolicySelected?: ?(string) => void,
  showGroupsColumn?: ?boolean,
  ...TableSelectionProps,
  ...TableDesignProps,
|}>;

export default function PermissionsPoliciesTable(props: Props) {
  const {policies, onPolicySelected, showGroupsColumn, ...tableProps} = props;
  const policiesTable = useMemo(() => policies.map(policy2PolicyTableRow), [
    policies,
  ]);

  const classes = useStyles();

  const columns = useMemo(() => {
    const cols = [
      {
        key: 'name',
        title: (
          <fbt desc="Policy Name column header in permission policies table">
            Policy Name
          </fbt>
        ),
        getSortingValue: PolicyRow => PolicyRow.name,
        render: PolicyRow => (
          <div
            className={classNames(classes.nameCell, {
              [classes.disabled]: PolicyRow.disabled,
            })}>
            {PolicyRow.isSystemDefault && <LockIcon color="inherit" />}
            <span>{PolicyRow.name}</span>
          </div>
        ),
        tooltip: PolicyRow => PolicyRow.name,
      },
      {
        key: 'description',
        title: (
          <fbt desc="Description column header in permission policies table">
            Description
          </fbt>
        ),
        getSortingValue: PolicyRow => PolicyRow.description,
        render: PolicyRow => (
          <>
            {PolicyRow.isSystemDefault && (
              <Text
                variant="body2"
                color="inherit"
                className={classes.defaultPolicyPrefix}>
                {SYSTEM_DEFAULT_POLICY_PREFIX}:
              </Text>
            )}
            {PolicyRow.description}
          </>
        ),
        tooltip: PolicyRow => PolicyRow.description,
        titleClassName: classes.wideColumn,
        className: classes.wideColumn,
      },
      {
        key: 'type',
        title: (
          <fbt desc="Policy Type column header in permission policies table">
            Policy Type
          </fbt>
        ),
        getSortingValue: getPolicyType,
        render: getPolicyType,
        tooltip: getPolicyType,
        titleClassName: classes.narrowColumn,
        className: classes.narrowColumn,
      },
    ];

    if (showGroupsColumn !== false) {
      cols.push({
        key: 'groups',
        title: (
          <fbt desc="Groups Applied column header in permission groups table">
            Groups Applied
          </fbt>
        ),
        getSortingValue: getPolicyUsersCount,
        render: getPolicyUsersCount,
        tooltip: getPolicyUsersCount,
        titleClassName: classes.narrowColumn,
        className: classes.narrowColumn,
      });
    }

    return cols;
  }, [
    classes.defaultPolicyPrefix,
    classes.disabled,
    classes.nameCell,
    classes.narrowColumn,
    classes.wideColumn,
    showGroupsColumn,
  ]);

  return (
    <Table
      data={policiesTable}
      onActiveRowIdChanged={
        onPolicySelected != null
          ? policyId =>
              (policyId != null && onPolicySelected(`${policyId}`)) || undefined
          : undefined
      }
      columns={columns}
      {...tableProps}
    />
  );
}
