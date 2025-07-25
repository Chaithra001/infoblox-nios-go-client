/*
Infoblox DHCP API

OpenAPI specification for Infoblox NIOS WAPI DHCP objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dhcp

import (
	"encoding/json"
	"fmt"
)

// CreateFilteroptionResponse - struct for CreateFilteroptionResponse
type CreateFilteroptionResponse struct {
	CreateFilteroptionResponseAsObject *CreateFilteroptionResponseAsObject
	String                             *string
}

// CreateFilteroptionResponseAsObjectAsCreateFilteroptionResponse is a convenience function that returns CreateFilteroptionResponseAsObject wrapped in CreateFilteroptionResponse
func CreateFilteroptionResponseAsObjectAsCreateFilteroptionResponse(v *CreateFilteroptionResponseAsObject) CreateFilteroptionResponse {
	return CreateFilteroptionResponse{
		CreateFilteroptionResponseAsObject: v,
	}
}

// stringAsCreateFilteroptionResponse is a convenience function that returns string wrapped in CreateFilteroptionResponse
func StringAsCreateFilteroptionResponse(v *string) CreateFilteroptionResponse {
	return CreateFilteroptionResponse{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateFilteroptionResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CreateFilteroptionResponseAsObject
	err = newStrictDecoder(data).Decode(&dst.CreateFilteroptionResponseAsObject)
	if err == nil {
		jsonCreateFilteroptionResponseAsObject, _ := json.Marshal(dst.CreateFilteroptionResponseAsObject)
		if string(jsonCreateFilteroptionResponseAsObject) == "{}" { // empty struct
			dst.CreateFilteroptionResponseAsObject = nil
		} else {
			match++
		}
	} else {
		dst.CreateFilteroptionResponseAsObject = nil
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
		dst.CreateFilteroptionResponseAsObject = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateFilteroptionResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateFilteroptionResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateFilteroptionResponse) MarshalJSON() ([]byte, error) {
	if src.CreateFilteroptionResponseAsObject != nil {
		return json.Marshal(&src.CreateFilteroptionResponseAsObject)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateFilteroptionResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CreateFilteroptionResponseAsObject != nil {
		return obj.CreateFilteroptionResponseAsObject
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableCreateFilteroptionResponse struct {
	value *CreateFilteroptionResponse
	isSet bool
}

func (v NullableCreateFilteroptionResponse) Get() *CreateFilteroptionResponse {
	return v.value
}

func (v *NullableCreateFilteroptionResponse) Set(val *CreateFilteroptionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFilteroptionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFilteroptionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFilteroptionResponse(val *CreateFilteroptionResponse) *NullableCreateFilteroptionResponse {
	return &NullableCreateFilteroptionResponse{value: val, isSet: true}
}

func (v NullableCreateFilteroptionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFilteroptionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
