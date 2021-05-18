/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * The version of the OpenAPI document: v0.6.3-alpha.1
 * Contact: hi@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// SubmitSelfServiceVerificationFlowWithLinkMethod : nolint:deadcode,unused



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct SubmitSelfServiceVerificationFlowWithLinkMethod {
    /// Sending the anti-csrf token is only required for browser login flows.
    #[serde(rename = "csrf_token", skip_serializing_if = "Option::is_none")]
    pub csrf_token: Option<String>,
    /// Email to Verify  Needs to be set when initiating the flow. If the email is a registered verification email, a verification link will be sent. If the email is not known, a email with details on what happened will be sent instead.  format: email in: body
    #[serde(rename = "email", skip_serializing_if = "Option::is_none")]
    pub email: Option<String>,
}

impl SubmitSelfServiceVerificationFlowWithLinkMethod {
    /// nolint:deadcode,unused
    pub fn new() -> SubmitSelfServiceVerificationFlowWithLinkMethod {
        SubmitSelfServiceVerificationFlowWithLinkMethod {
            csrf_token: None,
            email: None,
        }
    }
}


