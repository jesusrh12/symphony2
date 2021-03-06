#!/usr/bin/env python3
# @generated AUTOGENERATED file. Do not Change!

from dataclasses import dataclass, field as _field
from functools import partial
from ...config import custom_scalars, datetime
from numbers import Number
from typing import Any, AsyncGenerator, Dict, List, Generator, Optional

from dataclasses_json import DataClassJsonMixin, config

from ..input.variable_definition_input import VariableDefinitionInput


@dataclass(frozen=True)
class AddFlowDraftInput(DataClassJsonMixin):
    name: str
    endParamDefinitions: List[VariableDefinitionInput]
    description: Optional[str] = None
    flowID: Optional[str] = None
