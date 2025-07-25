/*
Infoblox DHCP API

OpenAPI specification for Infoblox NIOS WAPI DHCP objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dhcp

import (
	"encoding/json"
)

// checks if the SharednetworkNetworks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SharednetworkNetworks{}

// SharednetworkNetworks struct for SharednetworkNetworks
type SharednetworkNetworks struct {
	// Reference to the Network.
	Ref                  *string `json:"_ref,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SharednetworkNetworks SharednetworkNetworks

// NewSharednetworkNetworks instantiates a new SharednetworkNetworks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSharednetworkNetworks() *SharednetworkNetworks {
	this := SharednetworkNetworks{}
	return &this
}

// NewSharednetworkNetworksWithDefaults instantiates a new SharednetworkNetworks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSharednetworkNetworksWithDefaults() *SharednetworkNetworks {
	this := SharednetworkNetworks{}
	return &this
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *SharednetworkNetworks) GetRef() string {
	if o == nil || IsNil(o.Ref) {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharednetworkNetworks) GetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Ref) {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *SharednetworkNetworks) HasRef() bool {
	if o != nil && !IsNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *SharednetworkNetworks) SetRef(v string) {
	o.Ref = &v
}

func (o SharednetworkNetworks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SharednetworkNetworks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ref) {
		toSerialize["_ref"] = o.Ref
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SharednetworkNetworks) UnmarshalJSON(data []byte) (err error) {
	varSharednetworkNetworks := _SharednetworkNetworks{}

	err = json.Unmarshal(data, &varSharednetworkNetworks)

	if err != nil {
		return err
	}

	*o = SharednetworkNetworks(varSharednetworkNetworks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_ref")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSharednetworkNetworks struct {
	value *SharednetworkNetworks
	isSet bool
}

func (v NullableSharednetworkNetworks) Get() *SharednetworkNetworks {
	return v.value
}

func (v *NullableSharednetworkNetworks) Set(val *SharednetworkNetworks) {
	v.value = val
	v.isSet = true
}

func (v NullableSharednetworkNetworks) IsSet() bool {
	return v.isSet
}

func (v *NullableSharednetworkNetworks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSharednetworkNetworks(val *SharednetworkNetworks) *NullableSharednetworkNetworks {
	return &NullableSharednetworkNetworks{value: val, isSet: true}
}

func (v NullableSharednetworkNetworks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSharednetworkNetworks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
