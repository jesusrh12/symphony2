/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */

import type {RowsSeparationTypes} from './TableContent';
import type {SelectionType} from '../Checkbox/Checkbox';
import type {TableHeaderData} from './TableHeader';
import type {TableSettings, TableSortSettings} from './TableContext';

import * as React from 'react';
import TableContent from './TableContent';
import TableFooter from './TableFooter';
import TableHeader from './TableHeader';
import classNames from 'classnames';
import symphony from '@symphony/design-system/theme/symphony';
import useVerticalScrollingEffect from '../../hooks/useVerticalScrollingEffect';
import {AutoSizer} from 'react-virtualized';
import {TableContextProvider} from './TableContext';
import {TablePaginationContextProvider} from './TablePaginationContext';
import {TableSelectionContextProvider} from './TableSelectionContext';
import {makeStyles} from '@material-ui/styles';
import {useEffect, useMemo, useRef, useState} from 'react';

const borderRadius = '4px';
const useStyles = makeStyles(() => ({
  root: {
    display: 'flex',
    maxHeight: '100%',
  },
  standalone: {
    borderRadius: borderRadius,
    boxShadow: symphony.shadows.DP1,
    height: 'fit-content',
  },
  tableContainer: {
    display: 'flex',
    maxHeight: '100%',
    overflow: 'hidden',
    flexGrow: 0,
    flexBasis: '25%',
    borderRadius: borderRadius,
    '&$expanded': {
      flexGrow: 1,
    },
  },
  expanded: {},
  table: {
    display: 'flex',
    flexDirection: 'column',
    borderCollapse: 'collapse',
    overflow: 'hidden',
    '& tbody': {
      borderTop: `1px solid ${symphony.palette.separatorLight}`,
      overflowX: 'hidden',
      overflowY: 'auto',
    },
    '& tr': {
      tableLayout: 'fixed',
      display: 'table',
      width: '100%',
    },
  },
  embedded: {
    '& $cell': {
      '&:first-child': {
        paddingLeft: '0px',
      },
      '&:last-child': {
        paddingRight: '0px',
      },
    },
  },
  cell: {},
  detailsCardContainer: {
    backgroundColor: symphony.palette.white,
    borderLeft: '1px solid',
    borderColor: symphony.palette.separatorLight,
    flexBasis: '10px',
    flexGrow: 1,
    borderTopRightRadius: borderRadius,
    borderBottomRightRadius: borderRadius,
  },
}));

export type TableRowDataType<T> = $ReadOnly<{|
  key?: string,
  alwaysShowOnTop?: ?boolean,
  className?: ?string,
  disabled?: ?boolean,
  tooltip?: ?string,
  ...T,
|}>;

export type TableSelectionType = 'all' | 'none' | 'single_item_toggled';

export type TableRowId = string | number;
export type NullableTableRowId = TableRowId | null;

export type SelectionCallbackType = (
  selectedIds: $ReadOnlyArray<TableRowId>,
  selection: TableSelectionType,
  toggledItem?: ?{id: TableRowId, change: SelectionType},
) => void;
export type ActiveCallbackType = (activeId: NullableTableRowId) => void;

export const TABLE_VARIANT_TYPES = {
  standalone: 'standalone',
  embedded: 'embedded',
};
export type TableVariantTypes = $Keys<typeof TABLE_VARIANT_TYPES>;

export type TableDesignProps = $ReadOnly<{|
  showSelection?: boolean,
  className?: string,
  variant?: TableVariantTypes,
  dataRowsSeparator?: RowsSeparationTypes,
  dataRowClassName?: string,
|}>;

export type TableSelectionProps = $ReadOnly<{|
  selectedIds?: $ReadOnlyArray<TableRowId>,
  onSelectionChanged?: SelectionCallbackType,
|}>;

export type TablePaginationProps = $ReadOnly<{|
  loadNext?: ?(onCompleted: ?() => void) => void,
  pageSize: number,
  totalRowsCount: number,
|}>;

