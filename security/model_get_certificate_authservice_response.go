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

// GetCertificateAuthserviceResponse - struct for GetCertificateAuthserviceResponse
type GetCertificateAuthserviceResponse struct {
	CertificateAuthservice                          *CertificateAuthservice
	GetCertificateAuthserviceResponseObjectAsResult *GetCertificateAuthserviceResponseObjectAsResult
}

// CertificateAuthserviceAsGetCertificateAuthserviceResponse is a convenience function that returns CertificateAuthservice wrapped in GetCertificateAuthserviceResponse
func CertificateAuthserviceAsGetCertificateAuthserviceResponse(v *CertificateAuthservice) GetCertificateAuthserviceResponse {
	return GetCertificateAuthserviceResponse{
		CertificateAuthservice: v,
	}
}

// GetCertificateAuthserviceResponseObjectAsResultAsGetCertificateAuthserviceResponse is a convenience function that returns GetCertificateAuthserviceResponseObjectAsResult wrapped in GetCertificateAuthserviceResponse
func GetCertificateAuthserviceResponseObjectAsResultAsGetCertificateAuthserviceResponse(v *GetCertificateAuthserviceResponseObjectAsResult) GetCertificateAuthserviceResponse {
	return GetCertificateAuthserviceResponse{
		GetCertificateAuthserviceResponseObjectAsResult: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetCertificateAuthserviceResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CertificateAuthservice
	err = newStrictDecoder(data).Decode(&dst.CertificateAuthservice)
	if err == nil {
		jsonCertificateAuthservice, _ := json.Marshal(dst.CertificateAuthservice)
		if string(jsonCertificateAuthservice) == "{}" { // empty struct
			dst.CertificateAuthservice = nil
		} else {
			match++
		}
	} else {
		dst.CertificateAuthservice = nil
	}

	// try to unmarshal data into GetCertificateAuthserviceResponseObjectAsResult
	err = newStrictDecoder(data).Decode(&dst.GetCertificateAuthserviceResponseObjectAsResult)
	if err == nil {
		jsonGetCertificateAuthserviceResponseObjectAsResult, _ := json.Marshal(dst.GetCertificateAuthserviceResponseObjectAsResult)
		if string(jsonGetCertificateAuthserviceResponseObjectAsResult) == "{}" { // empty struct
			dst.GetCertificateAuthserviceResponseObjectAsResult = nil
		} else {
			match++
		}
	} else {
		dst.GetCertificateAuthserviceResponseObjectAsResult = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CertificateAuthservice = nil
		dst.GetCertificateAuthserviceResponseObjectAsResult = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetCertificateAuthserviceResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetCertificateAuthserviceResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetCertificateAuthserviceResponse) MarshalJSON() ([]byte, error) {
	if src.CertificateAuthservice != nil {
		return json.Marshal(&src.CertificateAuthservice)
	}

	if src.GetCertificateAuthserviceResponseObjectAsResult != nil {
		return json.Marshal(&src.GetCertificateAuthserviceResponseObjectAsResult)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetCertificateAuthserviceResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CertificateAuthservice != nil {
		return obj.CertificateAuthservice
	}

	if obj.GetCertificateAuthserviceResponseObjectAsResult != nil {
		return obj.GetCertificateAuthserviceResponseObjectAsResult
	}

	// all schemas are nil
	return nil
}

type NullableGetCertificateAuthserviceResponse struct {
	value *GetCertificateAuthserviceResponse
	isSet bool
}

func (v NullableGetCertificateAuthserviceResponse) Get() *GetCertificateAuthserviceResponse {
	return v.value
}

func (v *NullableGetCertificateAuthserviceResponse) Set(val *GetCertificateAuthserviceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCertificateAuthserviceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCertificateAuthserviceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCertificateAuthserviceResponse(val *GetCertificateAuthserviceResponse) *NullableGetCertificateAuthserviceResponse {
	return &NullableGetCertificateAuthserviceResponse{value: val, isSet: true}
}

func (v NullableGetCertificateAuthserviceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCertificateAuthserviceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
