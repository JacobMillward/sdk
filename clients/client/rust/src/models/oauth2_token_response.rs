/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.2.0-alpha.37
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// Oauth2TokenResponse : The Access Token Response



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct Oauth2TokenResponse {
    #[serde(rename = "access_token", skip_serializing_if = "Option::is_none")]
    pub access_token: Option<String>,
    #[serde(rename = "expires_in", skip_serializing_if = "Option::is_none")]
    pub expires_in: Option<i64>,
    #[serde(rename = "id_token", skip_serializing_if = "Option::is_none")]
    pub id_token: Option<String>,
    #[serde(rename = "refresh_token", skip_serializing_if = "Option::is_none")]
    pub refresh_token: Option<String>,
    #[serde(rename = "scope", skip_serializing_if = "Option::is_none")]
    pub scope: Option<String>,
    #[serde(rename = "token_type", skip_serializing_if = "Option::is_none")]
    pub token_type: Option<String>,
}

impl Default for Oauth2TokenResponse {
    fn default() -> Self {
        Self::new()
    }
}

impl Oauth2TokenResponse {
    /// The Access Token Response
    pub fn new() -> Oauth2TokenResponse {
        Oauth2TokenResponse {
                access_token: None,
                expires_in: None,
                id_token: None,
                refresh_token: None,
                scope: None,
                token_type: None,
        }
    }
}


