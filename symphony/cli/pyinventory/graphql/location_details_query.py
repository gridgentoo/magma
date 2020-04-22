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

from .location_fragment import LocationFragment, QUERY as LocationFragmentQuery

QUERY: List[str] = LocationFragmentQuery + ["""
query LocationDetailsQuery($id: ID!) {
  location: node(id: $id) {
    ... on Location {
      ...LocationFragment
    }
  }
}

"""]

@dataclass
class LocationDetailsQuery(DataClassJsonMixin):
    @dataclass
    class LocationDetailsQueryData(DataClassJsonMixin):
        @dataclass
        class Node(LocationFragment):
            pass

        location: Optional[Node]

    data: LocationDetailsQueryData

    @classmethod
    # fmt: off
    def execute(cls, client: GraphqlClient, id: str) -> LocationDetailsQueryData:
        # fmt: off
        variables = {"id": id}
        response_text = client.call(''.join(set(QUERY)), variables=variables)
        return cls.from_json(response_text).data
