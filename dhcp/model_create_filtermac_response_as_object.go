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

// checks if the CreateFiltermacResponseAsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateFiltermacResponseAsObject{}

// CreateFiltermacResponseAsObject The response format to create __Filtermac__ in object format.
type CreateFiltermacResponseAsObject struct {
	Result               *Filtermac `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateFiltermacResponseAsObject CreateFiltermacResponseAsObject

// NewCreateFiltermacResponseAsObject instantiates a new CreateFiltermacResponseAsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFiltermacResponseAsObject() *CreateFiltermacResponseAsObject {
	this := CreateFiltermacResponseAsObject{}
	return &this
}

// NewCreateFiltermacResponseAsObjectWithDefaults instantiates a new CreateFiltermacResponseAsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFiltermacResponseAsObjectWithDefaults() *CreateFiltermacResponseAsObject {
	this := CreateFiltermacResponseAsObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CreateFiltermacResponseAsObject) GetResult() Filtermac {
	if o == nil || IsNil(o.Result) {
		var ret Filtermac
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFiltermacResponseAsObject) GetResultOk() (*Filtermac, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CreateFiltermacResponseAsObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given Filtermac and assigns it to the Result field.
func (o *CreateFiltermacResponseAsObject) SetResult(v Filtermac) {
	o.Result = &v
}

func (o CreateFiltermacResponseAsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateFiltermacResponseAsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateFiltermacResponseAsObject) UnmarshalJSON(data []byte) (err error) {
	varCreateFiltermacResponseAsObject := _CreateFiltermacResponseAsObject{}

	err = json.Unmarshal(data, &varCreateFiltermacResponseAsObject)

	if err != nil {
		return err
	}

	*o = CreateFiltermacResponseAsObject(varCreateFiltermacResponseAsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateFiltermacResponseAsObject struct {
	value *CreateFiltermacResponseAsObject
	isSet bool
}

func (v NullableCreateFiltermacResponseAsObject) Get() *CreateFiltermacResponseAsObject {
	return v.value
}

func (v *NullableCreateFiltermacResponseAsObject) Set(val *CreateFiltermacResponseAsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFiltermacResponseAsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFiltermacResponseAsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFiltermacResponseAsObject(val *CreateFiltermacResponseAsObject) *NullableCreateFiltermacResponseAsObject {
	return &NullableCreateFiltermacResponseAsObject{value: val, isSet: true}
}

func (v NullableCreateFiltermacResponseAsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFiltermacResponseAsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
