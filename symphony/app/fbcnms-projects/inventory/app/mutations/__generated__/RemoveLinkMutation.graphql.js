/**
 * @generated
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 **/

 /**
 * @flow
 * @relayHash ae25bb864089314a5fece6dd880b3b10
 */

/* eslint-disable */

'use strict';

/*::
import type { ConcreteRequest } from 'relay-runtime';
type EquipmentBreadcrumbs_equipment$ref = any;
export type FutureState = "INSTALL" | "REMOVE" | "%future added value";
export type PropertyKind = "bool" | "date" | "datetime_local" | "email" | "enum" | "float" | "gps_location" | "int" | "node" | "range" | "string" | "%future added value";
export type WorkOrderStatus = "DONE" | "PENDING" | "PLANNED" | "%future added value";
export type RemoveLinkMutationVariables = {|
  id: string,
  workOrderId?: ?string,
|};
export type RemoveLinkMutationResponse = {|
  +removeLink: {|
    +id: string,
    +futureState: ?FutureState,
    +ports: $ReadOnlyArray<?{|
      +id: string,
      +definition: {|
        +id: string,
        +name: string,
        +visibleLabel: ?string,
        +portType: ?{|
          +linkPropertyTypes: $ReadOnlyArray<?{|
            +id: string,
            +name: string,
            +type: PropertyKind,
            +nodeType: ?string,
            +index: ?number,
            +stringValue: ?string,
            +intValue: ?number,
            +booleanValue: ?boolean,
            +floatValue: ?number,
            +latitudeValue: ?number,
            +longitudeValue: ?number,
            +rangeFromValue: ?number,
            +rangeToValue: ?number,
            +isEditable: ?boolean,
            +isInstanceProperty: ?boolean,
            +isMandatory: ?boolean,
            +category: ?string,
            +isDeleted: ?boolean,
          |}>
        |},
      |},
      +parentEquipment: {|
        +id: string,
        +name: string,
        +futureState: ?FutureState,
        +equipmentType: {|
          +id: string,
          +name: string,
          +portDefinitions: $ReadOnlyArray<?{|
            +id: string,
            +name: string,
            +visibleLabel: ?string,
            +bandwidth: ?string,
            +portType: ?{|
              +id: string,
              +name: string,
            |},
          |}>,
        |},
        +$fragmentRefs: EquipmentBreadcrumbs_equipment$ref,
      |},
      +serviceEndpoints: $ReadOnlyArray<{|
        +definition: {|
          +role: ?string
        |},
        +service: {|
          +name: string
        |},
      |}>,
    |}>,
    +workOrder: ?{|
      +id: string,
      +status: WorkOrderStatus,
    |},
    +properties: $ReadOnlyArray<?{|
      +id: string,
      +propertyType: {|
        +id: string,
        +name: string,
        +type: PropertyKind,
        +nodeType: ?string,
        +index: ?number,
        +stringValue: ?string,
        +intValue: ?number,
        +booleanValue: ?boolean,
        +floatValue: ?number,
        +latitudeValue: ?number,
        +longitudeValue: ?number,
        +rangeFromValue: ?number,
        +rangeToValue: ?number,
        +isEditable: ?boolean,
        +isInstanceProperty: ?boolean,
        +isMandatory: ?boolean,
        +category: ?string,
        +isDeleted: ?boolean,
      |},
      +stringValue: ?string,
      +intValue: ?number,
      +floatValue: ?number,
      +booleanValue: ?boolean,
      +latitudeValue: ?number,
      +longitudeValue: ?number,
      +rangeFromValue: ?number,
      +rangeToValue: ?number,
      +nodeValue: ?{|
        +id: string,
        +name: string,
      |},
    |}>,
    +services: $ReadOnlyArray<?{|
      +id: string,
      +name: string,
    |}>,
  |}
|};
export type RemoveLinkMutation = {|
  variables: RemoveLinkMutationVariables,
  response: RemoveLinkMutationResponse,
|};
*/


