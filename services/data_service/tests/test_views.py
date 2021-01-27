import json
import os
import unittest
from unittest.mock import patch

from project.app import MyMicroservice
from pyms.constants import CONFIGMAP_FILE_ENVIRONMENT


class ProjectTestCase(unittest.TestCase):
    BASE_DIR = os.path.dirname(os.path.abspath(__file__))

    def setUp(self):
        os.environ[CONFIGMAP_FILE_ENVIRONMENT] = os.path.join(self.BASE_DIR, "config-tests.yml")
        ms = MyMicroservice(path=os.path.join(os.path.dirname(os.path.dirname(__file__)), "test_views.py"))
        self.app = ms.create_app()
        self.base_url = self.app.config["APPLICATION_ROOT"]
        self.client = self.app.test_client()

    def test_home(self):
        response = self.client.get('/')
        self.assertEqual(404, response.status_code)

    def test_healthcheck(self):
        response = self.client.get('/healthcheck')
        self.assertEqual(200, response.status_code)

    @patch('project.views.country.get_countries')
    def test_list_view(self, mock_get_countries):
        mock_get_countries.return_value = [{
            'iso_code': 'usa',
            'vaccinations': '1',
            'cases': '2'
        }]
        response = self.client.get('/countries/usa?from_date=2021-01-25&to_date=2021-01-26')
        self.assertEqual(200, response.status_code)

