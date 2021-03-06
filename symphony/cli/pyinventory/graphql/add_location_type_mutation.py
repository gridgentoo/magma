#!/usr/bin/env python3
# @generated AUTOGENERATED file. Do not Change!

from dataclasses import dataclass
from datetime import datetime
from gql.gql.datetime_utils import DATETIME_FIELD
from gql.gql.graphql_client import GraphqlClient
from functools import partial
from numbers import Number
from typing import Any, Callable, List, Mapping, Optional

from dataclasses_json import DataClassJsonMixin

from .property_type_fragment import PropertyTypeFragment, QUERY as PropertyTypeFragmentQuery
from .add_location_type_input import AddLocationTypeInput


QUERY: List[str] = PropertyTypeFragmentQuery + ["""
mutation AddLocationTypeMutation($input: AddLocationTypeInput!) {
  addLocationType(input: $input) {
    id
    name
    propertyTypes {
      ...PropertyTypeFragment
    }
  }
}

"""]

@dataclass
class AddLocationTypeMutation(DataClassJsonMixin):
    @dataclass
    class AddLocationTypeMutationData(DataClassJsonMixin):
        @dataclass
        class LocationType(DataClassJsonMixin):
            @dataclass
            class PropertyType(PropertyTypeFragment):
                pass

            id: str
            name: str
            propertyTypes: List[PropertyType]

        addLocationType: LocationType

    data: AddLocationTypeMutationData

    @classmethod
    # fmt: off
    def execute(cls, client: GraphqlClient, input: AddLocationTypeInput) -> AddLocationTypeMutationData:
        # fmt: off
        variables = {"input": input}
        response_text = client.call(''.join(set(QUERY)), variables=variables)
        return cls.from_json(response_text).data
