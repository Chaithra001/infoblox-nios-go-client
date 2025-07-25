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

// GetNetworkcontainerResponse - struct for GetNetworkcontainerResponse
type GetNetworkcontainerResponse struct {
	GetNetworkcontainerResponseObjectAsResult *GetNetworkcontainerResponseObjectAsResult
	Networkcontainer                          *Networkcontainer
}

// GetNetworkcontainerResponseObjectAsResultAsGetNetworkcontainerResponse is a convenience function that returns GetNetworkcontainerResponseObjectAsResult wrapped in GetNetworkcontainerResponse
func GetNetworkcontainerResponseObjectAsResultAsGetNetworkcontainerResponse(v *GetNetworkcontainerResponseObjectAsResult) GetNetworkcontainerResponse {
	return GetNetworkcontainerResponse{
		GetNetworkcontainerResponseObjectAsResult: v,
	}
}

// NetworkcontainerAsGetNetworkcontainerResponse is a convenience function that returns Networkcontainer wrapped in GetNetworkcontainerResponse
func NetworkcontainerAsGetNetworkcontainerResponse(v *Networkcontainer) GetNetworkcontainerResponse {
	return GetNetworkcontainerResponse{
		Networkcontainer: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetNetworkcontainerResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GetNetworkcontainerResponseObjectAsResult
	err = newStrictDecoder(data).Decode(&dst.GetNetworkcontainerResponseObjectAsResult)
	if err == nil {
		jsonGetNetworkcontainerResponseObjectAsResult, _ := json.Marshal(dst.GetNetworkcontainerResponseObjectAsResult)
		if string(jsonGetNetworkcontainerResponseObjectAsResult) == "{}" { // empty struct
			dst.GetNetworkcontainerResponseObjectAsResult = nil
		} else {
			match++
		}
	} else {
		dst.GetNetworkcontainerResponseObjectAsResult = nil
	}

	// try to unmarshal data into Networkcontainer
	err = newStrictDecoder(data).Decode(&dst.Networkcontainer)
	if err == nil {
		jsonNetworkcontainer, _ := json.Marshal(dst.Networkcontainer)
		if string(jsonNetworkcontainer) == "{}" { // empty struct
			dst.Networkcontainer = nil
		} else {
			match++
		}
	} else {
		dst.Networkcontainer = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GetNetworkcontainerResponseObjectAsResult = nil
		dst.Networkcontainer = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetNetworkcontainerResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetNetworkcontainerResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetNetworkcontainerResponse) MarshalJSON() ([]byte, error) {
	if src.GetNetworkcontainerResponseObjectAsResult != nil {
		return json.Marshal(&src.GetNetworkcontainerResponseObjectAsResult)
	}

	if src.Networkcontainer != nil {
		return json.Marshal(&src.Networkcontainer)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetNetworkcontainerResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.GetNetworkcontainerResponseObjectAsResult != nil {
		return obj.GetNetworkcontainerResponseObjectAsResult
	}

	if obj.Networkcontainer != nil {
		return obj.Networkcontainer
	}

	// all schemas are nil
	return nil
}

type NullableGetNetworkcontainerResponse struct {
	value *GetNetworkcontainerResponse
	isSet bool
}

func (v NullableGetNetworkcontainerResponse) Get() *GetNetworkcontainerResponse {
	return v.value
}

func (v *NullableGetNetworkcontainerResponse) Set(val *GetNetworkcontainerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkcontainerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkcontainerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkcontainerResponse(val *GetNetworkcontainerResponse) *NullableGetNetworkcontainerResponse {
	return &NullableGetNetworkcontainerResponse{value: val, isSet: true}
}

func (v NullableGetNetworkcontainerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkcontainerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
