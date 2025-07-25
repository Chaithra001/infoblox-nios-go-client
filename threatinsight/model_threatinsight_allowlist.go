/*
Infoblox THREATINSIGHT API

OpenAPI specification for Infoblox NIOS WAPI THREATINSIGHT objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package threatinsight

import (
	"encoding/json"
)

// checks if the ThreatinsightAllowlist type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThreatinsightAllowlist{}

// ThreatinsightAllowlist struct for ThreatinsightAllowlist
type ThreatinsightAllowlist struct {
	// The reference to the object.
	Ref *string `json:"_ref,omitempty"`
	// The descriptive comment for the threat insight allowlist.
	Comment *string `json:"comment,omitempty"`
	// Determines whether the threat insight allowlist is disabled.
	Disable *bool `json:"disable,omitempty"`
	// The FQDN of the threat insight allowlist.
	Fqdn *string `json:"fqdn,omitempty"`
	// The type of the threat insight allowlist.
	Type *string `json:"type,omitempty"`
}

// NewThreatinsightAllowlist instantiates a new ThreatinsightAllowlist object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreatinsightAllowlist() *ThreatinsightAllowlist {
	this := ThreatinsightAllowlist{}
	return &this
}

// NewThreatinsightAllowlistWithDefaults instantiates a new ThreatinsightAllowlist object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreatinsightAllowlistWithDefaults() *ThreatinsightAllowlist {
	this := ThreatinsightAllowlist{}
	return &this
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *ThreatinsightAllowlist) GetRef() string {
	if o == nil || IsNil(o.Ref) {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatinsightAllowlist) GetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Ref) {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *ThreatinsightAllowlist) HasRef() bool {
	if o != nil && !IsNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *ThreatinsightAllowlist) SetRef(v string) {
	o.Ref = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ThreatinsightAllowlist) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatinsightAllowlist) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ThreatinsightAllowlist) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ThreatinsightAllowlist) SetComment(v string) {
	o.Comment = &v
}

// GetDisable returns the Disable field value if set, zero value otherwise.
func (o *ThreatinsightAllowlist) GetDisable() bool {
	if o == nil || IsNil(o.Disable) {
		var ret bool
		return ret
	}
	return *o.Disable
}

// GetDisableOk returns a tuple with the Disable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatinsightAllowlist) GetDisableOk() (*bool, bool) {
	if o == nil || IsNil(o.Disable) {
		return nil, false
	}
	return o.Disable, true
}

// HasDisable returns a boolean if a field has been set.
func (o *ThreatinsightAllowlist) HasDisable() bool {
	if o != nil && !IsNil(o.Disable) {
		return true
	}

	return false
}

// SetDisable gets a reference to the given bool and assigns it to the Disable field.
func (o *ThreatinsightAllowlist) SetDisable(v bool) {
	o.Disable = &v
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *ThreatinsightAllowlist) GetFqdn() string {
	if o == nil || IsNil(o.Fqdn) {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatinsightAllowlist) GetFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.Fqdn) {
		return nil, false
	}
	return o.Fqdn, true
}

// HasFqdn returns a boolean if a field has been set.
func (o *ThreatinsightAllowlist) HasFqdn() bool {
	if o != nil && !IsNil(o.Fqdn) {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *ThreatinsightAllowlist) SetFqdn(v string) {
	o.Fqdn = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ThreatinsightAllowlist) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatinsightAllowlist) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ThreatinsightAllowlist) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ThreatinsightAllowlist) SetType(v string) {
	o.Type = &v
}

func (o ThreatinsightAllowlist) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThreatinsightAllowlist) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ref) {
		toSerialize["_ref"] = o.Ref
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.Disable) {
		toSerialize["disable"] = o.Disable
	}
	if !IsNil(o.Fqdn) {
		toSerialize["fqdn"] = o.Fqdn
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableThreatinsightAllowlist struct {
	value *ThreatinsightAllowlist
	isSet bool
}

func (v NullableThreatinsightAllowlist) Get() *ThreatinsightAllowlist {
	return v.value
}

func (v *NullableThreatinsightAllowlist) Set(val *ThreatinsightAllowlist) {
	v.value = val
	v.isSet = true
}

func (v NullableThreatinsightAllowlist) IsSet() bool {
	return v.isSet
}

func (v *NullableThreatinsightAllowlist) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreatinsightAllowlist(val *ThreatinsightAllowlist) *NullableThreatinsightAllowlist {
	return &NullableThreatinsightAllowlist{value: val, isSet: true}
}

func (v NullableThreatinsightAllowlist) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreatinsightAllowlist) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
