/*
Infoblox CLOUD API

OpenAPI specification for Infoblox NIOS WAPI CLOUD objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloud

import (
	"encoding/json"
	"fmt"
)

// UpdateAwsrte53taskgroupResponse - struct for UpdateAwsrte53taskgroupResponse
type UpdateAwsrte53taskgroupResponse struct {
	UpdateAwsrte53taskgroupResponseAsObject *UpdateAwsrte53taskgroupResponseAsObject
	String                                  *string
}

// UpdateAwsrte53taskgroupResponseAsObjectAsUpdateAwsrte53taskgroupResponse is a convenience function that returns UpdateAwsrte53taskgroupResponseAsObject wrapped in UpdateAwsrte53taskgroupResponse
func UpdateAwsrte53taskgroupResponseAsObjectAsUpdateAwsrte53taskgroupResponse(v *UpdateAwsrte53taskgroupResponseAsObject) UpdateAwsrte53taskgroupResponse {
	return UpdateAwsrte53taskgroupResponse{
		UpdateAwsrte53taskgroupResponseAsObject: v,
	}
}

// stringAsUpdateAwsrte53taskgroupResponse is a convenience function that returns string wrapped in UpdateAwsrte53taskgroupResponse
func StringAsUpdateAwsrte53taskgroupResponse(v *string) UpdateAwsrte53taskgroupResponse {
	return UpdateAwsrte53taskgroupResponse{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *UpdateAwsrte53taskgroupResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into UpdateAwsrte53taskgroupResponseAsObject
	err = newStrictDecoder(data).Decode(&dst.UpdateAwsrte53taskgroupResponseAsObject)
	if err == nil {
		jsonUpdateAwsrte53taskgroupResponseAsObject, _ := json.Marshal(dst.UpdateAwsrte53taskgroupResponseAsObject)
		if string(jsonUpdateAwsrte53taskgroupResponseAsObject) == "{}" { // empty struct
			dst.UpdateAwsrte53taskgroupResponseAsObject = nil
		} else {
			match++
		}
	} else {
		dst.UpdateAwsrte53taskgroupResponseAsObject = nil
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
		dst.UpdateAwsrte53taskgroupResponseAsObject = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(UpdateAwsrte53taskgroupResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(UpdateAwsrte53taskgroupResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UpdateAwsrte53taskgroupResponse) MarshalJSON() ([]byte, error) {
	if src.UpdateAwsrte53taskgroupResponseAsObject != nil {
		return json.Marshal(&src.UpdateAwsrte53taskgroupResponseAsObject)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UpdateAwsrte53taskgroupResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.UpdateAwsrte53taskgroupResponseAsObject != nil {
		return obj.UpdateAwsrte53taskgroupResponseAsObject
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableUpdateAwsrte53taskgroupResponse struct {
	value *UpdateAwsrte53taskgroupResponse
	isSet bool
}

func (v NullableUpdateAwsrte53taskgroupResponse) Get() *UpdateAwsrte53taskgroupResponse {
	return v.value
}

func (v *NullableUpdateAwsrte53taskgroupResponse) Set(val *UpdateAwsrte53taskgroupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAwsrte53taskgroupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAwsrte53taskgroupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAwsrte53taskgroupResponse(val *UpdateAwsrte53taskgroupResponse) *NullableUpdateAwsrte53taskgroupResponse {
	return &NullableUpdateAwsrte53taskgroupResponse{value: val, isSet: true}
}

func (v NullableUpdateAwsrte53taskgroupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAwsrte53taskgroupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
