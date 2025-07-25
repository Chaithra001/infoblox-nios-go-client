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

// checks if the GetDtcMonitorPdpResponseObjectAsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDtcMonitorPdpResponseObjectAsResult{}

// GetDtcMonitorPdpResponseObjectAsResult The response format to retrieve __DtcMonitorPdp__ objects.
type GetDtcMonitorPdpResponseObjectAsResult struct {
	Result *DtcMonitorPdp `json:"result,omitempty"`
}

// NewGetDtcMonitorPdpResponseObjectAsResult instantiates a new GetDtcMonitorPdpResponseObjectAsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDtcMonitorPdpResponseObjectAsResult() *GetDtcMonitorPdpResponseObjectAsResult {
	this := GetDtcMonitorPdpResponseObjectAsResult{}
	return &this
}

// NewGetDtcMonitorPdpResponseObjectAsResultWithDefaults instantiates a new GetDtcMonitorPdpResponseObjectAsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDtcMonitorPdpResponseObjectAsResultWithDefaults() *GetDtcMonitorPdpResponseObjectAsResult {
	this := GetDtcMonitorPdpResponseObjectAsResult{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetDtcMonitorPdpResponseObjectAsResult) GetResult() DtcMonitorPdp {
	if o == nil || IsNil(o.Result) {
		var ret DtcMonitorPdp
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDtcMonitorPdpResponseObjectAsResult) GetResultOk() (*DtcMonitorPdp, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetDtcMonitorPdpResponseObjectAsResult) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given DtcMonitorPdp and assigns it to the Result field.
func (o *GetDtcMonitorPdpResponseObjectAsResult) SetResult(v DtcMonitorPdp) {
	o.Result = &v
}

func (o GetDtcMonitorPdpResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDtcMonitorPdpResponseObjectAsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetDtcMonitorPdpResponseObjectAsResult struct {
	value *GetDtcMonitorPdpResponseObjectAsResult
	isSet bool
}

func (v NullableGetDtcMonitorPdpResponseObjectAsResult) Get() *GetDtcMonitorPdpResponseObjectAsResult {
	return v.value
}

func (v *NullableGetDtcMonitorPdpResponseObjectAsResult) Set(val *GetDtcMonitorPdpResponseObjectAsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDtcMonitorPdpResponseObjectAsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDtcMonitorPdpResponseObjectAsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDtcMonitorPdpResponseObjectAsResult(val *GetDtcMonitorPdpResponseObjectAsResult) *NullableGetDtcMonitorPdpResponseObjectAsResult {
	return &NullableGetDtcMonitorPdpResponseObjectAsResult{value: val, isSet: true}
}

func (v NullableGetDtcMonitorPdpResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDtcMonitorPdpResponseObjectAsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
