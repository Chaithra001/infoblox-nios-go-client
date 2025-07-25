/*
Infoblox THREATPROTECTION API

OpenAPI specification for Infoblox NIOS WAPI THREATPROTECTION objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package threatprotection

import (
	"encoding/json"
)

// checks if the ThreatprotectionProfileRuleConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThreatprotectionProfileRuleConfig{}

// ThreatprotectionProfileRuleConfig struct for ThreatprotectionProfileRuleConfig
type ThreatprotectionProfileRuleConfig struct {
	// The rule action.
	Action *string `json:"action,omitempty"`
	// The rule log severity.
	LogSeverity *string `json:"log_severity,omitempty"`
	// The threat protection rule parameters.
	Params               []ThreatprotectionprofileruleconfigParams `json:"params,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ThreatprotectionProfileRuleConfig ThreatprotectionProfileRuleConfig

// NewThreatprotectionProfileRuleConfig instantiates a new ThreatprotectionProfileRuleConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreatprotectionProfileRuleConfig() *ThreatprotectionProfileRuleConfig {
	this := ThreatprotectionProfileRuleConfig{}
	return &this
}

// NewThreatprotectionProfileRuleConfigWithDefaults instantiates a new ThreatprotectionProfileRuleConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreatprotectionProfileRuleConfigWithDefaults() *ThreatprotectionProfileRuleConfig {
	this := ThreatprotectionProfileRuleConfig{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *ThreatprotectionProfileRuleConfig) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatprotectionProfileRuleConfig) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *ThreatprotectionProfileRuleConfig) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *ThreatprotectionProfileRuleConfig) SetAction(v string) {
	o.Action = &v
}

// GetLogSeverity returns the LogSeverity field value if set, zero value otherwise.
func (o *ThreatprotectionProfileRuleConfig) GetLogSeverity() string {
	if o == nil || IsNil(o.LogSeverity) {
		var ret string
		return ret
	}
	return *o.LogSeverity
}

// GetLogSeverityOk returns a tuple with the LogSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatprotectionProfileRuleConfig) GetLogSeverityOk() (*string, bool) {
	if o == nil || IsNil(o.LogSeverity) {
		return nil, false
	}
	return o.LogSeverity, true
}

// HasLogSeverity returns a boolean if a field has been set.
func (o *ThreatprotectionProfileRuleConfig) HasLogSeverity() bool {
	if o != nil && !IsNil(o.LogSeverity) {
		return true
	}

	return false
}

// SetLogSeverity gets a reference to the given string and assigns it to the LogSeverity field.
func (o *ThreatprotectionProfileRuleConfig) SetLogSeverity(v string) {
	o.LogSeverity = &v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *ThreatprotectionProfileRuleConfig) GetParams() []ThreatprotectionprofileruleconfigParams {
	if o == nil || IsNil(o.Params) {
		var ret []ThreatprotectionprofileruleconfigParams
		return ret
	}
	return o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatprotectionProfileRuleConfig) GetParamsOk() ([]ThreatprotectionprofileruleconfigParams, bool) {
	if o == nil || IsNil(o.Params) {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *ThreatprotectionProfileRuleConfig) HasParams() bool {
	if o != nil && !IsNil(o.Params) {
		return true
	}

	return false
}

// SetParams gets a reference to the given []ThreatprotectionprofileruleconfigParams and assigns it to the Params field.
func (o *ThreatprotectionProfileRuleConfig) SetParams(v []ThreatprotectionprofileruleconfigParams) {
	o.Params = v
}

func (o ThreatprotectionProfileRuleConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThreatprotectionProfileRuleConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.LogSeverity) {
		toSerialize["log_severity"] = o.LogSeverity
	}
	if !IsNil(o.Params) {
		toSerialize["params"] = o.Params
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ThreatprotectionProfileRuleConfig) UnmarshalJSON(data []byte) (err error) {
	varThreatprotectionProfileRuleConfig := _ThreatprotectionProfileRuleConfig{}

	err = json.Unmarshal(data, &varThreatprotectionProfileRuleConfig)

	if err != nil {
		return err
	}

	*o = ThreatprotectionProfileRuleConfig(varThreatprotectionProfileRuleConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "log_severity")
		delete(additionalProperties, "params")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableThreatprotectionProfileRuleConfig struct {
	value *ThreatprotectionProfileRuleConfig
	isSet bool
}

func (v NullableThreatprotectionProfileRuleConfig) Get() *ThreatprotectionProfileRuleConfig {
	return v.value
}

func (v *NullableThreatprotectionProfileRuleConfig) Set(val *ThreatprotectionProfileRuleConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableThreatprotectionProfileRuleConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableThreatprotectionProfileRuleConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreatprotectionProfileRuleConfig(val *ThreatprotectionProfileRuleConfig) *NullableThreatprotectionProfileRuleConfig {
	return &NullableThreatprotectionProfileRuleConfig{value: val, isSet: true}
}

func (v NullableThreatprotectionProfileRuleConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreatprotectionProfileRuleConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
