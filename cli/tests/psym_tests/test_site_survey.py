#!/usr/bin/env python3
# Copyright (c) 2004-present Facebook All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.


import os
from datetime import datetime
from unittest import skip

from psym.api.location import add_location
from psym.api.location_type import add_location_type
from psym.api.site_survey import (
    delete_site_survey,
    get_site_surveys,
    upload_site_survey,
)

from ..utils.base_test import BaseTest


@skip("site survey is deprecated")
class TestSiteSurvey(BaseTest):
    def setUp(self) -> None:
        super().setUp()
        add_location_type(self.client, "City Center", [])

    def test_site_survey_created(self) -> None:
        location = add_location(
            self.client, [("City Center", "Lima Downtown")], {}, 10, 20
        )
        self.assertEqual(0, len(get_site_surveys(self.client, location.id)))
        completion_date = datetime.strptime("25-7-2019", "%d-%m-%Y")
        upload_site_survey(
            self.client,
            location.id,
            "My site survey",
            completion_date,
            os.path.join(
                os.path.dirname(__file__), "resources/city_center_site_survey.xlsx"
            ),
            os.path.join(
                os.path.dirname(__file__), "resources/city_center_site_survey.json"
            ),
        )
        surveys = get_site_surveys(self.client, location.id)
        self.assertEqual(1, len(surveys))
        survey = surveys[0]
        self.assertEqual("My site survey", survey.name)
        self.assertEqual(completion_date, survey.completion_time)

        self.assertIsNotNone(survey.source_file_id)
        self.assertEqual(survey.source_file_name, "city_center_site_survey.xlsx")

        delete_site_survey(self.client, survey)
        surveys = get_site_surveys(self.client, location.id)
        self.assertEqual(0, len(surveys))
