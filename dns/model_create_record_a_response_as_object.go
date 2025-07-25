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

// checks if the CreateRecordAResponseAsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRecordAResponseAsObject{}

// CreateRecordAResponseAsObject The response format to create __RecordA__ in object format.
type CreateRecordAResponseAsObject struct {
	Result               *RecordA `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateRecordAResponseAsObject CreateRecordAResponseAsObject

// NewCreateRecordAResponseAsObject instantiates a new CreateRecordAResponseAsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRecordAResponseAsObject() *CreateRecordAResponseAsObject {
	this := CreateRecordAResponseAsObject{}
	return &this
}

// NewCreateRecordAResponseAsObjectWithDefaults instantiates a new CreateRecordAResponseAsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRecordAResponseAsObjectWithDefaults() *CreateRecordAResponseAsObject {
	this := CreateRecordAResponseAsObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CreateRecordAResponseAsObject) GetResult() RecordA {
	if o == nil || IsNil(o.Result) {
		var ret RecordA
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRecordAResponseAsObject) GetResultOk() (*RecordA, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CreateRecordAResponseAsObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given RecordA and assigns it to the Result field.
func (o *CreateRecordAResponseAsObject) SetResult(v RecordA) {
	o.Result = &v
}

func (o CreateRecordAResponseAsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRecordAResponseAsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateRecordAResponseAsObject) UnmarshalJSON(data []byte) (err error) {
	varCreateRecordAResponseAsObject := _CreateRecordAResponseAsObject{}

	err = json.Unmarshal(data, &varCreateRecordAResponseAsObject)

	if err != nil {
		return err
	}

	*o = CreateRecordAResponseAsObject(varCreateRecordAResponseAsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateRecordAResponseAsObject struct {
	value *CreateRecordAResponseAsObject
	isSet bool
}

func (v NullableCreateRecordAResponseAsObject) Get() *CreateRecordAResponseAsObject {
	return v.value
}

func (v *NullableCreateRecordAResponseAsObject) Set(val *CreateRecordAResponseAsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRecordAResponseAsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRecordAResponseAsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRecordAResponseAsObject(val *CreateRecordAResponseAsObject) *NullableCreateRecordAResponseAsObject {
	return &NullableCreateRecordAResponseAsObject{value: val, isSet: true}
}

func (v NullableCreateRecordAResponseAsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRecordAResponseAsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
