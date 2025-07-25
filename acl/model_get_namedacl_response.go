/*
Infoblox ACL API

OpenAPI specification for Infoblox NIOS WAPI ACL objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package acl

import (
	"encoding/json"
	"fmt"
)

// GetNamedaclResponse - struct for GetNamedaclResponse
type GetNamedaclResponse struct {
	GetNamedaclResponseObjectAsResult *GetNamedaclResponseObjectAsResult
	Namedacl                          *Namedacl
}

// GetNamedaclResponseObjectAsResultAsGetNamedaclResponse is a convenience function that returns GetNamedaclResponseObjectAsResult wrapped in GetNamedaclResponse
func GetNamedaclResponseObjectAsResultAsGetNamedaclResponse(v *GetNamedaclResponseObjectAsResult) GetNamedaclResponse {
	return GetNamedaclResponse{
		GetNamedaclResponseObjectAsResult: v,
	}
}

// NamedaclAsGetNamedaclResponse is a convenience function that returns Namedacl wrapped in GetNamedaclResponse
func NamedaclAsGetNamedaclResponse(v *Namedacl) GetNamedaclResponse {
	return GetNamedaclResponse{
		Namedacl: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetNamedaclResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GetNamedaclResponseObjectAsResult
	err = newStrictDecoder(data).Decode(&dst.GetNamedaclResponseObjectAsResult)
	if err == nil {
		jsonGetNamedaclResponseObjectAsResult, _ := json.Marshal(dst.GetNamedaclResponseObjectAsResult)
		if string(jsonGetNamedaclResponseObjectAsResult) == "{}" { // empty struct
			dst.GetNamedaclResponseObjectAsResult = nil
		} else {
			match++
		}
	} else {
		dst.GetNamedaclResponseObjectAsResult = nil
	}

	// try to unmarshal data into Namedacl
	err = newStrictDecoder(data).Decode(&dst.Namedacl)
	if err == nil {
		jsonNamedacl, _ := json.Marshal(dst.Namedacl)
		if string(jsonNamedacl) == "{}" { // empty struct
			dst.Namedacl = nil
		} else {
			match++
		}
	} else {
		dst.Namedacl = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GetNamedaclResponseObjectAsResult = nil
		dst.Namedacl = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetNamedaclResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetNamedaclResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetNamedaclResponse) MarshalJSON() ([]byte, error) {
	if src.GetNamedaclResponseObjectAsResult != nil {
		return json.Marshal(&src.GetNamedaclResponseObjectAsResult)
	}

	if src.Namedacl != nil {
		return json.Marshal(&src.Namedacl)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetNamedaclResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.GetNamedaclResponseObjectAsResult != nil {
		return obj.GetNamedaclResponseObjectAsResult
	}

	if obj.Namedacl != nil {
		return obj.Namedacl
	}

	// all schemas are nil
	return nil
}

type NullableGetNamedaclResponse struct {
	value *GetNamedaclResponse
	isSet bool
}

func (v NullableGetNamedaclResponse) Get() *GetNamedaclResponse {
	return v.value
}

func (v *NullableGetNamedaclResponse) Set(val *GetNamedaclResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNamedaclResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNamedaclResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNamedaclResponse(val *GetNamedaclResponse) *NullableGetNamedaclResponse {
	return &NullableGetNamedaclResponse{value: val, isSet: true}
}

func (v NullableGetNamedaclResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNamedaclResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
