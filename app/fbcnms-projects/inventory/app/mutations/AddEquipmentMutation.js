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
  AddEquipmentMutation,
  AddEquipmentMutationResponse,
  AddEquipmentMutationVariables,
} from './__generated__/AddEquipmentMutation.graphql';
import type {MutationCallbacks} from './MutationCallbacks.js';
import type {SelectorStoreUpdater} from 'relay-runtime';

import RelayEnvironment from '../common/RelayEnvironment.js';
import {commitMutation, graphql} from 'react-relay';

const mutation = graphql`
  mutation AddEquipmentMutation($input: AddEquipmentInput!) {
    addEquipment(input: $input) {
      ...EquipmentTable_equipments
    }
  }
`;

export default (
  variables: AddEquipmentMutationVariables,
  callbacks?: MutationCallbacks<AddEquipmentMutationResponse>,
  updater?: SelectorStoreUpdater,
) => {
  const {onCompleted, onError} = callbacks ? callbacks : {};
  commitMutation<AddEquipmentMutation>(RelayEnvironment, {
    mutation,
    variables,
    updater,
    onCompleted,
    onError,
  });
};
