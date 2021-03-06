#!/usr/bin/env python3
# Copyright (c) 2004-present Facebook All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

from typing import Dict, List, Tuple

from psym.client import SymphonyClient
from psym.common.data_class import Equipment, Location
from psym.common.data_enum import Entity
from psym.common.data_format import format_to_equipment

from ..exceptions import EntityNotFoundError
from ..graphql.query.equipment_positions import EquipmentPositionsQuery
from ..graphql.query.location_equipments import LocationEquipmentsQuery
from .equipment import copy_equipment, copy_equipment_in_position
from .link import add_link, get_all_links_and_port_names_of_equipment


def _get_one_level_attachments_of_equipment(
    client: SymphonyClient, equipment: Equipment
) -> List[Tuple[str, Equipment]]:
    equipment_with_positions = EquipmentPositionsQuery.execute(client, id=equipment.id)
    if not equipment_with_positions:
        raise EntityNotFoundError(entity=Entity.Equipment, entity_id=equipment.id)
    attachments = []
    for position in equipment_with_positions.positions:
        attached_equipment = position.attachedEquipment
        if attached_equipment is not None:
            attachments.append(
                (
                    position.definition.id,
                    format_to_equipment(equipment_fragment=attached_equipment),
                )
            )

    return attachments


def copy_equipment_with_all_attachments(
    client: SymphonyClient, equipment: Equipment, dest_location: Location
) -> Dict[Equipment, Equipment]:
    """Copy the equipment to the new location with all its attachments

    :param equipment: Equipment object to be copied, could be retrieved from

        * :meth:`~psym.api.equipment.get_equipment`
        * :meth:`~psym.api.equipment.get_equipment_in_position`
        * :meth:`~psym.api.equipment.add_equipment`
        * :meth:`~psym.api.equipment.add_equipment_to_position`

    :type equipment: :class:`~psym.common.data_class.Equipment`
    :param dest_location: Location to copy equipment to, could be retrieved from

        * :meth:`~psym.api.location.get_location`
        * :meth:`~psym.api.location.add_location`

    :type dest_location: :class:`~psym.common.data_class.Location`

    :raises:
        FailedOperationException: Internal symphony error

    :return: Dictionary of source equipment to new equipment,
        includes the equipment given as parameter and also all the equipments attached to it

        * source equipment - :class:`~psym.common.data_class.Equipment`
        * new equipment - :class:`~psym.common.data_class.Equipment`

    :rtype: Iterator[ Tuple[ :class:`~psym.common.data_class.Equipment`, str ] ]
    """

    result = {}

    new_equipment = copy_equipment(client, equipment, dest_location)
    equipments = [(equipment, new_equipment)]

    while len(equipments) != 0:
        old_equipment, new_equipment = equipments.pop()
        result[old_equipment] = new_equipment
        attachments = _get_one_level_attachments_of_equipment(client, old_equipment)
        for position_name, child_equipment in attachments:
            new_child_equipment = copy_equipment_in_position(
                client, child_equipment, new_equipment, position_name
            )
            equipments.append((child_equipment, new_child_equipment))
    return result


def apply_location_template_to_location(
    client: SymphonyClient, template_location: Location, location: Location
) -> None:

    location_with_equipments = LocationEquipmentsQuery.execute(
        client, id=template_location.id
    )
    if not location_with_equipments:
        raise EntityNotFoundError(
            entity=Entity.Location, entity_id=template_location.id
        )
    equipments = [
        format_to_equipment(equipment_fragment=equipment)
        for equipment in location_with_equipments.equipments
    ]
    equipments_to_new_equipments = {}
    for equipment in equipments:
        # return back all and gather link ids
        equipments_to_new_equipments.update(
            copy_equipment_with_all_attachments(client, equipment, location)
        )

    link_to_equipment_and_port = {}
    connected_links = []

    for equipment in equipments_to_new_equipments.keys():
        links_and_ports = get_all_links_and_port_names_of_equipment(client, equipment)
        for link, port_name in links_and_ports:
            if link not in link_to_equipment_and_port:
                link_to_equipment_and_port[link] = (port_name, equipment)
            else:
                other_port_name, other_equipment = link_to_equipment_and_port.pop(link)
                connected_links.append(
                    (equipment, port_name, other_equipment, other_port_name)
                )

    assert (
        len(link_to_equipment_and_port) == 0
    ), "Some equipments in location are connected to equipments outside the location"

    for equipment, port_name, other_equipment, other_port_name in connected_links:
        new_equipment = equipments_to_new_equipments[equipment]
        new_other_equipment = equipments_to_new_equipments[other_equipment]
        add_link(client, new_equipment, port_name, new_other_equipment, other_port_name)
