#!/usr/bin/env python3
# @generated AUTOGENERATED file. Do not Change!

from dataclasses import dataclass, field as _field
from functools import partial
from ...config import custom_scalars, datetime
from numbers import Number
from typing import Any, AsyncGenerator, Dict, List, Generator, Optional

from dataclasses_json import DataClassJsonMixin, config

from gql_client.runtime.enum_utils import enum_field_metadata
from ..enum.check_list_item_enum_selection_mode import CheckListItemEnumSelectionMode
from ..enum.check_list_item_type import CheckListItemType
from ..enum.yes_no_response import YesNoResponse

from ..input.file_input import FileInput
from ..input.survey_cell_scan_data import SurveyCellScanData
from ..input.survey_wi_fi_scan_data import SurveyWiFiScanData


@dataclass(frozen=True)
class CheckListItemInput(DataClassJsonMixin):
    title: str
    type: CheckListItemType = _field(metadata=enum_field_metadata(CheckListItemType))
    files: List[FileInput]
    wifiData: List[SurveyWiFiScanData]
    cellData: List[SurveyCellScanData]
    id: Optional[str] = None
    index: Optional[int] = None
    isMandatory: Optional[bool] = None
    helpText: Optional[str] = None
    enumValues: Optional[str] = None
    enumSelectionMode: Optional[CheckListItemEnumSelectionMode] = None
    selectedEnumValues: Optional[str] = None
    stringValue: Optional[str] = None
    checked: Optional[bool] = None
    yesNoResponse: Optional[YesNoResponse] = None
