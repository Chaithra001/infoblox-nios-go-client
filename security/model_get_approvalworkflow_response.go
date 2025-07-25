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

// GetApprovalworkflowResponse - struct for GetApprovalworkflowResponse
type GetApprovalworkflowResponse struct {
	Approvalworkflow                          *Approvalworkflow
	GetApprovalworkflowResponseObjectAsResult *GetApprovalworkflowResponseObjectAsResult
}

// ApprovalworkflowAsGetApprovalworkflowResponse is a convenience function that returns Approvalworkflow wrapped in GetApprovalworkflowResponse
func ApprovalworkflowAsGetApprovalworkflowResponse(v *Approvalworkflow) GetApprovalworkflowResponse {
	return GetApprovalworkflowResponse{
		Approvalworkflow: v,
	}
}

// GetApprovalworkflowResponseObjectAsResultAsGetApprovalworkflowResponse is a convenience function that returns GetApprovalworkflowResponseObjectAsResult wrapped in GetApprovalworkflowResponse
func GetApprovalworkflowResponseObjectAsResultAsGetApprovalworkflowResponse(v *GetApprovalworkflowResponseObjectAsResult) GetApprovalworkflowResponse {
	return GetApprovalworkflowResponse{
		GetApprovalworkflowResponseObjectAsResult: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetApprovalworkflowResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Approvalworkflow
	err = newStrictDecoder(data).Decode(&dst.Approvalworkflow)
	if err == nil {
		jsonApprovalworkflow, _ := json.Marshal(dst.Approvalworkflow)
		if string(jsonApprovalworkflow) == "{}" { // empty struct
			dst.Approvalworkflow = nil
		} else {
			match++
		}
	} else {
		dst.Approvalworkflow = nil
	}

	// try to unmarshal data into GetApprovalworkflowResponseObjectAsResult
	err = newStrictDecoder(data).Decode(&dst.GetApprovalworkflowResponseObjectAsResult)
	if err == nil {
		jsonGetApprovalworkflowResponseObjectAsResult, _ := json.Marshal(dst.GetApprovalworkflowResponseObjectAsResult)
		if string(jsonGetApprovalworkflowResponseObjectAsResult) == "{}" { // empty struct
			dst.GetApprovalworkflowResponseObjectAsResult = nil
		} else {
			match++
		}
	} else {
		dst.GetApprovalworkflowResponseObjectAsResult = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Approvalworkflow = nil
		dst.GetApprovalworkflowResponseObjectAsResult = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetApprovalworkflowResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetApprovalworkflowResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetApprovalworkflowResponse) MarshalJSON() ([]byte, error) {
	if src.Approvalworkflow != nil {
		return json.Marshal(&src.Approvalworkflow)
	}

	if src.GetApprovalworkflowResponseObjectAsResult != nil {
		return json.Marshal(&src.GetApprovalworkflowResponseObjectAsResult)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetApprovalworkflowResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Approvalworkflow != nil {
		return obj.Approvalworkflow
	}

	if obj.GetApprovalworkflowResponseObjectAsResult != nil {
		return obj.GetApprovalworkflowResponseObjectAsResult
	}

	// all schemas are nil
	return nil
}

type NullableGetApprovalworkflowResponse struct {
	value *GetApprovalworkflowResponse
	isSet bool
}

func (v NullableGetApprovalworkflowResponse) Get() *GetApprovalworkflowResponse {
	return v.value
}

func (v *NullableGetApprovalworkflowResponse) Set(val *GetApprovalworkflowResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetApprovalworkflowResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetApprovalworkflowResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetApprovalworkflowResponse(val *GetApprovalworkflowResponse) *NullableGetApprovalworkflowResponse {
	return &NullableGetApprovalworkflowResponse{value: val, isSet: true}
}

func (v NullableGetApprovalworkflowResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetApprovalworkflowResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
