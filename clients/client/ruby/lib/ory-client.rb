=begin
#Ory APIs

#Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

The version of the OpenAPI document: v0.2.0-alpha.37
Contact: support@ory.sh
Generated by: https://openapi-generator.tech
OpenAPI Generator version: 6.0.1

=end

# Common files
require 'ory-client/api_client'
require 'ory-client/api_error'
require 'ory-client/version'
require 'ory-client/configuration'

# Models
require 'ory-client/models/accept_consent_request'
require 'ory-client/models/accept_login_request'
require 'ory-client/models/active_project'
require 'ory-client/models/admin_create_identity_body'
require 'ory-client/models/admin_create_identity_import_credentials_oidc'
require 'ory-client/models/admin_create_identity_import_credentials_oidc_config'
require 'ory-client/models/admin_create_identity_import_credentials_oidc_provider'
require 'ory-client/models/admin_create_identity_import_credentials_password'
require 'ory-client/models/admin_create_identity_import_credentials_password_config'
require 'ory-client/models/admin_create_self_service_recovery_link_body'
require 'ory-client/models/admin_identity_import_credentials'
require 'ory-client/models/admin_update_identity_body'
require 'ory-client/models/authenticator_assurance_level'
require 'ory-client/models/cloud_account'
require 'ory-client/models/cname_settings'
require 'ory-client/models/completed_request'
require 'ory-client/models/consent_request'
require 'ory-client/models/consent_request_session'
require 'ory-client/models/create_custom_hostname_body'
require 'ory-client/models/create_project_api_key_request'
require 'ory-client/models/create_project_body'
require 'ory-client/models/create_subscription_payload'
require 'ory-client/models/error_authenticator_assurance_level_not_satisfied'
require 'ory-client/models/expand_tree'
require 'ory-client/models/flush_inactive_o_auth2_tokens_request'
require 'ory-client/models/flush_login_consent_request'
require 'ory-client/models/generic_error'
require 'ory-client/models/generic_error_content'
require 'ory-client/models/get_check_response'
require 'ory-client/models/get_managed_identity_schema_location'
require 'ory-client/models/get_relation_tuples_response'
require 'ory-client/models/get_version200_response'
require 'ory-client/models/health_not_ready_status'
require 'ory-client/models/health_status'
require 'ory-client/models/identity'
require 'ory-client/models/identity_credentials'
require 'ory-client/models/identity_credentials_oidc'
require 'ory-client/models/identity_credentials_oidc_provider'
require 'ory-client/models/identity_credentials_password'
require 'ory-client/models/identity_credentials_type'
require 'ory-client/models/identity_schema_container'
require 'ory-client/models/identity_schema_preset'
require 'ory-client/models/identity_state'
require 'ory-client/models/invite_payload'
require 'ory-client/models/is_owner_for_project_by_slug'
require 'ory-client/models/is_owner_for_project_by_slug_payload'
require 'ory-client/models/is_ready200_response'
require 'ory-client/models/is_ready503_response'
require 'ory-client/models/json_web_key'
require 'ory-client/models/json_web_key_set'
require 'ory-client/models/json_error'
require 'ory-client/models/json_patch'
require 'ory-client/models/json_web_key_set_generator_request'
require 'ory-client/models/keto_namespace'
require 'ory-client/models/login_request'
require 'ory-client/models/logout_request'
require 'ory-client/models/managed_identity_schema'
require 'ory-client/models/managed_identity_schema_validation_result'
require 'ory-client/models/needs_privileged_session_error'
require 'ory-client/models/normalized_project'
require 'ory-client/models/normalized_project_revision'
require 'ory-client/models/normalized_project_revision_hook'
require 'ory-client/models/normalized_project_revision_identity_schema'
require 'ory-client/models/normalized_project_revision_third_party_provider'
require 'ory-client/models/null_plan'
require 'ory-client/models/o_auth2_client'
require 'ory-client/models/o_auth2_token_introspection'
require 'ory-client/models/oauth2_token_response'
require 'ory-client/models/oauth_token_response'
require 'ory-client/models/open_id_connect_context'
require 'ory-client/models/pagination'
require 'ory-client/models/patch_delta'
require 'ory-client/models/patch_document'
require 'ory-client/models/previous_consent_session'
require 'ory-client/models/project'
require 'ory-client/models/project_api_key'
require 'ory-client/models/project_host'
require 'ory-client/models/project_invite'
require 'ory-client/models/project_metadata'
require 'ory-client/models/project_service_identity'
require 'ory-client/models/project_service_o_auth2'
require 'ory-client/models/project_service_permission'
require 'ory-client/models/project_services'
require 'ory-client/models/provision_mock_subscription_payload'
require 'ory-client/models/quota_custom_domains'
require 'ory-client/models/quota_project_member_seats'
require 'ory-client/models/recovery_address'
require 'ory-client/models/refresh_token_hook_request'
require 'ory-client/models/refresh_token_hook_response'
require 'ory-client/models/reject_request'
require 'ory-client/models/relation_query'
require 'ory-client/models/relation_tuple'
require 'ory-client/models/request_was_handled_response'
require 'ory-client/models/revoked_sessions'
require 'ory-client/models/schema_patch'
require 'ory-client/models/self_service_browser_location_change_required_error'
require 'ory-client/models/self_service_error'
require 'ory-client/models/self_service_flow_expired_error'
require 'ory-client/models/self_service_login_flow'
require 'ory-client/models/self_service_logout_url'
require 'ory-client/models/self_service_recovery_flow'
require 'ory-client/models/self_service_recovery_flow_state'
require 'ory-client/models/self_service_recovery_link'
require 'ory-client/models/self_service_registration_flow'
require 'ory-client/models/self_service_settings_flow'
require 'ory-client/models/self_service_settings_flow_state'
require 'ory-client/models/self_service_verification_flow'
require 'ory-client/models/self_service_verification_flow_state'
require 'ory-client/models/session'
require 'ory-client/models/session_authentication_method'
require 'ory-client/models/session_device'
require 'ory-client/models/settings_profile_form_config'
require 'ory-client/models/stripe_customer_response'
require 'ory-client/models/subject_set'
require 'ory-client/models/submit_self_service_flow_with_web_authn_registration_method'
require 'ory-client/models/submit_self_service_login_flow_body'
require 'ory-client/models/submit_self_service_login_flow_with_lookup_secret_method_body'
require 'ory-client/models/submit_self_service_login_flow_with_oidc_method_body'
require 'ory-client/models/submit_self_service_login_flow_with_password_method_body'
require 'ory-client/models/submit_self_service_login_flow_with_totp_method_body'
require 'ory-client/models/submit_self_service_login_flow_with_web_authn_method_body'
require 'ory-client/models/submit_self_service_logout_flow_without_browser_body'
require 'ory-client/models/submit_self_service_recovery_flow_body'
require 'ory-client/models/submit_self_service_recovery_flow_with_link_method_body'
require 'ory-client/models/submit_self_service_registration_flow_body'
require 'ory-client/models/submit_self_service_registration_flow_with_oidc_method_body'
require 'ory-client/models/submit_self_service_registration_flow_with_password_method_body'
require 'ory-client/models/submit_self_service_registration_flow_with_web_authn_method_body'
require 'ory-client/models/submit_self_service_settings_flow_body'
require 'ory-client/models/submit_self_service_settings_flow_with_lookup_method_body'
require 'ory-client/models/submit_self_service_settings_flow_with_oidc_method_body'
require 'ory-client/models/submit_self_service_settings_flow_with_password_method_body'
require 'ory-client/models/submit_self_service_settings_flow_with_profile_method_body'
require 'ory-client/models/submit_self_service_settings_flow_with_totp_method_body'
require 'ory-client/models/submit_self_service_settings_flow_with_web_authn_method_body'
require 'ory-client/models/submit_self_service_verification_flow_body'
require 'ory-client/models/submit_self_service_verification_flow_with_link_method_body'
require 'ory-client/models/subscription'
require 'ory-client/models/successful_project_update'
require 'ory-client/models/successful_self_service_login_without_browser'
require 'ory-client/models/successful_self_service_registration_without_browser'
require 'ory-client/models/token_pagination'
require 'ory-client/models/token_pagination_headers'
require 'ory-client/models/trust_jwt_grant_issuer_body'
require 'ory-client/models/trusted_json_web_key'
require 'ory-client/models/trusted_jwt_grant_issuer'
require 'ory-client/models/ui_container'
require 'ory-client/models/ui_node'
require 'ory-client/models/ui_node_anchor_attributes'
require 'ory-client/models/ui_node_attributes'
require 'ory-client/models/ui_node_image_attributes'
require 'ory-client/models/ui_node_input_attributes'
require 'ory-client/models/ui_node_meta'
require 'ory-client/models/ui_node_script_attributes'
require 'ory-client/models/ui_node_text_attributes'
require 'ory-client/models/ui_text'
require 'ory-client/models/update_custom_hostname_body'
require 'ory-client/models/update_project'
require 'ory-client/models/update_subscription_payload'
require 'ory-client/models/userinfo_response'
require 'ory-client/models/verifiable_identity_address'
require 'ory-client/models/version'
require 'ory-client/models/warning'
require 'ory-client/models/well_known'

# APIs
require 'ory-client/api/admin_api'
require 'ory-client/api/metadata_api'
require 'ory-client/api/public_api'
require 'ory-client/api/read_api'
require 'ory-client/api/v0alpha2_api'
require 'ory-client/api/write_api'

module OryClient
  class << self
    # Customize default settings for the SDK using block.
    #   OryClient.configure do |config|
    #     config.username = "xxx"
    #     config.password = "xxx"
    #   end
    # If no block given, return the default Configuration object.
    def configure
      if block_given?
        yield(Configuration.default)
      else
        Configuration.default
      end
    end
  end
end
