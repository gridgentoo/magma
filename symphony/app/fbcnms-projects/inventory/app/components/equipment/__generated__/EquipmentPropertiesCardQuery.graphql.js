/**
 * @generated
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 **/

 /**
 * @flow
 * @relayHash ae16f61354e6713de60a62f972db461f
 */

/* eslint-disable */

'use strict';

/*::
import type { ConcreteRequest } from 'relay-runtime';
type DynamicPropertiesGrid_properties$ref = any;
type DynamicPropertiesGrid_propertyTypes$ref = any;
type EquipmentBreadcrumbs_equipment$ref = any;
type EquipmentDocumentsCard_equipment$ref = any;
type EquipmentPortsTable_equipment$ref = any;
type EquipmentPositionsGrid_equipment$ref = any;
type PositionDefinitionsTable_positionDefinitions$ref = any;
type PropertyFormField_property$ref = any;
type PropertyTypeFormField_propertyType$ref = any;
export type EquipmentPropertiesCardQueryVariables = {|
  equipmentId: string
|};
export type EquipmentPropertiesCardQueryResponse = {|
  +equipment: ?{|
    +id?: string,
    +name?: string,
    +equipmentType?: {|
      +id: string,
      +name: string,
      +propertyTypes: $ReadOnlyArray<?{|
        +$fragmentRefs: PropertyTypeFormField_propertyType$ref & DynamicPropertiesGrid_propertyTypes$ref
      |}>,
      +positionDefinitions: $ReadOnlyArray<?{|
        +id: string,
        +$fragmentRefs: PositionDefinitionsTable_positionDefinitions$ref,
      |}>,
      +portDefinitions: $ReadOnlyArray<?{|
        +id: string
      |}>,
    |},
    +parentLocation?: ?{|
      +id: string,
      +name: string,
    |},
    +parentPosition?: ?{|
      +parentEquipment: {|
        +parentLocation: ?{|
          +id: string
        |}
      |}
    |},
    +positions?: $ReadOnlyArray<?{|
      +parentEquipment: {|
        +id: string
      |}
    |}>,
    +properties?: $ReadOnlyArray<?{|
      +$fragmentRefs: PropertyFormField_property$ref & DynamicPropertiesGrid_properties$ref
    |}>,
    +services?: $ReadOnlyArray<?{|
      +id: string,
      +name: string,
      +externalId: ?string,
      +customer: ?{|
        +name: string
      |},
      +serviceType: {|
        +id: string,
        +name: string,
      |},
    |}>,
    +$fragmentRefs: EquipmentPortsTable_equipment$ref & EquipmentBreadcrumbs_equipment$ref & EquipmentPositionsGrid_equipment$ref & EquipmentDocumentsCard_equipment$ref,
  |}
|};
export type EquipmentPropertiesCardQuery = {|
  variables: EquipmentPropertiesCardQueryVariables,
  response: EquipmentPropertiesCardQueryResponse,
|};
*/


/*
query EquipmentPropertiesCardQuery(
  $equipmentId: ID!
) {
  equipment: node(id: $equipmentId) {
    __typename
    ... on Equipment {
      id
      name
      ...EquipmentPortsTable_equipment
      equipmentType {
        id
        name
        propertyTypes {
          ...PropertyTypeFormField_propertyType
          ...DynamicPropertiesGrid_propertyTypes
          id
        }
        positionDefinitions {
          id
          ...PositionDefinitionsTable_positionDefinitions
        }
        portDefinitions {
          id
        }
      }
      ...EquipmentBreadcrumbs_equipment
      parentLocation {
        id
        name
      }
      parentPosition {
        parentEquipment {
          parentLocation {
            id
          }
          id
        }
        id
      }
      ...EquipmentPositionsGrid_equipment
      positions {
        parentEquipment {
          id
        }
        id
      }
      properties {
        ...PropertyFormField_property
        ...DynamicPropertiesGrid_properties
        id
      }
      services {
        id
        name
        externalId
        customer {
          name
          id
        }
        serviceType {
          id
          name
        }
      }
      ...EquipmentDocumentsCard_equipment
    }
    id
  }
}

fragment AddToEquipmentDialog_parentEquipment on Equipment {
  id
  locationHierarchy {
    id
  }
}

fragment DocumentTable_files on File {
  id
  fileName
  category
  ...FileAttachment_file
}

fragment DocumentTable_hyperlinks on Hyperlink {
  id
  category
  url
  displayName
  ...HyperlinkTableRow_hyperlink
}

fragment DynamicPropertiesGrid_properties on Property {
  ...PropertyFormField_property
  propertyType {
    id
    index
  }
}

fragment DynamicPropertiesGrid_propertyTypes on PropertyType {
  id
  name
  index
  isInstanceProperty
  type
  nodeType
  stringValue
  intValue
  booleanValue
  latitudeValue
  longitudeValue
  rangeFromValue
  rangeToValue
  floatValue
}

fragment EntityDocumentsTable_files on File {
  ...DocumentTable_files
}

fragment EntityDocumentsTable_hyperlinks on Hyperlink {
  ...DocumentTable_hyperlinks
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

fragment EquipmentDocumentsCard_equipment on Equipment {
  id
  images {
    ...EntityDocumentsTable_files
    id
  }
  files {
    ...EntityDocumentsTable_files
    id
  }
  hyperlinks {
    ...EntityDocumentsTable_hyperlinks
    id
  }
}

fragment EquipmentPortsTable_equipment on Equipment {
  id
  name
  equipmentType {
    id
    name
    portDefinitions {
      id
      name
      index
      visibleLabel
      portType {
        id
        name
        propertyTypes {
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
      }
    }
  }
  ports {
    id
    definition {
      id
      name
      index
      visibleLabel
      portType {
        id
        name
        propertyTypes {
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
      }
    }
    parentEquipment {
      id
      name
      equipmentType {
        id
        name
        portDefinitions {
          id
          name
          visibleLabel
          portType {
            id
            name
          }
          bandwidth
        }
      }
    }
    link {
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
  positions {
    attachedEquipment {
      id
      name
      ports {
        id
        definition {
          id
          name
          index
          visibleLabel
          portType {
            id
            name
            propertyTypes {
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
          }
        }
        parentEquipment {
          id
          name
          equipmentType {
            id
            name
            portDefinitions {
              id
              name
              visibleLabel
              portType {
                id
                name
              }
              bandwidth
            }
          }
        }
        link {
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
      equipmentType {
        portDefinitions {
          id
          name
          visibleLabel
          bandwidth
        }
        id
      }
      positions {
        attachedEquipment {
          id
          name
          ports {
            id
            definition {
              id
              name
              index
              visibleLabel
              portType {
                id
                name
                propertyTypes {
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
              }
            }
            parentEquipment {
              id
              name
              equipmentType {
                id
                name
                portDefinitions {
                  id
                  name
                  visibleLabel
                  portType {
                    id
                    name
                  }
                  bandwidth
                }
              }
            }
            link {
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
          equipmentType {
            portDefinitions {
              id
              name
              visibleLabel
              bandwidth
            }
            id
          }
          positions {
            attachedEquipment {
              id
              name
              ports {
                id
                definition {
                  id
                  name
                  index
                  visibleLabel
                  portType {
                    id
                    name
                    propertyTypes {
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
                  }
                }
                parentEquipment {
                  id
                  name
                  equipmentType {
                    id
                    name
                    portDefinitions {
                      id
                      name
                      visibleLabel
                      portType {
                        id
                        name
                      }
                      bandwidth
                    }
                  }
                }
                link {
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
              equipmentType {
                portDefinitions {
                  id
                  name
                  visibleLabel
                  bandwidth
                }
                id
              }
              positions {
                attachedEquipment {
                  id
                  name
                  ports {
                    id
                    definition {
                      id
                      name
                      index
                      visibleLabel
                      portType {
                        id
                        name
                        propertyTypes {
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
                      }
                    }
                    parentEquipment {
                      id
                      name
                      equipmentType {
                        id
                        name
                        portDefinitions {
                          id
                          name
                          visibleLabel
                          portType {
                            id
                            name
                          }
                          bandwidth
                        }
                      }
                    }
                    link {
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
                  equipmentType {
                    portDefinitions {
                      id
                      name
                      visibleLabel
                      bandwidth
                    }
                    id
                  }
                }
                id
              }
            }
            id
          }
        }
        id
      }
    }
    id
  }
}

fragment EquipmentPositionsGrid_equipment on Equipment {
  id
  ...AddToEquipmentDialog_parentEquipment
  positions {
    id
    definition {
      id
      name
      index
      visibleLabel
    }
    attachedEquipment {
      id
      name
      futureState
      services {
        id
      }
    }
    parentEquipment {
      id
    }
  }
  equipmentType {
    positionDefinitions {
      id
      name
      index
      visibleLabel
    }
    id
  }
}

fragment FileAttachment_file on File {
  id
  fileName
  sizeInBytes
  uploaded
  fileType
  storeKey
  category
  ...ImageDialog_img
}

fragment HyperlinkTableRow_hyperlink on Hyperlink {
  id
  category
  url
  displayName
  createTime
}

fragment ImageDialog_img on File {
  storeKey
  fileName
}

fragment PositionDefinitionsTable_positionDefinitions on EquipmentPositionDefinition {
  id
  name
  index
  visibleLabel
}

fragment PropertyFormField_property on Property {
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

fragment PropertyTypeFormField_propertyType on PropertyType {
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
*/

