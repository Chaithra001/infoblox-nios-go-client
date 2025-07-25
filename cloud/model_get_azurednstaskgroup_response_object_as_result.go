/*
Infoblox CLOUD API

OpenAPI specification for Infoblox NIOS WAPI CLOUD objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloud

import (
	"encoding/json"
)

// checks if the GetAzurednstaskgroupResponseObjectAsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAzurednstaskgroupResponseObjectAsResult{}

// GetAzurednstaskgroupResponseObjectAsResult The response format to retrieve __Azurednstaskgroup__ objects.
type GetAzurednstaskgroupResponseObjectAsResult struct {
	Result *Azurednstaskgroup `json:"result,omitempty"`
}

// NewGetAzurednstaskgroupResponseObjectAsResult instantiates a new GetAzurednstaskgroupResponseObjectAsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAzurednstaskgroupResponseObjectAsResult() *GetAzurednstaskgroupResponseObjectAsResult {
	this := GetAzurednstaskgroupResponseObjectAsResult{}
	return &this
}

// NewGetAzurednstaskgroupResponseObjectAsResultWithDefaults instantiates a new GetAzurednstaskgroupResponseObjectAsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAzurednstaskgroupResponseObjectAsResultWithDefaults() *GetAzurednstaskgroupResponseObjectAsResult {
	this := GetAzurednstaskgroupResponseObjectAsResult{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetAzurednstaskgroupResponseObjectAsResult) GetResult() Azurednstaskgroup {
	if o == nil || IsNil(o.Result) {
		var ret Azurednstaskgroup
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAzurednstaskgroupResponseObjectAsResult) GetResultOk() (*Azurednstaskgroup, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetAzurednstaskgroupResponseObjectAsResult) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given Azurednstaskgroup and assigns it to the Result field.
func (o *GetAzurednstaskgroupResponseObjectAsResult) SetResult(v Azurednstaskgroup) {
	o.Result = &v
}

func (o GetAzurednstaskgroupResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAzurednstaskgroupResponseObjectAsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetAzurednstaskgroupResponseObjectAsResult struct {
	value *GetAzurednstaskgroupResponseObjectAsResult
	isSet bool
}

func (v NullableGetAzurednstaskgroupResponseObjectAsResult) Get() *GetAzurednstaskgroupResponseObjectAsResult {
	return v.value
}

func (v *NullableGetAzurednstaskgroupResponseObjectAsResult) Set(val *GetAzurednstaskgroupResponseObjectAsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAzurednstaskgroupResponseObjectAsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAzurednstaskgroupResponseObjectAsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAzurednstaskgroupResponseObjectAsResult(val *GetAzurednstaskgroupResponseObjectAsResult) *NullableGetAzurednstaskgroupResponseObjectAsResult {
	return &NullableGetAzurednstaskgroupResponseObjectAsResult{value: val, isSet: true}
}

func (v NullableGetAzurednstaskgroupResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAzurednstaskgroupResponseObjectAsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
