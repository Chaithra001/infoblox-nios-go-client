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

// checks if the NetworkDiscoveryBlackoutSetting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkDiscoveryBlackoutSetting{}

// NetworkDiscoveryBlackoutSetting struct for NetworkDiscoveryBlackoutSetting
type NetworkDiscoveryBlackoutSetting struct {
	// Determines whether a blackout is enabled or not.
	EnableBlackout *bool `json:"enable_blackout,omitempty"`
	// The blackout duration in seconds; minimum value is 1 minute.
	BlackoutDuration     *int64                                           `json:"blackout_duration,omitempty"`
	BlackoutSchedule     *NetworkdiscoveryblackoutsettingBlackoutSchedule `json:"blackout_schedule,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkDiscoveryBlackoutSetting NetworkDiscoveryBlackoutSetting

// NewNetworkDiscoveryBlackoutSetting instantiates a new NetworkDiscoveryBlackoutSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkDiscoveryBlackoutSetting() *NetworkDiscoveryBlackoutSetting {
	this := NetworkDiscoveryBlackoutSetting{}
	return &this
}

// NewNetworkDiscoveryBlackoutSettingWithDefaults instantiates a new NetworkDiscoveryBlackoutSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkDiscoveryBlackoutSettingWithDefaults() *NetworkDiscoveryBlackoutSetting {
	this := NetworkDiscoveryBlackoutSetting{}
	return &this
}

// GetEnableBlackout returns the EnableBlackout field value if set, zero value otherwise.
func (o *NetworkDiscoveryBlackoutSetting) GetEnableBlackout() bool {
	if o == nil || IsNil(o.EnableBlackout) {
		var ret bool
		return ret
	}
	return *o.EnableBlackout
}

// GetEnableBlackoutOk returns a tuple with the EnableBlackout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDiscoveryBlackoutSetting) GetEnableBlackoutOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableBlackout) {
		return nil, false
	}
	return o.EnableBlackout, true
}

// HasEnableBlackout returns a boolean if a field has been set.
func (o *NetworkDiscoveryBlackoutSetting) HasEnableBlackout() bool {
	if o != nil && !IsNil(o.EnableBlackout) {
		return true
	}

	return false
}

// SetEnableBlackout gets a reference to the given bool and assigns it to the EnableBlackout field.
func (o *NetworkDiscoveryBlackoutSetting) SetEnableBlackout(v bool) {
	o.EnableBlackout = &v
}

// GetBlackoutDuration returns the BlackoutDuration field value if set, zero value otherwise.
func (o *NetworkDiscoveryBlackoutSetting) GetBlackoutDuration() int64 {
	if o == nil || IsNil(o.BlackoutDuration) {
		var ret int64
		return ret
	}
	return *o.BlackoutDuration
}

// GetBlackoutDurationOk returns a tuple with the BlackoutDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDiscoveryBlackoutSetting) GetBlackoutDurationOk() (*int64, bool) {
	if o == nil || IsNil(o.BlackoutDuration) {
		return nil, false
	}
	return o.BlackoutDuration, true
}

// HasBlackoutDuration returns a boolean if a field has been set.
func (o *NetworkDiscoveryBlackoutSetting) HasBlackoutDuration() bool {
	if o != nil && !IsNil(o.BlackoutDuration) {
		return true
	}

	return false
}

// SetBlackoutDuration gets a reference to the given int64 and assigns it to the BlackoutDuration field.
func (o *NetworkDiscoveryBlackoutSetting) SetBlackoutDuration(v int64) {
	o.BlackoutDuration = &v
}

// GetBlackoutSchedule returns the BlackoutSchedule field value if set, zero value otherwise.
func (o *NetworkDiscoveryBlackoutSetting) GetBlackoutSchedule() NetworkdiscoveryblackoutsettingBlackoutSchedule {
	if o == nil || IsNil(o.BlackoutSchedule) {
		var ret NetworkdiscoveryblackoutsettingBlackoutSchedule
		return ret
	}
	return *o.BlackoutSchedule
}

// GetBlackoutScheduleOk returns a tuple with the BlackoutSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDiscoveryBlackoutSetting) GetBlackoutScheduleOk() (*NetworkdiscoveryblackoutsettingBlackoutSchedule, bool) {
	if o == nil || IsNil(o.BlackoutSchedule) {
		return nil, false
	}
	return o.BlackoutSchedule, true
}

// HasBlackoutSchedule returns a boolean if a field has been set.
func (o *NetworkDiscoveryBlackoutSetting) HasBlackoutSchedule() bool {
	if o != nil && !IsNil(o.BlackoutSchedule) {
		return true
	}

	return false
}

// SetBlackoutSchedule gets a reference to the given NetworkdiscoveryblackoutsettingBlackoutSchedule and assigns it to the BlackoutSchedule field.
func (o *NetworkDiscoveryBlackoutSetting) SetBlackoutSchedule(v NetworkdiscoveryblackoutsettingBlackoutSchedule) {
	o.BlackoutSchedule = &v
}

func (o NetworkDiscoveryBlackoutSetting) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkDiscoveryBlackoutSetting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnableBlackout) {
		toSerialize["enable_blackout"] = o.EnableBlackout
	}
	if !IsNil(o.BlackoutDuration) {
		toSerialize["blackout_duration"] = o.BlackoutDuration
	}
	if !IsNil(o.BlackoutSchedule) {
		toSerialize["blackout_schedule"] = o.BlackoutSchedule
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NetworkDiscoveryBlackoutSetting) UnmarshalJSON(data []byte) (err error) {
	varNetworkDiscoveryBlackoutSetting := _NetworkDiscoveryBlackoutSetting{}

	err = json.Unmarshal(data, &varNetworkDiscoveryBlackoutSetting)

	if err != nil {
		return err
	}

	*o = NetworkDiscoveryBlackoutSetting(varNetworkDiscoveryBlackoutSetting)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enable_blackout")
		delete(additionalProperties, "blackout_duration")
		delete(additionalProperties, "blackout_schedule")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetworkDiscoveryBlackoutSetting struct {
	value *NetworkDiscoveryBlackoutSetting
	isSet bool
}

func (v NullableNetworkDiscoveryBlackoutSetting) Get() *NetworkDiscoveryBlackoutSetting {
	return v.value
}

func (v *NullableNetworkDiscoveryBlackoutSetting) Set(val *NetworkDiscoveryBlackoutSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkDiscoveryBlackoutSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkDiscoveryBlackoutSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkDiscoveryBlackoutSetting(val *NetworkDiscoveryBlackoutSetting) *NullableNetworkDiscoveryBlackoutSetting {
	return &NullableNetworkDiscoveryBlackoutSetting{value: val, isSet: true}
}

func (v NullableNetworkDiscoveryBlackoutSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkDiscoveryBlackoutSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
