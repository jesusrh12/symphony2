/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */

import type {ButtonProps} from '../Button';
import type {SelectMenuProps} from './SelectMenu';

import * as React from 'react';
import ArrowDropDownIcon from '@material-ui/icons/ArrowDropDown';
import BasePopoverTrigger from '../ContexualLayer/BasePopoverTrigger';
import Button from '../Button';
import SelectMenu from './SelectMenu';
import Text from '../Text';
import classNames from 'classnames';
import symphony from '@symphony/design-system/theme/symphony';
import {makeStyles} from '@material-ui/styles';
import {useFormElementContext} from '../Form/FormElementContext';
import {useMemo} from 'react';

const useStyles = makeStyles(() => ({
  root: {
    justifyContent: 'flex-start',
    border: `1px solid ${symphony.palette.D100}`,
    '&$disabled': {
      backgroundColor: symphony.palette.background,
      color: symphony.palette.secondary,
    },
  },
  disabled: {
    '&&': {
      '&&': {
        color: symphony.palette.secondary,
        fill: symphony.palette.secondary,
      },
    },
  },
  formValue: {
    ...symphony.typography.body2,
  },
  selectedValue: {
    fontWeight: 500,
  },
  menu: {
    margin: '8px 0px',
  },
  label: {
    fontWeight: symphony.typography.body2.fontWeight,
  },
}));

type Props<TValue> = $ReadOnly<{|
  className?: string,
  label?: React.Node,
  ...ButtonProps,
  ...SelectMenuProps<TValue>,
|}>;

const INVERTED_TEXT_SKINS = [
  'primary',
  'red',
  'orange',
  'green',
  'darkGray',
  'brightGray',
];

const Select = <TValue>(props: Props<TValue>) => {
  const {
    label,
    className,
    disabled: disabledProp,
    skin = 'regular',
    tooltip,
    useEllipsis = true,
    variant,
    ...selectMenuProps
  } = props;
  const {selectedValue, options} = selectMenuProps;
  const classes = useStyles();
  const {disabled: contextDisabled} = useFormElementContext();
  const disabled = useMemo(
    () => (disabledProp ? disabledProp : contextDisabled),
    [disabledProp, contextDisabled],
  );
  const isInverted = INVERTED_TEXT_SKINS.includes(skin);

  return (
    <BasePopoverTrigger
      popover={<SelectMenu {...selectMenuProps} className={classes.menu} />}>
      {(onShow, _onHide, contextRef) => (
        <Button
          className={classNames(classes.root, className, {
            [classes.disabled]: disabled,
          })}
          ref={contextRef}
          onClick={onShow}
          skin={skin}
          variant={variant}
          disabled={disabled}
          rightIcon={ArrowDropDownIcon}
          rightIconClass={classNames({[classes.disabled]: disabled})}
          tooltip={tooltip}
          useEllipsis={useEllipsis}>
          <Text color={isInverted ? 'light' : 'regular'} variant="body2">
            <span className={classes.label}>{label}</span>
            {selectedValue != null && !!label ? ': ' : null}
            {selectedValue != null ? (
              <span
                className={
                  classNames({
                    [classes.formValue]: !label,
                    [classes.selectedValue]: label,
                    [classes.disabled]: !label && disabled,
                  }) || null
                }>
                {options.find(option => option.value === selectedValue)
                  ?.label ?? ''}
              </span>
            ) : null}
          </Text>
        </Button>
      )}
    </BasePopoverTrigger>
  );
};

export default Select;
