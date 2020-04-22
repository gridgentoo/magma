#!/usr/bin/env python3

import urllib3
from gql.gql.graphql_client import GraphqlClient
from requests.sessions import Session

from .graphql.simple import VersionQuery


def get_version() -> VersionQuery.VersionQueryData:
    address = "http://localhost:8085/query"

    session = Session()
    session.verify = False
    urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)

    client = GraphqlClient(address, session, "Pyhub/0.0.1")

    return VersionQuery.execute(client)
