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

// CreateFingerprintResponse - struct for CreateFingerprintResponse
type CreateFingerprintResponse struct {
	CreateFingerprintResponseAsObject *CreateFingerprintResponseAsObject
	String                            *string
}

// CreateFingerprintResponseAsObjectAsCreateFingerprintResponse is a convenience function that returns CreateFingerprintResponseAsObject wrapped in CreateFingerprintResponse
func CreateFingerprintResponseAsObjectAsCreateFingerprintResponse(v *CreateFingerprintResponseAsObject) CreateFingerprintResponse {
	return CreateFingerprintResponse{
		CreateFingerprintResponseAsObject: v,
	}
}

// stringAsCreateFingerprintResponse is a convenience function that returns string wrapped in CreateFingerprintResponse
func StringAsCreateFingerprintResponse(v *string) CreateFingerprintResponse {
	return CreateFingerprintResponse{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateFingerprintResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CreateFingerprintResponseAsObject
	err = newStrictDecoder(data).Decode(&dst.CreateFingerprintResponseAsObject)
	if err == nil {
		jsonCreateFingerprintResponseAsObject, _ := json.Marshal(dst.CreateFingerprintResponseAsObject)
		if string(jsonCreateFingerprintResponseAsObject) == "{}" { // empty struct
			dst.CreateFingerprintResponseAsObject = nil
		} else {
			match++
		}
	} else {
		dst.CreateFingerprintResponseAsObject = nil
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
		dst.CreateFingerprintResponseAsObject = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateFingerprintResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateFingerprintResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateFingerprintResponse) MarshalJSON() ([]byte, error) {
	if src.CreateFingerprintResponseAsObject != nil {
		return json.Marshal(&src.CreateFingerprintResponseAsObject)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateFingerprintResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CreateFingerprintResponseAsObject != nil {
		return obj.CreateFingerprintResponseAsObject
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableCreateFingerprintResponse struct {
	value *CreateFingerprintResponse
	isSet bool
}

func (v NullableCreateFingerprintResponse) Get() *CreateFingerprintResponse {
	return v.value
}

func (v *NullableCreateFingerprintResponse) Set(val *CreateFingerprintResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFingerprintResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFingerprintResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFingerprintResponse(val *CreateFingerprintResponse) *NullableCreateFingerprintResponse {
	return &NullableCreateFingerprintResponse{value: val, isSet: true}
}

func (v NullableCreateFingerprintResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFingerprintResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
