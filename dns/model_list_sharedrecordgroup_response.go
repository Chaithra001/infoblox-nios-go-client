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

// ListSharedrecordgroupResponse - struct for ListSharedrecordgroupResponse
type ListSharedrecordgroupResponse struct {
	ListSharedrecordgroupResponseObject *ListSharedrecordgroupResponseObject
	ArrayOfSharedrecordgroup            *[]Sharedrecordgroup
}

// ListSharedrecordgroupResponseObjectAsListSharedrecordgroupResponse is a convenience function that returns ListSharedrecordgroupResponseObject wrapped in ListSharedrecordgroupResponse
func ListSharedrecordgroupResponseObjectAsListSharedrecordgroupResponse(v *ListSharedrecordgroupResponseObject) ListSharedrecordgroupResponse {
	return ListSharedrecordgroupResponse{
		ListSharedrecordgroupResponseObject: v,
	}
}

// []SharedrecordgroupAsListSharedrecordgroupResponse is a convenience function that returns []Sharedrecordgroup wrapped in ListSharedrecordgroupResponse
func ArrayOfSharedrecordgroupAsListSharedrecordgroupResponse(v *[]Sharedrecordgroup) ListSharedrecordgroupResponse {
	return ListSharedrecordgroupResponse{
		ArrayOfSharedrecordgroup: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListSharedrecordgroupResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ListSharedrecordgroupResponseObject
	err = newStrictDecoder(data).Decode(&dst.ListSharedrecordgroupResponseObject)
	if err == nil {
		jsonListSharedrecordgroupResponseObject, _ := json.Marshal(dst.ListSharedrecordgroupResponseObject)
		if string(jsonListSharedrecordgroupResponseObject) == "{}" { // empty struct
			dst.ListSharedrecordgroupResponseObject = nil
		} else {
			match++
		}
	} else {
		dst.ListSharedrecordgroupResponseObject = nil
	}

	// try to unmarshal data into ArrayOfSharedrecordgroup
	err = newStrictDecoder(data).Decode(&dst.ArrayOfSharedrecordgroup)
	if err == nil {
		jsonArrayOfSharedrecordgroup, _ := json.Marshal(dst.ArrayOfSharedrecordgroup)
		if string(jsonArrayOfSharedrecordgroup) == "{}" { // empty struct
			dst.ArrayOfSharedrecordgroup = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfSharedrecordgroup = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ListSharedrecordgroupResponseObject = nil
		dst.ArrayOfSharedrecordgroup = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ListSharedrecordgroupResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ListSharedrecordgroupResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListSharedrecordgroupResponse) MarshalJSON() ([]byte, error) {
	if src.ListSharedrecordgroupResponseObject != nil {
		return json.Marshal(&src.ListSharedrecordgroupResponseObject)
	}

	if src.ArrayOfSharedrecordgroup != nil {
		return json.Marshal(&src.ArrayOfSharedrecordgroup)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListSharedrecordgroupResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ListSharedrecordgroupResponseObject != nil {
		return obj.ListSharedrecordgroupResponseObject
	}

	if obj.ArrayOfSharedrecordgroup != nil {
		return obj.ArrayOfSharedrecordgroup
	}

	// all schemas are nil
	return nil
}

type NullableListSharedrecordgroupResponse struct {
	value *ListSharedrecordgroupResponse
	isSet bool
}

func (v NullableListSharedrecordgroupResponse) Get() *ListSharedrecordgroupResponse {
	return v.value
}

func (v *NullableListSharedrecordgroupResponse) Set(val *ListSharedrecordgroupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListSharedrecordgroupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListSharedrecordgroupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSharedrecordgroupResponse(val *ListSharedrecordgroupResponse) *NullableListSharedrecordgroupResponse {
	return &NullableListSharedrecordgroupResponse{value: val, isSet: true}
}

func (v NullableListSharedrecordgroupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSharedrecordgroupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
