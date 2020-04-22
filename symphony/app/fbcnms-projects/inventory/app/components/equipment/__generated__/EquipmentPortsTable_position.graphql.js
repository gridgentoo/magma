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
type EquipmentBreadcrumbs_equipment$ref = any;
export type FutureState = "INSTALL" | "REMOVE" | "%future added value";
export type PropertyKind = "bool" | "date" | "datetime_local" | "email" | "enum" | "float" | "gps_location" | "int" | "node" | "range" | "string" | "%future added value";
export type WorkOrderStatus = "DONE" | "PENDING" | "PLANNED" | "%future added value";
import type { FragmentReference } from "relay-runtime";
declare export opaque type EquipmentPortsTable_position$ref: FragmentReference;
declare export opaque type EquipmentPortsTable_position$fragmentType: EquipmentPortsTable_position$ref;
export type EquipmentPortsTable_position = {
  +attachedEquipment: ?{
    +id: string,
    +name: string,
    +ports: $ReadOnlyArray<?{
      +id: string,
      +definition: {
        +id: string,
        +name: string,
        +index: ?number,
        +visibleLabel: ?string,
        +portType: ?{
          +id: string,
          +name: string,
          +propertyTypes: $ReadOnlyArray<?{
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
            ...
          }>,
          +linkPropertyTypes: $ReadOnlyArray<?{
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
            ...
          }>,
          ...
        },
        ...
      },
      +parentEquipment: {
        +id: string,
        +name: string,
        +equipmentType: {
          +id: string,
          +name: string,
          +portDefinitions: $ReadOnlyArray<?{
            +id: string,
            +name: string,
            +visibleLabel: ?string,
            +portType: ?{
              +id: string,
              +name: string,
              ...
            },
            +bandwidth: ?string,
            ...
          }>,
          ...
        },
        ...
      },
      +link: ?{
        +id: string,
        +futureState: ?FutureState,
        +ports: $ReadOnlyArray<?{
          +id: string,
          +definition: {
            +id: string,
            +name: string,
            +visibleLabel: ?string,
            +portType: ?{
              +linkPropertyTypes: $ReadOnlyArray<?{
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
                ...
              }>,
              ...
            },
            ...
          },
          +parentEquipment: {
            +id: string,
            +name: string,
            +futureState: ?FutureState,
            +equipmentType: {
              +id: string,
              +name: string,
              +portDefinitions: $ReadOnlyArray<?{
                +id: string,
                +name: string,
                +visibleLabel: ?string,
                +bandwidth: ?string,
                +portType: ?{
                  +id: string,
                  +name: string,
                  ...
                },
                ...
              }>,
              ...
            },
            +$fragmentRefs: EquipmentBreadcrumbs_equipment$ref,
            ...
          },
          +serviceEndpoints: $ReadOnlyArray<{
            +definition: {
              +role: ?string,
              ...
            },
            +service: {
              +name: string,
              ...
            },
            ...
          }>,
          ...
        }>,
        +workOrder: ?{
          +id: string,
          +status: WorkOrderStatus,
          ...
        },
        +properties: $ReadOnlyArray<?{
          +id: string,
          +propertyType: {
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
            ...
          },
          +stringValue: ?string,
          +intValue: ?number,
          +floatValue: ?number,
          +booleanValue: ?boolean,
          +latitudeValue: ?number,
          +longitudeValue: ?number,
          +rangeFromValue: ?number,
          +rangeToValue: ?number,
          +nodeValue: ?{
            +id: string,
            +name: string,
            ...
          },
          ...
        }>,
        +services: $ReadOnlyArray<?{
          +id: string,
          +name: string,
          ...
        }>,
        ...
      },
      +properties: $ReadOnlyArray<{
        +id: string,
        +propertyType: {
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
          ...
        },
        +stringValue: ?string,
        +intValue: ?number,
        +floatValue: ?number,
        +booleanValue: ?boolean,
        +latitudeValue: ?number,
        +longitudeValue: ?number,
        +rangeFromValue: ?number,
        +rangeToValue: ?number,
        +nodeValue: ?{
          +id: string,
          +name: string,
          ...
        },
        ...
      }>,
      +serviceEndpoints: $ReadOnlyArray<{
        +definition: {
          +role: ?string,
          ...
        },
        +service: {
          +name: string,
          ...
        },
        ...
      }>,
      ...
    }>,
    +equipmentType: {
      +portDefinitions: $ReadOnlyArray<?{
        +id: string,
        +name: string,
        +visibleLabel: ?string,
        +bandwidth: ?string,
        ...
      }>,
      ...
    },
    +positions: $ReadOnlyArray<?{
      +attachedEquipment: ?{
        +id: string,
        +name: string,
        +ports: $ReadOnlyArray<?{
          +id: string,
          +definition: {
            +id: string,
            +name: string,
            +index: ?number,
            +visibleLabel: ?string,
            +portType: ?{
              +id: string,
              +name: string,
              +propertyTypes: $ReadOnlyArray<?{
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
                ...
              }>,
              +linkPropertyTypes: $ReadOnlyArray<?{
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
                ...
              }>,
              ...
            },
            ...
          },
          +parentEquipment: {
            +id: string,
            +name: string,
            +equipmentType: {
              +id: string,
              +name: string,
              +portDefinitions: $ReadOnlyArray<?{
                +id: string,
                +name: string,
                +visibleLabel: ?string,
                +portType: ?{
                  +id: string,
                  +name: string,
                  ...
                },
                +bandwidth: ?string,
                ...
              }>,
              ...
            },
            ...
          },
          +link: ?{
            +id: string,
            +futureState: ?FutureState,
            +ports: $ReadOnlyArray<?{
              +id: string,
              +definition: {
                +id: string,
                +name: string,
                +visibleLabel: ?string,
                +portType: ?{
                  +linkPropertyTypes: $ReadOnlyArray<?{
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
                    ...
                  }>,
                  ...
                },
                ...
              },
              +parentEquipment: {
                +id: string,
                +name: string,
                +futureState: ?FutureState,
                +equipmentType: {
                  +id: string,
                  +name: string,
                  +portDefinitions: $ReadOnlyArray<?{
                    +id: string,
                    +name: string,
                    +visibleLabel: ?string,
                    +bandwidth: ?string,
                    +portType: ?{
                      +id: string,
                      +name: string,
                      ...
                    },
                    ...
                  }>,
                  ...
                },
                +$fragmentRefs: EquipmentBreadcrumbs_equipment$ref,
                ...
              },
              +serviceEndpoints: $ReadOnlyArray<{
                +definition: {
                  +role: ?string,
                  ...
                },
                +service: {
                  +name: string,
                  ...
                },
                ...
              }>,
              ...
            }>,
            +workOrder: ?{
              +id: string,
              +status: WorkOrderStatus,
              ...
            },
            +properties: $ReadOnlyArray<?{
              +id: string,
              +propertyType: {
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
                ...
              },
              +stringValue: ?string,
              +intValue: ?number,
              +floatValue: ?number,
              +booleanValue: ?boolean,
              +latitudeValue: ?number,
              +longitudeValue: ?number,
              +rangeFromValue: ?number,
              +rangeToValue: ?number,
              +nodeValue: ?{
                +id: string,
                +name: string,
                ...
              },
              ...
            }>,
            +services: $ReadOnlyArray<?{
              +id: string,
              +name: string,
              ...
            }>,
            ...
          },
          +properties: $ReadOnlyArray<{
            +id: string,
            +propertyType: {
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
              ...
            },
            +stringValue: ?string,
            +intValue: ?number,
            +floatValue: ?number,
            +booleanValue: ?boolean,
            +latitudeValue: ?number,
            +longitudeValue: ?number,
            +rangeFromValue: ?number,
            +rangeToValue: ?number,
            +nodeValue: ?{
              +id: string,
              +name: string,
              ...
            },
            ...
          }>,
          +serviceEndpoints: $ReadOnlyArray<{
            +definition: {
              +role: ?string,
              ...
            },
            +service: {
              +name: string,
              ...
            },
            ...
          }>,
          ...
        }>,
        +equipmentType: {
          +portDefinitions: $ReadOnlyArray<?{
            +id: string,
            +name: string,
            +visibleLabel: ?string,
            +bandwidth: ?string,
            ...
          }>,
          ...
        },
        +positions: $ReadOnlyArray<?{
          +attachedEquipment: ?{
            +id: string,
            +name: string,
            +ports: $ReadOnlyArray<?{
              +id: string,
              +definition: {
                +id: string,
                +name: string,
                +index: ?number,
                +visibleLabel: ?string,
                +portType: ?{
                  +id: string,
                  +name: string,
                  +propertyTypes: $ReadOnlyArray<?{
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
                    ...
                  }>,
                  +linkPropertyTypes: $ReadOnlyArray<?{
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
                    ...
                  }>,
                  ...
                },
                ...
              },
              +parentEquipment: {
                +id: string,
                +name: string,
                +equipmentType: {
                  +id: string,
                  +name: string,
                  +portDefinitions: $ReadOnlyArray<?{
                    +id: string,
                    +name: string,
                    +visibleLabel: ?string,
                    +portType: ?{
                      +id: string,
                      +name: string,
                      ...
                    },
                    +bandwidth: ?string,
                    ...
                  }>,
                  ...
                },
                ...
              },
              +link: ?{
                +id: string,
                +futureState: ?FutureState,
                +ports: $ReadOnlyArray<?{
                  +id: string,
                  +definition: {
                    +id: string,
                    +name: string,
                    +visibleLabel: ?string,
                    +portType: ?{
                      +linkPropertyTypes: $ReadOnlyArray<?{
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
                        ...
                      }>,
                      ...
                    },
                    ...
                  },
                  +parentEquipment: {
                    +id: string,
                    +name: string,
                    +futureState: ?FutureState,
                    +equipmentType: {
                      +id: string,
                      +name: string,
                      +portDefinitions: $ReadOnlyArray<?{
                        +id: string,
                        +name: string,
                        +visibleLabel: ?string,
                        +bandwidth: ?string,
                        +portType: ?{
                          +id: string,
                          +name: string,
                          ...
                        },
                        ...
                      }>,
                      ...
                    },
                    +$fragmentRefs: EquipmentBreadcrumbs_equipment$ref,
                    ...
                  },
                  +serviceEndpoints: $ReadOnlyArray<{
                    +definition: {
                      +role: ?string,
                      ...
                    },
                    +service: {
                      +name: string,
                      ...
                    },
                    ...
                  }>,
                  ...
                }>,
                +workOrder: ?{
                  +id: string,
                  +status: WorkOrderStatus,
                  ...
                },
                +properties: $ReadOnlyArray<?{
                  +id: string,
                  +propertyType: {
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
                    ...
                  },
                  +stringValue: ?string,
                  +intValue: ?number,
                  +floatValue: ?number,
                  +booleanValue: ?boolean,
                  +latitudeValue: ?number,
                  +longitudeValue: ?number,
                  +rangeFromValue: ?number,
                  +rangeToValue: ?number,
                  +nodeValue: ?{
                    +id: string,
                    +name: string,
                    ...
                  },
                  ...
                }>,
                +services: $ReadOnlyArray<?{
                  +id: string,
                  +name: string,
                  ...
                }>,
                ...
              },
              +properties: $ReadOnlyArray<{
                +id: string,
                +propertyType: {
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
                  ...
                },
                +stringValue: ?string,
                +intValue: ?number,
                +floatValue: ?number,
                +booleanValue: ?boolean,
                +latitudeValue: ?number,
                +longitudeValue: ?number,
                +rangeFromValue: ?number,
                +rangeToValue: ?number,
                +nodeValue: ?{
                  +id: string,
                  +name: string,
                  ...
                },
                ...
              }>,
              +serviceEndpoints: $ReadOnlyArray<{
                +definition: {
                  +role: ?string,
                  ...
                },
                +service: {
                  +name: string,
                  ...
                },
                ...
              }>,
              ...
            }>,
            +equipmentType: {
              +portDefinitions: $ReadOnlyArray<?{
                +id: string,
                +name: string,
                +visibleLabel: ?string,
                +bandwidth: ?string,
                ...
              }>,
              ...
            },
            +positions: $ReadOnlyArray<?{
              +attachedEquipment: ?{
                +id: string,
                +name: string,
                +ports: $ReadOnlyArray<?{
                  +id: string,
                  +definition: {
                    +id: string,
                    +name: string,
                    +index: ?number,
                    +visibleLabel: ?string,
                    +portType: ?{
                      +id: string,
                      +name: string,
                      +propertyTypes: $ReadOnlyArray<?{
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
                        ...
                      }>,
                      +linkPropertyTypes: $ReadOnlyArray<?{
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
                        ...
                      }>,
                      ...
                    },
                    ...
                  },
                  +parentEquipment: {
                    +id: string,
                    +name: string,
                    +equipmentType: {
                      +id: string,
                      +name: string,
                      +portDefinitions: $ReadOnlyArray<?{
                        +id: string,
                        +name: string,
                        +visibleLabel: ?string,
                        +portType: ?{
                          +id: string,
                          +name: string,
                          ...
                        },
                        +bandwidth: ?string,
                        ...
                      }>,
                      ...
                    },
                    ...
                  },
                  +link: ?{
                    +id: string,
                    +futureState: ?FutureState,
                    +ports: $ReadOnlyArray<?{
                      +id: string,
                      +definition: {
                        +id: string,
                        +name: string,
                        +visibleLabel: ?string,
                        +portType: ?{
                          +linkPropertyTypes: $ReadOnlyArray<?{
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
                            ...
                          }>,
                          ...
                        },
                        ...
                      },
                      +parentEquipment: {
                        +id: string,
                        +name: string,
                        +futureState: ?FutureState,
                        +equipmentType: {
                          +id: string,
                          +name: string,
                          +portDefinitions: $ReadOnlyArray<?{
                            +id: string,
                            +name: string,
                            +visibleLabel: ?string,
                            +bandwidth: ?string,
                            +portType: ?{
                              +id: string,
                              +name: string,
                              ...
                            },
                            ...
                          }>,
                          ...
                        },
                        +$fragmentRefs: EquipmentBreadcrumbs_equipment$ref,
                        ...
                      },
                      +serviceEndpoints: $ReadOnlyArray<{
                        +definition: {
                          +role: ?string,
                          ...
                        },
                        +service: {
                          +name: string,
                          ...
                        },
                        ...
                      }>,
                      ...
                    }>,
                    +workOrder: ?{
                      +id: string,
                      +status: WorkOrderStatus,
                      ...
                    },
                    +properties: $ReadOnlyArray<?{
                      +id: string,
                      +propertyType: {
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
                        ...
                      },
                      +stringValue: ?string,
                      +intValue: ?number,
                      +floatValue: ?number,
                      +booleanValue: ?boolean,
                      +latitudeValue: ?number,
                      +longitudeValue: ?number,
                      +rangeFromValue: ?number,
                      +rangeToValue: ?number,
                      +nodeValue: ?{
                        +id: string,
                        +name: string,
                        ...
                      },
                      ...
                    }>,
                    +services: $ReadOnlyArray<?{
                      +id: string,
                      +name: string,
                      ...
                    }>,
                    ...
                  },
                  +properties: $ReadOnlyArray<{
                    +id: string,
                    +propertyType: {
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
                      ...
                    },
                    +stringValue: ?string,
                    +intValue: ?number,
                    +floatValue: ?number,
                    +booleanValue: ?boolean,
                    +latitudeValue: ?number,
                    +longitudeValue: ?number,
                    +rangeFromValue: ?number,
                    +rangeToValue: ?number,
                    +nodeValue: ?{
                      +id: string,
                      +name: string,
                      ...
                    },
                    ...
                  }>,
                  +serviceEndpoints: $ReadOnlyArray<{
                    +definition: {
                      +role: ?string,
                      ...
                    },
                    +service: {
                      +name: string,
                      ...
                    },
                    ...
                  }>,
                  ...
                }>,
                +equipmentType: {
                  +portDefinitions: $ReadOnlyArray<?{
                    +id: string,
                    +name: string,
                    +visibleLabel: ?string,
                    +bandwidth: ?string,
                    ...
                  }>,
                  ...
                },
                ...
              },
              ...
            }>,
            ...
          },
          ...
        }>,
        ...
      },
      ...
    }>,
    ...
  },
  ...
};
export type EquipmentPortsTable_position$data = EquipmentPortsTable_position;
export type EquipmentPortsTable_position$key = {
  +$data?: EquipmentPortsTable_position$data,
  +$fragmentRefs: EquipmentPortsTable_position$ref,
  ...
};
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
v2 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "index",
  "args": null,
  "storageKey": null
},
v3 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "visibleLabel",
  "args": null,
  "storageKey": null
},
v4 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "stringValue",
  "args": null,
  "storageKey": null
},
v5 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "intValue",
  "args": null,
  "storageKey": null
},
v6 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "booleanValue",
  "args": null,
  "storageKey": null
},
v7 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "floatValue",
  "args": null,
  "storageKey": null
},
v8 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "latitudeValue",
  "args": null,
  "storageKey": null
},
v9 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "longitudeValue",
  "args": null,
  "storageKey": null
},
v10 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "rangeFromValue",
  "args": null,
  "storageKey": null
},
v11 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "rangeToValue",
  "args": null,
  "storageKey": null
},
v12 = [
  (v0/*: any*/),
  (v1/*: any*/),
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
  (v2/*: any*/),
  (v4/*: any*/),
  (v5/*: any*/),
  (v6/*: any*/),
  (v7/*: any*/),
  (v8/*: any*/),
  (v9/*: any*/),
  (v10/*: any*/),
  (v11/*: any*/),
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
v13 = {
  "kind": "LinkedField",
  "alias": null,
  "name": "linkPropertyTypes",
  "storageKey": null,
  "args": null,
  "concreteType": "PropertyType",
  "plural": true,
  "selections": (v12/*: any*/)
},
v14 = [
  (v0/*: any*/),
  (v1/*: any*/)
],
v15 = {
  "kind": "LinkedField",
  "alias": null,
  "name": "portType",
  "storageKey": null,
  "args": null,
  "concreteType": "EquipmentPortType",
  "plural": false,
  "selections": (v14/*: any*/)
},
v16 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "bandwidth",
  "args": null,
  "storageKey": null
},
v17 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "futureState",
  "args": null,
  "storageKey": null
},
v18 = {
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
        }
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
        (v1/*: any*/)
      ]
    }
  ]
},
v19 = {
  "kind": "LinkedField",
  "alias": null,
  "name": "properties",
  "storageKey": null,
  "args": null,
  "concreteType": "Property",
  "plural": true,
  "selections": [
    (v0/*: any*/),
    {
      "kind": "LinkedField",
      "alias": null,
      "name": "propertyType",
      "storageKey": null,
      "args": null,
      "concreteType": "PropertyType",
      "plural": false,
      "selections": (v12/*: any*/)
    },
    (v4/*: any*/),
    (v5/*: any*/),
    (v7/*: any*/),
    (v6/*: any*/),
    (v8/*: any*/),
    (v9/*: any*/),
    (v10/*: any*/),
    (v11/*: any*/),
    {
      "kind": "LinkedField",
      "alias": null,
      "name": "nodeValue",
      "storageKey": null,
      "args": null,
      "concreteType": null,
      "plural": false,
      "selections": (v14/*: any*/)
    }
  ]
},
v20 = {
  "kind": "LinkedField",
  "alias": null,
  "name": "ports",
  "storageKey": null,
  "args": null,
  "concreteType": "EquipmentPort",
  "plural": true,
  "selections": [
    (v0/*: any*/),
    {
      "kind": "LinkedField",
      "alias": null,
      "name": "definition",
      "storageKey": null,
      "args": null,
      "concreteType": "EquipmentPortDefinition",
      "plural": false,
      "selections": [
        (v0/*: any*/),
        (v1/*: any*/),
        (v2/*: any*/),
        (v3/*: any*/),
        {
          "kind": "LinkedField",
          "alias": null,
          "name": "portType",
          "storageKey": null,
          "args": null,
          "concreteType": "EquipmentPortType",
          "plural": false,
          "selections": [
            (v0/*: any*/),
            (v1/*: any*/),
            {
              "kind": "LinkedField",
              "alias": null,
              "name": "propertyTypes",
              "storageKey": null,
              "args": null,
              "concreteType": "PropertyType",
              "plural": true,
              "selections": (v12/*: any*/)
            },
            (v13/*: any*/)
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
        (v0/*: any*/),
        (v1/*: any*/),
        {
          "kind": "LinkedField",
          "alias": null,
          "name": "equipmentType",
          "storageKey": null,
          "args": null,
          "concreteType": "EquipmentType",
          "plural": false,
          "selections": [
            (v0/*: any*/),
            (v1/*: any*/),
            {
              "kind": "LinkedField",
              "alias": null,
              "name": "portDefinitions",
              "storageKey": null,
              "args": null,
              "concreteType": "EquipmentPortDefinition",
              "plural": true,
              "selections": [
                (v0/*: any*/),
                (v1/*: any*/),
                (v3/*: any*/),
                (v15/*: any*/),
                (v16/*: any*/)
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
        (v0/*: any*/),
        (v17/*: any*/),
        {
          "kind": "LinkedField",
          "alias": null,
          "name": "ports",
          "storageKey": null,
          "args": null,
          "concreteType": "EquipmentPort",
          "plural": true,
          "selections": [
            (v0/*: any*/),
            {
              "kind": "LinkedField",
              "alias": null,
              "name": "definition",
              "storageKey": null,
              "args": null,
              "concreteType": "EquipmentPortDefinition",
              "plural": false,
              "selections": [
                (v0/*: any*/),
                (v1/*: any*/),
                (v3/*: any*/),
                {
                  "kind": "LinkedField",
                  "alias": null,
                  "name": "portType",
                  "storageKey": null,
                  "args": null,
                  "concreteType": "EquipmentPortType",
                  "plural": false,
                  "selections": [
                    (v13/*: any*/)
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
                (v0/*: any*/),
                (v1/*: any*/),
                (v17/*: any*/),
                {
                  "kind": "LinkedField",
                  "alias": null,
                  "name": "equipmentType",
                  "storageKey": null,
                  "args": null,
                  "concreteType": "EquipmentType",
                  "plural": false,
                  "selections": [
                    (v0/*: any*/),
                    (v1/*: any*/),
                    {
                      "kind": "LinkedField",
                      "alias": null,
                      "name": "portDefinitions",
                      "storageKey": null,
                      "args": null,
                      "concreteType": "EquipmentPortDefinition",
                      "plural": true,
                      "selections": [
                        (v0/*: any*/),
                        (v1/*: any*/),
                        (v3/*: any*/),
                        (v16/*: any*/),
                        (v15/*: any*/)
                      ]
                    }
                  ]
                },
                {
                  "kind": "FragmentSpread",
                  "name": "EquipmentBreadcrumbs_equipment",
                  "args": null
                }
              ]
            },
            (v18/*: any*/)
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
            (v0/*: any*/),
            {
              "kind": "ScalarField",
              "alias": null,
              "name": "status",
              "args": null,
              "storageKey": null
            }
          ]
        },
        (v19/*: any*/),
        {
          "kind": "LinkedField",
          "alias": null,
          "name": "services",
          "storageKey": null,
          "args": null,
          "concreteType": "Service",
          "plural": true,
          "selections": (v14/*: any*/)
        }
      ]
    },
    (v19/*: any*/),
    (v18/*: any*/)
  ]
},
v21 = {
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
        (v0/*: any*/),
        (v1/*: any*/),
        (v3/*: any*/),
        (v16/*: any*/)
      ]
    }
  ]
};
return {
  "kind": "Fragment",
  "name": "EquipmentPortsTable_position",
  "type": "EquipmentPosition",
  "metadata": {
    "mask": false
  },
  "argumentDefinitions": [],
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
        (v0/*: any*/),
        (v1/*: any*/),
        (v20/*: any*/),
        (v21/*: any*/),
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
                (v0/*: any*/),
                (v1/*: any*/),
                (v20/*: any*/),
                (v21/*: any*/),
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
                        (v0/*: any*/),
                        (v1/*: any*/),
                        (v20/*: any*/),
                        (v21/*: any*/),
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
                                (v0/*: any*/),
                                (v1/*: any*/),
                                (v20/*: any*/),
                                (v21/*: any*/)
                              ]
                            }
                          ]
                        }
                      ]
                    }
                  ]
                }
              ]
            }
          ]
        }
      ]
    }
  ]
};
})();
// prettier-ignore
(node/*: any*/).hash = 'b832b8a9e1093a7fc65b0764dfd2a8a5';
module.exports = node;
