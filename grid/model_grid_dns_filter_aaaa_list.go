/*
Infoblox GRID API

OpenAPI specification for Infoblox NIOS WAPI GRID objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package grid

import (
	"encoding/json"
)

// checks if the GridDnsFilterAaaaList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GridDnsFilterAaaaList{}

// GridDnsFilterAaaaList struct for GridDnsFilterAaaaList
type GridDnsFilterAaaaList struct {
	// The address this rule applies to or \"Any\".
	Address *string `json:"address,omitempty"`
	// The permission to use for this address.
	Permission           *string `json:"permission,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GridDnsFilterAaaaList GridDnsFilterAaaaList

// NewGridDnsFilterAaaaList instantiates a new GridDnsFilterAaaaList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGridDnsFilterAaaaList() *GridDnsFilterAaaaList {
	this := GridDnsFilterAaaaList{}
	return &this
}

// NewGridDnsFilterAaaaListWithDefaults instantiates a new GridDnsFilterAaaaList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGridDnsFilterAaaaListWithDefaults() *GridDnsFilterAaaaList {
	this := GridDnsFilterAaaaList{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *GridDnsFilterAaaaList) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDnsFilterAaaaList) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *GridDnsFilterAaaaList) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *GridDnsFilterAaaaList) SetAddress(v string) {
	o.Address = &v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *GridDnsFilterAaaaList) GetPermission() string {
	if o == nil || IsNil(o.Permission) {
		var ret string
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDnsFilterAaaaList) GetPermissionOk() (*string, bool) {
	if o == nil || IsNil(o.Permission) {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *GridDnsFilterAaaaList) HasPermission() bool {
	if o != nil && !IsNil(o.Permission) {
		return true
	}

	return false
}

// SetPermission gets a reference to the given string and assigns it to the Permission field.
func (o *GridDnsFilterAaaaList) SetPermission(v string) {
	o.Permission = &v
}

func (o GridDnsFilterAaaaList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GridDnsFilterAaaaList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Permission) {
		toSerialize["permission"] = o.Permission
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GridDnsFilterAaaaList) UnmarshalJSON(data []byte) (err error) {
	varGridDnsFilterAaaaList := _GridDnsFilterAaaaList{}

	err = json.Unmarshal(data, &varGridDnsFilterAaaaList)

	if err != nil {
		return err
	}

	*o = GridDnsFilterAaaaList(varGridDnsFilterAaaaList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "permission")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGridDnsFilterAaaaList struct {
	value *GridDnsFilterAaaaList
	isSet bool
}

func (v NullableGridDnsFilterAaaaList) Get() *GridDnsFilterAaaaList {
	return v.value
}

func (v *NullableGridDnsFilterAaaaList) Set(val *GridDnsFilterAaaaList) {
	v.value = val
	v.isSet = true
}

func (v NullableGridDnsFilterAaaaList) IsSet() bool {
	return v.isSet
}

func (v *NullableGridDnsFilterAaaaList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGridDnsFilterAaaaList(val *GridDnsFilterAaaaList) *NullableGridDnsFilterAaaaList {
	return &NullableGridDnsFilterAaaaList{value: val, isSet: true}
}

func (v NullableGridDnsFilterAaaaList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGridDnsFilterAaaaList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
