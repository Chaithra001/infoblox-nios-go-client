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

// checks if the SharednetworkLogicFilterRules type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SharednetworkLogicFilterRules{}

// SharednetworkLogicFilterRules struct for SharednetworkLogicFilterRules
type SharednetworkLogicFilterRules struct {
	// The filter name.
	Filter *string `json:"filter,omitempty"`
	// The filter type. Valid values are: * MAC * NAC * Option
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SharednetworkLogicFilterRules SharednetworkLogicFilterRules

// NewSharednetworkLogicFilterRules instantiates a new SharednetworkLogicFilterRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSharednetworkLogicFilterRules() *SharednetworkLogicFilterRules {
	this := SharednetworkLogicFilterRules{}
	return &this
}

// NewSharednetworkLogicFilterRulesWithDefaults instantiates a new SharednetworkLogicFilterRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSharednetworkLogicFilterRulesWithDefaults() *SharednetworkLogicFilterRules {
	this := SharednetworkLogicFilterRules{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *SharednetworkLogicFilterRules) GetFilter() string {
	if o == nil || IsNil(o.Filter) {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharednetworkLogicFilterRules) GetFilterOk() (*string, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *SharednetworkLogicFilterRules) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *SharednetworkLogicFilterRules) SetFilter(v string) {
	o.Filter = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SharednetworkLogicFilterRules) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SharednetworkLogicFilterRules) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SharednetworkLogicFilterRules) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SharednetworkLogicFilterRules) SetType(v string) {
	o.Type = &v
}

func (o SharednetworkLogicFilterRules) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SharednetworkLogicFilterRules) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SharednetworkLogicFilterRules) UnmarshalJSON(data []byte) (err error) {
	varSharednetworkLogicFilterRules := _SharednetworkLogicFilterRules{}

	err = json.Unmarshal(data, &varSharednetworkLogicFilterRules)

	if err != nil {
		return err
	}

	*o = SharednetworkLogicFilterRules(varSharednetworkLogicFilterRules)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filter")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSharednetworkLogicFilterRules struct {
	value *SharednetworkLogicFilterRules
	isSet bool
}

func (v NullableSharednetworkLogicFilterRules) Get() *SharednetworkLogicFilterRules {
	return v.value
}

func (v *NullableSharednetworkLogicFilterRules) Set(val *SharednetworkLogicFilterRules) {
	v.value = val
	v.isSet = true
}

func (v NullableSharednetworkLogicFilterRules) IsSet() bool {
	return v.isSet
}

func (v *NullableSharednetworkLogicFilterRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSharednetworkLogicFilterRules(val *SharednetworkLogicFilterRules) *NullableSharednetworkLogicFilterRules {
	return &NullableSharednetworkLogicFilterRules{value: val, isSet: true}
}

func (v NullableSharednetworkLogicFilterRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSharednetworkLogicFilterRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
