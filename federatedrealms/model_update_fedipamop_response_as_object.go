/*
Infoblox FEDERATEDREALMS API

OpenAPI specification for Infoblox NIOS WAPI FEDERATEDREALMS objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package federatedrealms

import (
	"encoding/json"
)

// checks if the UpdateFedipamopResponseAsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateFedipamopResponseAsObject{}

// UpdateFedipamopResponseAsObject The response format to update __Fedipamop__ in object format.
type UpdateFedipamopResponseAsObject struct {
	Result               *Fedipamop `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateFedipamopResponseAsObject UpdateFedipamopResponseAsObject

// NewUpdateFedipamopResponseAsObject instantiates a new UpdateFedipamopResponseAsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateFedipamopResponseAsObject() *UpdateFedipamopResponseAsObject {
	this := UpdateFedipamopResponseAsObject{}
	return &this
}

// NewUpdateFedipamopResponseAsObjectWithDefaults instantiates a new UpdateFedipamopResponseAsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateFedipamopResponseAsObjectWithDefaults() *UpdateFedipamopResponseAsObject {
	this := UpdateFedipamopResponseAsObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *UpdateFedipamopResponseAsObject) GetResult() Fedipamop {
	if o == nil || IsNil(o.Result) {
		var ret Fedipamop
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFedipamopResponseAsObject) GetResultOk() (*Fedipamop, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *UpdateFedipamopResponseAsObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given Fedipamop and assigns it to the Result field.
func (o *UpdateFedipamopResponseAsObject) SetResult(v Fedipamop) {
	o.Result = &v
}

func (o UpdateFedipamopResponseAsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateFedipamopResponseAsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateFedipamopResponseAsObject) UnmarshalJSON(data []byte) (err error) {
	varUpdateFedipamopResponseAsObject := _UpdateFedipamopResponseAsObject{}

	err = json.Unmarshal(data, &varUpdateFedipamopResponseAsObject)

	if err != nil {
		return err
	}

	*o = UpdateFedipamopResponseAsObject(varUpdateFedipamopResponseAsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateFedipamopResponseAsObject struct {
	value *UpdateFedipamopResponseAsObject
	isSet bool
}

func (v NullableUpdateFedipamopResponseAsObject) Get() *UpdateFedipamopResponseAsObject {
	return v.value
}

func (v *NullableUpdateFedipamopResponseAsObject) Set(val *UpdateFedipamopResponseAsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateFedipamopResponseAsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateFedipamopResponseAsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateFedipamopResponseAsObject(val *UpdateFedipamopResponseAsObject) *NullableUpdateFedipamopResponseAsObject {
	return &NullableUpdateFedipamopResponseAsObject{value: val, isSet: true}
}

func (v NullableUpdateFedipamopResponseAsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateFedipamopResponseAsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
