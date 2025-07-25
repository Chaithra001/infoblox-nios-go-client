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

// GetThreatprotectionRuletemplateResponse - struct for GetThreatprotectionRuletemplateResponse
type GetThreatprotectionRuletemplateResponse struct {
	GetThreatprotectionRuletemplateResponseObjectAsResult *GetThreatprotectionRuletemplateResponseObjectAsResult
	ThreatprotectionRuletemplate                          *ThreatprotectionRuletemplate
}

// GetThreatprotectionRuletemplateResponseObjectAsResultAsGetThreatprotectionRuletemplateResponse is a convenience function that returns GetThreatprotectionRuletemplateResponseObjectAsResult wrapped in GetThreatprotectionRuletemplateResponse
func GetThreatprotectionRuletemplateResponseObjectAsResultAsGetThreatprotectionRuletemplateResponse(v *GetThreatprotectionRuletemplateResponseObjectAsResult) GetThreatprotectionRuletemplateResponse {
	return GetThreatprotectionRuletemplateResponse{
		GetThreatprotectionRuletemplateResponseObjectAsResult: v,
	}
}

// ThreatprotectionRuletemplateAsGetThreatprotectionRuletemplateResponse is a convenience function that returns ThreatprotectionRuletemplate wrapped in GetThreatprotectionRuletemplateResponse
func ThreatprotectionRuletemplateAsGetThreatprotectionRuletemplateResponse(v *ThreatprotectionRuletemplate) GetThreatprotectionRuletemplateResponse {
	return GetThreatprotectionRuletemplateResponse{
		ThreatprotectionRuletemplate: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetThreatprotectionRuletemplateResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GetThreatprotectionRuletemplateResponseObjectAsResult
	err = newStrictDecoder(data).Decode(&dst.GetThreatprotectionRuletemplateResponseObjectAsResult)
	if err == nil {
		jsonGetThreatprotectionRuletemplateResponseObjectAsResult, _ := json.Marshal(dst.GetThreatprotectionRuletemplateResponseObjectAsResult)
		if string(jsonGetThreatprotectionRuletemplateResponseObjectAsResult) == "{}" { // empty struct
			dst.GetThreatprotectionRuletemplateResponseObjectAsResult = nil
		} else {
			match++
		}
	} else {
		dst.GetThreatprotectionRuletemplateResponseObjectAsResult = nil
	}

	// try to unmarshal data into ThreatprotectionRuletemplate
	err = newStrictDecoder(data).Decode(&dst.ThreatprotectionRuletemplate)
	if err == nil {
		jsonThreatprotectionRuletemplate, _ := json.Marshal(dst.ThreatprotectionRuletemplate)
		if string(jsonThreatprotectionRuletemplate) == "{}" { // empty struct
			dst.ThreatprotectionRuletemplate = nil
		} else {
			match++
		}
	} else {
		dst.ThreatprotectionRuletemplate = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GetThreatprotectionRuletemplateResponseObjectAsResult = nil
		dst.ThreatprotectionRuletemplate = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetThreatprotectionRuletemplateResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetThreatprotectionRuletemplateResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetThreatprotectionRuletemplateResponse) MarshalJSON() ([]byte, error) {
	if src.GetThreatprotectionRuletemplateResponseObjectAsResult != nil {
		return json.Marshal(&src.GetThreatprotectionRuletemplateResponseObjectAsResult)
	}

	if src.ThreatprotectionRuletemplate != nil {
		return json.Marshal(&src.ThreatprotectionRuletemplate)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetThreatprotectionRuletemplateResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.GetThreatprotectionRuletemplateResponseObjectAsResult != nil {
		return obj.GetThreatprotectionRuletemplateResponseObjectAsResult
	}

	if obj.ThreatprotectionRuletemplate != nil {
		return obj.ThreatprotectionRuletemplate
	}

	// all schemas are nil
	return nil
}

type NullableGetThreatprotectionRuletemplateResponse struct {
	value *GetThreatprotectionRuletemplateResponse
	isSet bool
}

func (v NullableGetThreatprotectionRuletemplateResponse) Get() *GetThreatprotectionRuletemplateResponse {
	return v.value
}

func (v *NullableGetThreatprotectionRuletemplateResponse) Set(val *GetThreatprotectionRuletemplateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetThreatprotectionRuletemplateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetThreatprotectionRuletemplateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetThreatprotectionRuletemplateResponse(val *GetThreatprotectionRuletemplateResponse) *NullableGetThreatprotectionRuletemplateResponse {
	return &NullableGetThreatprotectionRuletemplateResponse{value: val, isSet: true}
}

func (v NullableGetThreatprotectionRuletemplateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetThreatprotectionRuletemplateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
