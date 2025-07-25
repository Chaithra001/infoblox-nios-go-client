/*
Infoblox MISC API

OpenAPI specification for Infoblox NIOS WAPI MISC objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package misc

import (
	"encoding/json"
	"fmt"
)

// ListBfdtemplateResponse - struct for ListBfdtemplateResponse
type ListBfdtemplateResponse struct {
	ListBfdtemplateResponseObject *ListBfdtemplateResponseObject
	ArrayOfBfdtemplate            *[]Bfdtemplate
}

// ListBfdtemplateResponseObjectAsListBfdtemplateResponse is a convenience function that returns ListBfdtemplateResponseObject wrapped in ListBfdtemplateResponse
func ListBfdtemplateResponseObjectAsListBfdtemplateResponse(v *ListBfdtemplateResponseObject) ListBfdtemplateResponse {
	return ListBfdtemplateResponse{
		ListBfdtemplateResponseObject: v,
	}
}

// []BfdtemplateAsListBfdtemplateResponse is a convenience function that returns []Bfdtemplate wrapped in ListBfdtemplateResponse
func ArrayOfBfdtemplateAsListBfdtemplateResponse(v *[]Bfdtemplate) ListBfdtemplateResponse {
	return ListBfdtemplateResponse{
		ArrayOfBfdtemplate: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListBfdtemplateResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ListBfdtemplateResponseObject
	err = newStrictDecoder(data).Decode(&dst.ListBfdtemplateResponseObject)
	if err == nil {
		jsonListBfdtemplateResponseObject, _ := json.Marshal(dst.ListBfdtemplateResponseObject)
		if string(jsonListBfdtemplateResponseObject) == "{}" { // empty struct
			dst.ListBfdtemplateResponseObject = nil
		} else {
			match++
		}
	} else {
		dst.ListBfdtemplateResponseObject = nil
	}

	// try to unmarshal data into ArrayOfBfdtemplate
	err = newStrictDecoder(data).Decode(&dst.ArrayOfBfdtemplate)
	if err == nil {
		jsonArrayOfBfdtemplate, _ := json.Marshal(dst.ArrayOfBfdtemplate)
		if string(jsonArrayOfBfdtemplate) == "{}" { // empty struct
			dst.ArrayOfBfdtemplate = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfBfdtemplate = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ListBfdtemplateResponseObject = nil
		dst.ArrayOfBfdtemplate = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ListBfdtemplateResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ListBfdtemplateResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListBfdtemplateResponse) MarshalJSON() ([]byte, error) {
	if src.ListBfdtemplateResponseObject != nil {
		return json.Marshal(&src.ListBfdtemplateResponseObject)
	}

	if src.ArrayOfBfdtemplate != nil {
		return json.Marshal(&src.ArrayOfBfdtemplate)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListBfdtemplateResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ListBfdtemplateResponseObject != nil {
		return obj.ListBfdtemplateResponseObject
	}

	if obj.ArrayOfBfdtemplate != nil {
		return obj.ArrayOfBfdtemplate
	}

	// all schemas are nil
	return nil
}

type NullableListBfdtemplateResponse struct {
	value *ListBfdtemplateResponse
	isSet bool
}

func (v NullableListBfdtemplateResponse) Get() *ListBfdtemplateResponse {
	return v.value
}

func (v *NullableListBfdtemplateResponse) Set(val *ListBfdtemplateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListBfdtemplateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListBfdtemplateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListBfdtemplateResponse(val *ListBfdtemplateResponse) *NullableListBfdtemplateResponse {
	return &NullableListBfdtemplateResponse{value: val, isSet: true}
}

func (v NullableListBfdtemplateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListBfdtemplateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
