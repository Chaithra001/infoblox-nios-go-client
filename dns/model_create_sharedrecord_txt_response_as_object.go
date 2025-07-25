/*
Infoblox DNS API

OpenAPI specification for Infoblox NIOS WAPI DNS objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"encoding/json"
)

// checks if the CreateSharedrecordTxtResponseAsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSharedrecordTxtResponseAsObject{}

// CreateSharedrecordTxtResponseAsObject The response format to create __SharedrecordTxt__ in object format.
type CreateSharedrecordTxtResponseAsObject struct {
	Result               *SharedrecordTxt `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateSharedrecordTxtResponseAsObject CreateSharedrecordTxtResponseAsObject

// NewCreateSharedrecordTxtResponseAsObject instantiates a new CreateSharedrecordTxtResponseAsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSharedrecordTxtResponseAsObject() *CreateSharedrecordTxtResponseAsObject {
	this := CreateSharedrecordTxtResponseAsObject{}
	return &this
}

// NewCreateSharedrecordTxtResponseAsObjectWithDefaults instantiates a new CreateSharedrecordTxtResponseAsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSharedrecordTxtResponseAsObjectWithDefaults() *CreateSharedrecordTxtResponseAsObject {
	this := CreateSharedrecordTxtResponseAsObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CreateSharedrecordTxtResponseAsObject) GetResult() SharedrecordTxt {
	if o == nil || IsNil(o.Result) {
		var ret SharedrecordTxt
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSharedrecordTxtResponseAsObject) GetResultOk() (*SharedrecordTxt, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CreateSharedrecordTxtResponseAsObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given SharedrecordTxt and assigns it to the Result field.
func (o *CreateSharedrecordTxtResponseAsObject) SetResult(v SharedrecordTxt) {
	o.Result = &v
}

func (o CreateSharedrecordTxtResponseAsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSharedrecordTxtResponseAsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateSharedrecordTxtResponseAsObject) UnmarshalJSON(data []byte) (err error) {
	varCreateSharedrecordTxtResponseAsObject := _CreateSharedrecordTxtResponseAsObject{}

	err = json.Unmarshal(data, &varCreateSharedrecordTxtResponseAsObject)

	if err != nil {
		return err
	}

	*o = CreateSharedrecordTxtResponseAsObject(varCreateSharedrecordTxtResponseAsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateSharedrecordTxtResponseAsObject struct {
	value *CreateSharedrecordTxtResponseAsObject
	isSet bool
}

func (v NullableCreateSharedrecordTxtResponseAsObject) Get() *CreateSharedrecordTxtResponseAsObject {
	return v.value
}

func (v *NullableCreateSharedrecordTxtResponseAsObject) Set(val *CreateSharedrecordTxtResponseAsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSharedrecordTxtResponseAsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSharedrecordTxtResponseAsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSharedrecordTxtResponseAsObject(val *CreateSharedrecordTxtResponseAsObject) *NullableCreateSharedrecordTxtResponseAsObject {
	return &NullableCreateSharedrecordTxtResponseAsObject{value: val, isSet: true}
}

func (v NullableCreateSharedrecordTxtResponseAsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSharedrecordTxtResponseAsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