const node/*: ConcreteRequest*/ = (function(){
var v0 = [
  {
    "kind": "LocalArgument",
    "name": "equipmentId",
    "type": "ID!",
    "defaultValue": null
  }
],
v1 = [
  {
    "kind": "Variable",
    "name": "id",
    "variableName": "equipmentId"
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
  "name": "name",
  "args": null,
  "storageKey": null
},
v4 = [
  (v2/*: any*/)
],
v5 = [
  (v2/*: any*/),
  (v3/*: any*/)
],
v6 = {
  "kind": "LinkedField",
  "alias": null,
  "name": "parentLocation",
  "storageKey": null,
  "args": null,
  "concreteType": "Location",
  "plural": false,
  "selections": (v5/*: any*/)
},
v7 = {
  "kind": "LinkedField",
  "alias": null,
  "name": "parentLocation",
  "storageKey": null,
  "args": null,
  "concreteType": "Location",
  "plural": false,
  "selections": (v4/*: any*/)
},
v8 = {
  "kind": "LinkedField",
  "alias": null,
  "name": "parentEquipment",
  "storageKey": null,
  "args": null,
  "concreteType": "Equipment",
  "plural": false,
  "selections": (v4/*: any*/)
},
v9 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "externalId",
  "args": null,
  "storageKey": null
},
v10 = {
  "kind": "LinkedField",
  "alias": null,
  "name": "serviceType",
  "storageKey": null,
  "args": null,
  "concreteType": "ServiceType",
  "plural": false,
  "selections": (v5/*: any*/)
},
v11 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "__typename",
  "args": null,
  "storageKey": null
},
v12 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "index",
  "args": null,
  "storageKey": null
},
v13 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "visibleLabel",
  "args": null,
  "storageKey": null
},
v14 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "stringValue",
  "args": null,
  "storageKey": null
},
v15 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "intValue",
  "args": null,
  "storageKey": null
},
v16 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "booleanValue",
  "args": null,
  "storageKey": null
},
v17 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "floatValue",
  "args": null,
  "storageKey": null
},
v18 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "latitudeValue",
  "args": null,
  "storageKey": null
},
v19 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "longitudeValue",
  "args": null,
  "storageKey": null
},
v20 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "rangeFromValue",
  "args": null,
  "storageKey": null
},
v21 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "rangeToValue",
  "args": null,
  "storageKey": null
},
v22 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "category",
  "args": null,
  "storageKey": null
},
v23 = [
  (v2/*: any*/),
  (v3/*: any*/),
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
  (v12/*: any*/),
  (v14/*: any*/),
  (v15/*: any*/),
  (v16/*: any*/),
  (v17/*: any*/),
  (v18/*: any*/),
  (v19/*: any*/),
  (v20/*: any*/),
  (v21/*: any*/),
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
  (v22/*: any*/),
  {
    "kind": "ScalarField",
    "alias": null,
    "name": "isDeleted",
    "args": null,
    "storageKey": null
  }
],
v24 = {
  "kind": "LinkedField",
  "alias": null,
  "name": "propertyTypes",
  "storageKey": null,
  "args": null,
  "concreteType": "PropertyType",
  "plural": true,
  "selections": (v23/*: any*/)
},
v25 = {
  "kind": "LinkedField",
  "alias": null,
  "name": "linkPropertyTypes",
  "storageKey": null,
  "args": null,
  "concreteType": "PropertyType",
  "plural": true,
  "selections": (v23/*: any*/)
},
v26 = [
  (v2/*: any*/),
  (v3/*: any*/),
  (v12/*: any*/),
  (v13/*: any*/),
  {
    "kind": "LinkedField",
    "alias": null,
    "name": "portType",
    "storageKey": null,
    "args": null,
    "concreteType": "EquipmentPortType",
    "plural": false,
    "selections": [
      (v2/*: any*/),
      (v3/*: any*/),
      (v24/*: any*/),
      (v25/*: any*/)
    ]
  }
],
v27 = [
  (v2/*: any*/),
  (v3/*: any*/),
  (v12/*: any*/),
  (v13/*: any*/)
],
v28 = {
  "kind": "LinkedField",
  "alias": null,
  "name": "portType",
  "storageKey": null,
  "args": null,
  "concreteType": "EquipmentPortType",
  "plural": false,
  "selections": (v5/*: any*/)
},
v29 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "bandwidth",
  "args": null,
  "storageKey": null
},
v30 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "futureState",
  "args": null,
  "storageKey": null
},
v31 = [
  (v3/*: any*/),
  (v2/*: any*/)
],
v32 = {
  "kind": "LinkedField",
  "alias": null,
  "name": "locationHierarchy",
  "storageKey": null,
  "args": null,
  "concreteType": "Location",
  "plural": true,
  "selections": [
    (v2/*: any*/),
    (v3/*: any*/),
    {
      "kind": "LinkedField",
      "alias": null,
      "name": "locationType",
      "storageKey": null,
      "args": null,
      "concreteType": "LocationType",
      "plural": false,
      "selections": (v31/*: any*/)
    }
  ]
},
v33 = {
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
        (v3/*: any*/),
        (v13/*: any*/)
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
        (v3/*: any*/),
        {
          "kind": "LinkedField",
          "alias": null,
          "name": "equipmentType",
          "storageKey": null,
          "args": null,
          "concreteType": "EquipmentType",
          "plural": false,
          "selections": (v5/*: any*/)
        }
      ]
    }
  ]
},
v34 = {
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
        {
          "kind": "ScalarField",
          "alias": null,
          "name": "role",
          "args": null,
          "storageKey": null
        },
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
      "selections": (v31/*: any*/)
    },
    (v2/*: any*/)
  ]
},
v35 = {
  "kind": "LinkedField",
  "alias": null,
  "name": "properties",
  "storageKey": null,
  "args": null,
  "concreteType": "Property",
  "plural": true,
  "selections": [
    (v2/*: any*/),
    {
      "kind": "LinkedField",
      "alias": null,
      "name": "propertyType",
      "storageKey": null,
      "args": null,
      "concreteType": "PropertyType",
      "plural": false,
      "selections": (v23/*: any*/)
    },
    (v14/*: any*/),
    (v15/*: any*/),
    (v17/*: any*/),
    (v16/*: any*/),
    (v18/*: any*/),
    (v19/*: any*/),
    (v20/*: any*/),
    (v21/*: any*/),
    {
      "kind": "LinkedField",
      "alias": null,
      "name": "nodeValue",
      "storageKey": null,
      "args": null,
      "concreteType": null,
      "plural": false,
      "selections": [
        (v11/*: any*/),
        (v2/*: any*/),
        (v3/*: any*/)
      ]
    }
  ]
},
v36 = {
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
      "selections": (v26/*: any*/)
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
        (v3/*: any*/),
        {
          "kind": "LinkedField",
          "alias": null,
          "name": "equipmentType",
          "storageKey": null,
          "args": null,
          "concreteType": "EquipmentType",
          "plural": false,
          "selections": [
            (v2/*: any*/),
            (v3/*: any*/),
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
                (v3/*: any*/),
                (v13/*: any*/),
                (v28/*: any*/),
                (v29/*: any*/)
              ]
            }
          ]
        }
      ]
    },
    {
      "kind": "LinkedField",
      "alias": null,
      "name": "link",
      "storageKey": null,
      "args": null,
      "concreteType": "Link",
      "plural": false,
      "selections": [
        (v2/*: any*/),
        (v30/*: any*/),
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
                (v3/*: any*/),
                (v13/*: any*/),
                {
                  "kind": "LinkedField",
                  "alias": null,
                  "name": "portType",
                  "storageKey": null,
                  "args": null,
                  "concreteType": "EquipmentPortType",
                  "plural": false,
                  "selections": [
                    (v25/*: any*/),
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
                (v3/*: any*/),
                (v30/*: any*/),
                {
                  "kind": "LinkedField",
                  "alias": null,
                  "name": "equipmentType",
                  "storageKey": null,
                  "args": null,
                  "concreteType": "EquipmentType",
                  "plural": false,
                  "selections": [
                    (v2/*: any*/),
                    (v3/*: any*/),
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
                        (v3/*: any*/),
                        (v13/*: any*/),
                        (v29/*: any*/),
                        (v28/*: any*/)
                      ]
                    }
                  ]
                },
                (v32/*: any*/),
                (v33/*: any*/)
              ]
            },
            (v34/*: any*/)
          ]
        },
        {
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
        (v35/*: any*/),
        {
          "kind": "LinkedField",
          "alias": null,
          "name": "services",
          "storageKey": null,
          "args": null,
          "concreteType": "Service",
          "plural": true,
          "selections": (v5/*: any*/)
        }
      ]
    },
    (v35/*: any*/),
    (v34/*: any*/)
  ]
},
v37 = {
  "kind": "LinkedField",
  "alias": null,
  "name": "equipmentType",
  "storageKey": null,
  "args": null,
  "concreteType": "EquipmentType",
  "plural": false,
  "selections": [
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
        (v3/*: any*/),
        (v13/*: any*/),
        (v29/*: any*/)
      ]
    },
    (v2/*: any*/)
  ]
},
v38 = [
  (v2/*: any*/),
  {
    "kind": "ScalarField",
    "alias": null,
    "name": "fileName",
    "args": null,
    "storageKey": null
  },
  (v22/*: any*/),
  {
    "kind": "ScalarField",
    "alias": null,
    "name": "sizeInBytes",
    "args": null,
    "storageKey": null
  },
  {
    "kind": "ScalarField",
    "alias": null,
    "name": "uploaded",
    "args": null,
    "storageKey": null
  },
  {
    "kind": "ScalarField",
    "alias": null,
    "name": "fileType",
    "args": null,
    "storageKey": null
  },
  {
    "kind": "ScalarField",
    "alias": null,
    "name": "storeKey",
    "args": null,
    "storageKey": null
  }
];
return {
  "kind": "Request",
  "fragment": {
    "kind": "Fragment",
    "name": "EquipmentPropertiesCardQuery",
    "type": "Query",
    "metadata": null,
    "argumentDefinitions": (v0/*: any*/),
    "selections": [
      {
        "kind": "LinkedField",
        "alias": "equipment",
        "name": "node",
        "storageKey": null,
        "args": (v1/*: any*/),
        "concreteType": null,
        "plural": false,
        "selections": [
          {
            "kind": "InlineFragment",
            "type": "Equipment",
            "selections": [
              (v2/*: any*/),
              (v3/*: any*/),
              {
                "kind": "LinkedField",
                "alias": null,
                "name": "equipmentType",
                "storageKey": null,
                "args": null,
                "concreteType": "EquipmentType",
                "plural": false,
                "selections": [
                  (v2/*: any*/),
                  (v3/*: any*/),
                  {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "propertyTypes",
                    "storageKey": null,
                    "args": null,
                    "concreteType": "PropertyType",
                    "plural": true,
                    "selections": [
                      {
                        "kind": "FragmentSpread",
                        "name": "PropertyTypeFormField_propertyType",
                        "args": null
                      },
                      {
                        "kind": "FragmentSpread",
                        "name": "DynamicPropertiesGrid_propertyTypes",
                        "args": null
                      }
                    ]
                  },
                  {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "positionDefinitions",
                    "storageKey": null,
                    "args": null,
                    "concreteType": "EquipmentPositionDefinition",
                    "plural": true,
                    "selections": [
                      (v2/*: any*/),
                      {
                        "kind": "FragmentSpread",
                        "name": "PositionDefinitionsTable_positionDefinitions",
                        "args": null
                      }
                    ]
                  },
                  {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "portDefinitions",
                    "storageKey": null,
                    "args": null,
                    "concreteType": "EquipmentPortDefinition",
                    "plural": true,
                    "selections": (v4/*: any*/)
                  }
                ]
              },
              (v6/*: any*/),
              {
                "kind": "LinkedField",
                "alias": null,
                "name": "parentPosition",
                "storageKey": null,
                "args": null,
                "concreteType": "EquipmentPosition",
                "plural": false,
                "selections": [
                  {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "parentEquipment",
                    "storageKey": null,
                    "args": null,
                    "concreteType": "Equipment",
                    "plural": false,
                    "selections": [
                      (v7/*: any*/)
                    ]
                  }
                ]
              },
              {
                "kind": "LinkedField",
                "alias": null,
                "name": "positions",
                "storageKey": null,
                "args": null,
                "concreteType": "EquipmentPosition",
                "plural": true,
                "selections": [
                  (v8/*: any*/)
                ]
              },
              {
                "kind": "LinkedField",
                "alias": null,
                "name": "properties",
                "storageKey": null,
                "args": null,
                "concreteType": "Property",
                "plural": true,
                "selections": [
                  {
                    "kind": "FragmentSpread",
                    "name": "PropertyFormField_property",
                    "args": null
                  },
                  {
                    "kind": "FragmentSpread",
                    "name": "DynamicPropertiesGrid_properties",
                    "args": null
                  }
                ]
              },
              {
                "kind": "LinkedField",
                "alias": null,
                "name": "services",
                "storageKey": null,
                "args": null,
                "concreteType": "Service",
                "plural": true,
                "selections": [
                  (v2/*: any*/),
                  (v3/*: any*/),
                  (v9/*: any*/),
                  {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "customer",
                    "storageKey": null,
                    "args": null,
                    "concreteType": "Customer",
                    "plural": false,
                    "selections": [
                      (v3/*: any*/)
                    ]
                  },
                  (v10/*: any*/)
                ]
              },
              {
                "kind": "FragmentSpread",
                "name": "EquipmentPortsTable_equipment",
                "args": null
              },
              {
                "kind": "FragmentSpread",
                "name": "EquipmentBreadcrumbs_equipment",
                "args": null
              },
              {
                "kind": "FragmentSpread",
                "name": "EquipmentPositionsGrid_equipment",
                "args": null
              },
              {
                "kind": "FragmentSpread",
                "name": "EquipmentDocumentsCard_equipment",
                "args": null
              }
            ]
          }
        ]
      }
    ]
  },
  "operation": {
    "kind": "Operation",
    "name": "EquipmentPropertiesCardQuery",
    "argumentDefinitions": (v0/*: any*/),
    "selections": [
      {
        "kind": "LinkedField",
        "alias": "equipment",
        "name": "node",
        "storageKey": null,
        "args": (v1/*: any*/),
        "concreteType": null,
        "plural": false,
        "selections": [
          (v11/*: any*/),
          (v2/*: any*/),
          {
            "kind": "InlineFragment",
            "type": "Equipment",
            "selections": [
              (v3/*: any*/),
              {
                "kind": "LinkedField",
                "alias": null,
                "name": "equipmentType",
                "storageKey": null,
                "args": null,
                "concreteType": "EquipmentType",
                "plural": false,
                "selections": [
                  (v2/*: any*/),
                  (v3/*: any*/),
                  {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "portDefinitions",
                    "storageKey": null,
                    "args": null,
                    "concreteType": "EquipmentPortDefinition",
                    "plural": true,
                    "selections": (v26/*: any*/)
                  },
                  (v24/*: any*/),
                  {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "positionDefinitions",
                    "storageKey": null,
                    "args": null,
                    "concreteType": "EquipmentPositionDefinition",
                    "plural": true,
                    "selections": (v27/*: any*/)
                  }
                ]
              },
              (v36/*: any*/),
              {
                "kind": "LinkedField",
                "alias": null,
                "name": "positions",
                "storageKey": null,
                "args": null,
                "concreteType": "EquipmentPosition",
                "plural": true,
                "selections": [
                  {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "attachedEquipment",
                    "storageKey": null,
                    "args": null,
                    "concreteType": "Equipment",
                    "plural": false,
                    "selections": [
                      (v2/*: any*/),
                      (v3/*: any*/),
                      (v36/*: any*/),
                      (v37/*: any*/),
                      {
                        "kind": "LinkedField",
                        "alias": null,
                        "name": "positions",
                        "storageKey": null,
                        "args": null,
                        "concreteType": "EquipmentPosition",
                        "plural": true,
                        "selections": [
                          {
                            "kind": "LinkedField",
                            "alias": null,
                            "name": "attachedEquipment",
                            "storageKey": null,
                            "args": null,
                            "concreteType": "Equipment",
                            "plural": false,
                            "selections": [
                              (v2/*: any*/),
                              (v3/*: any*/),
                              (v36/*: any*/),
                              (v37/*: any*/),
                              {
                                "kind": "LinkedField",
                                "alias": null,
                                "name": "positions",
                                "storageKey": null,
                                "args": null,
                                "concreteType": "EquipmentPosition",
                                "plural": true,
                                "selections": [
                                  {
                                    "kind": "LinkedField",
                                    "alias": null,
                                    "name": "attachedEquipment",
                                    "storageKey": null,
                                    "args": null,
                                    "concreteType": "Equipment",
                                    "plural": false,
                                    "selections": [
                                      (v2/*: any*/),
                                      (v3/*: any*/),
                                      (v36/*: any*/),
                                      (v37/*: any*/),
                                      {
                                        "kind": "LinkedField",
                                        "alias": null,
                                        "name": "positions",
                                        "storageKey": null,
                                        "args": null,
                                        "concreteType": "EquipmentPosition",
                                        "plural": true,
                                        "selections": [
                                          {
                                            "kind": "LinkedField",
                                            "alias": null,
                                            "name": "attachedEquipment",
                                            "storageKey": null,
                                            "args": null,
                                            "concreteType": "Equipment",
                                            "plural": false,
                                            "selections": [
                                              (v2/*: any*/),
                                              (v3/*: any*/),
                                              (v36/*: any*/),
                                              (v37/*: any*/)
                                            ]
                                          },
                                          (v2/*: any*/)
                                        ]
                                      }
                                    ]
                                  },
                                  (v2/*: any*/)
                                ]
                              }
                            ]
                          },
                          (v2/*: any*/)
                        ]
                      },
                      (v30/*: any*/),
                      {
                        "kind": "LinkedField",
                        "alias": null,
                        "name": "services",
                        "storageKey": null,
                        "args": null,
                        "concreteType": "Service",
                        "plural": true,
                        "selections": (v4/*: any*/)
                      }
                    ]
                  },
                  (v2/*: any*/),
                  {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "definition",
                    "storageKey": null,
                    "args": null,
                    "concreteType": "EquipmentPositionDefinition",
                    "plural": false,
                    "selections": (v27/*: any*/)
                  },
                  (v8/*: any*/)
                ]
              },
              (v32/*: any*/),
              (v33/*: any*/),
              (v6/*: any*/),
              {
                "kind": "LinkedField",
                "alias": null,
                "name": "parentPosition",
                "storageKey": null,
                "args": null,
                "concreteType": "EquipmentPosition",
                "plural": false,
                "selections": [
                  {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "parentEquipment",
                    "storageKey": null,
                    "args": null,
                    "concreteType": "Equipment",
                    "plural": false,
                    "selections": [
                      (v7/*: any*/),
                      (v2/*: any*/)
                    ]
                  },
                  (v2/*: any*/)
                ]
              },
              (v35/*: any*/),
              {
                "kind": "LinkedField",
                "alias": null,
                "name": "services",
                "storageKey": null,
                "args": null,
                "concreteType": "Service",
                "plural": true,
                "selections": [
                  (v2/*: any*/),
                  (v3/*: any*/),
                  (v9/*: any*/),
                  {
                    "kind": "LinkedField",
                    "alias": null,
                    "name": "customer",
                    "storageKey": null,
                    "args": null,
                    "concreteType": "Customer",
                    "plural": false,
                    "selections": (v31/*: any*/)
                  },
                  (v10/*: any*/)
                ]
              },
              {
                "kind": "LinkedField",
                "alias": null,
                "name": "images",
                "storageKey": null,
                "args": null,
                "concreteType": "File",
                "plural": true,
                "selections": (v38/*: any*/)
              },
              {
                "kind": "LinkedField",
                "alias": null,
                "name": "files",
                "storageKey": null,
                "args": null,
                "concreteType": "File",
                "plural": true,
                "selections": (v38/*: any*/)
              },
              {
                "kind": "LinkedField",
                "alias": null,
                "name": "hyperlinks",
                "storageKey": null,
                "args": null,
                "concreteType": "Hyperlink",
                "plural": true,
                "selections": [
                  (v2/*: any*/),
                  (v22/*: any*/),
                  {
                    "kind": "ScalarField",
                    "alias": null,
                    "name": "url",
                    "args": null,
                    "storageKey": null
                  },
                  {
                    "kind": "ScalarField",
                    "alias": null,
                    "name": "displayName",
                    "args": null,
                    "storageKey": null
                  },
                  {
                    "kind": "ScalarField",
                    "alias": null,
                    "name": "createTime",
                    "args": null,
                    "storageKey": null
                  }
                ]
              }
            ]
          }
        ]
      }
    ]
  },
  "params": {
    "operationKind": "query",
    "name": "EquipmentPropertiesCardQuery",
    "id": null,
    "text": "query EquipmentPropertiesCardQuery(\n  $equipmentId: ID!\n) {\n  equipment: node(id: $equipmentId) {\n    __typename\n    ... on Equipment {\n      id\n      name\n      ...EquipmentPortsTable_equipment\n      equipmentType {\n        id\n        name\n        propertyTypes {\n          ...PropertyTypeFormField_propertyType\n          ...DynamicPropertiesGrid_propertyTypes\n          id\n        }\n        positionDefinitions {\n          id\n          ...PositionDefinitionsTable_positionDefinitions\n        }\n        portDefinitions {\n          id\n        }\n      }\n      ...EquipmentBreadcrumbs_equipment\n      parentLocation {\n        id\n        name\n      }\n      parentPosition {\n        parentEquipment {\n          parentLocation {\n            id\n          }\n          id\n        }\n        id\n      }\n      ...EquipmentPositionsGrid_equipment\n      positions {\n        parentEquipment {\n          id\n        }\n        id\n      }\n      properties {\n        ...PropertyFormField_property\n        ...DynamicPropertiesGrid_properties\n        id\n      }\n      services {\n        id\n        name\n        externalId\n        customer {\n          name\n          id\n        }\n        serviceType {\n          id\n          name\n        }\n      }\n      ...EquipmentDocumentsCard_equipment\n    }\n    id\n  }\n}\n\nfragment AddToEquipmentDialog_parentEquipment on Equipment {\n  id\n  locationHierarchy {\n    id\n  }\n}\n\nfragment DocumentTable_files on File {\n  id\n  fileName\n  category\n  ...FileAttachment_file\n}\n\nfragment DocumentTable_hyperlinks on Hyperlink {\n  id\n  category\n  url\n  displayName\n  ...HyperlinkTableRow_hyperlink\n}\n\nfragment DynamicPropertiesGrid_properties on Property {\n  ...PropertyFormField_property\n  propertyType {\n    id\n    index\n  }\n}\n\nfragment DynamicPropertiesGrid_propertyTypes on PropertyType {\n  id\n  name\n  index\n  isInstanceProperty\n  type\n  nodeType\n  stringValue\n  intValue\n  booleanValue\n  latitudeValue\n  longitudeValue\n  rangeFromValue\n  rangeToValue\n  floatValue\n}\n\nfragment EntityDocumentsTable_files on File {\n  ...DocumentTable_files\n}\n\nfragment EntityDocumentsTable_hyperlinks on Hyperlink {\n  ...DocumentTable_hyperlinks\n}\n\nfragment EquipmentBreadcrumbs_equipment on Equipment {\n  id\n  name\n  equipmentType {\n    id\n    name\n  }\n  locationHierarchy {\n    id\n    name\n    locationType {\n      name\n      id\n    }\n  }\n  positionHierarchy {\n    id\n    definition {\n      id\n      name\n      visibleLabel\n    }\n    parentEquipment {\n      id\n      name\n      equipmentType {\n        id\n        name\n      }\n    }\n  }\n}\n\nfragment EquipmentDocumentsCard_equipment on Equipment {\n  id\n  images {\n    ...EntityDocumentsTable_files\n    id\n  }\n  files {\n    ...EntityDocumentsTable_files\n    id\n  }\n  hyperlinks {\n    ...EntityDocumentsTable_hyperlinks\n    id\n  }\n}\n\nfragment EquipmentPortsTable_equipment on Equipment {\n  id\n  name\n  equipmentType {\n    id\n    name\n    portDefinitions {\n      id\n      name\n      index\n      visibleLabel\n      portType {\n        id\n        name\n        propertyTypes {\n          id\n          name\n          type\n          nodeType\n          index\n          stringValue\n          intValue\n          booleanValue\n          floatValue\n          latitudeValue\n          longitudeValue\n          rangeFromValue\n          rangeToValue\n          isEditable\n          isInstanceProperty\n          isMandatory\n          category\n          isDeleted\n        }\n        linkPropertyTypes {\n          id\n          name\n          type\n          nodeType\n          index\n          stringValue\n          intValue\n          booleanValue\n          floatValue\n          latitudeValue\n          longitudeValue\n          rangeFromValue\n          rangeToValue\n          isEditable\n          isInstanceProperty\n          isMandatory\n          category\n          isDeleted\n        }\n      }\n    }\n  }\n  ports {\n    id\n    definition {\n      id\n      name\n      index\n      visibleLabel\n      portType {\n        id\n        name\n        propertyTypes {\n          id\n          name\n          type\n          nodeType\n          index\n          stringValue\n          intValue\n          booleanValue\n          floatValue\n          latitudeValue\n          longitudeValue\n          rangeFromValue\n          rangeToValue\n          isEditable\n          isInstanceProperty\n          isMandatory\n          category\n          isDeleted\n        }\n        linkPropertyTypes {\n          id\n          name\n          type\n          nodeType\n          index\n          stringValue\n          intValue\n          booleanValue\n          floatValue\n          latitudeValue\n          longitudeValue\n          rangeFromValue\n          rangeToValue\n          isEditable\n          isInstanceProperty\n          isMandatory\n          category\n          isDeleted\n        }\n      }\n    }\n    parentEquipment {\n      id\n      name\n      equipmentType {\n        id\n        name\n        portDefinitions {\n          id\n          name\n          visibleLabel\n          portType {\n            id\n            name\n          }\n          bandwidth\n        }\n      }\n    }\n    link {\n      id\n      futureState\n      ports {\n        id\n        definition {\n          id\n          name\n          visibleLabel\n          portType {\n            linkPropertyTypes {\n              id\n              name\n              type\n              nodeType\n              index\n              stringValue\n              intValue\n              booleanValue\n              floatValue\n              latitudeValue\n              longitudeValue\n              rangeFromValue\n              rangeToValue\n              isEditable\n              isInstanceProperty\n              isMandatory\n              category\n              isDeleted\n            }\n            id\n          }\n        }\n        parentEquipment {\n          id\n          name\n          futureState\n          equipmentType {\n            id\n            name\n            portDefinitions {\n              id\n              name\n              visibleLabel\n              bandwidth\n              portType {\n                id\n                name\n              }\n            }\n          }\n          ...EquipmentBreadcrumbs_equipment\n        }\n        serviceEndpoints {\n          definition {\n            role\n            id\n          }\n          service {\n            name\n            id\n          }\n          id\n        }\n      }\n      workOrder {\n        id\n        status\n      }\n      properties {\n        id\n        propertyType {\n          id\n          name\n          type\n          nodeType\n          index\n          stringValue\n          intValue\n          booleanValue\n          floatValue\n          latitudeValue\n          longitudeValue\n          rangeFromValue\n          rangeToValue\n          isEditable\n          isInstanceProperty\n          isMandatory\n          category\n          isDeleted\n        }\n        stringValue\n        intValue\n        floatValue\n        booleanValue\n        latitudeValue\n        longitudeValue\n        rangeFromValue\n        rangeToValue\n        nodeValue {\n          __typename\n          id\n          name\n        }\n      }\n      services {\n        id\n        name\n      }\n    }\n    properties {\n      id\n      propertyType {\n        id\n        name\n        type\n        nodeType\n        index\n        stringValue\n        intValue\n        booleanValue\n        floatValue\n        latitudeValue\n        longitudeValue\n        rangeFromValue\n        rangeToValue\n        isEditable\n        isInstanceProperty\n        isMandatory\n        category\n        isDeleted\n      }\n      stringValue\n      intValue\n      floatValue\n      booleanValue\n      latitudeValue\n      longitudeValue\n      rangeFromValue\n      rangeToValue\n      nodeValue {\n        __typename\n        id\n        name\n      }\n    }\n    serviceEndpoints {\n      definition {\n        role\n        id\n      }\n      service {\n        name\n        id\n      }\n      id\n    }\n  }\n  positions {\n    attachedEquipment {\n      id\n      name\n      ports {\n        id\n        definition {\n          id\n          name\n          index\n          visibleLabel\n          portType {\n            id\n            name\n            propertyTypes {\n              id\n              name\n              type\n              nodeType\n              index\n              stringValue\n              intValue\n              booleanValue\n              floatValue\n              latitudeValue\n              longitudeValue\n              rangeFromValue\n              rangeToValue\n              isEditable\n              isInstanceProperty\n              isMandatory\n              category\n              isDeleted\n            }\n            linkPropertyTypes {\n              id\n              name\n              type\n              nodeType\n              index\n              stringValue\n              intValue\n              booleanValue\n              floatValue\n              latitudeValue\n              longitudeValue\n              rangeFromValue\n              rangeToValue\n              isEditable\n              isInstanceProperty\n              isMandatory\n              category\n              isDeleted\n            }\n          }\n        }\n        parentEquipment {\n          id\n          name\n          equipmentType {\n            id\n            name\n            portDefinitions {\n              id\n              name\n              visibleLabel\n              portType {\n                id\n                name\n              }\n              bandwidth\n            }\n          }\n        }\n        link {\n          id\n          futureState\n          ports {\n            id\n            definition {\n              id\n              name\n              visibleLabel\n              portType {\n                linkPropertyTypes {\n                  id\n                  name\n                  type\n                  nodeType\n                  index\n                  stringValue\n                  intValue\n                  booleanValue\n                  floatValue\n                  latitudeValue\n                  longitudeValue\n                  rangeFromValue\n                  rangeToValue\n                  isEditable\n                  isInstanceProperty\n                  isMandatory\n                  category\n                  isDeleted\n                }\n                id\n              }\n            }\n            parentEquipment {\n              id\n              name\n              futureState\n              equipmentType {\n                id\n                name\n                portDefinitions {\n                  id\n                  name\n                  visibleLabel\n                  bandwidth\n                  portType {\n                    id\n                    name\n                  }\n                }\n              }\n              ...EquipmentBreadcrumbs_equipment\n            }\n            serviceEndpoints {\n              definition {\n                role\n                id\n              }\n              service {\n                name\n                id\n              }\n              id\n            }\n          }\n          workOrder {\n            id\n            status\n          }\n          properties {\n            id\n            propertyType {\n              id\n              name\n              type\n              nodeType\n              index\n              stringValue\n              intValue\n              booleanValue\n              floatValue\n              latitudeValue\n              longitudeValue\n              rangeFromValue\n              rangeToValue\n              isEditable\n              isInstanceProperty\n              isMandatory\n              category\n              isDeleted\n            }\n            stringValue\n            intValue\n            floatValue\n            booleanValue\n            latitudeValue\n            longitudeValue\n            rangeFromValue\n            rangeToValue\n            nodeValue {\n              __typename\n              id\n              name\n            }\n          }\n          services {\n            id\n            name\n          }\n        }\n        properties {\n          id\n          propertyType {\n            id\n            name\n            type\n            nodeType\n            index\n            stringValue\n            intValue\n            booleanValue\n            floatValue\n            latitudeValue\n            longitudeValue\n            rangeFromValue\n            rangeToValue\n            isEditable\n            isInstanceProperty\n            isMandatory\n            category\n            isDeleted\n          }\n          stringValue\n          intValue\n          floatValue\n          booleanValue\n          latitudeValue\n          longitudeValue\n          rangeFromValue\n          rangeToValue\n          nodeValue {\n            __typename\n            id\n            name\n          }\n        }\n        serviceEndpoints {\n          definition {\n            role\n            id\n          }\n          service {\n            name\n            id\n          }\n          id\n        }\n      }\n      equipmentType {\n        portDefinitions {\n          id\n          name\n          visibleLabel\n          bandwidth\n        }\n        id\n      }\n      positions {\n        attachedEquipment {\n          id\n          name\n          ports {\n            id\n            definition {\n              id\n              name\n              index\n              visibleLabel\n              portType {\n                id\n                name\n                propertyTypes {\n                  id\n                  name\n                  type\n                  nodeType\n                  index\n                  stringValue\n                  intValue\n                  booleanValue\n                  floatValue\n                  latitudeValue\n                  longitudeValue\n                  rangeFromValue\n                  rangeToValue\n                  isEditable\n                  isInstanceProperty\n                  isMandatory\n                  category\n                  isDeleted\n                }\n                linkPropertyTypes {\n                  id\n                  name\n                  type\n                  nodeType\n                  index\n                  stringValue\n                  intValue\n                  booleanValue\n                  floatValue\n                  latitudeValue\n                  longitudeValue\n                  rangeFromValue\n                  rangeToValue\n                  isEditable\n                  isInstanceProperty\n                  isMandatory\n                  category\n                  isDeleted\n                }\n              }\n            }\n            parentEquipment {\n              id\n              name\n              equipmentType {\n                id\n                name\n                portDefinitions {\n                  id\n                  name\n                  visibleLabel\n                  portType {\n                    id\n                    name\n                  }\n                  bandwidth\n                }\n              }\n            }\n            link {\n              id\n              futureState\n              ports {\n                id\n                definition {\n                  id\n                  name\n                  visibleLabel\n                  portType {\n                    linkPropertyTypes {\n                      id\n                      name\n                      type\n                      nodeType\n                      index\n                      stringValue\n                      intValue\n                      booleanValue\n                      floatValue\n                      latitudeValue\n                      longitudeValue\n                      rangeFromValue\n                      rangeToValue\n                      isEditable\n                      isInstanceProperty\n                      isMandatory\n                      category\n                      isDeleted\n                    }\n                    id\n                  }\n                }\n                parentEquipment {\n                  id\n                  name\n                  futureState\n                  equipmentType {\n                    id\n                    name\n                    portDefinitions {\n                      id\n                      name\n                      visibleLabel\n                      bandwidth\n                      portType {\n                        id\n                        name\n                      }\n                    }\n                  }\n                  ...EquipmentBreadcrumbs_equipment\n                }\n                serviceEndpoints {\n                  definition {\n                    role\n                    id\n                  }\n                  service {\n                    name\n                    id\n                  }\n                  id\n                }\n              }\n              workOrder {\n                id\n                status\n              }\n              properties {\n                id\n                propertyType {\n                  id\n                  name\n                  type\n                  nodeType\n                  index\n                  stringValue\n                  intValue\n                  booleanValue\n                  floatValue\n                  latitudeValue\n                  longitudeValue\n                  rangeFromValue\n                  rangeToValue\n                  isEditable\n                  isInstanceProperty\n                  isMandatory\n                  category\n                  isDeleted\n                }\n                stringValue\n                intValue\n                floatValue\n                booleanValue\n                latitudeValue\n                longitudeValue\n                rangeFromValue\n                rangeToValue\n                nodeValue {\n                  __typename\n                  id\n                  name\n                }\n              }\n              services {\n                id\n                name\n              }\n            }\n            properties {\n              id\n              propertyType {\n                id\n                name\n                type\n                nodeType\n                index\n                stringValue\n                intValue\n                booleanValue\n                floatValue\n                latitudeValue\n                longitudeValue\n                rangeFromValue\n                rangeToValue\n                isEditable\n                isInstanceProperty\n                isMandatory\n                category\n                isDeleted\n              }\n              stringValue\n              intValue\n              floatValue\n              booleanValue\n              latitudeValue\n              longitudeValue\n              rangeFromValue\n              rangeToValue\n              nodeValue {\n                __typename\n                id\n                name\n              }\n            }\n            serviceEndpoints {\n              definition {\n                role\n                id\n              }\n              service {\n                name\n                id\n              }\n              id\n            }\n          }\n          equipmentType {\n            portDefinitions {\n              id\n              name\n              visibleLabel\n              bandwidth\n            }\n            id\n          }\n          positions {\n            attachedEquipment {\n              id\n              name\n              ports {\n                id\n                definition {\n                  id\n                  name\n                  index\n                  visibleLabel\n                  portType {\n                    id\n                    name\n                    propertyTypes {\n                      id\n                      name\n                      type\n                      nodeType\n                      index\n                      stringValue\n                      intValue\n                      booleanValue\n                      floatValue\n                      latitudeValue\n                      longitudeValue\n                      rangeFromValue\n                      rangeToValue\n                      isEditable\n                      isInstanceProperty\n                      isMandatory\n                      category\n                      isDeleted\n                    }\n                    linkPropertyTypes {\n                      id\n                      name\n                      type\n                      nodeType\n                      index\n                      stringValue\n                      intValue\n                      booleanValue\n                      floatValue\n                      latitudeValue\n                      longitudeValue\n                      rangeFromValue\n                      rangeToValue\n                      isEditable\n                      isInstanceProperty\n                      isMandatory\n                      category\n                      isDeleted\n                    }\n                  }\n                }\n                parentEquipment {\n                  id\n                  name\n                  equipmentType {\n                    id\n                    name\n                    portDefinitions {\n                      id\n                      name\n                      visibleLabel\n                      portType {\n                        id\n                        name\n                      }\n                      bandwidth\n                    }\n                  }\n                }\n                link {\n                  id\n                  futureState\n                  ports {\n                    id\n                    definition {\n                      id\n                      name\n                      visibleLabel\n                      portType {\n                        linkPropertyTypes {\n                          id\n                          name\n                          type\n                          nodeType\n                          index\n                          stringValue\n                          intValue\n                          booleanValue\n                          floatValue\n                          latitudeValue\n                          longitudeValue\n                          rangeFromValue\n                          rangeToValue\n                          isEditable\n                          isInstanceProperty\n                          isMandatory\n                          category\n                          isDeleted\n                        }\n                        id\n                      }\n                    }\n                    parentEquipment {\n                      id\n                      name\n                      futureState\n                      equipmentType {\n                        id\n                        name\n                        portDefinitions {\n                          id\n                          name\n                          visibleLabel\n                          bandwidth\n                          portType {\n                            id\n                            name\n                          }\n                        }\n                      }\n                      ...EquipmentBreadcrumbs_equipment\n                    }\n                    serviceEndpoints {\n                      definition {\n                        role\n                        id\n                      }\n                      service {\n                        name\n                        id\n                      }\n                      id\n                    }\n                  }\n                  workOrder {\n                    id\n                    status\n                  }\n                  properties {\n                    id\n                    propertyType {\n                      id\n                      name\n                      type\n                      nodeType\n                      index\n                      stringValue\n                      intValue\n                      booleanValue\n                      floatValue\n                      latitudeValue\n                      longitudeValue\n                      rangeFromValue\n                      rangeToValue\n                      isEditable\n                      isInstanceProperty\n                      isMandatory\n                      category\n                      isDeleted\n                    }\n                    stringValue\n                    intValue\n                    floatValue\n                    booleanValue\n                    latitudeValue\n                    longitudeValue\n                    rangeFromValue\n                    rangeToValue\n                    nodeValue {\n                      __typename\n                      id\n                      name\n                    }\n                  }\n                  services {\n                    id\n                    name\n                  }\n                }\n                properties {\n                  id\n                  propertyType {\n                    id\n                    name\n                    type\n                    nodeType\n                    index\n                    stringValue\n                    intValue\n                    booleanValue\n                    floatValue\n                    latitudeValue\n                    longitudeValue\n                    rangeFromValue\n                    rangeToValue\n                    isEditable\n                    isInstanceProperty\n                    isMandatory\n                    category\n                    isDeleted\n                  }\n                  stringValue\n                  intValue\n                  floatValue\n                  booleanValue\n                  latitudeValue\n                  longitudeValue\n                  rangeFromValue\n                  rangeToValue\n                  nodeValue {\n                    __typename\n                    id\n                    name\n                  }\n                }\n                serviceEndpoints {\n                  definition {\n                    role\n                    id\n                  }\n                  service {\n                    name\n                    id\n                  }\n                  id\n                }\n              }\n              equipmentType {\n                portDefinitions {\n                  id\n                  name\n                  visibleLabel\n                  bandwidth\n                }\n                id\n              }\n              positions {\n                attachedEquipment {\n                  id\n                  name\n                  ports {\n                    id\n                    definition {\n                      id\n                      name\n                      index\n                      visibleLabel\n                      portType {\n                        id\n                        name\n                        propertyTypes {\n                          id\n                          name\n                          type\n                          nodeType\n                          index\n                          stringValue\n                          intValue\n                          booleanValue\n                          floatValue\n                          latitudeValue\n                          longitudeValue\n                          rangeFromValue\n                          rangeToValue\n                          isEditable\n                          isInstanceProperty\n                          isMandatory\n                          category\n                          isDeleted\n                        }\n                        linkPropertyTypes {\n                          id\n                          name\n                          type\n                          nodeType\n                          index\n                          stringValue\n                          intValue\n                          booleanValue\n                          floatValue\n                          latitudeValue\n                          longitudeValue\n                          rangeFromValue\n                          rangeToValue\n                          isEditable\n                          isInstanceProperty\n                          isMandatory\n                          category\n                          isDeleted\n                        }\n                      }\n                    }\n                    parentEquipment {\n                      id\n                      name\n                      equipmentType {\n                        id\n                        name\n                        portDefinitions {\n                          id\n                          name\n                          visibleLabel\n                          portType {\n                            id\n                            name\n                          }\n                          bandwidth\n                        }\n                      }\n                    }\n                    link {\n                      id\n                      futureState\n                      ports {\n                        id\n                        definition {\n                          id\n                          name\n                          visibleLabel\n                          portType {\n                            linkPropertyTypes {\n                              id\n                              name\n                              type\n                              nodeType\n                              index\n                              stringValue\n                              intValue\n                              booleanValue\n                              floatValue\n                              latitudeValue\n                              longitudeValue\n                              rangeFromValue\n                              rangeToValue\n                              isEditable\n                              isInstanceProperty\n                              isMandatory\n                              category\n                              isDeleted\n                            }\n                            id\n                          }\n                        }\n                        parentEquipment {\n                          id\n                          name\n                          futureState\n                          equipmentType {\n                            id\n                            name\n                            portDefinitions {\n                              id\n                              name\n                              visibleLabel\n                              bandwidth\n                              portType {\n                                id\n                                name\n                              }\n                            }\n                          }\n                          ...EquipmentBreadcrumbs_equipment\n                        }\n                        serviceEndpoints {\n                          definition {\n                            role\n                            id\n                          }\n                          service {\n                            name\n                            id\n                          }\n                          id\n                        }\n                      }\n                      workOrder {\n                        id\n                        status\n                      }\n                      properties {\n                        id\n                        propertyType {\n                          id\n                          name\n                          type\n                          nodeType\n                          index\n                          stringValue\n                          intValue\n                          booleanValue\n                          floatValue\n                          latitudeValue\n                          longitudeValue\n                          rangeFromValue\n                          rangeToValue\n                          isEditable\n                          isInstanceProperty\n                          isMandatory\n                          category\n                          isDeleted\n                        }\n                        stringValue\n                        intValue\n                        floatValue\n                        booleanValue\n                        latitudeValue\n                        longitudeValue\n                        rangeFromValue\n                        rangeToValue\n                        nodeValue {\n                          __typename\n                          id\n                          name\n                        }\n                      }\n                      services {\n                        id\n                        name\n                      }\n                    }\n                    properties {\n                      id\n                      propertyType {\n                        id\n                        name\n                        type\n                        nodeType\n                        index\n                        stringValue\n                        intValue\n                        booleanValue\n                        floatValue\n                        latitudeValue\n                        longitudeValue\n                        rangeFromValue\n                        rangeToValue\n                        isEditable\n                        isInstanceProperty\n                        isMandatory\n                        category\n                        isDeleted\n                      }\n                      stringValue\n                      intValue\n                      floatValue\n                      booleanValue\n                      latitudeValue\n                      longitudeValue\n                      rangeFromValue\n                      rangeToValue\n                      nodeValue {\n                        __typename\n                        id\n                        name\n                      }\n                    }\n                    serviceEndpoints {\n                      definition {\n                        role\n                        id\n                      }\n                      service {\n                        name\n                        id\n                      }\n                      id\n                    }\n                  }\n                  equipmentType {\n                    portDefinitions {\n                      id\n                      name\n                      visibleLabel\n                      bandwidth\n                    }\n                    id\n                  }\n                }\n                id\n              }\n            }\n            id\n          }\n        }\n        id\n      }\n    }\n    id\n  }\n}\n\nfragment EquipmentPositionsGrid_equipment on Equipment {\n  id\n  ...AddToEquipmentDialog_parentEquipment\n  positions {\n    id\n    definition {\n      id\n      name\n      index\n      visibleLabel\n    }\n    attachedEquipment {\n      id\n      name\n      futureState\n      services {\n        id\n      }\n    }\n    parentEquipment {\n      id\n    }\n  }\n  equipmentType {\n    positionDefinitions {\n      id\n      name\n      index\n      visibleLabel\n    }\n    id\n  }\n}\n\nfragment FileAttachment_file on File {\n  id\n  fileName\n  sizeInBytes\n  uploaded\n  fileType\n  storeKey\n  category\n  ...ImageDialog_img\n}\n\nfragment HyperlinkTableRow_hyperlink on Hyperlink {\n  id\n  category\n  url\n  displayName\n  createTime\n}\n\nfragment ImageDialog_img on File {\n  storeKey\n  fileName\n}\n\nfragment PositionDefinitionsTable_positionDefinitions on EquipmentPositionDefinition {\n  id\n  name\n  index\n  visibleLabel\n}\n\nfragment PropertyFormField_property on Property {\n  id\n  propertyType {\n    id\n    name\n    type\n    nodeType\n    index\n    stringValue\n    intValue\n    booleanValue\n    floatValue\n    latitudeValue\n    longitudeValue\n    rangeFromValue\n    rangeToValue\n    isEditable\n    isInstanceProperty\n    isMandatory\n    category\n    isDeleted\n  }\n  stringValue\n  intValue\n  floatValue\n  booleanValue\n  latitudeValue\n  longitudeValue\n  rangeFromValue\n  rangeToValue\n  nodeValue {\n    __typename\n    id\n    name\n  }\n}\n\nfragment PropertyTypeFormField_propertyType on PropertyType {\n  id\n  name\n  type\n  nodeType\n  index\n  stringValue\n  intValue\n  booleanValue\n  floatValue\n  latitudeValue\n  longitudeValue\n  rangeFromValue\n  rangeToValue\n  isEditable\n  isInstanceProperty\n  isMandatory\n  category\n  isDeleted\n}\n",
    "metadata": {}
  }
};
})();
// prettier-ignore
(node/*: any*/).hash = '45289b80018793a427858ef59546d84c';
module.exports = node;
