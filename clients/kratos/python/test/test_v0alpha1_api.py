"""
    Ory Kratos API

    Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests.   # noqa: E501

    The version of the OpenAPI document: v0.7.6-alpha.5
    Contact: hi@ory.sh
    Generated by: https://openapi-generator.tech
"""


import unittest

import ory_kratos_client
from ory_kratos_client.api.v0alpha1_api import V0alpha1Api  # noqa: E501


class TestV0alpha1Api(unittest.TestCase):
    """V0alpha1Api unit test stubs"""

    def setUp(self):
        self.api = V0alpha1Api()  # noqa: E501

    def tearDown(self):
        pass

    def test_admin_create_identity(self):
        """Test case for admin_create_identity

        Create an Identity  # noqa: E501
        """
        pass

    def test_admin_create_self_service_recovery_link(self):
        """Test case for admin_create_self_service_recovery_link

        Create a Recovery Link  # noqa: E501
        """
        pass

    def test_admin_delete_identity(self):
        """Test case for admin_delete_identity

        Delete an Identity  # noqa: E501
        """
        pass

    def test_admin_get_identity(self):
        """Test case for admin_get_identity

        Get an Identity  # noqa: E501
        """
        pass

    def test_admin_list_identities(self):
        """Test case for admin_list_identities

        List Identities  # noqa: E501
        """
        pass

    def test_admin_update_identity(self):
        """Test case for admin_update_identity

        Update an Identity  # noqa: E501
        """
        pass

    def test_create_self_service_logout_flow_url_for_browsers(self):
        """Test case for create_self_service_logout_flow_url_for_browsers

        Create a Logout URL for Browsers  # noqa: E501
        """
        pass

    def test_get_json_schema(self):
        """Test case for get_json_schema

        """
        pass

    def test_get_self_service_error(self):
        """Test case for get_self_service_error

        Get Self-Service Errors  # noqa: E501
        """
        pass

    def test_get_self_service_login_flow(self):
        """Test case for get_self_service_login_flow

        Get Login Flow  # noqa: E501
        """
        pass

    def test_get_self_service_recovery_flow(self):
        """Test case for get_self_service_recovery_flow

        Get Recovery Flow  # noqa: E501
        """
        pass

    def test_get_self_service_registration_flow(self):
        """Test case for get_self_service_registration_flow

        Get Registration Flow  # noqa: E501
        """
        pass

    def test_get_self_service_settings_flow(self):
        """Test case for get_self_service_settings_flow

        Get Settings Flow  # noqa: E501
        """
        pass

    def test_get_self_service_verification_flow(self):
        """Test case for get_self_service_verification_flow

        Get Verification Flow  # noqa: E501
        """
        pass

    def test_initialize_self_service_login_flow_for_browsers(self):
        """Test case for initialize_self_service_login_flow_for_browsers

        Initialize Login Flow for Browsers  # noqa: E501
        """
        pass

    def test_initialize_self_service_login_flow_without_browser(self):
        """Test case for initialize_self_service_login_flow_without_browser

        Initialize Login Flow for APIs, Services, Apps, ...  # noqa: E501
        """
        pass

    def test_initialize_self_service_recovery_flow_for_browsers(self):
        """Test case for initialize_self_service_recovery_flow_for_browsers

        Initialize Recovery Flow for Browsers  # noqa: E501
        """
        pass

    def test_initialize_self_service_recovery_flow_without_browser(self):
        """Test case for initialize_self_service_recovery_flow_without_browser

        Initialize Recovery Flow for APIs, Services, Apps, ...  # noqa: E501
        """
        pass

    def test_initialize_self_service_registration_flow_for_browsers(self):
        """Test case for initialize_self_service_registration_flow_for_browsers

        Initialize Registration Flow for Browsers  # noqa: E501
        """
        pass

    def test_initialize_self_service_registration_flow_without_browser(self):
        """Test case for initialize_self_service_registration_flow_without_browser

        Initialize Registration Flow for APIs, Services, Apps, ...  # noqa: E501
        """
        pass

    def test_initialize_self_service_settings_flow_for_browsers(self):
        """Test case for initialize_self_service_settings_flow_for_browsers

        Initialize Settings Flow for Browsers  # noqa: E501
        """
        pass

    def test_initialize_self_service_settings_flow_without_browser(self):
        """Test case for initialize_self_service_settings_flow_without_browser

        Initialize Settings Flow for APIs, Services, Apps, ...  # noqa: E501
        """
        pass

    def test_initialize_self_service_verification_flow_for_browsers(self):
        """Test case for initialize_self_service_verification_flow_for_browsers

        Initialize Verification Flow for Browser Clients  # noqa: E501
        """
        pass

    def test_initialize_self_service_verification_flow_without_browser(self):
        """Test case for initialize_self_service_verification_flow_without_browser

        Initialize Verification Flow for APIs, Services, Apps, ...  # noqa: E501
        """
        pass

    def test_submit_self_service_login_flow(self):
        """Test case for submit_self_service_login_flow

        Submit a Login Flow  # noqa: E501
        """
        pass

    def test_submit_self_service_logout_flow(self):
        """Test case for submit_self_service_logout_flow

        Complete Self-Service Logout  # noqa: E501
        """
        pass

    def test_submit_self_service_logout_flow_without_browser(self):
        """Test case for submit_self_service_logout_flow_without_browser

        Perform Logout for APIs, Services, Apps, ...  # noqa: E501
        """
        pass

    def test_submit_self_service_recovery_flow(self):
        """Test case for submit_self_service_recovery_flow

        Complete Recovery Flow  # noqa: E501
        """
        pass

    def test_submit_self_service_registration_flow(self):
        """Test case for submit_self_service_registration_flow

        Submit a Registration Flow  # noqa: E501
        """
        pass

    def test_submit_self_service_settings_flow(self):
        """Test case for submit_self_service_settings_flow

        Complete Settings Flow  # noqa: E501
        """
        pass

    def test_submit_self_service_verification_flow(self):
        """Test case for submit_self_service_verification_flow

        Complete Verification Flow  # noqa: E501
        """
        pass

    def test_to_session(self):
        """Test case for to_session

        Check Who the Current HTTP Session Belongs To  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()
