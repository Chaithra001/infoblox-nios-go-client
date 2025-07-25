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

// checks if the GridsnmpsettingTrapReceivers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GridsnmpsettingTrapReceivers{}

// GridsnmpsettingTrapReceivers struct for GridsnmpsettingTrapReceivers
type GridsnmpsettingTrapReceivers struct {
	// The address of the trap receiver.
	Address *string `json:"address,omitempty"`
	// The SNMPv3 user for this trap receiver.
	User *string `json:"user,omitempty"`
	// A descriptive comment for this trap receiver.
	Comment              *string `json:"comment,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GridsnmpsettingTrapReceivers GridsnmpsettingTrapReceivers

// NewGridsnmpsettingTrapReceivers instantiates a new GridsnmpsettingTrapReceivers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGridsnmpsettingTrapReceivers() *GridsnmpsettingTrapReceivers {
	this := GridsnmpsettingTrapReceivers{}
	return &this
}

// NewGridsnmpsettingTrapReceiversWithDefaults instantiates a new GridsnmpsettingTrapReceivers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGridsnmpsettingTrapReceiversWithDefaults() *GridsnmpsettingTrapReceivers {
	this := GridsnmpsettingTrapReceivers{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *GridsnmpsettingTrapReceivers) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridsnmpsettingTrapReceivers) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *GridsnmpsettingTrapReceivers) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *GridsnmpsettingTrapReceivers) SetAddress(v string) {
	o.Address = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *GridsnmpsettingTrapReceivers) GetUser() string {
	if o == nil || IsNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridsnmpsettingTrapReceivers) GetUserOk() (*string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *GridsnmpsettingTrapReceivers) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *GridsnmpsettingTrapReceivers) SetUser(v string) {
	o.User = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *GridsnmpsettingTrapReceivers) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridsnmpsettingTrapReceivers) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *GridsnmpsettingTrapReceivers) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *GridsnmpsettingTrapReceivers) SetComment(v string) {
	o.Comment = &v
}

func (o GridsnmpsettingTrapReceivers) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GridsnmpsettingTrapReceivers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GridsnmpsettingTrapReceivers) UnmarshalJSON(data []byte) (err error) {
	varGridsnmpsettingTrapReceivers := _GridsnmpsettingTrapReceivers{}

	err = json.Unmarshal(data, &varGridsnmpsettingTrapReceivers)

	if err != nil {
		return err
	}

	*o = GridsnmpsettingTrapReceivers(varGridsnmpsettingTrapReceivers)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "user")
		delete(additionalProperties, "comment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGridsnmpsettingTrapReceivers struct {
	value *GridsnmpsettingTrapReceivers
	isSet bool
}

func (v NullableGridsnmpsettingTrapReceivers) Get() *GridsnmpsettingTrapReceivers {
	return v.value
}

func (v *NullableGridsnmpsettingTrapReceivers) Set(val *GridsnmpsettingTrapReceivers) {
	v.value = val
	v.isSet = true
}

func (v NullableGridsnmpsettingTrapReceivers) IsSet() bool {
	return v.isSet
}

func (v *NullableGridsnmpsettingTrapReceivers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGridsnmpsettingTrapReceivers(val *GridsnmpsettingTrapReceivers) *NullableGridsnmpsettingTrapReceivers {
	return &NullableGridsnmpsettingTrapReceivers{value: val, isSet: true}
}

func (v NullableGridsnmpsettingTrapReceivers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGridsnmpsettingTrapReceivers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
