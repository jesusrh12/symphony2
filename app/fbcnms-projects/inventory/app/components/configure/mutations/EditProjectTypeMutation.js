/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */

import type {
  EditProjectTypeMutation,
  EditProjectTypeMutationResponse,
  EditProjectTypeMutationVariables,
} from './__generated__/EditProjectTypeMutation.graphql';
import type {MutationCallbacks} from '../../../mutations/MutationCallbacks.js';
import type {SelectorStoreUpdater} from 'relay-runtime';

import RelayEnvironment from '../../../common/RelayEnvironment.js';
import {commitMutation, graphql} from 'react-relay';

const mutation = graphql`
  mutation EditProjectTypeMutation($input: EditProjectTypeInput!) {
    editProjectType(input: $input) {
      ...ProjectTypeCard_projectType
      ...AddEditProjectTypeCard_editingProjectType
    }
  }
`;

export default (
  variables: EditProjectTypeMutationVariables,
  callbacks?: MutationCallbacks<EditProjectTypeMutationResponse>,
  updater?: SelectorStoreUpdater,
) => {
  const {onCompleted, onError} = callbacks ? callbacks : {};
  commitMutation<EditProjectTypeMutation>(RelayEnvironment, {
    mutation,
    variables,
    updater,
    onCompleted,
    onError,
  });
};
