/*
Infoblox SMARTFOLDER API

OpenAPI specification for Infoblox NIOS WAPI SMARTFOLDER objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smartfolder

import (
	"encoding/json"
	"fmt"
)

// GetSmartfolderGlobalResponse - struct for GetSmartfolderGlobalResponse
type GetSmartfolderGlobalResponse struct {
	GetSmartfolderGlobalResponseObjectAsResult *GetSmartfolderGlobalResponseObjectAsResult
	SmartfolderGlobal                          *SmartfolderGlobal
}

// GetSmartfolderGlobalResponseObjectAsResultAsGetSmartfolderGlobalResponse is a convenience function that returns GetSmartfolderGlobalResponseObjectAsResult wrapped in GetSmartfolderGlobalResponse
func GetSmartfolderGlobalResponseObjectAsResultAsGetSmartfolderGlobalResponse(v *GetSmartfolderGlobalResponseObjectAsResult) GetSmartfolderGlobalResponse {
	return GetSmartfolderGlobalResponse{
		GetSmartfolderGlobalResponseObjectAsResult: v,
	}
}

// SmartfolderGlobalAsGetSmartfolderGlobalResponse is a convenience function that returns SmartfolderGlobal wrapped in GetSmartfolderGlobalResponse
func SmartfolderGlobalAsGetSmartfolderGlobalResponse(v *SmartfolderGlobal) GetSmartfolderGlobalResponse {
	return GetSmartfolderGlobalResponse{
		SmartfolderGlobal: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetSmartfolderGlobalResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GetSmartfolderGlobalResponseObjectAsResult
	err = newStrictDecoder(data).Decode(&dst.GetSmartfolderGlobalResponseObjectAsResult)
	if err == nil {
		jsonGetSmartfolderGlobalResponseObjectAsResult, _ := json.Marshal(dst.GetSmartfolderGlobalResponseObjectAsResult)
		if string(jsonGetSmartfolderGlobalResponseObjectAsResult) == "{}" { // empty struct
			dst.GetSmartfolderGlobalResponseObjectAsResult = nil
		} else {
			match++
		}
	} else {
		dst.GetSmartfolderGlobalResponseObjectAsResult = nil
	}

	// try to unmarshal data into SmartfolderGlobal
	err = newStrictDecoder(data).Decode(&dst.SmartfolderGlobal)
	if err == nil {
		jsonSmartfolderGlobal, _ := json.Marshal(dst.SmartfolderGlobal)
		if string(jsonSmartfolderGlobal) == "{}" { // empty struct
			dst.SmartfolderGlobal = nil
		} else {
			match++
		}
	} else {
		dst.SmartfolderGlobal = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GetSmartfolderGlobalResponseObjectAsResult = nil
		dst.SmartfolderGlobal = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetSmartfolderGlobalResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetSmartfolderGlobalResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetSmartfolderGlobalResponse) MarshalJSON() ([]byte, error) {
	if src.GetSmartfolderGlobalResponseObjectAsResult != nil {
		return json.Marshal(&src.GetSmartfolderGlobalResponseObjectAsResult)
	}

	if src.SmartfolderGlobal != nil {
		return json.Marshal(&src.SmartfolderGlobal)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetSmartfolderGlobalResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.GetSmartfolderGlobalResponseObjectAsResult != nil {
		return obj.GetSmartfolderGlobalResponseObjectAsResult
	}

	if obj.SmartfolderGlobal != nil {
		return obj.SmartfolderGlobal
	}

	// all schemas are nil
	return nil
}

type NullableGetSmartfolderGlobalResponse struct {
	value *GetSmartfolderGlobalResponse
	isSet bool
}

func (v NullableGetSmartfolderGlobalResponse) Get() *GetSmartfolderGlobalResponse {
	return v.value
}

func (v *NullableGetSmartfolderGlobalResponse) Set(val *GetSmartfolderGlobalResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSmartfolderGlobalResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSmartfolderGlobalResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSmartfolderGlobalResponse(val *GetSmartfolderGlobalResponse) *NullableGetSmartfolderGlobalResponse {
	return &NullableGetSmartfolderGlobalResponse{value: val, isSet: true}
}

func (v NullableGetSmartfolderGlobalResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSmartfolderGlobalResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
