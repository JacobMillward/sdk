/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.68
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InvitePayload struct for InvitePayload
type InvitePayload struct {
	InviteeEmail *string `json:"invitee_email,omitempty"`
}

// NewInvitePayload instantiates a new InvitePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvitePayload() *InvitePayload {
	this := InvitePayload{}
	return &this
}

// NewInvitePayloadWithDefaults instantiates a new InvitePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvitePayloadWithDefaults() *InvitePayload {
	this := InvitePayload{}
	return &this
}

// GetInviteeEmail returns the InviteeEmail field value if set, zero value otherwise.
func (o *InvitePayload) GetInviteeEmail() string {
	if o == nil || o.InviteeEmail == nil {
		var ret string
		return ret
	}
	return *o.InviteeEmail
}

// GetInviteeEmailOk returns a tuple with the InviteeEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitePayload) GetInviteeEmailOk() (*string, bool) {
	if o == nil || o.InviteeEmail == nil {
		return nil, false
	}
	return o.InviteeEmail, true
}

// HasInviteeEmail returns a boolean if a field has been set.
func (o *InvitePayload) HasInviteeEmail() bool {
	if o != nil && o.InviteeEmail != nil {
		return true
	}

	return false
}

// SetInviteeEmail gets a reference to the given string and assigns it to the InviteeEmail field.
func (o *InvitePayload) SetInviteeEmail(v string) {
	o.InviteeEmail = &v
}

func (o InvitePayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.InviteeEmail != nil {
		toSerialize["invitee_email"] = o.InviteeEmail
	}
	return json.Marshal(toSerialize)
}

type NullableInvitePayload struct {
	value *InvitePayload
	isSet bool
}

func (v NullableInvitePayload) Get() *InvitePayload {
	return v.value
}

func (v *NullableInvitePayload) Set(val *InvitePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableInvitePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableInvitePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvitePayload(val *InvitePayload) *NullableInvitePayload {
	return &NullableInvitePayload{value: val, isSet: true}
}

func (v NullableInvitePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvitePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

