#!/usr/bin/env python3

from typing import Any, Dict, Optional

import requests
from gql.gql import gql
from gql.gql.client import Client
from requests.auth import AuthBase

from .reporter import DUMMY_REPORTER, Reporter
from .transport.session import RequestsHTTPSessionTransport


class GraphqlClient:
    def __init__(
        self,
        graphql_endpoint_address: str,
        session: requests.Session,
        client_name: str,
        auth: Optional[AuthBase] = None,
        reporter: Reporter = DUMMY_REPORTER,
    ) -> None:

        """This is the class to use for working with graphql server

            Args:
                graphql_endpoint_address (str): The graphql server address
                auth (Optional[requests.auth.AuthBase], optional): Auth used
                    to authenticate to graphql server      
                reporter (object, optional): Use reporter.InventoryReporter to
                            store reports on all successful and failed mutations
                            in inventory. The default is DummyReporter that
                            discards reports

        """

        self.reporter = reporter
        self.client = Client(
            transport=RequestsHTTPSessionTransport(
                session,
                graphql_endpoint_address,
                headers={
                    "Accept": "application/json",
                    "Content-Type": "application/json",
                    "User-Agent": client_name,
                },
                auth=auth,
            ),
            fetch_schema_from_transport=True,
        )

    def call(self, query: str, variables: Dict[str, Any]) -> str:
        return self.client.execute(gql(query), variable_values=variables)
