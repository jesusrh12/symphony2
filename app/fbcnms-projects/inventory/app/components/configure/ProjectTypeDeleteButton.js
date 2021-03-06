/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow strict-local
 * @format
 */

import type {MutationCallbacks} from '../../mutations/MutationCallbacks.js';
import type {
  RemoveProjectTypeMutationResponse,
  RemoveProjectTypeMutationVariables,
} from '../../mutations/__generated__/RemoveProjectTypeMutation.graphql';
import type {WithAlert} from '@fbcnms/ui/components/Alert/withAlert';
import type {WithSnackbarProps} from 'notistack';

import Button from '@symphony/design-system/components/Button';
import DeleteOutlineIcon from '@material-ui/icons/DeleteOutline';
import FormActionWithPermissions from '../../common/FormActionWithPermissions';
import React from 'react';
import RemoveProjectTypeMutation from '../../mutations/RemoveProjectTypeMutation';
import nullthrows from '@fbcnms/util/nullthrows';
import withAlert from '@fbcnms/ui/components/Alert/withAlert';
import {LogEvents, ServerLogger} from '../../common/LoggingUtils';
import {withSnackbar} from 'notistack';

type Props = $ReadOnly<{|
  className?: string,
  disabled?: boolean,
  tooltip?: ?string,
  projectType: {id: string, name: string},
  ...WithAlert,
  ...WithSnackbarProps,
|}>;

class ProjectTypeDeleteButton extends React.Component<Props> {
  render() {
    return (
      <FormActionWithPermissions
        disabled={this.props.disabled}
        tooltip={this.props.tooltip}
        permissions={{
          entity: 'projectTemplate',
          action: 'delete',
        }}>
        <Button
          className={this.props.className}
          variant="text"
          skin="primary"
          onClick={this.removeProject}>
          <DeleteOutlineIcon />
        </Button>
      </FormActionWithPermissions>
    );
  }

  removeProject = () => {
    ServerLogger.info(LogEvents.DELETE_PROJECT_TYPE_BUTTON_CLICKED, {
      source: 'project_templates',
    });
    const {projectType} = this.props;
    const projectTypeId = projectType.id;
    this.props
      .confirm({
        message: 'Are you sure you want to delete this project template?',
        confirmLabel: 'Delete',
      })
      .then(confirmed => {
        if (!confirmed) {
          return;
        }

        const variables: RemoveProjectTypeMutationVariables = {
          id: nullthrows(projectTypeId),
        };

        const updater = store => {
          store.delete(projectTypeId);
        };

        const callbacks: MutationCallbacks<RemoveProjectTypeMutationResponse> = {
          onCompleted: (response, errors) => {
            if (errors && errors[0]) {
              this.props.alert('Failed removing project template');
            }
          },
          onError: (_error: Error) => {
            this.props.alert('Failed removing project template');
          },
        };

        RemoveProjectTypeMutation(variables, callbacks, updater);
      });
  };
}

export default withAlert(withSnackbar(ProjectTypeDeleteButton));
