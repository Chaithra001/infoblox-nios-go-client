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

// ListRadiusAuthserviceResponse - struct for ListRadiusAuthserviceResponse
type ListRadiusAuthserviceResponse struct {
	ListRadiusAuthserviceResponseObject *ListRadiusAuthserviceResponseObject
	ArrayOfRadiusAuthservice            *[]RadiusAuthservice
}

// ListRadiusAuthserviceResponseObjectAsListRadiusAuthserviceResponse is a convenience function that returns ListRadiusAuthserviceResponseObject wrapped in ListRadiusAuthserviceResponse
func ListRadiusAuthserviceResponseObjectAsListRadiusAuthserviceResponse(v *ListRadiusAuthserviceResponseObject) ListRadiusAuthserviceResponse {
	return ListRadiusAuthserviceResponse{
		ListRadiusAuthserviceResponseObject: v,
	}
}

// []RadiusAuthserviceAsListRadiusAuthserviceResponse is a convenience function that returns []RadiusAuthservice wrapped in ListRadiusAuthserviceResponse
func ArrayOfRadiusAuthserviceAsListRadiusAuthserviceResponse(v *[]RadiusAuthservice) ListRadiusAuthserviceResponse {
	return ListRadiusAuthserviceResponse{
		ArrayOfRadiusAuthservice: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListRadiusAuthserviceResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ListRadiusAuthserviceResponseObject
	err = newStrictDecoder(data).Decode(&dst.ListRadiusAuthserviceResponseObject)
	if err == nil {
		jsonListRadiusAuthserviceResponseObject, _ := json.Marshal(dst.ListRadiusAuthserviceResponseObject)
		if string(jsonListRadiusAuthserviceResponseObject) == "{}" { // empty struct
			dst.ListRadiusAuthserviceResponseObject = nil
		} else {
			match++
		}
	} else {
		dst.ListRadiusAuthserviceResponseObject = nil
	}

	// try to unmarshal data into ArrayOfRadiusAuthservice
	err = newStrictDecoder(data).Decode(&dst.ArrayOfRadiusAuthservice)
	if err == nil {
		jsonArrayOfRadiusAuthservice, _ := json.Marshal(dst.ArrayOfRadiusAuthservice)
		if string(jsonArrayOfRadiusAuthservice) == "{}" { // empty struct
			dst.ArrayOfRadiusAuthservice = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfRadiusAuthservice = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ListRadiusAuthserviceResponseObject = nil
		dst.ArrayOfRadiusAuthservice = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ListRadiusAuthserviceResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ListRadiusAuthserviceResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListRadiusAuthserviceResponse) MarshalJSON() ([]byte, error) {
	if src.ListRadiusAuthserviceResponseObject != nil {
		return json.Marshal(&src.ListRadiusAuthserviceResponseObject)
	}

	if src.ArrayOfRadiusAuthservice != nil {
		return json.Marshal(&src.ArrayOfRadiusAuthservice)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListRadiusAuthserviceResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ListRadiusAuthserviceResponseObject != nil {
		return obj.ListRadiusAuthserviceResponseObject
	}

	if obj.ArrayOfRadiusAuthservice != nil {
		return obj.ArrayOfRadiusAuthservice
	}

	// all schemas are nil
	return nil
}

type NullableListRadiusAuthserviceResponse struct {
	value *ListRadiusAuthserviceResponse
	isSet bool
}

func (v NullableListRadiusAuthserviceResponse) Get() *ListRadiusAuthserviceResponse {
	return v.value
}

func (v *NullableListRadiusAuthserviceResponse) Set(val *ListRadiusAuthserviceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListRadiusAuthserviceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListRadiusAuthserviceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListRadiusAuthserviceResponse(val *ListRadiusAuthserviceResponse) *NullableListRadiusAuthserviceResponse {
	return &NullableListRadiusAuthserviceResponse{value: val, isSet: true}
}

func (v NullableListRadiusAuthserviceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListRadiusAuthserviceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
