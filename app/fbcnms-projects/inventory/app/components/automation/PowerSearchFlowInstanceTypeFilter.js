/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */

import type {FilterProps} from '../comparison_view/ComparisonViewTypes';

import PowerSearchFilter from '../comparison_view/PowerSearchFilter';
import React, {useState} from 'react';
import Tokenizer from '@fbcnms/ui/components/Tokenizer';
import withSuspense from '../../common/withSuspense';
import {useFlowInstanceTemplateNodes} from '../../common/Flow';

const PowerSearchFlowInstanceTypeFilter = (props: FilterProps) => {
  const [searchEntries, setSearchEntries] = useState([]);
  const [tokens, setTokens] = useState([]);

  const flowInstanceTypes = useFlowInstanceTemplateNodes();

  const {
    value,
    onInputBlurred,
    onValueChanged,
    onRemoveFilter,
    editMode,
  } = props;
  return (
    <PowerSearchFilter
      name="Flow Instance Type"
      operator={value.operator}
      editMode={editMode}
      value={(value.idSet ?? [])
        .map(id => flowInstanceTypes.find(type => type.id === id)?.name)
        .join(', ')}
      onRemoveFilter={onRemoveFilter}
      input={
        <Tokenizer
          searchSource="Options"
          tokens={tokens}
          onEntriesRequested={searchTerm =>
            setSearchEntries(
              flowInstanceTypes
                .filter(type =>
                  type.name.toLowerCase().includes(searchTerm.toLowerCase()),
                )
                .map(type => ({id: type.id, label: type.name})),
            )
          }
          searchEntries={searchEntries}
          onBlur={onInputBlurred}
          onSubmit={onInputBlurred}
          onChange={newEntries => {
            setTokens(newEntries);
            onValueChanged({
              id: value.id,
              key: value.key,
              name: value.name,
              operator: value.operator,
              idSet: newEntries.map(entry => entry.id),
            });
          }}
        />
      }
    />
  );
};

export default withSuspense(PowerSearchFlowInstanceTypeFilter);
