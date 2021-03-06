#!/usr/bin/env python3
# @generated AUTOGENERATED file. Do not Change!

from dataclasses import dataclass, field as _field
from functools import partial
from ...config import custom_scalars, datetime
from numbers import Number
from typing import Any, AsyncGenerator, Dict, List, Generator, Optional

from dataclasses_json import DataClassJsonMixin, config

from gql_client.runtime.enum_utils import enum_field_metadata
from ..enum.discovery_method import DiscoveryMethod

from ..input.property_type_input import PropertyTypeInput
from ..input.service_endpoint_definition_input import ServiceEndpointDefinitionInput


@dataclass(frozen=True)
class ServiceTypeCreateData(DataClassJsonMixin):
    name: str
    hasCustomer: bool
    properties: Optional[List[PropertyTypeInput]] = None
    endpoints: Optional[List[ServiceEndpointDefinitionInput]] = None
    discoveryMethod: Optional[DiscoveryMethod] = None
