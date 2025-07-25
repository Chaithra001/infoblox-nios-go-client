/*
Infoblox DISCOVERY API

OpenAPI specification for Infoblox NIOS WAPI DISCOVERY objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package discovery

import (
	"encoding/json"
	"fmt"
)

// UpdateDiscoveryDiagnostictaskResponse - struct for UpdateDiscoveryDiagnostictaskResponse
type UpdateDiscoveryDiagnostictaskResponse struct {
	UpdateDiscoveryDiagnostictaskResponseAsObject *UpdateDiscoveryDiagnostictaskResponseAsObject
	String                                        *string
}

// UpdateDiscoveryDiagnostictaskResponseAsObjectAsUpdateDiscoveryDiagnostictaskResponse is a convenience function that returns UpdateDiscoveryDiagnostictaskResponseAsObject wrapped in UpdateDiscoveryDiagnostictaskResponse
func UpdateDiscoveryDiagnostictaskResponseAsObjectAsUpdateDiscoveryDiagnostictaskResponse(v *UpdateDiscoveryDiagnostictaskResponseAsObject) UpdateDiscoveryDiagnostictaskResponse {
	return UpdateDiscoveryDiagnostictaskResponse{
		UpdateDiscoveryDiagnostictaskResponseAsObject: v,
	}
}

// stringAsUpdateDiscoveryDiagnostictaskResponse is a convenience function that returns string wrapped in UpdateDiscoveryDiagnostictaskResponse
func StringAsUpdateDiscoveryDiagnostictaskResponse(v *string) UpdateDiscoveryDiagnostictaskResponse {
	return UpdateDiscoveryDiagnostictaskResponse{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *UpdateDiscoveryDiagnostictaskResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into UpdateDiscoveryDiagnostictaskResponseAsObject
	err = newStrictDecoder(data).Decode(&dst.UpdateDiscoveryDiagnostictaskResponseAsObject)
	if err == nil {
		jsonUpdateDiscoveryDiagnostictaskResponseAsObject, _ := json.Marshal(dst.UpdateDiscoveryDiagnostictaskResponseAsObject)
		if string(jsonUpdateDiscoveryDiagnostictaskResponseAsObject) == "{}" { // empty struct
			dst.UpdateDiscoveryDiagnostictaskResponseAsObject = nil
		} else {
			match++
		}
	} else {
		dst.UpdateDiscoveryDiagnostictaskResponseAsObject = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.UpdateDiscoveryDiagnostictaskResponseAsObject = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(UpdateDiscoveryDiagnostictaskResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(UpdateDiscoveryDiagnostictaskResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UpdateDiscoveryDiagnostictaskResponse) MarshalJSON() ([]byte, error) {
	if src.UpdateDiscoveryDiagnostictaskResponseAsObject != nil {
		return json.Marshal(&src.UpdateDiscoveryDiagnostictaskResponseAsObject)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UpdateDiscoveryDiagnostictaskResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.UpdateDiscoveryDiagnostictaskResponseAsObject != nil {
		return obj.UpdateDiscoveryDiagnostictaskResponseAsObject
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableUpdateDiscoveryDiagnostictaskResponse struct {
	value *UpdateDiscoveryDiagnostictaskResponse
	isSet bool
}

func (v NullableUpdateDiscoveryDiagnostictaskResponse) Get() *UpdateDiscoveryDiagnostictaskResponse {
	return v.value
}

func (v *NullableUpdateDiscoveryDiagnostictaskResponse) Set(val *UpdateDiscoveryDiagnostictaskResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDiscoveryDiagnostictaskResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDiscoveryDiagnostictaskResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDiscoveryDiagnostictaskResponse(val *UpdateDiscoveryDiagnostictaskResponse) *NullableUpdateDiscoveryDiagnostictaskResponse {
	return &NullableUpdateDiscoveryDiagnostictaskResponse{value: val, isSet: true}
}

func (v NullableUpdateDiscoveryDiagnostictaskResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDiscoveryDiagnostictaskResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
