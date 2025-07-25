/*
Infoblox DNS API

OpenAPI specification for Infoblox NIOS WAPI DNS objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"encoding/json"
	"fmt"
)

// GetSharedrecordAResponse - struct for GetSharedrecordAResponse
type GetSharedrecordAResponse struct {
	GetSharedrecordAResponseObjectAsResult *GetSharedrecordAResponseObjectAsResult
	SharedrecordA                          *SharedrecordA
}

// GetSharedrecordAResponseObjectAsResultAsGetSharedrecordAResponse is a convenience function that returns GetSharedrecordAResponseObjectAsResult wrapped in GetSharedrecordAResponse
func GetSharedrecordAResponseObjectAsResultAsGetSharedrecordAResponse(v *GetSharedrecordAResponseObjectAsResult) GetSharedrecordAResponse {
	return GetSharedrecordAResponse{
		GetSharedrecordAResponseObjectAsResult: v,
	}
}

// SharedrecordAAsGetSharedrecordAResponse is a convenience function that returns SharedrecordA wrapped in GetSharedrecordAResponse
func SharedrecordAAsGetSharedrecordAResponse(v *SharedrecordA) GetSharedrecordAResponse {
	return GetSharedrecordAResponse{
		SharedrecordA: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetSharedrecordAResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GetSharedrecordAResponseObjectAsResult
	err = newStrictDecoder(data).Decode(&dst.GetSharedrecordAResponseObjectAsResult)
	if err == nil {
		jsonGetSharedrecordAResponseObjectAsResult, _ := json.Marshal(dst.GetSharedrecordAResponseObjectAsResult)
		if string(jsonGetSharedrecordAResponseObjectAsResult) == "{}" { // empty struct
			dst.GetSharedrecordAResponseObjectAsResult = nil
		} else {
			match++
		}
	} else {
		dst.GetSharedrecordAResponseObjectAsResult = nil
	}

	// try to unmarshal data into SharedrecordA
	err = newStrictDecoder(data).Decode(&dst.SharedrecordA)
	if err == nil {
		jsonSharedrecordA, _ := json.Marshal(dst.SharedrecordA)
		if string(jsonSharedrecordA) == "{}" { // empty struct
			dst.SharedrecordA = nil
		} else {
			match++
		}
	} else {
		dst.SharedrecordA = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GetSharedrecordAResponseObjectAsResult = nil
		dst.SharedrecordA = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetSharedrecordAResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetSharedrecordAResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetSharedrecordAResponse) MarshalJSON() ([]byte, error) {
	if src.GetSharedrecordAResponseObjectAsResult != nil {
		return json.Marshal(&src.GetSharedrecordAResponseObjectAsResult)
	}

	if src.SharedrecordA != nil {
		return json.Marshal(&src.SharedrecordA)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetSharedrecordAResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.GetSharedrecordAResponseObjectAsResult != nil {
		return obj.GetSharedrecordAResponseObjectAsResult
	}

	if obj.SharedrecordA != nil {
		return obj.SharedrecordA
	}

	// all schemas are nil
	return nil
}

type NullableGetSharedrecordAResponse struct {
	value *GetSharedrecordAResponse
	isSet bool
}

func (v NullableGetSharedrecordAResponse) Get() *GetSharedrecordAResponse {
	return v.value
}

func (v *NullableGetSharedrecordAResponse) Set(val *GetSharedrecordAResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSharedrecordAResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSharedrecordAResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSharedrecordAResponse(val *GetSharedrecordAResponse) *NullableGetSharedrecordAResponse {
	return &NullableGetSharedrecordAResponse{value: val, isSet: true}
}

func (v NullableGetSharedrecordAResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSharedrecordAResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
