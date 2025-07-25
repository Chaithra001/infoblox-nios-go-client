/*
Infoblox MICROSOFTSERVER API

OpenAPI specification for Infoblox NIOS WAPI MICROSOFTSERVER objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package microsoftserver

import (
	"encoding/json"
	"fmt"
)

// UpdateMsserverResponse - struct for UpdateMsserverResponse
type UpdateMsserverResponse struct {
	UpdateMsserverResponseAsObject *UpdateMsserverResponseAsObject
	String                         *string
}

// UpdateMsserverResponseAsObjectAsUpdateMsserverResponse is a convenience function that returns UpdateMsserverResponseAsObject wrapped in UpdateMsserverResponse
func UpdateMsserverResponseAsObjectAsUpdateMsserverResponse(v *UpdateMsserverResponseAsObject) UpdateMsserverResponse {
	return UpdateMsserverResponse{
		UpdateMsserverResponseAsObject: v,
	}
}

// stringAsUpdateMsserverResponse is a convenience function that returns string wrapped in UpdateMsserverResponse
func StringAsUpdateMsserverResponse(v *string) UpdateMsserverResponse {
	return UpdateMsserverResponse{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *UpdateMsserverResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into UpdateMsserverResponseAsObject
	err = newStrictDecoder(data).Decode(&dst.UpdateMsserverResponseAsObject)
	if err == nil {
		jsonUpdateMsserverResponseAsObject, _ := json.Marshal(dst.UpdateMsserverResponseAsObject)
		if string(jsonUpdateMsserverResponseAsObject) == "{}" { // empty struct
			dst.UpdateMsserverResponseAsObject = nil
		} else {
			match++
		}
	} else {
		dst.UpdateMsserverResponseAsObject = nil
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
		dst.UpdateMsserverResponseAsObject = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(UpdateMsserverResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(UpdateMsserverResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UpdateMsserverResponse) MarshalJSON() ([]byte, error) {
	if src.UpdateMsserverResponseAsObject != nil {
		return json.Marshal(&src.UpdateMsserverResponseAsObject)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UpdateMsserverResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.UpdateMsserverResponseAsObject != nil {
		return obj.UpdateMsserverResponseAsObject
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableUpdateMsserverResponse struct {
	value *UpdateMsserverResponse
	isSet bool
}

func (v NullableUpdateMsserverResponse) Get() *UpdateMsserverResponse {
	return v.value
}

func (v *NullableUpdateMsserverResponse) Set(val *UpdateMsserverResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMsserverResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMsserverResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMsserverResponse(val *UpdateMsserverResponse) *NullableUpdateMsserverResponse {
	return &NullableUpdateMsserverResponse{value: val, isSet: true}
}

func (v NullableUpdateMsserverResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMsserverResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
