/*
Infoblox IPAM API

OpenAPI specification for Infoblox NIOS WAPI IPAM objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"encoding/json"
)

// checks if the NetworkFederatedRealms type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkFederatedRealms{}

// NetworkFederatedRealms struct for NetworkFederatedRealms
type NetworkFederatedRealms struct {
	// The federated realm name
	Name *string `json:"name,omitempty"`
	// The federated realm id
	Id                   *string `json:"id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkFederatedRealms NetworkFederatedRealms

// NewNetworkFederatedRealms instantiates a new NetworkFederatedRealms object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkFederatedRealms() *NetworkFederatedRealms {
	this := NetworkFederatedRealms{}
	return &this
}

// NewNetworkFederatedRealmsWithDefaults instantiates a new NetworkFederatedRealms object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkFederatedRealmsWithDefaults() *NetworkFederatedRealms {
	this := NetworkFederatedRealms{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworkFederatedRealms) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkFederatedRealms) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworkFederatedRealms) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworkFederatedRealms) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NetworkFederatedRealms) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkFederatedRealms) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NetworkFederatedRealms) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NetworkFederatedRealms) SetId(v string) {
	o.Id = &v
}

func (o NetworkFederatedRealms) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkFederatedRealms) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NetworkFederatedRealms) UnmarshalJSON(data []byte) (err error) {
	varNetworkFederatedRealms := _NetworkFederatedRealms{}

	err = json.Unmarshal(data, &varNetworkFederatedRealms)

	if err != nil {
		return err
	}

	*o = NetworkFederatedRealms(varNetworkFederatedRealms)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetworkFederatedRealms struct {
	value *NetworkFederatedRealms
	isSet bool
}

func (v NullableNetworkFederatedRealms) Get() *NetworkFederatedRealms {
	return v.value
}

func (v *NullableNetworkFederatedRealms) Set(val *NetworkFederatedRealms) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkFederatedRealms) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkFederatedRealms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkFederatedRealms(val *NetworkFederatedRealms) *NullableNetworkFederatedRealms {
	return &NullableNetworkFederatedRealms{value: val, isSet: true}
}

func (v NullableNetworkFederatedRealms) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkFederatedRealms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
