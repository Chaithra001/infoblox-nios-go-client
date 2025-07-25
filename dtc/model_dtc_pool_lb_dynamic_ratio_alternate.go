/*
Infoblox DTC API

OpenAPI specification for Infoblox NIOS WAPI DTC objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dtc

import (
	"encoding/json"
)

// checks if the DtcPoolLbDynamicRatioAlternate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DtcPoolLbDynamicRatioAlternate{}

// DtcPoolLbDynamicRatioAlternate struct for DtcPoolLbDynamicRatioAlternate
type DtcPoolLbDynamicRatioAlternate struct {
	// The method of the DTC dynamic ratio load balancing.
	Method *string `json:"method,omitempty"`
	// The DTC monitor output of which will be used for dynamic ratio load balancing.
	Monitor *string `json:"monitor,omitempty"`
	// The metric of the DTC SNMP monitor that will be used for dynamic weighing.
	MonitorMetric *string `json:"monitor_metric,omitempty"`
	// The DTC monitor weight. 'PRIORITY' means that all clients will be forwarded to the least loaded server. 'RATIO' means that distribution will be calculated based on dynamic weights.
	MonitorWeighing *string `json:"monitor_weighing,omitempty"`
	// Determines whether the inverted values of the DTC SNMP monitor metric will be used.
	InvertMonitorMetric  *bool `json:"invert_monitor_metric,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DtcPoolLbDynamicRatioAlternate DtcPoolLbDynamicRatioAlternate

// NewDtcPoolLbDynamicRatioAlternate instantiates a new DtcPoolLbDynamicRatioAlternate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDtcPoolLbDynamicRatioAlternate() *DtcPoolLbDynamicRatioAlternate {
	this := DtcPoolLbDynamicRatioAlternate{}
	return &this
}

// NewDtcPoolLbDynamicRatioAlternateWithDefaults instantiates a new DtcPoolLbDynamicRatioAlternate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDtcPoolLbDynamicRatioAlternateWithDefaults() *DtcPoolLbDynamicRatioAlternate {
	this := DtcPoolLbDynamicRatioAlternate{}
	return &this
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *DtcPoolLbDynamicRatioAlternate) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtcPoolLbDynamicRatioAlternate) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *DtcPoolLbDynamicRatioAlternate) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *DtcPoolLbDynamicRatioAlternate) SetMethod(v string) {
	o.Method = &v
}

// GetMonitor returns the Monitor field value if set, zero value otherwise.
func (o *DtcPoolLbDynamicRatioAlternate) GetMonitor() string {
	if o == nil || IsNil(o.Monitor) {
		var ret string
		return ret
	}
	return *o.Monitor
}

// GetMonitorOk returns a tuple with the Monitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtcPoolLbDynamicRatioAlternate) GetMonitorOk() (*string, bool) {
	if o == nil || IsNil(o.Monitor) {
		return nil, false
	}
	return o.Monitor, true
}

// HasMonitor returns a boolean if a field has been set.
func (o *DtcPoolLbDynamicRatioAlternate) HasMonitor() bool {
	if o != nil && !IsNil(o.Monitor) {
		return true
	}

	return false
}

// SetMonitor gets a reference to the given string and assigns it to the Monitor field.
func (o *DtcPoolLbDynamicRatioAlternate) SetMonitor(v string) {
	o.Monitor = &v
}

// GetMonitorMetric returns the MonitorMetric field value if set, zero value otherwise.
func (o *DtcPoolLbDynamicRatioAlternate) GetMonitorMetric() string {
	if o == nil || IsNil(o.MonitorMetric) {
		var ret string
		return ret
	}
	return *o.MonitorMetric
}

// GetMonitorMetricOk returns a tuple with the MonitorMetric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtcPoolLbDynamicRatioAlternate) GetMonitorMetricOk() (*string, bool) {
	if o == nil || IsNil(o.MonitorMetric) {
		return nil, false
	}
	return o.MonitorMetric, true
}

// HasMonitorMetric returns a boolean if a field has been set.
func (o *DtcPoolLbDynamicRatioAlternate) HasMonitorMetric() bool {
	if o != nil && !IsNil(o.MonitorMetric) {
		return true
	}

	return false
}

// SetMonitorMetric gets a reference to the given string and assigns it to the MonitorMetric field.
func (o *DtcPoolLbDynamicRatioAlternate) SetMonitorMetric(v string) {
	o.MonitorMetric = &v
}

// GetMonitorWeighing returns the MonitorWeighing field value if set, zero value otherwise.
func (o *DtcPoolLbDynamicRatioAlternate) GetMonitorWeighing() string {
	if o == nil || IsNil(o.MonitorWeighing) {
		var ret string
		return ret
	}
	return *o.MonitorWeighing
}

// GetMonitorWeighingOk returns a tuple with the MonitorWeighing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtcPoolLbDynamicRatioAlternate) GetMonitorWeighingOk() (*string, bool) {
	if o == nil || IsNil(o.MonitorWeighing) {
		return nil, false
	}
	return o.MonitorWeighing, true
}

// HasMonitorWeighing returns a boolean if a field has been set.
func (o *DtcPoolLbDynamicRatioAlternate) HasMonitorWeighing() bool {
	if o != nil && !IsNil(o.MonitorWeighing) {
		return true
	}

	return false
}

// SetMonitorWeighing gets a reference to the given string and assigns it to the MonitorWeighing field.
func (o *DtcPoolLbDynamicRatioAlternate) SetMonitorWeighing(v string) {
	o.MonitorWeighing = &v
}

// GetInvertMonitorMetric returns the InvertMonitorMetric field value if set, zero value otherwise.
func (o *DtcPoolLbDynamicRatioAlternate) GetInvertMonitorMetric() bool {
	if o == nil || IsNil(o.InvertMonitorMetric) {
		var ret bool
		return ret
	}
	return *o.InvertMonitorMetric
}

// GetInvertMonitorMetricOk returns a tuple with the InvertMonitorMetric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtcPoolLbDynamicRatioAlternate) GetInvertMonitorMetricOk() (*bool, bool) {
	if o == nil || IsNil(o.InvertMonitorMetric) {
		return nil, false
	}
	return o.InvertMonitorMetric, true
}

// HasInvertMonitorMetric returns a boolean if a field has been set.
func (o *DtcPoolLbDynamicRatioAlternate) HasInvertMonitorMetric() bool {
	if o != nil && !IsNil(o.InvertMonitorMetric) {
		return true
	}

	return false
}

// SetInvertMonitorMetric gets a reference to the given bool and assigns it to the InvertMonitorMetric field.
func (o *DtcPoolLbDynamicRatioAlternate) SetInvertMonitorMetric(v bool) {
	o.InvertMonitorMetric = &v
}

func (o DtcPoolLbDynamicRatioAlternate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DtcPoolLbDynamicRatioAlternate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.Monitor) {
		toSerialize["monitor"] = o.Monitor
	}
	if !IsNil(o.MonitorMetric) {
		toSerialize["monitor_metric"] = o.MonitorMetric
	}
	if !IsNil(o.MonitorWeighing) {
		toSerialize["monitor_weighing"] = o.MonitorWeighing
	}
	if !IsNil(o.InvertMonitorMetric) {
		toSerialize["invert_monitor_metric"] = o.InvertMonitorMetric
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DtcPoolLbDynamicRatioAlternate) UnmarshalJSON(data []byte) (err error) {
	varDtcPoolLbDynamicRatioAlternate := _DtcPoolLbDynamicRatioAlternate{}

	err = json.Unmarshal(data, &varDtcPoolLbDynamicRatioAlternate)

	if err != nil {
		return err
	}

	*o = DtcPoolLbDynamicRatioAlternate(varDtcPoolLbDynamicRatioAlternate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "method")
		delete(additionalProperties, "monitor")
		delete(additionalProperties, "monitor_metric")
		delete(additionalProperties, "monitor_weighing")
		delete(additionalProperties, "invert_monitor_metric")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDtcPoolLbDynamicRatioAlternate struct {
	value *DtcPoolLbDynamicRatioAlternate
	isSet bool
}

func (v NullableDtcPoolLbDynamicRatioAlternate) Get() *DtcPoolLbDynamicRatioAlternate {
	return v.value
}

func (v *NullableDtcPoolLbDynamicRatioAlternate) Set(val *DtcPoolLbDynamicRatioAlternate) {
	v.value = val
	v.isSet = true
}

func (v NullableDtcPoolLbDynamicRatioAlternate) IsSet() bool {
	return v.isSet
}

func (v *NullableDtcPoolLbDynamicRatioAlternate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDtcPoolLbDynamicRatioAlternate(val *DtcPoolLbDynamicRatioAlternate) *NullableDtcPoolLbDynamicRatioAlternate {
	return &NullableDtcPoolLbDynamicRatioAlternate{value: val, isSet: true}
}

func (v NullableDtcPoolLbDynamicRatioAlternate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDtcPoolLbDynamicRatioAlternate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
