/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow strict-local
 * @format
 */

import type {TableColumnType} from './TableHeader';

import * as React from 'react';
import emptyFunction from '@fbcnms/util/emptyFunction';
import {createContext, useContext, useEffect, useState} from 'react';
import {useCallback} from 'react';

export const TABLE_SORT_ORDER = {
  ascending: 'ascending',
  descending: 'descending',
};

export const MIN_COLUMN_WIDTH_PX = 80;

export type TableColumnWidths = Array<{|
  key: string,
  width: number,
|}>;

export type TableSortOrders = $Keys<typeof TABLE_SORT_ORDER>;

export type TableSortSettings = $ReadOnly<{|
  columnKey: string,
  order: TableSortOrders,
  overrideSorting?: ?boolean,
|}>;

export type TableSettings = $ReadOnly<{|
  showSelection: boolean,
  clickableRows: boolean,
  sort?: ?TableSortSettings,
  columnWidths?: ?TableColumnWidths,
  resizableColumns?: ?boolean,
|}>;

export type TableContextValue = $ReadOnly<{|
  settings: TableSettings,
  setSortSettings: (?TableSortSettings) => void,
  isLoading: boolean,
  setIsLoading: (isLoading: boolean) => void,
  changeColumnWidthByDelta: (colIndex: number, deltaX: number) => void,
  width: ?number,
|}>;

const TableContext = createContext<TableContextValue>({
  settings: {
    showSelection: false,
    clickableRows: false,
    resizableColumns: false,
    columnWidths: null,
  },
  setSortSettings: emptyFunction,
  isLoading: false,
  setIsLoading: emptyFunction,
  changeColumnWidthByDelta: emptyFunction,
  width: null,
});

export function useTable() {
  return useContext(TableContext);
}

type Props<T> = $ReadOnly<{|
  settings: TableSettings,
  columns: Array<TableColumnType<T>>,
  width: ?number,
  children: React.Node,
|}>;

const calculateColumnWidthPixels = <T>(
  tableWidth: ?number,
  columns: Array<TableColumnType<T>>,
): ?TableColumnWidths => {
  if (tableWidth == null) {
    return null;
  }

  const initialColumnWidths = columns
    .filter(c => c.hidden != true)
    .map(c => ({key: c.key, width: c.width}));

  const pixelColumnWidths: TableColumnWidths = [];
  let tableAvailableWidth = tableWidth;

  // Handle columns with pixels width
  initialColumnWidths.forEach((col, i) => {
    const colWidth = col.width;
    if (colWidth != null && colWidth >= 1) {
      pixelColumnWidths[i] = {
        key: col.key,
        width: colWidth,
      };
      tableAvailableWidth -= colWidth;
    }
  });

  // Handle columns with ratio width (e.g. 0.33)
  let widthLeftoverAfterRatioColumns = tableAvailableWidth;
  initialColumnWidths.forEach((col, i) => {
    const colWidthRatio = col.width;
    if (colWidthRatio != null && colWidthRatio < 1) {
      const colWidthPx = tableAvailableWidth * colWidthRatio;
      pixelColumnWidths[i] = {
        key: col.key,
        width: colWidthPx,
      };
      widthLeftoverAfterRatioColumns -= colWidthPx;
    }
  });
  tableAvailableWidth = widthLeftoverAfterRatioColumns;

  const numColumnsWithUndefinedWidth = initialColumnWidths.filter(
    col => col.width == null,
  ).length;

  initialColumnWidths.forEach((col, i) => {
    if (col.width == null) {
      pixelColumnWidths[i] = {
        key: col.key,
        width: tableAvailableWidth / numColumnsWithUndefinedWidth,
      };
    }
  });

  return pixelColumnWidths;
};

export function TableContextProvider<T>(props: Props<T>) {
  const {children, settings, width, columns} = props;
  const [isLoading, setIsLoading] = useState(false);
  const [sortSettings, setSortSettings] = useState<?TableSortSettings>(null);
  const [columnWidths, setColumnWidths] = useState<?TableColumnWidths>(null);
  useEffect(() => setSortSettings(settings.sort), [settings.sort]);
  useEffect(() => setColumnWidths(calculateColumnWidthPixels(width, columns)), [
    columns,
    width,
  ]);

  const changeColumnWidthByDelta = useCallback(
    (colIndex: number, deltaX: number) => {
      if (settings.resizableColumns == false || width == null) {
        return;
      }

      setColumnWidths(prevWidths => {
        if (prevWidths == null) {
          return prevWidths;
        }

        const newColWidths = prevWidths.slice().map(col => ({...col}));

        const newColWidth = prevWidths[colIndex].width + deltaX;
        if (newColWidth < MIN_COLUMN_WIDTH_PX) {
          return prevWidths;
        }

        newColWidths[colIndex].width = newColWidth;

        if (colIndex !== prevWidths.length - 1) {
          const nextColKey = colIndex + 1;
          const nextColWidth = prevWidths[nextColKey].width - deltaX;

          if (nextColWidth < MIN_COLUMN_WIDTH_PX) {
            return prevWidths;
          }
          newColWidths[nextColKey].width = nextColWidth;
        }

        return newColWidths;
      });
    },
    [settings.resizableColumns, width],
  );

  return (
    <TableContext.Provider
      value={{
        settings: {
          ...settings,
          sort: sortSettings,
          columnWidths,
        },
        setSortSettings,
        isLoading,
        setIsLoading,
        changeColumnWidthByDelta,
        width,
      }}>
      {children}
    </TableContext.Provider>
  );
}

export default TableContext;
