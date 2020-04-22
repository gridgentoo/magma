/**
 * @generated
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 **/

 /**
 * @flow
 * @relayHash 32e7294c82aabb473f50d055ce4193c6
 */

/* eslint-disable */

'use strict';

/*::
import type { ConcreteRequest } from 'relay-runtime';
export type ProjectTypesListQueryVariables = {||};
export type ProjectTypesListQueryResponse = {|
  +projectTypes: ?{|
    +edges: $ReadOnlyArray<{|
      +node: ?{|
        +id: string,
        +name: string,
      |}
    |}>
  |}
|};
export type ProjectTypesListQuery = {|
  variables: ProjectTypesListQueryVariables,
  response: ProjectTypesListQueryResponse,
|};
*/


/*
query ProjectTypesListQuery {
  projectTypes(first: 500) {
    edges {
      node {
        id
        name
        __typename
      }
      cursor
    }
    pageInfo {
      endCursor
      hasNextPage
    }
  }
}
*/

const node/*: ConcreteRequest*/ = (function(){
var v0 = [
  {
    "kind": "LinkedField",
    "alias": null,
    "name": "edges",
    "storageKey": null,
    "args": null,
    "concreteType": "ProjectTypeEdge",
    "plural": true,
    "selections": [
      {
        "kind": "LinkedField",
        "alias": null,
        "name": "node",
        "storageKey": null,
        "args": null,
        "concreteType": "ProjectType",
        "plural": false,
        "selections": [
          {
            "kind": "ScalarField",
            "alias": null,
            "name": "id",
            "args": null,
            "storageKey": null
          },
          {
            "kind": "ScalarField",
            "alias": null,
            "name": "name",
            "args": null,
            "storageKey": null
          },
          {
            "kind": "ScalarField",
            "alias": null,
            "name": "__typename",
            "args": null,
            "storageKey": null
          }
        ]
      },
      {
        "kind": "ScalarField",
        "alias": null,
        "name": "cursor",
        "args": null,
        "storageKey": null
      }
    ]
  },
  {
    "kind": "LinkedField",
    "alias": null,
    "name": "pageInfo",
    "storageKey": null,
    "args": null,
    "concreteType": "PageInfo",
    "plural": false,
    "selections": [
      {
        "kind": "ScalarField",
        "alias": null,
        "name": "endCursor",
        "args": null,
        "storageKey": null
      },
      {
        "kind": "ScalarField",
        "alias": null,
        "name": "hasNextPage",
        "args": null,
        "storageKey": null
      }
    ]
  }
],
v1 = [
  {
    "kind": "Literal",
    "name": "first",
    "value": 500
  }
];
return {
  "kind": "Request",
  "fragment": {
    "kind": "Fragment",
    "name": "ProjectTypesListQuery",
    "type": "Query",
    "metadata": null,
    "argumentDefinitions": [],
    "selections": [
      {
        "kind": "LinkedField",
        "alias": "projectTypes",
        "name": "__ProjectTypesListQuery_projectTypes_connection",
        "storageKey": null,
        "args": null,
        "concreteType": "ProjectTypeConnection",
        "plural": false,
        "selections": (v0/*: any*/)
      }
    ]
  },
  "operation": {
    "kind": "Operation",
    "name": "ProjectTypesListQuery",
    "argumentDefinitions": [],
    "selections": [
      {
        "kind": "LinkedField",
        "alias": null,
        "name": "projectTypes",
        "storageKey": "projectTypes(first:500)",
        "args": (v1/*: any*/),
        "concreteType": "ProjectTypeConnection",
        "plural": false,
        "selections": (v0/*: any*/)
      },
      {
        "kind": "LinkedHandle",
        "alias": null,
        "name": "projectTypes",
        "args": (v1/*: any*/),
        "handle": "connection",
        "key": "ProjectTypesListQuery_projectTypes",
        "filters": null
      }
    ]
  },
  "params": {
    "operationKind": "query",
    "name": "ProjectTypesListQuery",
    "id": null,
    "text": "query ProjectTypesListQuery {\n  projectTypes(first: 500) {\n    edges {\n      node {\n        id\n        name\n        __typename\n      }\n      cursor\n    }\n    pageInfo {\n      endCursor\n      hasNextPage\n    }\n  }\n}\n",
    "metadata": {
      "connection": [
        {
          "count": null,
          "cursor": null,
          "direction": "forward",
          "path": [
            "projectTypes"
          ]
        }
      ]
    }
  }
};
})();
// prettier-ignore
(node/*: any*/).hash = '8ef4b6ab3858672b2cf381844a9f8757';
module.exports = node;
