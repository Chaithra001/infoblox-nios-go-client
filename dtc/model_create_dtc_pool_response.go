/*
Infoblox DTC API

OpenAPI specification for Infoblox NIOS WAPI DTC objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dtc

import (
	"encoding/json"
	"fmt"
)

// CreateDtcPoolResponse - struct for CreateDtcPoolResponse
type CreateDtcPoolResponse struct {
	CreateDtcPoolResponseAsObject *CreateDtcPoolResponseAsObject
	String                        *string
}

// CreateDtcPoolResponseAsObjectAsCreateDtcPoolResponse is a convenience function that returns CreateDtcPoolResponseAsObject wrapped in CreateDtcPoolResponse
func CreateDtcPoolResponseAsObjectAsCreateDtcPoolResponse(v *CreateDtcPoolResponseAsObject) CreateDtcPoolResponse {
	return CreateDtcPoolResponse{
		CreateDtcPoolResponseAsObject: v,
	}
}

// stringAsCreateDtcPoolResponse is a convenience function that returns string wrapped in CreateDtcPoolResponse
func StringAsCreateDtcPoolResponse(v *string) CreateDtcPoolResponse {
	return CreateDtcPoolResponse{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateDtcPoolResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CreateDtcPoolResponseAsObject
	err = newStrictDecoder(data).Decode(&dst.CreateDtcPoolResponseAsObject)
	if err == nil {
		jsonCreateDtcPoolResponseAsObject, _ := json.Marshal(dst.CreateDtcPoolResponseAsObject)
		if string(jsonCreateDtcPoolResponseAsObject) == "{}" { // empty struct
			dst.CreateDtcPoolResponseAsObject = nil
		} else {
			match++
		}
	} else {
		dst.CreateDtcPoolResponseAsObject = nil
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
		dst.CreateDtcPoolResponseAsObject = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateDtcPoolResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateDtcPoolResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateDtcPoolResponse) MarshalJSON() ([]byte, error) {
	if src.CreateDtcPoolResponseAsObject != nil {
		return json.Marshal(&src.CreateDtcPoolResponseAsObject)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateDtcPoolResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CreateDtcPoolResponseAsObject != nil {
		return obj.CreateDtcPoolResponseAsObject
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableCreateDtcPoolResponse struct {
	value *CreateDtcPoolResponse
	isSet bool
}

func (v NullableCreateDtcPoolResponse) Get() *CreateDtcPoolResponse {
	return v.value
}

func (v *NullableCreateDtcPoolResponse) Set(val *CreateDtcPoolResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDtcPoolResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDtcPoolResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDtcPoolResponse(val *CreateDtcPoolResponse) *NullableCreateDtcPoolResponse {
	return &NullableCreateDtcPoolResponse{value: val, isSet: true}
}

func (v NullableCreateDtcPoolResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDtcPoolResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
