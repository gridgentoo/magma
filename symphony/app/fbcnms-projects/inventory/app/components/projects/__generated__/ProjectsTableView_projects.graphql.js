/**
 * @generated
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 **/

 /**
 * @flow
 */

/* eslint-disable */

'use strict';

/*::
import type { ReaderFragment } from 'relay-runtime';
import type { FragmentReference } from "relay-runtime";
declare export opaque type ProjectsTableView_projects$ref: FragmentReference;
declare export opaque type ProjectsTableView_projects$fragmentType: ProjectsTableView_projects$ref;
export type ProjectsTableView_projects = $ReadOnlyArray<{|
  +id: string,
  +name: string,
  +createdBy: ?{|
    +email: string
  |},
  +location: ?{|
    +id: string,
    +name: string,
  |},
  +type: {|
    +id: string,
    +name: string,
  |},
  +$refType: ProjectsTableView_projects$ref,
|}>;
export type ProjectsTableView_projects$data = ProjectsTableView_projects;
export type ProjectsTableView_projects$key = $ReadOnlyArray<{
  +$data?: ProjectsTableView_projects$data,
  +$fragmentRefs: ProjectsTableView_projects$ref,
  ...
}>;
*/


const node/*: ReaderFragment*/ = (function(){
var v0 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "id",
  "args": null,
  "storageKey": null
},
v1 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "name",
  "args": null,
  "storageKey": null
},
v2 = [
  (v0/*: any*/),
  (v1/*: any*/)
];
return {
  "kind": "Fragment",
  "name": "ProjectsTableView_projects",
  "type": "Project",
  "metadata": {
    "plural": true
  },
  "argumentDefinitions": [],
  "selections": [
    (v0/*: any*/),
    (v1/*: any*/),
    {
      "kind": "LinkedField",
      "alias": null,
      "name": "createdBy",
      "storageKey": null,
      "args": null,
      "concreteType": "User",
      "plural": false,
      "selections": [
        {
          "kind": "ScalarField",
          "alias": null,
          "name": "email",
          "args": null,
          "storageKey": null
        }
      ]
    },
    {
      "kind": "LinkedField",
      "alias": null,
      "name": "location",
      "storageKey": null,
      "args": null,
      "concreteType": "Location",
      "plural": false,
      "selections": (v2/*: any*/)
    },
    {
      "kind": "LinkedField",
      "alias": null,
      "name": "type",
      "storageKey": null,
      "args": null,
      "concreteType": "ProjectType",
      "plural": false,
      "selections": (v2/*: any*/)
    }
  ]
};
})();
// prettier-ignore
(node/*: any*/).hash = 'fd90de9f841444fa7177ad7e440fad84';
module.exports = node;