/*
mutation RemoveLinkMutation(
  $id: ID!
  $workOrderId: ID
) {
  removeLink(id: $id, workOrderId: $workOrderId) {
    id
    futureState
    ports {
      id
      definition {
        id
        name
        visibleLabel
        portType {
          linkPropertyTypes {
            id
            name
            type
            nodeType
            index
            stringValue
            intValue
            booleanValue
            floatValue
            latitudeValue
            longitudeValue
            rangeFromValue
            rangeToValue
            isEditable
            isInstanceProperty
            isMandatory
            category
            isDeleted
          }
          id
        }
      }
      parentEquipment {
        id
        name
        futureState
        equipmentType {
          id
          name
          portDefinitions {
            id
            name
            visibleLabel
            bandwidth
            portType {
              id
              name
            }
          }
        }
        ...EquipmentBreadcrumbs_equipment
      }
      serviceEndpoints {
        definition {
          role
          id
        }
        service {
          name
          id
        }
        id
      }
    }
    workOrder {
      id
      status
    }
    properties {
      id
      propertyType {
        id
        name
        type
        nodeType
        index
        stringValue
        intValue
        booleanValue
        floatValue
        latitudeValue
        longitudeValue
        rangeFromValue
        rangeToValue
        isEditable
        isInstanceProperty
        isMandatory
        category
        isDeleted
      }
      stringValue
      intValue
      floatValue
      booleanValue
      latitudeValue
      longitudeValue
      rangeFromValue
      rangeToValue
      nodeValue {
        __typename
        id
        name
      }
    }
    services {
      id
      name
    }
  }
}

fragment EquipmentBreadcrumbs_equipment on Equipment {
  id
  name
  equipmentType {
    id
    name
  }
  locationHierarchy {
    id
    name
    locationType {
      name
      id
    }
  }
  positionHierarchy {
    id
    definition {
      id
      name
      visibleLabel
    }
    parentEquipment {
      id
      name
      equipmentType {
        id
        name
      }
    }
  }
}
*/

