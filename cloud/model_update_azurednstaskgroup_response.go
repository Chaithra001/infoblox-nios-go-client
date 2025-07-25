/*
Infoblox CLOUD API

OpenAPI specification for Infoblox NIOS WAPI CLOUD objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloud

import (
	"encoding/json"
	"fmt"
)

// UpdateAzurednstaskgroupResponse - struct for UpdateAzurednstaskgroupResponse
type UpdateAzurednstaskgroupResponse struct {
	UpdateAzurednstaskgroupResponseAsObject *UpdateAzurednstaskgroupResponseAsObject
	String                                  *string
}

// UpdateAzurednstaskgroupResponseAsObjectAsUpdateAzurednstaskgroupResponse is a convenience function that returns UpdateAzurednstaskgroupResponseAsObject wrapped in UpdateAzurednstaskgroupResponse
func UpdateAzurednstaskgroupResponseAsObjectAsUpdateAzurednstaskgroupResponse(v *UpdateAzurednstaskgroupResponseAsObject) UpdateAzurednstaskgroupResponse {
	return UpdateAzurednstaskgroupResponse{
		UpdateAzurednstaskgroupResponseAsObject: v,
	}
}

// stringAsUpdateAzurednstaskgroupResponse is a convenience function that returns string wrapped in UpdateAzurednstaskgroupResponse
func StringAsUpdateAzurednstaskgroupResponse(v *string) UpdateAzurednstaskgroupResponse {
	return UpdateAzurednstaskgroupResponse{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *UpdateAzurednstaskgroupResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into UpdateAzurednstaskgroupResponseAsObject
	err = newStrictDecoder(data).Decode(&dst.UpdateAzurednstaskgroupResponseAsObject)
	if err == nil {
		jsonUpdateAzurednstaskgroupResponseAsObject, _ := json.Marshal(dst.UpdateAzurednstaskgroupResponseAsObject)
		if string(jsonUpdateAzurednstaskgroupResponseAsObject) == "{}" { // empty struct
			dst.UpdateAzurednstaskgroupResponseAsObject = nil
		} else {
			match++
		}
	} else {
		dst.UpdateAzurednstaskgroupResponseAsObject = nil
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
		dst.UpdateAzurednstaskgroupResponseAsObject = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(UpdateAzurednstaskgroupResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(UpdateAzurednstaskgroupResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UpdateAzurednstaskgroupResponse) MarshalJSON() ([]byte, error) {
	if src.UpdateAzurednstaskgroupResponseAsObject != nil {
		return json.Marshal(&src.UpdateAzurednstaskgroupResponseAsObject)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UpdateAzurednstaskgroupResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.UpdateAzurednstaskgroupResponseAsObject != nil {
		return obj.UpdateAzurednstaskgroupResponseAsObject
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableUpdateAzurednstaskgroupResponse struct {
	value *UpdateAzurednstaskgroupResponse
	isSet bool
}

func (v NullableUpdateAzurednstaskgroupResponse) Get() *UpdateAzurednstaskgroupResponse {
	return v.value
}

func (v *NullableUpdateAzurednstaskgroupResponse) Set(val *UpdateAzurednstaskgroupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAzurednstaskgroupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAzurednstaskgroupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAzurednstaskgroupResponse(val *UpdateAzurednstaskgroupResponse) *NullableUpdateAzurednstaskgroupResponse {
	return &NullableUpdateAzurednstaskgroupResponse{value: val, isSet: true}
}

func (v NullableUpdateAzurednstaskgroupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAzurednstaskgroupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
