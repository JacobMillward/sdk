/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * The version of the OpenAPI document: v0.7.6-alpha.5
 * Contact: hi@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// SubmitSelfServiceRegistrationFlowWithPasswordMethodBody : SubmitSelfServiceRegistrationFlowWithPasswordMethodBody is used to decode the registration form payload when using the password method.



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct SubmitSelfServiceRegistrationFlowWithPasswordMethodBody {
    /// The CSRF Token
    #[serde(rename = "csrf_token", skip_serializing_if = "Option::is_none")]
    pub csrf_token: Option<String>,
    /// Method to use  This field must be set to `password` when using the password method.
    #[serde(rename = "method")]
    pub method: String,
    /// Password to sign the user up with
    #[serde(rename = "password")]
    pub password: String,
    /// The identity's traits
    #[serde(rename = "traits")]
    pub traits: serde_json::Value,
}

impl SubmitSelfServiceRegistrationFlowWithPasswordMethodBody {
    /// SubmitSelfServiceRegistrationFlowWithPasswordMethodBody is used to decode the registration form payload when using the password method.
    pub fn new(method: String, password: String, traits: serde_json::Value) -> SubmitSelfServiceRegistrationFlowWithPasswordMethodBody {
        SubmitSelfServiceRegistrationFlowWithPasswordMethodBody {
            csrf_token: None,
            method,
            password,
            traits,
        }
    }
}


