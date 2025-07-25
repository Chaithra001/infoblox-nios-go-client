/*
Infoblox DNS API

OpenAPI specification for Infoblox NIOS WAPI DNS objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"encoding/json"
)

// checks if the Dns64group type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dns64group{}

// Dns64group struct for Dns64group
type Dns64group struct {
	// The reference to the object.
	Ref *string `json:"_ref,omitempty"`
	// Access Control settings that contain IPv4 and IPv6 DNS clients and networks to which the DNS server is allowed to send synthesized AAAA records with the specified IPv6 prefix.
	Clients []Dns64groupClients `json:"clients,omitempty"`
	// The descriptive comment for the DNS64 synthesis group object.
	Comment *string `json:"comment,omitempty"`
	// Determines whether the DNS64 synthesis group is disabled.
	Disable *bool `json:"disable,omitempty"`
	// Determines whether the DNS64 synthesis of AAAA records is enabled for DNS64 synthesis groups that request DNSSEC data.
	EnableDnssecDns64 *bool `json:"enable_dnssec_dns64,omitempty"`
	// Access Control settings that contain IPv6 addresses or prefix ranges that cannot be used by IPv6-only hosts, such as IP addresses in the ::ffff:0:0/96 network. When DNS server retrieves an AAAA record that contains an IPv6 address that matches an excluded address, it does not return the AAAA record. Instead it synthesizes an AAAA record from the A record.
	Exclude []Dns64groupExclude `json:"exclude,omitempty"`
	// Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}.
	ExtAttrs *map[string]ExtAttrs `json:"extattrs,omitempty"`
	// Access Control settings that contain IPv4 addresses and networks for which the DNS server can synthesize AAAA records with the specified prefix.
	Mapped []Dns64groupMapped `json:"mapped,omitempty"`
	// The name of the DNS64 synthesis group object.
	Name *string `json:"name,omitempty"`
	// The IPv6 prefix used for the synthesized AAAA records. The prefix length must be /32, /40, /48, /56, /64 or /96, and all bits beyond the specified length must be zero.
	Prefix *string `json:"prefix,omitempty"`
}

// NewDns64group instantiates a new Dns64group object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDns64group() *Dns64group {
	this := Dns64group{}
	return &this
}

// NewDns64groupWithDefaults instantiates a new Dns64group object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDns64groupWithDefaults() *Dns64group {
	this := Dns64group{}
	return &this
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *Dns64group) GetRef() string {
	if o == nil || IsNil(o.Ref) {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dns64group) GetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Ref) {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *Dns64group) HasRef() bool {
	if o != nil && !IsNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *Dns64group) SetRef(v string) {
	o.Ref = &v
}

// GetClients returns the Clients field value if set, zero value otherwise.
func (o *Dns64group) GetClients() []Dns64groupClients {
	if o == nil || IsNil(o.Clients) {
		var ret []Dns64groupClients
		return ret
	}
	return o.Clients
}

// GetClientsOk returns a tuple with the Clients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dns64group) GetClientsOk() ([]Dns64groupClients, bool) {
	if o == nil || IsNil(o.Clients) {
		return nil, false
	}
	return o.Clients, true
}

// HasClients returns a boolean if a field has been set.
func (o *Dns64group) HasClients() bool {
	if o != nil && !IsNil(o.Clients) {
		return true
	}

	return false
}

// SetClients gets a reference to the given []Dns64groupClients and assigns it to the Clients field.
func (o *Dns64group) SetClients(v []Dns64groupClients) {
	o.Clients = v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *Dns64group) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dns64group) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *Dns64group) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *Dns64group) SetComment(v string) {
	o.Comment = &v
}

// GetDisable returns the Disable field value if set, zero value otherwise.
func (o *Dns64group) GetDisable() bool {
	if o == nil || IsNil(o.Disable) {
		var ret bool
		return ret
	}
	return *o.Disable
}

// GetDisableOk returns a tuple with the Disable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dns64group) GetDisableOk() (*bool, bool) {
	if o == nil || IsNil(o.Disable) {
		return nil, false
	}
	return o.Disable, true
}

// HasDisable returns a boolean if a field has been set.
func (o *Dns64group) HasDisable() bool {
	if o != nil && !IsNil(o.Disable) {
		return true
	}

	return false
}

// SetDisable gets a reference to the given bool and assigns it to the Disable field.
func (o *Dns64group) SetDisable(v bool) {
	o.Disable = &v
}

// GetEnableDnssecDns64 returns the EnableDnssecDns64 field value if set, zero value otherwise.
func (o *Dns64group) GetEnableDnssecDns64() bool {
	if o == nil || IsNil(o.EnableDnssecDns64) {
		var ret bool
		return ret
	}
	return *o.EnableDnssecDns64
}

// GetEnableDnssecDns64Ok returns a tuple with the EnableDnssecDns64 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dns64group) GetEnableDnssecDns64Ok() (*bool, bool) {
	if o == nil || IsNil(o.EnableDnssecDns64) {
		return nil, false
	}
	return o.EnableDnssecDns64, true
}

// HasEnableDnssecDns64 returns a boolean if a field has been set.
func (o *Dns64group) HasEnableDnssecDns64() bool {
	if o != nil && !IsNil(o.EnableDnssecDns64) {
		return true
	}

	return false
}

// SetEnableDnssecDns64 gets a reference to the given bool and assigns it to the EnableDnssecDns64 field.
func (o *Dns64group) SetEnableDnssecDns64(v bool) {
	o.EnableDnssecDns64 = &v
}

// GetExclude returns the Exclude field value if set, zero value otherwise.
func (o *Dns64group) GetExclude() []Dns64groupExclude {
	if o == nil || IsNil(o.Exclude) {
		var ret []Dns64groupExclude
		return ret
	}
	return o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dns64group) GetExcludeOk() ([]Dns64groupExclude, bool) {
	if o == nil || IsNil(o.Exclude) {
		return nil, false
	}
	return o.Exclude, true
}

// HasExclude returns a boolean if a field has been set.
func (o *Dns64group) HasExclude() bool {
	if o != nil && !IsNil(o.Exclude) {
		return true
	}

	return false
}

// SetExclude gets a reference to the given []Dns64groupExclude and assigns it to the Exclude field.
func (o *Dns64group) SetExclude(v []Dns64groupExclude) {
	o.Exclude = v
}

// GetExtAttrs returns the ExtAttrs field value if set, zero value otherwise.
func (o *Dns64group) GetExtAttrs() map[string]ExtAttrs {
	if o == nil || IsNil(o.ExtAttrs) {
		var ret map[string]ExtAttrs
		return ret
	}
	return *o.ExtAttrs
}

// GetExtAttrsOk returns a tuple with the ExtAttrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dns64group) GetExtAttrsOk() (*map[string]ExtAttrs, bool) {
	if o == nil || IsNil(o.ExtAttrs) {
		return nil, false
	}
	return o.ExtAttrs, true
}

// HasExtAttrs returns a boolean if a field has been set.
func (o *Dns64group) HasExtAttrs() bool {
	if o != nil && !IsNil(o.ExtAttrs) {
		return true
	}

	return false
}

// SetExtAttrs gets a reference to the given map[string]ExtAttrs and assigns it to the ExtAttrs field.
func (o *Dns64group) SetExtAttrs(v map[string]ExtAttrs) {
	o.ExtAttrs = &v
}

// GetMapped returns the Mapped field value if set, zero value otherwise.
func (o *Dns64group) GetMapped() []Dns64groupMapped {
	if o == nil || IsNil(o.Mapped) {
		var ret []Dns64groupMapped
		return ret
	}
	return o.Mapped
}

// GetMappedOk returns a tuple with the Mapped field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dns64group) GetMappedOk() ([]Dns64groupMapped, bool) {
	if o == nil || IsNil(o.Mapped) {
		return nil, false
	}
	return o.Mapped, true
}

// HasMapped returns a boolean if a field has been set.
func (o *Dns64group) HasMapped() bool {
	if o != nil && !IsNil(o.Mapped) {
		return true
	}

	return false
}

// SetMapped gets a reference to the given []Dns64groupMapped and assigns it to the Mapped field.
func (o *Dns64group) SetMapped(v []Dns64groupMapped) {
	o.Mapped = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Dns64group) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dns64group) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Dns64group) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Dns64group) SetName(v string) {
	o.Name = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *Dns64group) GetPrefix() string {
	if o == nil || IsNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dns64group) GetPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *Dns64group) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *Dns64group) SetPrefix(v string) {
	o.Prefix = &v
}

func (o Dns64group) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dns64group) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ref) {
		toSerialize["_ref"] = o.Ref
	}
	if !IsNil(o.Clients) {
		toSerialize["clients"] = o.Clients
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.Disable) {
		toSerialize["disable"] = o.Disable
	}
	if !IsNil(o.EnableDnssecDns64) {
		toSerialize["enable_dnssec_dns64"] = o.EnableDnssecDns64
	}
	if !IsNil(o.Exclude) {
		toSerialize["exclude"] = o.Exclude
	}
	if !IsNil(o.ExtAttrs) {
		toSerialize["extattrs"] = o.ExtAttrs
	}
	if !IsNil(o.Mapped) {
		toSerialize["mapped"] = o.Mapped
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	return toSerialize, nil
}

type NullableDns64group struct {
	value *Dns64group
	isSet bool
}

func (v NullableDns64group) Get() *Dns64group {
	return v.value
}

func (v *NullableDns64group) Set(val *Dns64group) {
	v.value = val
	v.isSet = true
}

func (v NullableDns64group) IsSet() bool {
	return v.isSet
}

func (v *NullableDns64group) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDns64group(val *Dns64group) *NullableDns64group {
	return &NullableDns64group{value: val, isSet: true}
}

func (v NullableDns64group) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDns64group) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
