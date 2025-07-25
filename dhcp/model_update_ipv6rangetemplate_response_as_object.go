/*
Infoblox DHCP API

OpenAPI specification for Infoblox NIOS WAPI DHCP objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dhcp

import (
	"encoding/json"
)

// checks if the UpdateIpv6rangetemplateResponseAsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateIpv6rangetemplateResponseAsObject{}

// UpdateIpv6rangetemplateResponseAsObject The response format to update __Ipv6rangetemplate__ in object format.
type UpdateIpv6rangetemplateResponseAsObject struct {
	Result               *Ipv6rangetemplate `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateIpv6rangetemplateResponseAsObject UpdateIpv6rangetemplateResponseAsObject

// NewUpdateIpv6rangetemplateResponseAsObject instantiates a new UpdateIpv6rangetemplateResponseAsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateIpv6rangetemplateResponseAsObject() *UpdateIpv6rangetemplateResponseAsObject {
	this := UpdateIpv6rangetemplateResponseAsObject{}
	return &this
}

// NewUpdateIpv6rangetemplateResponseAsObjectWithDefaults instantiates a new UpdateIpv6rangetemplateResponseAsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateIpv6rangetemplateResponseAsObjectWithDefaults() *UpdateIpv6rangetemplateResponseAsObject {
	this := UpdateIpv6rangetemplateResponseAsObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *UpdateIpv6rangetemplateResponseAsObject) GetResult() Ipv6rangetemplate {
	if o == nil || IsNil(o.Result) {
		var ret Ipv6rangetemplate
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIpv6rangetemplateResponseAsObject) GetResultOk() (*Ipv6rangetemplate, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *UpdateIpv6rangetemplateResponseAsObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given Ipv6rangetemplate and assigns it to the Result field.
func (o *UpdateIpv6rangetemplateResponseAsObject) SetResult(v Ipv6rangetemplate) {
	o.Result = &v
}

func (o UpdateIpv6rangetemplateResponseAsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateIpv6rangetemplateResponseAsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateIpv6rangetemplateResponseAsObject) UnmarshalJSON(data []byte) (err error) {
	varUpdateIpv6rangetemplateResponseAsObject := _UpdateIpv6rangetemplateResponseAsObject{}

	err = json.Unmarshal(data, &varUpdateIpv6rangetemplateResponseAsObject)

	if err != nil {
		return err
	}

	*o = UpdateIpv6rangetemplateResponseAsObject(varUpdateIpv6rangetemplateResponseAsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateIpv6rangetemplateResponseAsObject struct {
	value *UpdateIpv6rangetemplateResponseAsObject
	isSet bool
}

func (v NullableUpdateIpv6rangetemplateResponseAsObject) Get() *UpdateIpv6rangetemplateResponseAsObject {
	return v.value
}

func (v *NullableUpdateIpv6rangetemplateResponseAsObject) Set(val *UpdateIpv6rangetemplateResponseAsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateIpv6rangetemplateResponseAsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateIpv6rangetemplateResponseAsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateIpv6rangetemplateResponseAsObject(val *UpdateIpv6rangetemplateResponseAsObject) *NullableUpdateIpv6rangetemplateResponseAsObject {
	return &NullableUpdateIpv6rangetemplateResponseAsObject{value: val, isSet: true}
}

func (v NullableUpdateIpv6rangetemplateResponseAsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateIpv6rangetemplateResponseAsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
