/*
Infoblox GRID API

OpenAPI specification for Infoblox NIOS WAPI GRID objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package grid

import (
	"encoding/json"
	"fmt"
)

// GetGridFiledistributionResponse - struct for GetGridFiledistributionResponse
type GetGridFiledistributionResponse struct {
	GetGridFiledistributionResponseObjectAsResult *GetGridFiledistributionResponseObjectAsResult
	GridFiledistribution                          *GridFiledistribution
}

// GetGridFiledistributionResponseObjectAsResultAsGetGridFiledistributionResponse is a convenience function that returns GetGridFiledistributionResponseObjectAsResult wrapped in GetGridFiledistributionResponse
func GetGridFiledistributionResponseObjectAsResultAsGetGridFiledistributionResponse(v *GetGridFiledistributionResponseObjectAsResult) GetGridFiledistributionResponse {
	return GetGridFiledistributionResponse{
		GetGridFiledistributionResponseObjectAsResult: v,
	}
}

// GridFiledistributionAsGetGridFiledistributionResponse is a convenience function that returns GridFiledistribution wrapped in GetGridFiledistributionResponse
func GridFiledistributionAsGetGridFiledistributionResponse(v *GridFiledistribution) GetGridFiledistributionResponse {
	return GetGridFiledistributionResponse{
		GridFiledistribution: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetGridFiledistributionResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GetGridFiledistributionResponseObjectAsResult
	err = newStrictDecoder(data).Decode(&dst.GetGridFiledistributionResponseObjectAsResult)
	if err == nil {
		jsonGetGridFiledistributionResponseObjectAsResult, _ := json.Marshal(dst.GetGridFiledistributionResponseObjectAsResult)
		if string(jsonGetGridFiledistributionResponseObjectAsResult) == "{}" { // empty struct
			dst.GetGridFiledistributionResponseObjectAsResult = nil
		} else {
			match++
		}
	} else {
		dst.GetGridFiledistributionResponseObjectAsResult = nil
	}

	// try to unmarshal data into GridFiledistribution
	err = newStrictDecoder(data).Decode(&dst.GridFiledistribution)
	if err == nil {
		jsonGridFiledistribution, _ := json.Marshal(dst.GridFiledistribution)
		if string(jsonGridFiledistribution) == "{}" { // empty struct
			dst.GridFiledistribution = nil
		} else {
			match++
		}
	} else {
		dst.GridFiledistribution = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GetGridFiledistributionResponseObjectAsResult = nil
		dst.GridFiledistribution = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetGridFiledistributionResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetGridFiledistributionResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetGridFiledistributionResponse) MarshalJSON() ([]byte, error) {
	if src.GetGridFiledistributionResponseObjectAsResult != nil {
		return json.Marshal(&src.GetGridFiledistributionResponseObjectAsResult)
	}

	if src.GridFiledistribution != nil {
		return json.Marshal(&src.GridFiledistribution)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetGridFiledistributionResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.GetGridFiledistributionResponseObjectAsResult != nil {
		return obj.GetGridFiledistributionResponseObjectAsResult
	}

	if obj.GridFiledistribution != nil {
		return obj.GridFiledistribution
	}

	// all schemas are nil
	return nil
}

type NullableGetGridFiledistributionResponse struct {
	value *GetGridFiledistributionResponse
	isSet bool
}

func (v NullableGetGridFiledistributionResponse) Get() *GetGridFiledistributionResponse {
	return v.value
}

func (v *NullableGetGridFiledistributionResponse) Set(val *GetGridFiledistributionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetGridFiledistributionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetGridFiledistributionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetGridFiledistributionResponse(val *GetGridFiledistributionResponse) *NullableGetGridFiledistributionResponse {
	return &NullableGetGridFiledistributionResponse{value: val, isSet: true}
}

func (v NullableGetGridFiledistributionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetGridFiledistributionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
