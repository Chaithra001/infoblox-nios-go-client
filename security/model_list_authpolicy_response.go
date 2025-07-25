/*
Infoblox SECURITY API

OpenAPI specification for Infoblox NIOS WAPI SECURITY objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package security

import (
	"encoding/json"
	"fmt"
)

// ListAuthpolicyResponse - struct for ListAuthpolicyResponse
type ListAuthpolicyResponse struct {
	ListAuthpolicyResponseObject *ListAuthpolicyResponseObject
	ArrayOfAuthpolicy            *[]Authpolicy
}

// ListAuthpolicyResponseObjectAsListAuthpolicyResponse is a convenience function that returns ListAuthpolicyResponseObject wrapped in ListAuthpolicyResponse
func ListAuthpolicyResponseObjectAsListAuthpolicyResponse(v *ListAuthpolicyResponseObject) ListAuthpolicyResponse {
	return ListAuthpolicyResponse{
		ListAuthpolicyResponseObject: v,
	}
}

// []AuthpolicyAsListAuthpolicyResponse is a convenience function that returns []Authpolicy wrapped in ListAuthpolicyResponse
func ArrayOfAuthpolicyAsListAuthpolicyResponse(v *[]Authpolicy) ListAuthpolicyResponse {
	return ListAuthpolicyResponse{
		ArrayOfAuthpolicy: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListAuthpolicyResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ListAuthpolicyResponseObject
	err = newStrictDecoder(data).Decode(&dst.ListAuthpolicyResponseObject)
	if err == nil {
		jsonListAuthpolicyResponseObject, _ := json.Marshal(dst.ListAuthpolicyResponseObject)
		if string(jsonListAuthpolicyResponseObject) == "{}" { // empty struct
			dst.ListAuthpolicyResponseObject = nil
		} else {
			match++
		}
	} else {
		dst.ListAuthpolicyResponseObject = nil
	}

	// try to unmarshal data into ArrayOfAuthpolicy
	err = newStrictDecoder(data).Decode(&dst.ArrayOfAuthpolicy)
	if err == nil {
		jsonArrayOfAuthpolicy, _ := json.Marshal(dst.ArrayOfAuthpolicy)
		if string(jsonArrayOfAuthpolicy) == "{}" { // empty struct
			dst.ArrayOfAuthpolicy = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfAuthpolicy = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ListAuthpolicyResponseObject = nil
		dst.ArrayOfAuthpolicy = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ListAuthpolicyResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ListAuthpolicyResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListAuthpolicyResponse) MarshalJSON() ([]byte, error) {
	if src.ListAuthpolicyResponseObject != nil {
		return json.Marshal(&src.ListAuthpolicyResponseObject)
	}

	if src.ArrayOfAuthpolicy != nil {
		return json.Marshal(&src.ArrayOfAuthpolicy)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListAuthpolicyResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ListAuthpolicyResponseObject != nil {
		return obj.ListAuthpolicyResponseObject
	}

	if obj.ArrayOfAuthpolicy != nil {
		return obj.ArrayOfAuthpolicy
	}

	// all schemas are nil
	return nil
}

type NullableListAuthpolicyResponse struct {
	value *ListAuthpolicyResponse
	isSet bool
}

func (v NullableListAuthpolicyResponse) Get() *ListAuthpolicyResponse {
	return v.value
}

func (v *NullableListAuthpolicyResponse) Set(val *ListAuthpolicyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListAuthpolicyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListAuthpolicyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAuthpolicyResponse(val *ListAuthpolicyResponse) *NullableListAuthpolicyResponse {
	return &NullableListAuthpolicyResponse{value: val, isSet: true}
}

func (v NullableListAuthpolicyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAuthpolicyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
