#!/usr/bin/env python3
# Copyright (c) 2004-present Facebook All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

from typing import Dict, List, Optional

from gql.gql.client import OperationException
from gql.gql.reporter import FailedOperationException

from .._utils import format_property_definitions, get_graphql_property_type_inputs
from ..client import SymphonyClient
from ..consts import Entity, EquipmentPortType, PropertyDefinition, PropertyValue
from ..exceptions import EntityNotFoundError
from ..graphql.add_equipment_port_type_mutation import (
    AddEquipmentPortTypeInput,
    AddEquipmentPortTypeMutation,
)
from ..graphql.edit_equipment_port_type_mutation import (
    EditEquipmentPortTypeInput,
    EditEquipmentPortTypeMutation,
)
from ..graphql.equipment_port_type_query import EquipmentPortTypeQuery
from ..graphql.remove_equipment_port_type_mutation import (
    RemoveEquipmentPortTypeMutation,
)


ADD_EQUIPMENT_PORT_TYPE_MUTATION_NAME = "addEquipmentPortType"
EDIT_EQUIPMENT_PORT_TYPE_MUTATION_NAME = "editEquipmentPortType"


def add_equipment_port_type(
    client: SymphonyClient,
    name: str,
    properties: List[PropertyDefinition],
    link_properties: List[PropertyDefinition],
) -> EquipmentPortType:
    """This function creates an equipment port type.

        Args:
            name (str): equipment port type name
            properties: (List[ `pyinventory.consts.PropertyDefinition` ]): list of property definitions
            link_properties: (List[ `pyinventory.consts.PropertyDefinition` ]): list of property definitions

        Returns:
            `pyinventory.consts.EquipmentPortType` object

        Raises:
            FailedOperationException: internal inventory error

        Example:
            ```
            from pyinventory.consts import PropertyDefinition
            from pyinventory.graphql.property_kind_enum import PropertyKind
            port_type1 = client.add_equipment_port_type(
                name="port type 1",
                properties=[PropertyDefinition(
                    property_name="port property",
                    property_kind=PropertyKind.string,
                    default_value=None,
                    is_fixed=True)],
                link_properties=[PropertyDefinition(
                    property_name="link port property",
                    property_kind=PropertyKind.string,
                    default_value=None,
                    is_fixed=True)],
            )
            ```
    """

    formated_property_types = format_property_definitions(properties)
    formated_link_property_types = format_property_definitions(link_properties)
    add_equipment_port_type_input = {
        "name": name,
        "properties": formated_property_types,
        "linkProperties": formated_link_property_types,
    }

    try:
        result = AddEquipmentPortTypeMutation.execute(
            client,
            AddEquipmentPortTypeInput(
                name=name,
                properties=formated_property_types,
                linkProperties=formated_link_property_types,
            ),
        ).__dict__[ADD_EQUIPMENT_PORT_TYPE_MUTATION_NAME]
        client.reporter.log_successful_operation(
            ADD_EQUIPMENT_PORT_TYPE_MUTATION_NAME, add_equipment_port_type_input
        )
    except OperationException as e:
        raise FailedOperationException(
            client.reporter,
            e.err_msg,
            e.err_id,
            ADD_EQUIPMENT_PORT_TYPE_MUTATION_NAME,
            add_equipment_port_type_input,
        )

    added = EquipmentPortType(
        id=result.id,
        name=result.name,
        property_types=result.propertyTypes,
        link_property_types=result.linkPropertyTypes,
    )
    client.portTypes[added.name] = added
    return added


def get_equipment_port_type(
    client: SymphonyClient, equipment_port_type_id: str
) -> EquipmentPortType:
    """This function returns an equipment port type.
        It can get only the requested equipment port type ID

        Args:
            equipment_port_type_id (str): equipment port type ID

        Returns:
            `pyinventory.consts.EquipmentPortType` object

        Raises:
            `pyinventory.exceptions.EntityNotFoundError`: equipment port type does not found

        Example:
            ```
            port_type = client.get_equipment_port_type(equipment_port_type_id=port_type1.id)
            ```
    """
    result = EquipmentPortTypeQuery.execute(client, id=equipment_port_type_id).port_type
    if not result:
        raise EntityNotFoundError(
            entity=Entity.EquipmentPortType, entity_id=equipment_port_type_id
        )

    return EquipmentPortType(
        id=result.id,
        name=result.name,
        property_types=result.propertyTypes,
        link_property_types=result.linkPropertyTypes,
    )


def edit_equipment_port_type(
    client: SymphonyClient,
    port_type: EquipmentPortType,
    new_name: Optional[str] = None,
    new_properties: Optional[Dict[str, PropertyValue]] = None,
    new_link_properties: Optional[Dict[str, PropertyValue]] = None,
) -> EquipmentPortType:
    """This function edits an existing equipment port type.

        Args:
            port_type ( `pyinventory.consts.EquipmentPortType` ): existing eqipment port type object
            new_name (str): new name
            new_properties: (Dict[str, PropertyValue]): dictionary
            - str - property type name
            - PropertyValue - new value of the same type for this property

            new_link_properties: (Dict[str, PropertyValue]): dictionary
            - str - link property type name
            - PropertyValue - new value of the same type for this link property

        Returns:
            `pyinventory.consts.EquipmentPortType` object

        Raises:
            FailedOperationException: internal inventory error

        Example:
            ```
            port_type1 = client.edit_equipment_port_type(
                port_type=equipment_port_type,
                new_name="new port type name",
                new_properties={"existing property name": "new value"},
                new_link_properties={"existing link property name": "new value"},
            )
            ```
    """
    new_name = port_type.name if new_name is None else new_name

    new_property_type_inputs = []
    if new_properties:
        property_types = client.portTypes[port_type.name].property_types
        new_property_type_inputs = get_graphql_property_type_inputs(
            property_types, new_properties
        )

    new_link_property_type_inputs = []
    if new_link_properties:
        link_property_types = client.portTypes[port_type.name].link_property_types
        new_link_property_type_inputs = get_graphql_property_type_inputs(
            link_property_types, new_link_properties
        )

    edit_equipment_port_type_input = {
        "name": new_name,
        "properties": new_property_type_inputs,
        "linkProperties": new_link_properties,
    }

    try:
        result = EditEquipmentPortTypeMutation.execute(
            client,
            EditEquipmentPortTypeInput(
                id=port_type.id,
                name=new_name,
                properties=new_property_type_inputs,
                linkProperties=new_link_property_type_inputs,
            ),
        ).__dict__[EDIT_EQUIPMENT_PORT_TYPE_MUTATION_NAME]
        client.reporter.log_successful_operation(
            EDIT_EQUIPMENT_PORT_TYPE_MUTATION_NAME, edit_equipment_port_type_input
        )
    except OperationException as e:
        raise FailedOperationException(
            client.reporter,
            e.err_msg,
            e.err_id,
            EDIT_EQUIPMENT_PORT_TYPE_MUTATION_NAME,
            edit_equipment_port_type_input,
        )
    return EquipmentPortType(
        id=result.id,
        name=result.name,
        property_types=result.propertyTypes,
        link_property_types=result.linkPropertyTypes,
    )


def delete_equipment_port_type(
    client: SymphonyClient, equipment_port_type_id: str
) -> None:
    """This function deletes an equipment port type.
        It can get only the requested equipment port type ID

        Args:
            equipment_port_type_id (str): equipment port type ID

        Example:
            ```
            client.delete_equipment_port_type(equipment_port_type_id=port_type1.id)
            ```
    """
    RemoveEquipmentPortTypeMutation.execute(client, id=equipment_port_type_id)
