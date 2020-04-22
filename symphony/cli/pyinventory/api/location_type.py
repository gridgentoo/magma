#!/usr/bin/env python3

from typing import List, Optional, Tuple

from gql.gql.client import OperationException
from gql.gql.reporter import FailedOperationException

from .._utils import format_properties
from ..client import SymphonyClient
from ..consts import Entity, Location, LocationType, PropertyValue
from ..exceptions import EntityNotFoundError
from ..graphql.add_location_type_input import AddLocationTypeInput
from ..graphql.add_location_type_mutation import AddLocationTypeMutation
from ..graphql.location_type_locations_query import LocationTypeLocationsQuery
from ..graphql.location_types_query import LocationTypesQuery
from ..graphql.remove_location_type_mutation import RemoveLocationTypeMutation
from .location import delete_location


ADD_LOCATION_TYPE_MUTATION_NAME = "addLocationType"


def _populate_location_types(client: SymphonyClient) -> None:
    location_types = LocationTypesQuery.execute(client).locationTypes
    if not location_types:
        return
    edges = location_types.edges
    for edge in edges:
        node = edge.node
        if node:
            client.locationTypes[node.name] = LocationType(
                name=node.name, id=node.id, property_types=node.propertyTypes
            )


def add_location_type(
    client: SymphonyClient,
    name: str,
    properties: List[Tuple[str, str, Optional[PropertyValue], Optional[bool]]],
    map_zoom_level: int = 8,
) -> LocationType:
    """This function creates new location type.

        Args:
            name (str): location type name
            properties (List[Tuple[str, str, Optional[PropertyValue], Optional[bool]]]):
            - str - type name
            - str - enum["string", "int", "bool", "float", "date", "enum", "range",
            "email", "gps_location", "equipment", "location", "service", "datetime_local"]
            - PropertyValue - default property value
            - bool - fixed value flag

            map_zoom_level (int): map zoom level

        Returns:
            `pyinventory.consts.LocationType` object

        Raises:
            FailedOperationException: internal inventory error

        Example:
            ```
            location_type = client.add_location_type(
                name="city",
                properties=[("Contact", "email", None, True)],
                map_zoom_level=5,
            )
            ```
    """
    new_property_types = format_properties(properties)
    add_location_type_variables = {
        "name": name,
        "mapZoomLevel": map_zoom_level,
        "properties": new_property_types,
        "surveyTemplateCategories": [],
    }
    try:
        result = AddLocationTypeMutation.execute(
            client,
            AddLocationTypeInput(
                name=name,
                mapZoomLevel=map_zoom_level,
                properties=new_property_types,
                surveyTemplateCategories=[],
            ),
        ).__dict__[ADD_LOCATION_TYPE_MUTATION_NAME]
        client.reporter.log_successful_operation(
            ADD_LOCATION_TYPE_MUTATION_NAME, add_location_type_variables
        )
    except OperationException as e:
        raise FailedOperationException(
            client.reporter,
            e.err_msg,
            e.err_id,
            ADD_LOCATION_TYPE_MUTATION_NAME,
            add_location_type_variables,
        )

    location_type = LocationType(
        name=result.name, id=result.id, property_types=result.propertyTypes
    )
    client.locationTypes[result.name] = location_type
    return location_type


def delete_locations_by_location_type(
    client: SymphonyClient, location_type: LocationType
) -> None:
    """Delete locatons by location type.

        Args:
            location_type ( `pyinventory.consts.LocationType` ): location type object

        Raises:
            `pyinventory.exceptions.EntityNotFoundError`: if location_type does not exist

        Example:
            ```
            client.delete_locations_by_location_type(location_type=location_type)
            ```
    """
    location_type_with_locations = LocationTypeLocationsQuery.execute(
        client, id=location_type.id
    ).locationType
    if location_type_with_locations is None:
        raise EntityNotFoundError(
            entity=Entity.LocationType, entity_id=location_type.id
        )
    locations = location_type_with_locations.locations
    if locations is None:
        return
    for location in locations.edges:
        node = location.node
        if node:
            delete_location(
                client,
                Location(
                    id=node.id,
                    name=node.name,
                    latitude=node.latitude,
                    longitude=node.longitude,
                    externalId=node.externalId,
                    locationTypeName=node.locationType.name,
                ),
            )


def delete_location_type_with_locations(
    client: SymphonyClient, location_type: LocationType
) -> None:
    """Delete locaton type with existing locations.

        Args:
            location_type (`pyinventory.consts.LocationType`): location type object

        Raises:
            `pyinventory.exceptions.EntityNotFoundError`: if location_type does not exist

        Example:
            ```
            client.delete_location_type_with_locations(location_type=location_type)
            ```
    """
    delete_locations_by_location_type(client, location_type)
    RemoveLocationTypeMutation.execute(client, id=location_type.id)
