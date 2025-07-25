/*
Infoblox THREATPROTECTION API

OpenAPI specification for Infoblox NIOS WAPI THREATPROTECTION objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package threatprotection

import (
	"encoding/json"
	"fmt"
)

// GetThreatprotectionProfileRuleResponse - struct for GetThreatprotectionProfileRuleResponse
type GetThreatprotectionProfileRuleResponse struct {
	GetThreatprotectionProfileRuleResponseObjectAsResult *GetThreatprotectionProfileRuleResponseObjectAsResult
	ThreatprotectionProfileRule                          *ThreatprotectionProfileRule
}

// GetThreatprotectionProfileRuleResponseObjectAsResultAsGetThreatprotectionProfileRuleResponse is a convenience function that returns GetThreatprotectionProfileRuleResponseObjectAsResult wrapped in GetThreatprotectionProfileRuleResponse
func GetThreatprotectionProfileRuleResponseObjectAsResultAsGetThreatprotectionProfileRuleResponse(v *GetThreatprotectionProfileRuleResponseObjectAsResult) GetThreatprotectionProfileRuleResponse {
	return GetThreatprotectionProfileRuleResponse{
		GetThreatprotectionProfileRuleResponseObjectAsResult: v,
	}
}

// ThreatprotectionProfileRuleAsGetThreatprotectionProfileRuleResponse is a convenience function that returns ThreatprotectionProfileRule wrapped in GetThreatprotectionProfileRuleResponse
func ThreatprotectionProfileRuleAsGetThreatprotectionProfileRuleResponse(v *ThreatprotectionProfileRule) GetThreatprotectionProfileRuleResponse {
	return GetThreatprotectionProfileRuleResponse{
		ThreatprotectionProfileRule: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetThreatprotectionProfileRuleResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GetThreatprotectionProfileRuleResponseObjectAsResult
	err = newStrictDecoder(data).Decode(&dst.GetThreatprotectionProfileRuleResponseObjectAsResult)
	if err == nil {
		jsonGetThreatprotectionProfileRuleResponseObjectAsResult, _ := json.Marshal(dst.GetThreatprotectionProfileRuleResponseObjectAsResult)
		if string(jsonGetThreatprotectionProfileRuleResponseObjectAsResult) == "{}" { // empty struct
			dst.GetThreatprotectionProfileRuleResponseObjectAsResult = nil
		} else {
			match++
		}
	} else {
		dst.GetThreatprotectionProfileRuleResponseObjectAsResult = nil
	}

	// try to unmarshal data into ThreatprotectionProfileRule
	err = newStrictDecoder(data).Decode(&dst.ThreatprotectionProfileRule)
	if err == nil {
		jsonThreatprotectionProfileRule, _ := json.Marshal(dst.ThreatprotectionProfileRule)
		if string(jsonThreatprotectionProfileRule) == "{}" { // empty struct
			dst.ThreatprotectionProfileRule = nil
		} else {
			match++
		}
	} else {
		dst.ThreatprotectionProfileRule = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GetThreatprotectionProfileRuleResponseObjectAsResult = nil
		dst.ThreatprotectionProfileRule = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetThreatprotectionProfileRuleResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetThreatprotectionProfileRuleResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetThreatprotectionProfileRuleResponse) MarshalJSON() ([]byte, error) {
	if src.GetThreatprotectionProfileRuleResponseObjectAsResult != nil {
		return json.Marshal(&src.GetThreatprotectionProfileRuleResponseObjectAsResult)
	}

	if src.ThreatprotectionProfileRule != nil {
		return json.Marshal(&src.ThreatprotectionProfileRule)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetThreatprotectionProfileRuleResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.GetThreatprotectionProfileRuleResponseObjectAsResult != nil {
		return obj.GetThreatprotectionProfileRuleResponseObjectAsResult
	}

	if obj.ThreatprotectionProfileRule != nil {
		return obj.ThreatprotectionProfileRule
	}

	// all schemas are nil
	return nil
}

type NullableGetThreatprotectionProfileRuleResponse struct {
	value *GetThreatprotectionProfileRuleResponse
	isSet bool
}

func (v NullableGetThreatprotectionProfileRuleResponse) Get() *GetThreatprotectionProfileRuleResponse {
	return v.value
}

func (v *NullableGetThreatprotectionProfileRuleResponse) Set(val *GetThreatprotectionProfileRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetThreatprotectionProfileRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetThreatprotectionProfileRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetThreatprotectionProfileRuleResponse(val *GetThreatprotectionProfileRuleResponse) *NullableGetThreatprotectionProfileRuleResponse {
	return &NullableGetThreatprotectionProfileRuleResponse{value: val, isSet: true}
}

func (v NullableGetThreatprotectionProfileRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetThreatprotectionProfileRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
