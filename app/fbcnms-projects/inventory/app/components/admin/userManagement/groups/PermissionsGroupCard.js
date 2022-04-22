/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow strict-local
 * @format
 */

import type {UsersGroup} from '../data/UsersGroups';
import type {WithAlert} from '@fbcnms/ui/components/Alert/withAlert';

import * as React from 'react';
import Breadcrumbs from '@fbcnms/ui/components/Breadcrumbs';
import Button from '@symphony/design-system/components/Button';
import DeleteIcon from '@symphony/design-system/icons/Actions/DeleteIcon';
import FormAction from '@symphony/design-system/components/Form/FormAction';
import Grid from '@material-ui/core/Grid';
import IconButton from '@symphony/design-system/components/IconButton';
import InventoryErrorBoundary from '../../../../common/InventoryErrorBoundary';
import PermissionsGroupDetailsPane from './PermissionsGroupDetailsPane';
import PermissionsGroupMembersPane from './PermissionsGroupMembersPane';
import PermissionsGroupPoliciesPane from './PermissionsGroupPoliciesPane';
import Strings from '@fbcnms/strings/Strings';
import ViewContainer from '@symphony/design-system/components/View/ViewContainer';
import classNames from 'classnames';
import fbt from 'fbt';
import withAlert from '@fbcnms/ui/components/Alert/withAlert';
import withSuspense from '../../../../common/withSuspense';
import {GROUP_STATUSES, NEW_DIALOG_PARAM} from '../utils/UserManagementUtils';
import {PERMISSION_GROUPS_VIEW_NAME} from './PermissionsGroupsView';
import {addGroup, deleteGroup, editGroup} from '../data/UsersGroups';
import {generateTempId} from '../../../../common/EntUtils';
import {makeStyles} from '@material-ui/styles';
import {useCallback, useEffect, useMemo, useState} from 'react';
import {useEnqueueSnackbar} from '@fbcnms/ui/hooks/useSnackbar';
import {useRouteMatch} from 'react-router-dom';
import {useUsersGroup} from '../data/UsersGroups';

const useStyles = makeStyles(() => ({
  container: {
    maxHeight: '100%',
  },
  vertical: {
    '&>:not(:first-child)': {
      marginTop: '16px',
    },
  },
}));

type Props = $ReadOnly<{|
  redirectToGroupsView: () => void,
  onClose: () => void,
  ...WithAlert,
|}>;

const initialNewGroup: UsersGroup = {
  id: generateTempId(),
  name: '',
  description: '',
  status: GROUP_STATUSES.ACTIVE.key,
  members: [],
  policies: [],
};

function PermissionsGroupCard(props: Props) {
  const {redirectToGroupsView, onClose} = props;
  const classes = useStyles();
  const match = useRouteMatch();

  const groupId = match.params.id;
  const fetchedGroup = useUsersGroup(groupId || '');

  const isOnNewGroup = groupId === NEW_DIALOG_PARAM;
  const [group, setGroup] = useState<?UsersGroup>(
    isOnNewGroup ? {...initialNewGroup} : null,
  );

  const enqueueSnackbar = useEnqueueSnackbar();
  const handleError = useCallback(
    (error: string) => {
      enqueueSnackbar(error, {variant: 'error'});
    },
    [enqueueSnackbar],
  );

  useEffect(() => {
    if (isOnNewGroup || group != null) {
      return;
    }
    if (fetchedGroup == null) {
      if (groupId != null) {
        handleError(
          `${fbt(
            `Group with id ${fbt.param(
              'group id url param',
              groupId,
            )} does not exist.`,
            '',
          )}`,
        );
      }
      redirectToGroupsView();
    }
    setGroup(fetchedGroup);
  }, [
    fetchedGroup,
    group,
    groupId,
    handleError,
    isOnNewGroup,
    redirectToGroupsView,
  ]);

  const header = useMemo(() => {
    const breadcrumbs = [
      {
        id: 'groups',
        name: `${PERMISSION_GROUPS_VIEW_NAME}`,
        onClick: redirectToGroupsView,
      },
      {
        id: 'groupName',
        name: isOnNewGroup ? `${fbt('New Group', '')}` : group?.name || '',
      },
    ];
    const actions = [
      <FormAction ignorePermissions={true}>
        <Button skin="regular" onClick={onClose}>
          {Strings.common.cancelButton}
        </Button>
      </FormAction>,
      <FormAction disableOnFromError={true}>
        <Button
          onClick={() => {
            if (group == null) {
              return;
            }
            const saveAction = isOnNewGroup ? addGroup : editGroup;
            saveAction(group).then(onClose).catch(handleError);
          }}>
          {Strings.common.saveButton}
        </Button>
      </FormAction>,
    ];
    if (!isOnNewGroup) {
      actions.unshift(
        <FormAction>
          <IconButton
            skin="gray"
            icon={DeleteIcon}
            onClick={() => {
              if (group == null) {
                return;
              }
              props
                .confirm(
                  <fbt desc="">
                    Are you sure you want to delete this group?
                  </fbt>,
                )
                .then(confirm => {
                  if (!confirm) {
                    return;
                  }
                  return deleteGroup(group.id).then(onClose);
                })
                .catch(handleError);
            }}
          />
        </FormAction>,
      );
    }
    return {
      title: <Breadcrumbs breadcrumbs={breadcrumbs} />,
      subtitle: fbt('Manage group details, members and policies', ''),
      actionButtons: actions,
    };
  }, [group, handleError, isOnNewGroup, onClose, props, redirectToGroupsView]);

  if (group == null) {
    return null;
  }
  return (
    <InventoryErrorBoundary>
      <ViewContainer header={header} useBodyScrollingEffect={false}>
        <Grid container spacing={2} className={classes.container}>
          <Grid
            item
            xs={8}
            className={classNames(classes.container, classes.vertical)}>
            <PermissionsGroupDetailsPane group={group} onChange={setGroup} />
            <PermissionsGroupPoliciesPane group={group} onChange={setGroup} />
          </Grid>
          <Grid item xs={4} className={classes.container}>
            <PermissionsGroupMembersPane group={group} onChange={setGroup} />
          </Grid>
        </Grid>
      </ViewContainer>
    </InventoryErrorBoundary>
  );
}

export default withSuspense(withAlert(PermissionsGroupCard));
