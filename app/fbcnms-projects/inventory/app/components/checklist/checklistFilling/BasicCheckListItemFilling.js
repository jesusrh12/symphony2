/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */

import type {CheckListItemFillingProps} from './CheckListItemFilling';

import * as React from 'react';
import Button from '@symphony/design-system/components/Button';
import fbt from 'fbt';
import {makeStyles} from '@material-ui/styles';
import {useFormContext} from '../../../common/FormContext';

const useStyles = makeStyles(() => ({
  container: {
    display: 'flex',
    flexDirection: 'row',
  },
}));

const BasicCheckListItemFilling = ({
  item,
  onChange,
}: CheckListItemFillingProps): React.Node => {
  const classes = useStyles();
  const form = useFormContext();

  const updateOnChange = () => {
    if (!onChange) {
      return;
    }
    const modifiedItem = {
      ...item,
      checked: !item.checked,
    };
    onChange(modifiedItem);
  };

  return (
    <div className={classes.container}>
      <Button
        disabled={form.alerts.missingPermissions.detected}
        onClick={updateOnChange}
        variant="text">
        {item.checked
          ? fbt(
              'Mark as Undone',
              'Caption of the simple checkbox item Uncheck button',
            )
          : fbt(
              'Mark as done',
              'Caption of the simple checkbox item Check button',
            )}
      </Button>
    </div>
  );
};

export default BasicCheckListItemFilling;
