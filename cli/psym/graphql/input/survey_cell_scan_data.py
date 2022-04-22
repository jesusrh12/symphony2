#!/usr/bin/env python3
# @generated AUTOGENERATED file. Do not Change!

from dataclasses import dataclass, field as _field
from functools import partial
from ...config import custom_scalars, datetime
from numbers import Number
from typing import Any, AsyncGenerator, Dict, List, Generator, Optional

from dataclasses_json import DataClassJsonMixin, config

from gql_client.runtime.enum_utils import enum_field_metadata
from ..enum.cellular_network_type import CellularNetworkType


@dataclass(frozen=True)
class SurveyCellScanData(DataClassJsonMixin):
    networkType: CellularNetworkType = _field(metadata=enum_field_metadata(CellularNetworkType))
    signalStrength: int
    timestamp: Optional[int] = None
    baseStationID: Optional[str] = None
    networkID: Optional[str] = None
    systemID: Optional[str] = None
    cellID: Optional[str] = None
    locationAreaCode: Optional[str] = None
    mobileCountryCode: Optional[str] = None
    mobileNetworkCode: Optional[str] = None
    primaryScramblingCode: Optional[str] = None
    operator: Optional[str] = None
    arfcn: Optional[int] = None
    physicalCellID: Optional[str] = None
    trackingAreaCode: Optional[str] = None
    timingAdvance: Optional[int] = None
    earfcn: Optional[int] = None
    uarfcn: Optional[int] = None
    latitude: Optional[Number] = None
    longitude: Optional[Number] = None
    altitude: Optional[Number] = None
    heading: Optional[Number] = None
    rssi: Optional[Number] = None