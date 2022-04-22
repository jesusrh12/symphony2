/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */

import type {FiltersQuery} from './ComparisonViewTypes';

import ComparisonViewNoResults from './ComparisonViewNoResults';
import InventoryQueryRenderer from '../InventoryQueryRenderer';
import PowerSearchLinksResultsTable from './PowerSearchLinksResultsTable';
import React from 'react';

import {graphql} from 'relay-runtime';
import {makeStyles} from '@material-ui/styles';

const useStyles = makeStyles(_theme => ({
  searchResults: {
    flexGrow: 1,
  },
}));

type Props = $ReadOnly<{|
  filters: FiltersQuery,
  limit?: number,
  onQueryReturn: number => void,
|}>;

const linkSearchQuery = graphql`
  query LinkViewQueryRendererSearchQuery(
    $limit: Int
    $filters: [LinkFilterInput!]!
  ) {
    links(first: $limit, filterBy: $filters) {
      edges {
        node {
          ...PowerSearchLinksResultsTable_links
        }
      }
      totalCount
    }
  }
`;

const LinkViewQueryRenderer = (props: Props) => {
  const classes = useStyles();
  const {limit, filters, onQueryReturn} = props;

  return (
    <InventoryQueryRenderer
      query={linkSearchQuery}
      variables={{
        limit: limit,
        filters: filters.map(f => ({
          filterType: f.name.toUpperCase(),
          operator: f.operator.toUpperCase(),
          stringValue: f.stringValue,
          propertyValue: f.propertyValue,
          idSet: f.idSet,
          stringSet: f.stringSet,
        })),
      }}
      render={props => {
        const {totalCount, edges} = props.links;
        onQueryReturn(totalCount);
        if (totalCount === 0) {
          return <ComparisonViewNoResults />;
        }
        return (
          <div className={classes.searchResults}>
            <PowerSearchLinksResultsTable
              links={edges.map(edge => edge.node)}
            />
          </div>
        );
      }}
    />
  );
};

export default LinkViewQueryRenderer;
