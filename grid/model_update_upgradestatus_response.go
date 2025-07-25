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

// UpdateUpgradestatusResponse - struct for UpdateUpgradestatusResponse
type UpdateUpgradestatusResponse struct {
	UpdateUpgradestatusResponseAsObject *UpdateUpgradestatusResponseAsObject
	String                              *string
}

// UpdateUpgradestatusResponseAsObjectAsUpdateUpgradestatusResponse is a convenience function that returns UpdateUpgradestatusResponseAsObject wrapped in UpdateUpgradestatusResponse
func UpdateUpgradestatusResponseAsObjectAsUpdateUpgradestatusResponse(v *UpdateUpgradestatusResponseAsObject) UpdateUpgradestatusResponse {
	return UpdateUpgradestatusResponse{
		UpdateUpgradestatusResponseAsObject: v,
	}
}

// stringAsUpdateUpgradestatusResponse is a convenience function that returns string wrapped in UpdateUpgradestatusResponse
func StringAsUpdateUpgradestatusResponse(v *string) UpdateUpgradestatusResponse {
	return UpdateUpgradestatusResponse{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *UpdateUpgradestatusResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into UpdateUpgradestatusResponseAsObject
	err = newStrictDecoder(data).Decode(&dst.UpdateUpgradestatusResponseAsObject)
	if err == nil {
		jsonUpdateUpgradestatusResponseAsObject, _ := json.Marshal(dst.UpdateUpgradestatusResponseAsObject)
		if string(jsonUpdateUpgradestatusResponseAsObject) == "{}" { // empty struct
			dst.UpdateUpgradestatusResponseAsObject = nil
		} else {
			match++
		}
	} else {
		dst.UpdateUpgradestatusResponseAsObject = nil
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
		dst.UpdateUpgradestatusResponseAsObject = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(UpdateUpgradestatusResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(UpdateUpgradestatusResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UpdateUpgradestatusResponse) MarshalJSON() ([]byte, error) {
	if src.UpdateUpgradestatusResponseAsObject != nil {
		return json.Marshal(&src.UpdateUpgradestatusResponseAsObject)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UpdateUpgradestatusResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.UpdateUpgradestatusResponseAsObject != nil {
		return obj.UpdateUpgradestatusResponseAsObject
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableUpdateUpgradestatusResponse struct {
	value *UpdateUpgradestatusResponse
	isSet bool
}

func (v NullableUpdateUpgradestatusResponse) Get() *UpdateUpgradestatusResponse {
	return v.value
}

func (v *NullableUpdateUpgradestatusResponse) Set(val *UpdateUpgradestatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUpgradestatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUpgradestatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUpgradestatusResponse(val *UpdateUpgradestatusResponse) *NullableUpdateUpgradestatusResponse {
	return &NullableUpdateUpgradestatusResponse{value: val, isSet: true}
}

func (v NullableUpdateUpgradestatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUpgradestatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
