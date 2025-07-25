/*
Infoblox DTC API

OpenAPI specification for Infoblox NIOS WAPI DTC objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dtc

import (
	"encoding/json"
)

// checks if the CreateDtcRecordAResponseAsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDtcRecordAResponseAsObject{}

// CreateDtcRecordAResponseAsObject The response format to create __DtcRecordA__ in object format.
type CreateDtcRecordAResponseAsObject struct {
	Result               *DtcRecordA `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateDtcRecordAResponseAsObject CreateDtcRecordAResponseAsObject

// NewCreateDtcRecordAResponseAsObject instantiates a new CreateDtcRecordAResponseAsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDtcRecordAResponseAsObject() *CreateDtcRecordAResponseAsObject {
	this := CreateDtcRecordAResponseAsObject{}
	return &this
}

// NewCreateDtcRecordAResponseAsObjectWithDefaults instantiates a new CreateDtcRecordAResponseAsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDtcRecordAResponseAsObjectWithDefaults() *CreateDtcRecordAResponseAsObject {
	this := CreateDtcRecordAResponseAsObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CreateDtcRecordAResponseAsObject) GetResult() DtcRecordA {
	if o == nil || IsNil(o.Result) {
		var ret DtcRecordA
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDtcRecordAResponseAsObject) GetResultOk() (*DtcRecordA, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CreateDtcRecordAResponseAsObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given DtcRecordA and assigns it to the Result field.
func (o *CreateDtcRecordAResponseAsObject) SetResult(v DtcRecordA) {
	o.Result = &v
}

func (o CreateDtcRecordAResponseAsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDtcRecordAResponseAsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateDtcRecordAResponseAsObject) UnmarshalJSON(data []byte) (err error) {
	varCreateDtcRecordAResponseAsObject := _CreateDtcRecordAResponseAsObject{}

	err = json.Unmarshal(data, &varCreateDtcRecordAResponseAsObject)

	if err != nil {
		return err
	}

	*o = CreateDtcRecordAResponseAsObject(varCreateDtcRecordAResponseAsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateDtcRecordAResponseAsObject struct {
	value *CreateDtcRecordAResponseAsObject
	isSet bool
}

func (v NullableCreateDtcRecordAResponseAsObject) Get() *CreateDtcRecordAResponseAsObject {
	return v.value
}

func (v *NullableCreateDtcRecordAResponseAsObject) Set(val *CreateDtcRecordAResponseAsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDtcRecordAResponseAsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDtcRecordAResponseAsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDtcRecordAResponseAsObject(val *CreateDtcRecordAResponseAsObject) *NullableCreateDtcRecordAResponseAsObject {
	return &NullableCreateDtcRecordAResponseAsObject{value: val, isSet: true}
}

func (v NullableCreateDtcRecordAResponseAsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDtcRecordAResponseAsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
