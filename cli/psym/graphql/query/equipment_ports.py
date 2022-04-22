#!/usr/bin/env python3
# @generated AUTOGENERATED file. Do not Change!

from dataclasses import dataclass, field as _field
from ...config import custom_scalars, datetime
from gql_client.runtime.variables import encode_variables
from gql import gql, Client
from gql.transport.exceptions import TransportQueryError
from functools import partial
from numbers import Number
from typing import Any, AsyncGenerator, Dict, List, Generator, Optional
from time import perf_counter
from dataclasses_json import DataClassJsonMixin, config

from ..fragment.equipment_port import EquipmentPortFragment, QUERY as EquipmentPortFragmentQuery


# fmt: off
QUERY: List[str] = EquipmentPortFragmentQuery + ["""
query EquipmentPortsQuery($id: ID!) {
  equipment: node(id: $id) {
    ... on Equipment {
      ports {
        ...EquipmentPortFragment
      }
    }
  }
}

"""
]


class EquipmentPortsQuery:
    @dataclass(frozen=True)
    class EquipmentPortsQueryData(DataClassJsonMixin):
        @dataclass(frozen=True)
        class Node(DataClassJsonMixin):
            @dataclass(frozen=True)
            class EquipmentPort(EquipmentPortFragment):
                pass

            ports: List[EquipmentPort]

        equipment: Optional[Node]

    # fmt: off
    @classmethod
    def execute(cls, client: Client, id: str) -> Optional[EquipmentPortsQueryData.Node]:
        variables: Dict[str, Any] = {"id": id}
        new_variables = encode_variables(variables, custom_scalars)
        response_text = client.execute(
            gql("".join(set(QUERY))), variable_values=new_variables
        )
        res = cls.EquipmentPortsQueryData.from_dict(response_text)
        return res.equipment

    # fmt: off
    @classmethod
    async def execute_async(cls, client: Client, id: str) -> Optional[EquipmentPortsQueryData.Node]:
        variables: Dict[str, Any] = {"id": id}
        new_variables = encode_variables(variables, custom_scalars)
        response_text = await client.execute_async(
            gql("".join(set(QUERY))), variable_values=new_variables
        )
        res = cls.EquipmentPortsQueryData.from_dict(response_text)
        return res.equipment