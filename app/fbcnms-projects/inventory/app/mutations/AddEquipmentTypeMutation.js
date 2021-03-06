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
  AddEquipmentTypeMutation,
  AddEquipmentTypeMutationResponse,
  AddEquipmentTypeMutationVariables,
} from './__generated__/AddEquipmentTypeMutation.graphql';
import type {MutationCallbacks} from './MutationCallbacks.js';
import type {SelectorStoreUpdater} from 'relay-runtime';

import RelayEnvironment from '../common/RelayEnvironment.js';
import {commitMutation, graphql} from 'react-relay';

const mutation = graphql`
  mutation AddEquipmentTypeMutation($input: AddEquipmentTypeInput!) {
    addEquipmentType(input: $input) {
      id
      name
      ...EquipmentTypeItem_equipmentType
    }
  }
`;

export default (
  variables: AddEquipmentTypeMutationVariables,
  callbacks?: MutationCallbacks<AddEquipmentTypeMutationResponse>,
  updater?: SelectorStoreUpdater,
) => {
  const {onCompleted, onError} = callbacks ? callbacks : {};
  commitMutation<AddEquipmentTypeMutation>(RelayEnvironment, {
    mutation,
    variables,
    updater,
    onCompleted,
    onError,
  });
};
