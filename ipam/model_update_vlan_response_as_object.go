/*
Infoblox IPAM API

OpenAPI specification for Infoblox NIOS WAPI IPAM objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"encoding/json"
)

// checks if the UpdateVlanResponseAsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateVlanResponseAsObject{}

// UpdateVlanResponseAsObject The response format to update __Vlan__ in object format.
type UpdateVlanResponseAsObject struct {
	Result               *Vlan `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateVlanResponseAsObject UpdateVlanResponseAsObject

// NewUpdateVlanResponseAsObject instantiates a new UpdateVlanResponseAsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateVlanResponseAsObject() *UpdateVlanResponseAsObject {
	this := UpdateVlanResponseAsObject{}
	return &this
}

// NewUpdateVlanResponseAsObjectWithDefaults instantiates a new UpdateVlanResponseAsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateVlanResponseAsObjectWithDefaults() *UpdateVlanResponseAsObject {
	this := UpdateVlanResponseAsObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *UpdateVlanResponseAsObject) GetResult() Vlan {
	if o == nil || IsNil(o.Result) {
		var ret Vlan
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVlanResponseAsObject) GetResultOk() (*Vlan, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *UpdateVlanResponseAsObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given Vlan and assigns it to the Result field.
func (o *UpdateVlanResponseAsObject) SetResult(v Vlan) {
	o.Result = &v
}

func (o UpdateVlanResponseAsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateVlanResponseAsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateVlanResponseAsObject) UnmarshalJSON(data []byte) (err error) {
	varUpdateVlanResponseAsObject := _UpdateVlanResponseAsObject{}

	err = json.Unmarshal(data, &varUpdateVlanResponseAsObject)

	if err != nil {
		return err
	}

	*o = UpdateVlanResponseAsObject(varUpdateVlanResponseAsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateVlanResponseAsObject struct {
	value *UpdateVlanResponseAsObject
	isSet bool
}

func (v NullableUpdateVlanResponseAsObject) Get() *UpdateVlanResponseAsObject {
	return v.value
}

func (v *NullableUpdateVlanResponseAsObject) Set(val *UpdateVlanResponseAsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateVlanResponseAsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateVlanResponseAsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateVlanResponseAsObject(val *UpdateVlanResponseAsObject) *NullableUpdateVlanResponseAsObject {
	return &NullableUpdateVlanResponseAsObject{value: val, isSet: true}
}

func (v NullableUpdateVlanResponseAsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateVlanResponseAsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
