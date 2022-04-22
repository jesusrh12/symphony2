/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow strict-local
 * @format
 */

import * as React from 'react';
import Grid from '@material-ui/core/Grid';
import Text from '@symphony/design-system/components/Text';
import symphony from '@symphony/design-system/theme/symphony';
import {makeStyles} from '@material-ui/styles';

const useStyles = makeStyles(() => ({
  propValue: {
    color: symphony.palette.D300,
  },
}));

type Props = $ReadOnly<{|
  title: React.Node,
  value: string,
|}>;

const LocationDetailsCardProperty = (props: Props) => {
  const {title, value} = props;
  const classes = useStyles();
  return (
    <>
      <Grid item xs={4}>
        <Text variant="subtitle2" weight="regular">
          {title}
        </Text>
      </Grid>
      <Grid item xs={8}>
        <Text
          variant="subtitle2"
          weight="regular"
          className={classes.propValue}>
          {value}
        </Text>
      </Grid>
    </>
  );
};

export default LocationDetailsCardProperty;
