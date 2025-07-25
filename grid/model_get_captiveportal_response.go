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

// GetCaptiveportalResponse - struct for GetCaptiveportalResponse
type GetCaptiveportalResponse struct {
	Captiveportal                          *Captiveportal
	GetCaptiveportalResponseObjectAsResult *GetCaptiveportalResponseObjectAsResult
}

// CaptiveportalAsGetCaptiveportalResponse is a convenience function that returns Captiveportal wrapped in GetCaptiveportalResponse
func CaptiveportalAsGetCaptiveportalResponse(v *Captiveportal) GetCaptiveportalResponse {
	return GetCaptiveportalResponse{
		Captiveportal: v,
	}
}

// GetCaptiveportalResponseObjectAsResultAsGetCaptiveportalResponse is a convenience function that returns GetCaptiveportalResponseObjectAsResult wrapped in GetCaptiveportalResponse
func GetCaptiveportalResponseObjectAsResultAsGetCaptiveportalResponse(v *GetCaptiveportalResponseObjectAsResult) GetCaptiveportalResponse {
	return GetCaptiveportalResponse{
		GetCaptiveportalResponseObjectAsResult: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetCaptiveportalResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Captiveportal
	err = newStrictDecoder(data).Decode(&dst.Captiveportal)
	if err == nil {
		jsonCaptiveportal, _ := json.Marshal(dst.Captiveportal)
		if string(jsonCaptiveportal) == "{}" { // empty struct
			dst.Captiveportal = nil
		} else {
			match++
		}
	} else {
		dst.Captiveportal = nil
	}

	// try to unmarshal data into GetCaptiveportalResponseObjectAsResult
	err = newStrictDecoder(data).Decode(&dst.GetCaptiveportalResponseObjectAsResult)
	if err == nil {
		jsonGetCaptiveportalResponseObjectAsResult, _ := json.Marshal(dst.GetCaptiveportalResponseObjectAsResult)
		if string(jsonGetCaptiveportalResponseObjectAsResult) == "{}" { // empty struct
			dst.GetCaptiveportalResponseObjectAsResult = nil
		} else {
			match++
		}
	} else {
		dst.GetCaptiveportalResponseObjectAsResult = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Captiveportal = nil
		dst.GetCaptiveportalResponseObjectAsResult = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetCaptiveportalResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetCaptiveportalResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetCaptiveportalResponse) MarshalJSON() ([]byte, error) {
	if src.Captiveportal != nil {
		return json.Marshal(&src.Captiveportal)
	}

	if src.GetCaptiveportalResponseObjectAsResult != nil {
		return json.Marshal(&src.GetCaptiveportalResponseObjectAsResult)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetCaptiveportalResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Captiveportal != nil {
		return obj.Captiveportal
	}

	if obj.GetCaptiveportalResponseObjectAsResult != nil {
		return obj.GetCaptiveportalResponseObjectAsResult
	}

	// all schemas are nil
	return nil
}

type NullableGetCaptiveportalResponse struct {
	value *GetCaptiveportalResponse
	isSet bool
}

func (v NullableGetCaptiveportalResponse) Get() *GetCaptiveportalResponse {
	return v.value
}

func (v *NullableGetCaptiveportalResponse) Set(val *GetCaptiveportalResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCaptiveportalResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCaptiveportalResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCaptiveportalResponse(val *GetCaptiveportalResponse) *NullableGetCaptiveportalResponse {
	return &NullableGetCaptiveportalResponse{value: val, isSet: true}
}

func (v NullableGetCaptiveportalResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCaptiveportalResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
