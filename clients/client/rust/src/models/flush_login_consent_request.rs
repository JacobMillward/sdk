/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.2.0-alpha.37
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct FlushLoginConsentRequest {
    /// NotAfter sets after which point tokens should not be flushed. This is useful when you want to keep a history of recent login and consent database entries for auditing.
    #[serde(rename = "notAfter", skip_serializing_if = "Option::is_none")]
    pub not_after: Option<String>,
}

impl Default for FlushLoginConsentRequest {
    fn default() -> Self {
        Self::new()
    }
}

impl FlushLoginConsentRequest {
    pub fn new() -> FlushLoginConsentRequest {
        FlushLoginConsentRequest {
                not_after: None,
        }
    }
}


