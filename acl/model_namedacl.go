/*
Infoblox ACL API

OpenAPI specification for Infoblox NIOS WAPI ACL objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package acl

import (
	"encoding/json"
)

// checks if the Namedacl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Namedacl{}

// Namedacl struct for Namedacl
type Namedacl struct {
	// The reference to the object.
	Ref *string `json:"_ref,omitempty"`
	// The access control list of IPv4/IPv6 addresses, networks, TSIG-based anonymous access controls, and other named ACLs.
	AccessList []NamedaclAccessList `json:"access_list,omitempty"`
	// Comment for the named ACL; maximum 256 characters.
	Comment *string `json:"comment,omitempty"`
	// The exploded access list for the named ACL. This list displays all the access control entries in a named ACL and its nested named ACLs, if applicable.
	ExplodedAccessList []NamedaclExplodedAccessList `json:"exploded_access_list,omitempty"`
	// Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}.
	ExtAttrs *map[string]ExtAttrs `json:"extattrs,omitempty"`
	// The name of the named ACL.
	Name *string `json:"name,omitempty"`
}

// NewNamedacl instantiates a new Namedacl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNamedacl() *Namedacl {
	this := Namedacl{}
	return &this
}

// NewNamedaclWithDefaults instantiates a new Namedacl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNamedaclWithDefaults() *Namedacl {
	this := Namedacl{}
	return &this
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *Namedacl) GetRef() string {
	if o == nil || IsNil(o.Ref) {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Namedacl) GetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Ref) {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *Namedacl) HasRef() bool {
	if o != nil && !IsNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *Namedacl) SetRef(v string) {
	o.Ref = &v
}

// GetAccessList returns the AccessList field value if set, zero value otherwise.
func (o *Namedacl) GetAccessList() []NamedaclAccessList {
	if o == nil || IsNil(o.AccessList) {
		var ret []NamedaclAccessList
		return ret
	}
	return o.AccessList
}

// GetAccessListOk returns a tuple with the AccessList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Namedacl) GetAccessListOk() ([]NamedaclAccessList, bool) {
	if o == nil || IsNil(o.AccessList) {
		return nil, false
	}
	return o.AccessList, true
}

// HasAccessList returns a boolean if a field has been set.
func (o *Namedacl) HasAccessList() bool {
	if o != nil && !IsNil(o.AccessList) {
		return true
	}

	return false
}

// SetAccessList gets a reference to the given []NamedaclAccessList and assigns it to the AccessList field.
func (o *Namedacl) SetAccessList(v []NamedaclAccessList) {
	o.AccessList = v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *Namedacl) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Namedacl) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *Namedacl) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *Namedacl) SetComment(v string) {
	o.Comment = &v
}

// GetExplodedAccessList returns the ExplodedAccessList field value if set, zero value otherwise.
func (o *Namedacl) GetExplodedAccessList() []NamedaclExplodedAccessList {
	if o == nil || IsNil(o.ExplodedAccessList) {
		var ret []NamedaclExplodedAccessList
		return ret
	}
	return o.ExplodedAccessList
}

// GetExplodedAccessListOk returns a tuple with the ExplodedAccessList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Namedacl) GetExplodedAccessListOk() ([]NamedaclExplodedAccessList, bool) {
	if o == nil || IsNil(o.ExplodedAccessList) {
		return nil, false
	}
	return o.ExplodedAccessList, true
}

// HasExplodedAccessList returns a boolean if a field has been set.
func (o *Namedacl) HasExplodedAccessList() bool {
	if o != nil && !IsNil(o.ExplodedAccessList) {
		return true
	}

	return false
}

// SetExplodedAccessList gets a reference to the given []NamedaclExplodedAccessList and assigns it to the ExplodedAccessList field.
func (o *Namedacl) SetExplodedAccessList(v []NamedaclExplodedAccessList) {
	o.ExplodedAccessList = v
}

// GetExtAttrs returns the ExtAttrs field value if set, zero value otherwise.
func (o *Namedacl) GetExtAttrs() map[string]ExtAttrs {
	if o == nil || IsNil(o.ExtAttrs) {
		var ret map[string]ExtAttrs
		return ret
	}
	return *o.ExtAttrs
}

// GetExtAttrsOk returns a tuple with the ExtAttrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Namedacl) GetExtAttrsOk() (*map[string]ExtAttrs, bool) {
	if o == nil || IsNil(o.ExtAttrs) {
		return nil, false
	}
	return o.ExtAttrs, true
}

// HasExtAttrs returns a boolean if a field has been set.
func (o *Namedacl) HasExtAttrs() bool {
	if o != nil && !IsNil(o.ExtAttrs) {
		return true
	}

	return false
}

// SetExtAttrs gets a reference to the given map[string]ExtAttrs and assigns it to the ExtAttrs field.
func (o *Namedacl) SetExtAttrs(v map[string]ExtAttrs) {
	o.ExtAttrs = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Namedacl) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Namedacl) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Namedacl) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Namedacl) SetName(v string) {
	o.Name = &v
}

func (o Namedacl) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Namedacl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ref) {
		toSerialize["_ref"] = o.Ref
	}
	if !IsNil(o.AccessList) {
		toSerialize["access_list"] = o.AccessList
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.ExplodedAccessList) {
		toSerialize["exploded_access_list"] = o.ExplodedAccessList
	}
	if !IsNil(o.ExtAttrs) {
		toSerialize["extattrs"] = o.ExtAttrs
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableNamedacl struct {
	value *Namedacl
	isSet bool
}

func (v NullableNamedacl) Get() *Namedacl {
	return v.value
}

func (v *NullableNamedacl) Set(val *Namedacl) {
	v.value = val
	v.isSet = true
}

func (v NullableNamedacl) IsSet() bool {
	return v.isSet
}

func (v *NullableNamedacl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNamedacl(val *Namedacl) *NullableNamedacl {
	return &NullableNamedacl{value: val, isSet: true}
}

func (v NullableNamedacl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNamedacl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
