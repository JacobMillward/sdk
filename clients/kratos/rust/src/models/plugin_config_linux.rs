/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * The version of the OpenAPI document: v0.6.3-alpha.1
 * Contact: hi@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// PluginConfigLinux : PluginConfigLinux plugin config linux



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct PluginConfigLinux {
    /// allow all devices
    #[serde(rename = "AllowAllDevices")]
    pub allow_all_devices: bool,
    /// capabilities
    #[serde(rename = "Capabilities")]
    pub capabilities: Vec<String>,
    /// devices
    #[serde(rename = "Devices")]
    pub devices: Vec<crate::models::PluginDevice>,
}

impl PluginConfigLinux {
    /// PluginConfigLinux plugin config linux
    pub fn new(allow_all_devices: bool, capabilities: Vec<String>, devices: Vec<crate::models::PluginDevice>) -> PluginConfigLinux {
        PluginConfigLinux {
            allow_all_devices,
            capabilities,
            devices,
        }
    }
}


