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

// CreateRecordSrvResponse - struct for CreateRecordSrvResponse
type CreateRecordSrvResponse struct {
	CreateRecordSrvResponseAsObject *CreateRecordSrvResponseAsObject
	String                          *string
}

// CreateRecordSrvResponseAsObjectAsCreateRecordSrvResponse is a convenience function that returns CreateRecordSrvResponseAsObject wrapped in CreateRecordSrvResponse
func CreateRecordSrvResponseAsObjectAsCreateRecordSrvResponse(v *CreateRecordSrvResponseAsObject) CreateRecordSrvResponse {
	return CreateRecordSrvResponse{
		CreateRecordSrvResponseAsObject: v,
	}
}

// stringAsCreateRecordSrvResponse is a convenience function that returns string wrapped in CreateRecordSrvResponse
func StringAsCreateRecordSrvResponse(v *string) CreateRecordSrvResponse {
	return CreateRecordSrvResponse{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateRecordSrvResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CreateRecordSrvResponseAsObject
	err = newStrictDecoder(data).Decode(&dst.CreateRecordSrvResponseAsObject)
	if err == nil {
		jsonCreateRecordSrvResponseAsObject, _ := json.Marshal(dst.CreateRecordSrvResponseAsObject)
		if string(jsonCreateRecordSrvResponseAsObject) == "{}" { // empty struct
			dst.CreateRecordSrvResponseAsObject = nil
		} else {
			match++
		}
	} else {
		dst.CreateRecordSrvResponseAsObject = nil
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
		dst.CreateRecordSrvResponseAsObject = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateRecordSrvResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateRecordSrvResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateRecordSrvResponse) MarshalJSON() ([]byte, error) {
	if src.CreateRecordSrvResponseAsObject != nil {
		return json.Marshal(&src.CreateRecordSrvResponseAsObject)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateRecordSrvResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CreateRecordSrvResponseAsObject != nil {
		return obj.CreateRecordSrvResponseAsObject
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableCreateRecordSrvResponse struct {
	value *CreateRecordSrvResponse
	isSet bool
}

func (v NullableCreateRecordSrvResponse) Get() *CreateRecordSrvResponse {
	return v.value
}

func (v *NullableCreateRecordSrvResponse) Set(val *CreateRecordSrvResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRecordSrvResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRecordSrvResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRecordSrvResponse(val *CreateRecordSrvResponse) *NullableCreateRecordSrvResponse {
	return &NullableCreateRecordSrvResponse{value: val, isSet: true}
}

func (v NullableCreateRecordSrvResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRecordSrvResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