const node/*: ConcreteRequest*/ = (function(){
var v0 = [
  {
    "kind": "LocalArgument",
    "name": "id",
    "type": "ID!",
    "defaultValue": null
  },
  {
    "kind": "LocalArgument",
    "name": "workOrderId",
    "type": "ID",
    "defaultValue": null
  }
],
v1 = [
  {
    "kind": "Variable",
    "name": "id",
    "variableName": "id"
  },
  {
    "kind": "Variable",
    "name": "workOrderId",
    "variableName": "workOrderId"
  }
],
v2 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "id",
  "args": null,
  "storageKey": null
},
v3 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "futureState",
  "args": null,
  "storageKey": null
},
v4 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "name",
  "args": null,
  "storageKey": null
},
v5 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "visibleLabel",
  "args": null,
  "storageKey": null
},
v6 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "stringValue",
  "args": null,
  "storageKey": null
},
v7 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "intValue",
  "args": null,
  "storageKey": null
},
v8 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "booleanValue",
  "args": null,
  "storageKey": null
},
v9 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "floatValue",
  "args": null,
  "storageKey": null
},
v10 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "latitudeValue",
  "args": null,
  "storageKey": null
},
v11 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "longitudeValue",
  "args": null,
  "storageKey": null
},
v12 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "rangeFromValue",
  "args": null,
  "storageKey": null
},
v13 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "rangeToValue",
  "args": null,
  "storageKey": null
},
v14 = [
  (v2/*: any*/),
  (v4/*: any*/),
  {
    "kind": "ScalarField",
    "alias": null,
    "name": "type",
    "args": null,
    "storageKey": null
  },
  {
    "kind": "ScalarField",
    "alias": null,
    "name": "nodeType",
    "args": null,
    "storageKey": null
  },
  {
    "kind": "ScalarField",
    "alias": null,
    "name": "index",
    "args": null,
    "storageKey": null
  },
  (v6/*: any*/),
  (v7/*: any*/),
  (v8/*: any*/),
  (v9/*: any*/),
  (v10/*: any*/),
  (v11/*: any*/),
  (v12/*: any*/),
  (v13/*: any*/),
  {
    "kind": "ScalarField",
    "alias": null,
    "name": "isEditable",
    "args": null,
    "storageKey": null
  },
  {
    "kind": "ScalarField",
    "alias": null,
    "name": "isInstanceProperty",
    "args": null,
    "storageKey": null
  },
  {
    "kind": "ScalarField",
    "alias": null,
    "name": "isMandatory",
    "args": null,
    "storageKey": null
  },
  {
    "kind": "ScalarField",
    "alias": null,
    "name": "category",
    "args": null,
    "storageKey": null
  },
  {
    "kind": "ScalarField",
    "alias": null,
    "name": "isDeleted",
    "args": null,
    "storageKey": null
  }
],
v15 = {
  "kind": "LinkedField",
  "alias": null,
  "name": "linkPropertyTypes",
  "storageKey": null,
  "args": null,
  "concreteType": "PropertyType",
  "plural": true,
  "selections": (v14/*: any*/)
},
v16 = [
  (v2/*: any*/),
  (v4/*: any*/)
],
v17 = {
  "kind": "LinkedField",
  "alias": null,
  "name": "equipmentType",
  "storageKey": null,
  "args": null,
  "concreteType": "EquipmentType",
  "plural": false,
  "selections": [
    (v2/*: any*/),
    (v4/*: any*/),
    {
      "kind": "LinkedField",
      "alias": null,
      "name": "portDefinitions",
      "storageKey": null,
      "args": null,
      "concreteType": "EquipmentPortDefinition",
      "plural": true,
      "selections": [
        (v2/*: any*/),
        (v4/*: any*/),
        (v5/*: any*/),
        {
          "kind": "ScalarField",
          "alias": null,
          "name": "bandwidth",
          "args": null,
          "storageKey": null
        },
        {
          "kind": "LinkedField",
          "alias": null,
          "name": "portType",
          "storageKey": null,
          "args": null,
          "concreteType": "EquipmentPortType",
          "plural": false,
          "selections": (v16/*: any*/)
        }
      ]
    }
  ]
},
v18 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "role",
  "args": null,
  "storageKey": null
},
v19 = {
  "kind": "LinkedField",
  "alias": null,
  "name": "workOrder",
  "storageKey": null,
  "args": null,
  "concreteType": "WorkOrder",
  "plural": false,
  "selections": [
    (v2/*: any*/),
    {
      "kind": "ScalarField",
      "alias": null,
      "name": "status",
      "args": null,
      "storageKey": null
    }
  ]
},
v20 = {
  "kind": "LinkedField",
  "alias": null,
  "name": "propertyType",
  "storageKey": null,
  "args": null,
  "concreteType": "PropertyType",
  "plural": false,
  "selections": (v14/*: any*/)
},
v21 = {
  "kind": "LinkedField",
  "alias": null,
  "name": "services",
  "storageKey": null,
  "args": null,
  "concreteType": "Service",
  "plural": true,
  "selections": (v16/*: any*/)
},
v22 = [
  (v4/*: any*/),
  (v2/*: any*/)
];
return {
  "kind": "Request",
  "fragment": {
    "kind": "Fragment",
    "name": "RemoveLinkMutation",
    "type": "Mutation",
    "metadata": null,
    "argumentDefinitions": (v0/*: any*/),
    "selections": [
      {
        "kind": "LinkedField",
        "alias": null,
        "name": "removeLink",
        "storageKey": null,
        "args": (v1/*: any*/),
        "concreteType": "Link",
        "plural": false,
        "selections": [
          (v2/*: any*/),
          (v3/*: any*/),
          {
            "kind": "LinkedField",
            "alias": null,
            "name": "ports",
            "storageKey": null,
            "args": null,
            "concreteType": "EquipmentPort",
            "plural": true,
            "selections": [
              (v2/*: any*/),
              {
                "kind": "LinkedField",
                "alias": null,
                "name": "definition",
                "storageKey": null,
                "args": null,
                "concreteType": "EquipmentPortDefinition",
                "plural": false,
                "selections": [
                  (v2/*: any*/),
                  (v4/*: any*/),
                  (v5/*: any*/),
                  {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "portType",
                    "storageKey": null,
                    "args": null,
                    "concreteType": "EquipmentPortType",
                    "plural": false,
                    "selections": [
                      (v15/*: any*/)
                    ]
                  }
                ]
              },
              {
                "kind": "LinkedField",
                "alias": null,
                "name": "parentEquipment",
                "storageKey": null,
                "args": null,
                "concreteType": "Equipment",
                "plural": false,
                "selections": [
                  (v2/*: any*/),
                  (v4/*: any*/),
                  (v3/*: any*/),
                  (v17/*: any*/),
                  {
                    "kind": "FragmentSpread",
                    "name": "EquipmentBreadcrumbs_equipment",
                    "args": null
                  }
                ]
              },
              {
                "kind": "LinkedField",
                "alias": null,
                "name": "serviceEndpoints",
                "storageKey": null,
                "args": null,
                "concreteType": "ServiceEndpoint",
                "plural": true,
                "selections": [
                  {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "definition",
                    "storageKey": null,
                    "args": null,
                    "concreteType": "ServiceEndpointDefinition",
                    "plural": false,
                    "selections": [
                      (v18/*: any*/)
                    ]
                  },
                  {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "service",
                    "storageKey": null,
                    "args": null,
                    "concreteType": "Service",
                    "plural": false,
                    "selections": [
                      (v4/*: any*/)
                    ]
                  }
                ]
              }
            ]
          },
          (v19/*: any*/),
          {
            "kind": "LinkedField",
            "alias": null,
            "name": "properties",
            "storageKey": null,
            "args": null,
            "concreteType": "Property",
            "plural": true,
            "selections": [
              (v2/*: any*/),
              (v20/*: any*/),
              (v6/*: any*/),
              (v7/*: any*/),
              (v9/*: any*/),
              (v8/*: any*/),
              (v10/*: any*/),
              (v11/*: any*/),
              (v12/*: any*/),
              (v13/*: any*/),
              {
                "kind": "LinkedField",
                "alias": null,
                "name": "nodeValue",
                "storageKey": null,
                "args": null,
                "concreteType": null,
                "plural": false,
                "selections": (v16/*: any*/)
              }
            ]
          },
          (v21/*: any*/)
        ]
      }
    ]
  },
  "operation": {
    "kind": "Operation",
    "name": "RemoveLinkMutation",
    "argumentDefinitions": (v0/*: any*/),
    "selections": [
      {
        "kind": "LinkedField",
        "alias": null,
        "name": "removeLink",
        "storageKey": null,
        "args": (v1/*: any*/),
        "concreteType": "Link",
        "plural": false,
        "selections": [
          (v2/*: any*/),
          (v3/*: any*/),
          {
            "kind": "LinkedField",
            "alias": null,
            "name": "ports",
            "storageKey": null,
            "args": null,
            "concreteType": "EquipmentPort",
            "plural": true,
            "selections": [
              (v2/*: any*/),
              {
                "kind": "LinkedField",
                "alias": null,
                "name": "definition",
                "storageKey": null,
                "args": null,
                "concreteType": "EquipmentPortDefinition",
                "plural": false,
                "selections": [
                  (v2/*: any*/),
                  (v4/*: any*/),
                  (v5/*: any*/),
                  {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "portType",
                    "storageKey": null,
                    "args": null,
                    "concreteType": "EquipmentPortType",
                    "plural": false,
                    "selections": [
                      (v15/*: any*/),
                      (v2/*: any*/)
                    ]
                  }
                ]
              },
              {
                "kind": "LinkedField",
                "alias": null,
                "name": "parentEquipment",
                "storageKey": null,
                "args": null,
                "concreteType": "Equipment",
                "plural": false,
                "selections": [
                  (v2/*: any*/),
                  (v4/*: any*/),
                  (v3/*: any*/),
                  (v17/*: any*/),
                  {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "locationHierarchy",
                    "storageKey": null,
                    "args": null,
                    "concreteType": "Location",
                    "plural": true,
                    "selections": [
                      (v2/*: any*/),
                      (v4/*: any*/),
                      {
                        "kind": "LinkedField",
                        "alias": null,
                        "name": "locationType",
                        "storageKey": null,
                        "args": null,
                        "concreteType": "LocationType",
                        "plural": false,
                        "selections": (v22/*: any*/)
                      }
                    ]
                  },
                  {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "positionHierarchy",
                    "storageKey": null,
                    "args": null,
                    "concreteType": "EquipmentPosition",
                    "plural": true,
                    "selections": [
                      (v2/*: any*/),
                      {
                        "kind": "LinkedField",
                        "alias": null,
                        "name": "definition",
                        "storageKey": null,
                        "args": null,
                        "concreteType": "EquipmentPositionDefinition",
                        "plural": false,
                        "selections": [
                          (v2/*: any*/),
                          (v4/*: any*/),
                          (v5/*: any*/)
                        ]
                      },
                      {
                        "kind": "LinkedField",
                        "alias": null,
                        "name": "parentEquipment",
                        "storageKey": null,
                        "args": null,
                        "concreteType": "Equipment",
                        "plural": false,
                        "selections": [
                          (v2/*: any*/),
                          (v4/*: any*/),
                          {
                            "kind": "LinkedField",
                            "alias": null,
                            "name": "equipmentType",
                            "storageKey": null,
                            "args": null,
                            "concreteType": "EquipmentType",
                            "plural": false,
                            "selections": (v16/*: any*/)
                          }
                        ]
                      }
                    ]
                  }
                ]
              },
              {
                "kind": "LinkedField",
                "alias": null,
                "name": "serviceEndpoints",
                "storageKey": null,
                "args": null,
                "concreteType": "ServiceEndpoint",
                "plural": true,
                "selections": [
                  {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "definition",
                    "storageKey": null,
                    "args": null,
                    "concreteType": "ServiceEndpointDefinition",
                    "plural": false,
                    "selections": [
                      (v18/*: any*/),
                      (v2/*: any*/)
                    ]
                  },
                  {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "service",
                    "storageKey": null,
                    "args": null,
                    "concreteType": "Service",
                    "plural": false,
                    "selections": (v22/*: any*/)
                  },
                  (v2/*: any*/)
                ]
              }
            ]
          },
          (v19/*: any*/),
          {
            "kind": "LinkedField",
            "alias": null,
            "name": "properties",
            "storageKey": null,
            "args": null,
            "concreteType": "Property",
            "plural": true,
            "selections": [
              (v2/*: any*/),
              (v20/*: any*/),
              (v6/*: any*/),
              (v7/*: any*/),
              (v9/*: any*/),
              (v8/*: any*/),
              (v10/*: any*/),
              (v11/*: any*/),
              (v12/*: any*/),
              (v13/*: any*/),
              {
                "kind": "LinkedField",
                "alias": null,
                "name": "nodeValue",
                "storageKey": null,
                "args": null,
                "concreteType": null,
                "plural": false,
                "selections": [
                  {
                    "kind": "ScalarField",
                    "alias": null,
                    "name": "__typename",
                    "args": null,
                    "storageKey": null
                  },
                  (v2/*: any*/),
                  (v4/*: any*/)
                ]
              }
            ]
          },
          (v21/*: any*/)
        ]
      }
    ]
  },
  "params": {
    "operationKind": "mutation",
    "name": "RemoveLinkMutation",
    "id": null,
    "text": "mutation RemoveLinkMutation(\n  $id: ID!\n  $workOrderId: ID\n) {\n  removeLink(id: $id, workOrderId: $workOrderId) {\n    id\n    futureState\n    ports {\n      id\n      definition {\n        id\n        name\n        visibleLabel\n        portType {\n          linkPropertyTypes {\n            id\n            name\n            type\n            nodeType\n            index\n            stringValue\n            intValue\n            booleanValue\n            floatValue\n            latitudeValue\n            longitudeValue\n            rangeFromValue\n            rangeToValue\n            isEditable\n            isInstanceProperty\n            isMandatory\n            category\n            isDeleted\n          }\n          id\n        }\n      }\n      parentEquipment {\n        id\n        name\n        futureState\n        equipmentType {\n          id\n          name\n          portDefinitions {\n            id\n            name\n            visibleLabel\n            bandwidth\n            portType {\n              id\n              name\n            }\n          }\n        }\n        ...EquipmentBreadcrumbs_equipment\n      }\n      serviceEndpoints {\n        definition {\n          role\n          id\n        }\n        service {\n          name\n          id\n        }\n        id\n      }\n    }\n    workOrder {\n      id\n      status\n    }\n    properties {\n      id\n      propertyType {\n        id\n        name\n        type\n        nodeType\n        index\n        stringValue\n        intValue\n        booleanValue\n        floatValue\n        latitudeValue\n        longitudeValue\n        rangeFromValue\n        rangeToValue\n        isEditable\n        isInstanceProperty\n        isMandatory\n        category\n        isDeleted\n      }\n      stringValue\n      intValue\n      floatValue\n      booleanValue\n      latitudeValue\n      longitudeValue\n      rangeFromValue\n      rangeToValue\n      nodeValue {\n        __typename\n        id\n        name\n      }\n    }\n    services {\n      id\n      name\n    }\n  }\n}\n\nfragment EquipmentBreadcrumbs_equipment on Equipment {\n  id\n  name\n  equipmentType {\n    id\n    name\n  }\n  locationHierarchy {\n    id\n    name\n    locationType {\n      name\n      id\n    }\n  }\n  positionHierarchy {\n    id\n    definition {\n      id\n      name\n      visibleLabel\n    }\n    parentEquipment {\n      id\n      name\n      equipmentType {\n        id\n        name\n      }\n    }\n  }\n}\n",
    "metadata": {}
  }
};
})();
// prettier-ignore
(node/*: any*/).hash = '779f6813617bbf5b040592f2e8c2c3e3';
module.exports = node;
