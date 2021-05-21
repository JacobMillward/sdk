/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.6
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PluginSettings struct for PluginSettings
type PluginSettings struct {
	// args
	Args []string `json:"Args"`
	// devices
	Devices []PluginDevice `json:"Devices"`
	// env
	Env []string `json:"Env"`
	// mounts
	Mounts []PluginMount `json:"Mounts"`
}

// NewPluginSettings instantiates a new PluginSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPluginSettings(args []string, devices []PluginDevice, env []string, mounts []PluginMount) *PluginSettings {
	this := PluginSettings{}
	this.Args = args
	this.Devices = devices
	this.Env = env
	this.Mounts = mounts
	return &this
}

// NewPluginSettingsWithDefaults instantiates a new PluginSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginSettingsWithDefaults() *PluginSettings {
	this := PluginSettings{}
	return &this
}

// GetArgs returns the Args field value
func (o *PluginSettings) GetArgs() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Args
}

// GetArgsOk returns a tuple with the Args field value
// and a boolean to check if the value has been set.
func (o *PluginSettings) GetArgsOk() ([]string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Args, true
}

// SetArgs sets field value
func (o *PluginSettings) SetArgs(v []string) {
	o.Args = v
}

// GetDevices returns the Devices field value
func (o *PluginSettings) GetDevices() []PluginDevice {
	if o == nil {
		var ret []PluginDevice
		return ret
	}

	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value
// and a boolean to check if the value has been set.
func (o *PluginSettings) GetDevicesOk() ([]PluginDevice, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Devices, true
}

// SetDevices sets field value
func (o *PluginSettings) SetDevices(v []PluginDevice) {
	o.Devices = v
}

// GetEnv returns the Env field value
func (o *PluginSettings) GetEnv() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Env
}

// GetEnvOk returns a tuple with the Env field value
// and a boolean to check if the value has been set.
func (o *PluginSettings) GetEnvOk() ([]string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Env, true
}

// SetEnv sets field value
func (o *PluginSettings) SetEnv(v []string) {
	o.Env = v
}

// GetMounts returns the Mounts field value
func (o *PluginSettings) GetMounts() []PluginMount {
	if o == nil {
		var ret []PluginMount
		return ret
	}

	return o.Mounts
}

// GetMountsOk returns a tuple with the Mounts field value
// and a boolean to check if the value has been set.
func (o *PluginSettings) GetMountsOk() ([]PluginMount, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Mounts, true
}

// SetMounts sets field value
func (o *PluginSettings) SetMounts(v []PluginMount) {
	o.Mounts = v
}

func (o PluginSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["Args"] = o.Args
	}
	if true {
		toSerialize["Devices"] = o.Devices
	}
	if true {
		toSerialize["Env"] = o.Env
	}
	if true {
		toSerialize["Mounts"] = o.Mounts
	}
	return json.Marshal(toSerialize)
}

type NullablePluginSettings struct {
	value *PluginSettings
	isSet bool
}

func (v NullablePluginSettings) Get() *PluginSettings {
	return v.value
}

func (v *NullablePluginSettings) Set(val *PluginSettings) {
	v.value = val
	v.isSet = true
}

func (v NullablePluginSettings) IsSet() bool {
	return v.isSet
}

func (v *NullablePluginSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePluginSettings(val *PluginSettings) *NullablePluginSettings {
	return &NullablePluginSettings{value: val, isSet: true}
}

func (v NullablePluginSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePluginSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

