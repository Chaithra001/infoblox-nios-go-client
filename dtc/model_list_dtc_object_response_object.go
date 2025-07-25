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

// checks if the ListDtcObjectResponseObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListDtcObjectResponseObject{}

// ListDtcObjectResponseObject The response format to retrieve __DtcObject__ objects.
type ListDtcObjectResponseObject struct {
	Result               []DtcObject `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListDtcObjectResponseObject ListDtcObjectResponseObject

// NewListDtcObjectResponseObject instantiates a new ListDtcObjectResponseObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDtcObjectResponseObject() *ListDtcObjectResponseObject {
	this := ListDtcObjectResponseObject{}
	return &this
}

// NewListDtcObjectResponseObjectWithDefaults instantiates a new ListDtcObjectResponseObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDtcObjectResponseObjectWithDefaults() *ListDtcObjectResponseObject {
	this := ListDtcObjectResponseObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ListDtcObjectResponseObject) GetResult() []DtcObject {
	if o == nil || IsNil(o.Result) {
		var ret []DtcObject
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDtcObjectResponseObject) GetResultOk() ([]DtcObject, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ListDtcObjectResponseObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given []DtcObject and assigns it to the Result field.
func (o *ListDtcObjectResponseObject) SetResult(v []DtcObject) {
	o.Result = v
}

func (o ListDtcObjectResponseObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListDtcObjectResponseObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListDtcObjectResponseObject) UnmarshalJSON(data []byte) (err error) {
	varListDtcObjectResponseObject := _ListDtcObjectResponseObject{}

	err = json.Unmarshal(data, &varListDtcObjectResponseObject)

	if err != nil {
		return err
	}

	*o = ListDtcObjectResponseObject(varListDtcObjectResponseObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListDtcObjectResponseObject struct {
	value *ListDtcObjectResponseObject
	isSet bool
}

func (v NullableListDtcObjectResponseObject) Get() *ListDtcObjectResponseObject {
	return v.value
}

func (v *NullableListDtcObjectResponseObject) Set(val *ListDtcObjectResponseObject) {
	v.value = val
	v.isSet = true
}

func (v NullableListDtcObjectResponseObject) IsSet() bool {
	return v.isSet
}

func (v *NullableListDtcObjectResponseObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDtcObjectResponseObject(val *ListDtcObjectResponseObject) *NullableListDtcObjectResponseObject {
	return &NullableListDtcObjectResponseObject{value: val, isSet: true}
}

func (v NullableListDtcObjectResponseObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDtcObjectResponseObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