/*
  detailsCard:
    When passed, will be shown on as part of the table content.
    Excepts for the first column, all columns will get hidden.
    The card will cover 75% of the table width.
*/
type Props<T> = $ReadOnly<{|
  ...TableHeaderData<T>,
  data: $ReadOnlyArray<TableRowDataType<T>>,
  sortSettings?: ?TableSortSettings,
  activeRowId?: NullableTableRowId,
  onActiveRowIdChanged?: ActiveCallbackType,
  detailsCard?: ?React.Node,
  resizableColumns?: boolean,
  paginationSettings?: ?TablePaginationProps,
  ...TableDesignProps,
  ...TableSelectionProps,
|}>;

const Table = <T>(props: Props<T>) => {
  const {
    className,
    variant = TABLE_VARIANT_TYPES.standalone,
    data,
    showSelection,
    activeRowId,
    onActiveRowIdChanged,
    selectedIds = [],
    onSelectionChanged,
    columns,
    sortSettings: propSortSettings,
    onSortChanged,
    dataRowClassName,
    dataRowsSeparator,
    detailsCard,
    resizableColumns = true,
    paginationSettings,
  } = props;
  const classes = useStyles();
  const [dataColumns, setDataColumns] = useState([]);
  useEffect(() => {
    if (detailsCard == null) {
      setDataColumns(columns);
      return;
    }
    let singleColumnToBeShown = columns.findIndex(col => !col.hidden);
    if (singleColumnToBeShown == -1) {
      singleColumnToBeShown = 0;
    }
    setDataColumns(
      columns.map((col, index) => {
        if (index === singleColumnToBeShown) {
          return {
            ...col,
          };
        }
        return {
          ...col,
          hidden: true,
        };
      }),
    );
  }, [detailsCard, columns]);

  const [tableHeaderPaddingRight, setTableHeaderPaddingRight] = useState(0);
  const bodyRef = useRef(null);
  useVerticalScrollingEffect(
    bodyRef,
    scrollArgs => setTableHeaderPaddingRight(scrollArgs.scrollbarWidth),
    false,
  );

  const renderChildren = (width: number) => (
    <div
      className={classNames(classes.root, classes[variant], className)}
      style={{width}}>
      <div
        className={classNames(classes.tableContainer, {
          [classes.expanded]: !detailsCard,
        })}>
        <table className={classes.table}>
          <TableHeader
            columns={dataColumns}
            onSortChanged={onSortChanged}
            cellClassName={classes.cell}
            paddingRight={tableHeaderPaddingRight}
          />
          <TableContent
            columns={dataColumns}
            data={data}
            dataRowClassName={dataRowClassName}
            rowsSeparator={dataRowsSeparator}
            cellClassName={classes.cell}
            fwdRef={bodyRef}
          />
          {paginationSettings != null ? (
            <TableFooter {...paginationSettings} />
          ) : null}
        </table>
      </div>
      {detailsCard ? (
        <div className={classes.detailsCardContainer}>{detailsCard}</div>
      ) : null}
    </div>
  );

  const allIds = useMemo(
    () => data.filter(d => !d.disabled).map((d, i) => d.key ?? i),
    [data],
  );
  const contextValue: TableSettings = useMemo(
    () => ({
      showSelection: showSelection ?? false,
      clickableRows: !!onActiveRowIdChanged,
      sort: propSortSettings,
      resizableColumns,
    }),
    [showSelection, onActiveRowIdChanged, propSortSettings, resizableColumns],
  );

  return (
    <AutoSizer disableHeight={true}>
      {({width}: {width: number}) => (
        <TableContextProvider
          width={width}
          settings={contextValue}
          columns={columns}>
          <TablePaginationContextProvider
            {...(paginationSettings ?? {
              pageSize: data.length,
              totalRowsCount: data.length,
            })}
            loadedDataCount={data.length}>
            {contextValue.showSelection || contextValue.clickableRows ? (
              <TableSelectionContextProvider
                allIds={allIds}
                activeId={activeRowId}
                onActiveChanged={onActiveRowIdChanged}
                selectedIds={selectedIds ?? []}
                onSelectionChanged={onSelectionChanged}>
                {renderChildren(width)}
              </TableSelectionContextProvider>
            ) : (
              renderChildren(width)
            )}
          </TablePaginationContextProvider>
        </TableContextProvider>
      )}
    </AutoSizer>
  );
};

export default Table;
