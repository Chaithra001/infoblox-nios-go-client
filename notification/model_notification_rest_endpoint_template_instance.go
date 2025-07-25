/*
Infoblox NOTIFICATION API

OpenAPI specification for Infoblox NIOS WAPI NOTIFICATION objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package notification

import (
	"encoding/json"
)

// checks if the NotificationRestEndpointTemplateInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationRestEndpointTemplateInstance{}

// NotificationRestEndpointTemplateInstance struct for NotificationRestEndpointTemplateInstance
type NotificationRestEndpointTemplateInstance struct {
	// The name of the REST API template parameter.
	Template *string `json:"template,omitempty"`
	// The notification REST template parameters.
	Parameters           []NotificationrestendpointtemplateinstanceParameters `json:"parameters,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NotificationRestEndpointTemplateInstance NotificationRestEndpointTemplateInstance

// NewNotificationRestEndpointTemplateInstance instantiates a new NotificationRestEndpointTemplateInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationRestEndpointTemplateInstance() *NotificationRestEndpointTemplateInstance {
	this := NotificationRestEndpointTemplateInstance{}
	return &this
}

// NewNotificationRestEndpointTemplateInstanceWithDefaults instantiates a new NotificationRestEndpointTemplateInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationRestEndpointTemplateInstanceWithDefaults() *NotificationRestEndpointTemplateInstance {
	this := NotificationRestEndpointTemplateInstance{}
	return &this
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *NotificationRestEndpointTemplateInstance) GetTemplate() string {
	if o == nil || IsNil(o.Template) {
		var ret string
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationRestEndpointTemplateInstance) GetTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *NotificationRestEndpointTemplateInstance) HasTemplate() bool {
	if o != nil && !IsNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given string and assigns it to the Template field.
func (o *NotificationRestEndpointTemplateInstance) SetTemplate(v string) {
	o.Template = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *NotificationRestEndpointTemplateInstance) GetParameters() []NotificationrestendpointtemplateinstanceParameters {
	if o == nil || IsNil(o.Parameters) {
		var ret []NotificationrestendpointtemplateinstanceParameters
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationRestEndpointTemplateInstance) GetParametersOk() ([]NotificationrestendpointtemplateinstanceParameters, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *NotificationRestEndpointTemplateInstance) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []NotificationrestendpointtemplateinstanceParameters and assigns it to the Parameters field.
func (o *NotificationRestEndpointTemplateInstance) SetParameters(v []NotificationrestendpointtemplateinstanceParameters) {
	o.Parameters = v
}

func (o NotificationRestEndpointTemplateInstance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationRestEndpointTemplateInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Template) {
		toSerialize["template"] = o.Template
	}
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NotificationRestEndpointTemplateInstance) UnmarshalJSON(data []byte) (err error) {
	varNotificationRestEndpointTemplateInstance := _NotificationRestEndpointTemplateInstance{}

	err = json.Unmarshal(data, &varNotificationRestEndpointTemplateInstance)

	if err != nil {
		return err
	}

	*o = NotificationRestEndpointTemplateInstance(varNotificationRestEndpointTemplateInstance)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "template")
		delete(additionalProperties, "parameters")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNotificationRestEndpointTemplateInstance struct {
	value *NotificationRestEndpointTemplateInstance
	isSet bool
}

func (v NullableNotificationRestEndpointTemplateInstance) Get() *NotificationRestEndpointTemplateInstance {
	return v.value
}

func (v *NullableNotificationRestEndpointTemplateInstance) Set(val *NotificationRestEndpointTemplateInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationRestEndpointTemplateInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationRestEndpointTemplateInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationRestEndpointTemplateInstance(val *NotificationRestEndpointTemplateInstance) *NullableNotificationRestEndpointTemplateInstance {
	return &NullableNotificationRestEndpointTemplateInstance{value: val, isSet: true}
}

func (v NullableNotificationRestEndpointTemplateInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationRestEndpointTemplateInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
