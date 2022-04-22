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

from ..fragment.work_order_type import WorkOrderTypeFragment, QUERY as WorkOrderTypeFragmentQuery

from ..input.edit_work_order_type_input import EditWorkOrderTypeInput


# fmt: off
QUERY: List[str] = WorkOrderTypeFragmentQuery + ["""
mutation EditWorkOrderTypeMutation($input: EditWorkOrderTypeInput!) {
  editWorkOrderType(input: $input) {
    ...WorkOrderTypeFragment
  }
}

"""
]


class EditWorkOrderTypeMutation:
    @dataclass(frozen=True)
    class EditWorkOrderTypeMutationData(DataClassJsonMixin):
        @dataclass(frozen=True)
        class WorkOrderType(WorkOrderTypeFragment):
            pass

        editWorkOrderType: WorkOrderType

    # fmt: off
    @classmethod
    def execute(cls, client: Client, input: EditWorkOrderTypeInput) -> EditWorkOrderTypeMutationData.WorkOrderType:
        variables: Dict[str, Any] = {"input": input}
        new_variables = encode_variables(variables, custom_scalars)
        response_text = client.execute(
            gql("".join(set(QUERY))), variable_values=new_variables
        )
        res = cls.EditWorkOrderTypeMutationData.from_dict(response_text)
        return res.editWorkOrderType

    # fmt: off
    @classmethod
    async def execute_async(cls, client: Client, input: EditWorkOrderTypeInput) -> EditWorkOrderTypeMutationData.WorkOrderType:
        variables: Dict[str, Any] = {"input": input}
        new_variables = encode_variables(variables, custom_scalars)
        response_text = await client.execute_async(
            gql("".join(set(QUERY))), variable_values=new_variables
        )
        res = cls.EditWorkOrderTypeMutationData.from_dict(response_text)
        return res.editWorkOrderType
