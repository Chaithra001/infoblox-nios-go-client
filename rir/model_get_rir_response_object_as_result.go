/*
Infoblox RIR API

OpenAPI specification for Infoblox NIOS WAPI RIR objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rir

import (
	"encoding/json"
)

// checks if the GetRirResponseObjectAsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRirResponseObjectAsResult{}

// GetRirResponseObjectAsResult The response format to retrieve __Rir__ objects.
type GetRirResponseObjectAsResult struct {
	Result *Rir `json:"result,omitempty"`
}

// NewGetRirResponseObjectAsResult instantiates a new GetRirResponseObjectAsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRirResponseObjectAsResult() *GetRirResponseObjectAsResult {
	this := GetRirResponseObjectAsResult{}
	return &this
}

// NewGetRirResponseObjectAsResultWithDefaults instantiates a new GetRirResponseObjectAsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRirResponseObjectAsResultWithDefaults() *GetRirResponseObjectAsResult {
	this := GetRirResponseObjectAsResult{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetRirResponseObjectAsResult) GetResult() Rir {
	if o == nil || IsNil(o.Result) {
		var ret Rir
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRirResponseObjectAsResult) GetResultOk() (*Rir, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetRirResponseObjectAsResult) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given Rir and assigns it to the Result field.
func (o *GetRirResponseObjectAsResult) SetResult(v Rir) {
	o.Result = &v
}

func (o GetRirResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRirResponseObjectAsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetRirResponseObjectAsResult struct {
	value *GetRirResponseObjectAsResult
	isSet bool
}

func (v NullableGetRirResponseObjectAsResult) Get() *GetRirResponseObjectAsResult {
	return v.value
}

func (v *NullableGetRirResponseObjectAsResult) Set(val *GetRirResponseObjectAsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRirResponseObjectAsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRirResponseObjectAsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRirResponseObjectAsResult(val *GetRirResponseObjectAsResult) *NullableGetRirResponseObjectAsResult {
	return &NullableGetRirResponseObjectAsResult{value: val, isSet: true}
}

func (v NullableGetRirResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRirResponseObjectAsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
