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

// UpdateMemberFiledistributionResponse - struct for UpdateMemberFiledistributionResponse
type UpdateMemberFiledistributionResponse struct {
	UpdateMemberFiledistributionResponseAsObject *UpdateMemberFiledistributionResponseAsObject
	String                                       *string
}

// UpdateMemberFiledistributionResponseAsObjectAsUpdateMemberFiledistributionResponse is a convenience function that returns UpdateMemberFiledistributionResponseAsObject wrapped in UpdateMemberFiledistributionResponse
func UpdateMemberFiledistributionResponseAsObjectAsUpdateMemberFiledistributionResponse(v *UpdateMemberFiledistributionResponseAsObject) UpdateMemberFiledistributionResponse {
	return UpdateMemberFiledistributionResponse{
		UpdateMemberFiledistributionResponseAsObject: v,
	}
}

// stringAsUpdateMemberFiledistributionResponse is a convenience function that returns string wrapped in UpdateMemberFiledistributionResponse
func StringAsUpdateMemberFiledistributionResponse(v *string) UpdateMemberFiledistributionResponse {
	return UpdateMemberFiledistributionResponse{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *UpdateMemberFiledistributionResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into UpdateMemberFiledistributionResponseAsObject
	err = newStrictDecoder(data).Decode(&dst.UpdateMemberFiledistributionResponseAsObject)
	if err == nil {
		jsonUpdateMemberFiledistributionResponseAsObject, _ := json.Marshal(dst.UpdateMemberFiledistributionResponseAsObject)
		if string(jsonUpdateMemberFiledistributionResponseAsObject) == "{}" { // empty struct
			dst.UpdateMemberFiledistributionResponseAsObject = nil
		} else {
			match++
		}
	} else {
		dst.UpdateMemberFiledistributionResponseAsObject = nil
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
		dst.UpdateMemberFiledistributionResponseAsObject = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(UpdateMemberFiledistributionResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(UpdateMemberFiledistributionResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UpdateMemberFiledistributionResponse) MarshalJSON() ([]byte, error) {
	if src.UpdateMemberFiledistributionResponseAsObject != nil {
		return json.Marshal(&src.UpdateMemberFiledistributionResponseAsObject)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UpdateMemberFiledistributionResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.UpdateMemberFiledistributionResponseAsObject != nil {
		return obj.UpdateMemberFiledistributionResponseAsObject
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableUpdateMemberFiledistributionResponse struct {
	value *UpdateMemberFiledistributionResponse
	isSet bool
}

func (v NullableUpdateMemberFiledistributionResponse) Get() *UpdateMemberFiledistributionResponse {
	return v.value
}

func (v *NullableUpdateMemberFiledistributionResponse) Set(val *UpdateMemberFiledistributionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMemberFiledistributionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMemberFiledistributionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMemberFiledistributionResponse(val *UpdateMemberFiledistributionResponse) *NullableUpdateMemberFiledistributionResponse {
	return &NullableUpdateMemberFiledistributionResponse{value: val, isSet: true}
}

func (v NullableUpdateMemberFiledistributionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMemberFiledistributionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
