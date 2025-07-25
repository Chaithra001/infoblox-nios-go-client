/*
Infoblox IPAM API

OpenAPI specification for Infoblox NIOS WAPI IPAM objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"encoding/json"
	"fmt"
)

// ListVlanviewResponse - struct for ListVlanviewResponse
type ListVlanviewResponse struct {
	ListVlanviewResponseObject *ListVlanviewResponseObject
	ArrayOfVlanview            *[]Vlanview
}

// ListVlanviewResponseObjectAsListVlanviewResponse is a convenience function that returns ListVlanviewResponseObject wrapped in ListVlanviewResponse
func ListVlanviewResponseObjectAsListVlanviewResponse(v *ListVlanviewResponseObject) ListVlanviewResponse {
	return ListVlanviewResponse{
		ListVlanviewResponseObject: v,
	}
}

// []VlanviewAsListVlanviewResponse is a convenience function that returns []Vlanview wrapped in ListVlanviewResponse
func ArrayOfVlanviewAsListVlanviewResponse(v *[]Vlanview) ListVlanviewResponse {
	return ListVlanviewResponse{
		ArrayOfVlanview: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListVlanviewResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ListVlanviewResponseObject
	err = newStrictDecoder(data).Decode(&dst.ListVlanviewResponseObject)
	if err == nil {
		jsonListVlanviewResponseObject, _ := json.Marshal(dst.ListVlanviewResponseObject)
		if string(jsonListVlanviewResponseObject) == "{}" { // empty struct
			dst.ListVlanviewResponseObject = nil
		} else {
			match++
		}
	} else {
		dst.ListVlanviewResponseObject = nil
	}

	// try to unmarshal data into ArrayOfVlanview
	err = newStrictDecoder(data).Decode(&dst.ArrayOfVlanview)
	if err == nil {
		jsonArrayOfVlanview, _ := json.Marshal(dst.ArrayOfVlanview)
		if string(jsonArrayOfVlanview) == "{}" { // empty struct
			dst.ArrayOfVlanview = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfVlanview = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ListVlanviewResponseObject = nil
		dst.ArrayOfVlanview = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ListVlanviewResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ListVlanviewResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListVlanviewResponse) MarshalJSON() ([]byte, error) {
	if src.ListVlanviewResponseObject != nil {
		return json.Marshal(&src.ListVlanviewResponseObject)
	}

	if src.ArrayOfVlanview != nil {
		return json.Marshal(&src.ArrayOfVlanview)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListVlanviewResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ListVlanviewResponseObject != nil {
		return obj.ListVlanviewResponseObject
	}

	if obj.ArrayOfVlanview != nil {
		return obj.ArrayOfVlanview
	}

	// all schemas are nil
	return nil
}

type NullableListVlanviewResponse struct {
	value *ListVlanviewResponse
	isSet bool
}

func (v NullableListVlanviewResponse) Get() *ListVlanviewResponse {
	return v.value
}

func (v *NullableListVlanviewResponse) Set(val *ListVlanviewResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListVlanviewResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListVlanviewResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListVlanviewResponse(val *ListVlanviewResponse) *NullableListVlanviewResponse {
	return &NullableListVlanviewResponse{value: val, isSet: true}
}

func (v NullableListVlanviewResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListVlanviewResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
