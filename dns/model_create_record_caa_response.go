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

// CreateRecordCaaResponse - struct for CreateRecordCaaResponse
type CreateRecordCaaResponse struct {
	CreateRecordCaaResponseAsObject *CreateRecordCaaResponseAsObject
	String                          *string
}

// CreateRecordCaaResponseAsObjectAsCreateRecordCaaResponse is a convenience function that returns CreateRecordCaaResponseAsObject wrapped in CreateRecordCaaResponse
func CreateRecordCaaResponseAsObjectAsCreateRecordCaaResponse(v *CreateRecordCaaResponseAsObject) CreateRecordCaaResponse {
	return CreateRecordCaaResponse{
		CreateRecordCaaResponseAsObject: v,
	}
}

// stringAsCreateRecordCaaResponse is a convenience function that returns string wrapped in CreateRecordCaaResponse
func StringAsCreateRecordCaaResponse(v *string) CreateRecordCaaResponse {
	return CreateRecordCaaResponse{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateRecordCaaResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CreateRecordCaaResponseAsObject
	err = newStrictDecoder(data).Decode(&dst.CreateRecordCaaResponseAsObject)
	if err == nil {
		jsonCreateRecordCaaResponseAsObject, _ := json.Marshal(dst.CreateRecordCaaResponseAsObject)
		if string(jsonCreateRecordCaaResponseAsObject) == "{}" { // empty struct
			dst.CreateRecordCaaResponseAsObject = nil
		} else {
			match++
		}
	} else {
		dst.CreateRecordCaaResponseAsObject = nil
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
		dst.CreateRecordCaaResponseAsObject = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateRecordCaaResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateRecordCaaResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateRecordCaaResponse) MarshalJSON() ([]byte, error) {
	if src.CreateRecordCaaResponseAsObject != nil {
		return json.Marshal(&src.CreateRecordCaaResponseAsObject)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateRecordCaaResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CreateRecordCaaResponseAsObject != nil {
		return obj.CreateRecordCaaResponseAsObject
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableCreateRecordCaaResponse struct {
	value *CreateRecordCaaResponse
	isSet bool
}

func (v NullableCreateRecordCaaResponse) Get() *CreateRecordCaaResponse {
	return v.value
}

func (v *NullableCreateRecordCaaResponse) Set(val *CreateRecordCaaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRecordCaaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRecordCaaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRecordCaaResponse(val *CreateRecordCaaResponse) *NullableCreateRecordCaaResponse {
	return &NullableCreateRecordCaaResponse{value: val, isSet: true}
}

func (v NullableCreateRecordCaaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRecordCaaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
