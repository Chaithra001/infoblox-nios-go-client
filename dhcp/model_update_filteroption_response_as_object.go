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

// checks if the UpdateFilteroptionResponseAsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateFilteroptionResponseAsObject{}

// UpdateFilteroptionResponseAsObject The response format to update __Filteroption__ in object format.
type UpdateFilteroptionResponseAsObject struct {
	Result               *Filteroption `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateFilteroptionResponseAsObject UpdateFilteroptionResponseAsObject

// NewUpdateFilteroptionResponseAsObject instantiates a new UpdateFilteroptionResponseAsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateFilteroptionResponseAsObject() *UpdateFilteroptionResponseAsObject {
	this := UpdateFilteroptionResponseAsObject{}
	return &this
}

// NewUpdateFilteroptionResponseAsObjectWithDefaults instantiates a new UpdateFilteroptionResponseAsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateFilteroptionResponseAsObjectWithDefaults() *UpdateFilteroptionResponseAsObject {
	this := UpdateFilteroptionResponseAsObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *UpdateFilteroptionResponseAsObject) GetResult() Filteroption {
	if o == nil || IsNil(o.Result) {
		var ret Filteroption
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFilteroptionResponseAsObject) GetResultOk() (*Filteroption, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *UpdateFilteroptionResponseAsObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given Filteroption and assigns it to the Result field.
func (o *UpdateFilteroptionResponseAsObject) SetResult(v Filteroption) {
	o.Result = &v
}

func (o UpdateFilteroptionResponseAsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateFilteroptionResponseAsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateFilteroptionResponseAsObject) UnmarshalJSON(data []byte) (err error) {
	varUpdateFilteroptionResponseAsObject := _UpdateFilteroptionResponseAsObject{}

	err = json.Unmarshal(data, &varUpdateFilteroptionResponseAsObject)

	if err != nil {
		return err
	}

	*o = UpdateFilteroptionResponseAsObject(varUpdateFilteroptionResponseAsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateFilteroptionResponseAsObject struct {
	value *UpdateFilteroptionResponseAsObject
	isSet bool
}

func (v NullableUpdateFilteroptionResponseAsObject) Get() *UpdateFilteroptionResponseAsObject {
	return v.value
}

func (v *NullableUpdateFilteroptionResponseAsObject) Set(val *UpdateFilteroptionResponseAsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateFilteroptionResponseAsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateFilteroptionResponseAsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateFilteroptionResponseAsObject(val *UpdateFilteroptionResponseAsObject) *NullableUpdateFilteroptionResponseAsObject {
	return &NullableUpdateFilteroptionResponseAsObject{value: val, isSet: true}
}

func (v NullableUpdateFilteroptionResponseAsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateFilteroptionResponseAsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
