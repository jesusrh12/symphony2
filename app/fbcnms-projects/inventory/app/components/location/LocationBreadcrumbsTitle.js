/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */

import type {ContextRouter} from 'react-router-dom';
import type {Location} from '../../common/Location';

import * as React from 'react';
import Breadcrumbs from '@fbcnms/ui/components/Breadcrumbs';
import {InventoryAPIUrls} from '../../common/InventoryAPI';
import {LogEvents, ServerLogger} from '../../common/LoggingUtils';
import {createFragmentContainer, graphql} from 'react-relay';
import {withRouter} from 'react-router-dom';

export const NAVIGATION_OPTIONS = {
  NEW_TAB: 'NEW_TAB',
  SAME_TAB: 'SAME_TAB',
  NONE: 'NONE',
};

type NavigationOption = $Keys<typeof NAVIGATION_OPTIONS>;

type Props = ContextRouter & {|
  locationDetails: Location,
  hideTypes?: boolean,
  navigateOnClick?: ?NavigationOption,
  size?: 'default' | 'small' | 'large',
|};

const LocationBreadcrumbsTitle = (props: Props) => {
  const {
    locationDetails,
    hideTypes,
    size,
    navigateOnClick = NAVIGATION_OPTIONS.SAME_TAB,
  } = props;

  const navigateToLocation = React.useCallback(
    (selectedLocationId: string) => {
      ServerLogger.info(LogEvents.NAVIGATE_TO_LOCATION, {
        locationId: selectedLocationId,
      });
      if (navigateOnClick == NAVIGATION_OPTIONS.SAME_TAB) {
        props.history.push(InventoryAPIUrls.location(selectedLocationId));
      } else if (navigateOnClick == NAVIGATION_OPTIONS.NEW_TAB) {
        window.open(InventoryAPIUrls.location(selectedLocationId));
      }
    },
    [props.history],
  );

  const onBreadcrumbClicked = React.useCallback(
    id => {
      ServerLogger.info(LogEvents.LOCATION_CARD_BREADCRUMB_CLICKED, {
        locationId: id,
      });
      if (id && navigateOnClick != NAVIGATION_OPTIONS.NONE) {
        navigateToLocation(id);
      }
    },
    [navigateOnClick, navigateToLocation],
  );

  return (
    <Breadcrumbs
      breadcrumbs={[...locationDetails.locationHierarchy, locationDetails].map(
        l => ({
          id: l.id,
          name: l.name,
          subtext: hideTypes ? null : l.locationType.name,
          onClick: () => onBreadcrumbClicked(l.id),
        }),
      )}
      size={size}
    />
  );
};

LocationBreadcrumbsTitle.defaultProps = {
  size: 'default',
};

export default withRouter(
  createFragmentContainer(LocationBreadcrumbsTitle, {
    locationDetails: graphql`
      fragment LocationBreadcrumbsTitle_locationDetails on Location {
        id
        name
        locationType {
          name
        }
        locationHierarchy {
          id
          name
          locationType {
            name
          }
        }
      }
    `,
  }),
);
